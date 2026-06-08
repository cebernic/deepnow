// api_gateway.go - 算力中转与规则化
// 2025/12/07 HEAD Revision
// Only for test
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ChatRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

func handleGatewayCompletions(w http.ResponseWriter, r *http.Request) {
	// 检查主干大模型是否可用
	activePri := getActiveModel("LM", "primary")
	if activePri == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, `{"error": "AIBase_Server_Error: 系统未配置完整或无可用主干"}`)
		return
	}

	// 鉴权扣费
	authHeader := r.Header.Get("Authorization")
	clientKey := strings.TrimPrefix(authHeader, "Bearer ")
	used, exists := authKeys.Load(clientKey)
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"error":"Invalid API Key"}`)
		return
	}

	newUsed := used.(int) + 1
	authKeys.Store(clientKey, newUsed)
	keyDB.Exec(`UPDATE access_keys SET used=? WHERE api_key=?`, newUsed, clientKey)

	// 解析请求体，合并客户传来的 System Prompt，提取 User 提问
	var chatReq ChatRequest
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &chatReq)

	userPrompt := ""
	clientSysPrompt := ""
	for _, msg := range chatReq.Messages {
		if msg.Role == "system" {
			clientSysPrompt += msg.Content + "\n"
		}
		if msg.Role == "user" {
			userPrompt = msg.Content
		}
	}

	// 获取系统的 RAG 滑块挡位
	var ragSlider int
	err := sysDB.QueryRow(`SELECT value FROM system_rules WHERE key='rag_slider'`).Scan(&ragSlider)
	if err != nil {
		ragSlider = 1
	}

	if getActiveModel("Embedding", "") == nil {
		ragSlider = 0
	}

	// 执行 RAG 向量召回
	ragContext := ""
	if ragSlider > 0 {
		chunk, score := searchKnowledge(userPrompt)
		// 严苛过滤：相似度低于 0.4 直接清空
		if score > 0.4 {
			ragContext = chunk
		}
	}

	// 如果是强制 RAG (2档)，且没查到知识库，直接网关截停返回 _UNKNOWN_
	if ragSlider == 2 && ragContext == "" {
		w.Header().Set("Content-Type", "application/json")
		escapedContent, _ := json.Marshal("_UNKNOWN_")
		simulatedReply := fmt.Sprintf(`{"choices":[{"message":{"content":%s}}]}`, string(escapedContent))
		fmt.Fprint(w, simulatedReply)
		return
	}

	// 生成底层核弹级强制系统指令
	sysPrompt := buildRAGSystemPrompt(ragSlider, ragContext)

	if strings.TrimSpace(clientSysPrompt) != "" {
		sysPrompt = strings.TrimSpace(clientSysPrompt) + "\n\n--------------------\n" + sysPrompt
	}

	// 转交真实大模型引擎处理
	replyContent, err := callRealLM(sysPrompt, userPrompt, activePri)

	// 格式化输出为标准 OpenAI 格式
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Fprintf(w, `{"error": "AIBase_API_Error: %s"}`, err.Error())
	} else {
		escapedContent, _ := json.Marshal(replyContent)
		simulatedReply := fmt.Sprintf(`{"choices":[{"message":{"content":%s}}]}`, string(escapedContent))
		fmt.Fprint(w, simulatedReply)
	}
}
