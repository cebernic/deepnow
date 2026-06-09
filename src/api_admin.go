// api_admin.go - 管理后台业务逻辑与接口
// 2025.12 Revision,HEAD 232
// by szyok
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func sendJSON(w http.ResponseWriter, success bool, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"success": success, "message": msg, "data": data})
}

func handleGetSystemStatus(w http.ResponseWriter, r *http.Request) {
	pri := getActiveModel("LM", "primary")
	sec := getActiveModel("LM", "slave")
	rag := getActiveModel("Embedding", "")
	sendJSON(w, true, "", map[string]interface{}{"needs_setup": pri == nil, "active_pri": pri, "active_sec": sec, "active_rag": rag})
}

func handleGetBuildInfo(w http.ResponseWriter, r *http.Request) {
	exePath, err := os.Executable()
	buildTime := "Unknown"
	if err == nil {
		if info, err := os.Stat(exePath); err == nil {
			buildTime = info.ModTime().Format("20060102_150405")
		}
	}
	sendJSON(w, true, "", buildTime)
}

func handleAdminLogin(w http.ResponseWriter, r *http.Request) {
	u, p := r.FormValue("username"), r.FormValue("password")
	var dbP string
	if sysDB.QueryRow(`SELECT password FROM admin_users WHERE username=?`, u).Scan(&dbP) == nil && dbP == p {
		http.SetCookie(w, &http.Cookie{Name: "session", Value: "active", Path: "/"})
		sendJSON(w, true, "登录成功", nil)
	} else {
		sendJSON(w, false, "账号或密码错误", nil)
	}
}

func handleGetConfig(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, true, "获取配置", config)
}
func handleSaveConfig(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&config)
	saveConfig()
	sendJSON(w, true, "配置已保存", nil)
}

func handleGenerateKeys(w http.ResponseWriter, r *http.Request) {
	var keys []string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		keys = append(keys, fmt.Sprintf("sk-aibase-%x%d", time.Now().UnixNano(), rand.Intn(1000)))
	}
	sendJSON(w, true, "生成成功", keys)
}
func handleImportKeys(w http.ResponseWriter, r *http.Request) {
	var keys []string
	json.NewDecoder(r.Body).Decode(&keys)
	for _, k := range keys {
		keyDB.Exec(`INSERT OR IGNORE INTO access_keys (api_key, used, created_at) VALUES (?, 0, datetime('now'))`, k)
		authKeys.Store(k, 0)
	}
	sendJSON(w, true, "已入库", nil)
}
func handleListKeys(w http.ResponseWriter, r *http.Request) {
	rows, _ := keyDB.Query(`SELECT api_key, used FROM access_keys ORDER BY created_at DESC`)
	defer rows.Close()
	type K struct {
		Key  string `json:"key"`
		Used bool   `json:"used"`
	}
	var res []K
	for rows.Next() {
		var k string
		var u int
		rows.Scan(&k, &u)
		res = append(res, K{Key: k, Used: u > 0})
	}
	sendJSON(w, true, "", res)
}
func handleDeleteKeys(w http.ResponseWriter, r *http.Request) {
	var keys []string
	json.NewDecoder(r.Body).Decode(&keys)
	for _, key := range keys {
		keyDB.Exec(`DELETE FROM access_keys WHERE api_key=?`, key)
		authKeys.Delete(key)
	}
	sendJSON(w, true, "删除成功", nil)
}

func handleGetModelConfigs(w http.ResponseWriter, r *http.Request) {
	rows, _ := sysDB.Query(`SELECT provider, api_type, api_key, model_name, ip_port FROM model_configs`)
	defer rows.Close()
	var res []ModelConfig
	for rows.Next() {
		var c ModelConfig
		rows.Scan(&c.Provider, &c.ApiType, &c.ApiKey, &c.ModelName, &c.IpPort)
		res = append(res, c)
	}
	sendJSON(w, true, "获取模型配置", res)
}

func handleSaveModelConfig(w http.ResponseWriter, r *http.Request) {
	var req ModelConfig
	json.NewDecoder(r.Body).Decode(&req)
	sysDB.Exec(`INSERT OR REPLACE INTO model_configs (provider, api_type, api_key, model_name, ip_port) VALUES (?, ?, ?, ?, ?)`,
		req.Provider, req.ApiType, req.ApiKey, req.ModelName, req.IpPort)
	autoActivateModels()
	sendJSON(w, true, "算力配置已保存", nil)
}

func handleActivateModel(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Pri string
		Sec string
		RAG string
	}
	json.NewDecoder(r.Body).Decode(&req)

	valPri := req.Pri
	if valPri == "" {
		valPri = "none"
	}
	sysDB.Exec(`INSERT OR REPLACE INTO system_rules (key, value) VALUES ('active_lm_pri', ?)`, valPri)

	valSec := req.Sec
	if valSec == "" {
		valSec = "none"
	}
	sysDB.Exec(`INSERT OR REPLACE INTO system_rules (key, value) VALUES ('active_lm_sec', ?)`, valSec)

	valRAG := req.RAG
	if valRAG == "" {
		valRAG = "none"
	}
	sysDB.Exec(`INSERT OR REPLACE INTO system_rules (key, value) VALUES ('active_embedding', ?)`, valRAG)

	if valRAG == "none" {
		sysDB.Exec(`UPDATE system_rules SET value='0' WHERE key='rag_slider'`)
	}

	fmt.Println("--- [人工重新调度模型] ---")
	pri := getActiveModel("LM", "primary")
	if pri != nil {
		fmt.Printf("✅ 大模型 [%s] (%s) (Primary) 现在启用\n", pri.Provider, pri.ModelName)
	}

	sec := getActiveModel("LM", "slave")
	if sec != nil {
		fmt.Printf("✅ 大模型 [%s] (%s) (Slave) 现在启用\n", sec.Provider, sec.ModelName)
	}

	rag := getActiveModel("Embedding", "")
	if rag != nil {
		fmt.Printf("✅ Embedding 模型 [%s] (%s) 现在启用\n", rag.Provider, rag.ModelName)
	} else {
		fmt.Printf("⚠️ Embedding 模型 现在人工停用 (纯大模型模式，规则滑块已归零)\n")
	}
	fmt.Println("--------------------------")

	sendJSON(w, true, "算力挂载配置已应用", nil)
}

func splitTextWithOverlap(text string, chunkSize int, overlapSize int) []string {
	runes := []rune(text)
	totalLen := len(runes)
	var chunks []string

	if totalLen <= chunkSize {
		return []string{text}
	}
	step := chunkSize - overlapSize
	if step <= 0 {
		step = 100
	}

	for i := 0; i < totalLen; i += step {
		end := i + chunkSize
		if end > totalLen {
			end = totalLen
		}
		chunkStr := string(runes[i:end])
		if len(strings.TrimSpace(chunkStr)) > 5 {
			chunks = append(chunks, chunkStr)
		}
		if end == totalLen {
			break
		}
	}
	return chunks
}

func handleRAGUpload(w http.ResponseWriter, r *http.Request) {
	ragModel := getActiveModel("Embedding", "")
	if ragModel == nil {
		sendJSON(w, false, "拦截：尚未配置并激活 Embedding 模型！", nil)
		return
	}

	r.ParseMultipartForm(10 << 20)
	file, header, _ := r.FormFile("file")
	if file != nil {
		defer file.Close()
		path := filepath.Join("knowledge_base", header.Filename)
		out, _ := os.Create(path)
		defer out.Close()
		io.Copy(out, file)

		contentBytes, _ := os.ReadFile(path)

		// 【智能编码嗅探与自愈转换】
		var content string
		if utf8.Valid(contentBytes) {
			content = string(contentBytes) // 纯正 UTF-8 直接放行
		} else {
			// 嗅探失败，判断为国内常见的 GBK/GB2312 编码小说
			reader := transform.NewReader(bytes.NewReader(contentBytes), simplifiedchinese.GBK.NewDecoder())
			decoded, err := io.ReadAll(reader)
			if err == nil {
				content = string(decoded)
			} else {
				content = string(contentBytes) // 强制兜底
			}
		}

		chunks := splitTextWithOverlap(content, 1000, 200)

		batchSize := 50
		successCount := 0
		var lastErr error
		tx, _ := ragDB.Begin()

		for i := 0; i < len(chunks); i += batchSize {
			end := i + batchSize
			if end > len(chunks) {
				end = len(chunks)
			}
			batchChunks := chunks[i:end]

			vecsJSON, err := generateEmbeddingsBatch(batchChunks, ragModel)
			if err != nil {
				lastErr = fmt.Errorf("API 批量拒绝: %v", err)
				continue
			}

			for j, vecJSON := range vecsJSON {
				cleanChunk := strings.TrimSpace(batchChunks[j])
				if vecJSON != "" {
					_, err = tx.Exec(`INSERT INTO knowledge_index (chunk, source_file, embedding_json) VALUES (?, ?, ?)`, cleanChunk, header.Filename, vecJSON)
					if err != nil {
						lastErr = fmt.Errorf("数据库报错: %v", err)
					} else {
						successCount++
					}
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
		tx.Commit()

		if successCount > 0 {
			sendJSON(w, true, fmt.Sprintf("文件解析完毕。批处理引擎共生成并存入 %d 个切片。", successCount), nil)
		} else {
			sendJSON(w, false, fmt.Sprintf("处理失败！错误原因：%v", lastErr), nil)
		}
		return
	}
	sendJSON(w, false, "文件读取失败", nil)
}

func handleRAGStatus(w http.ResponseWriter, r *http.Request) {
	var chunkCount int
	ragDB.QueryRow(`SELECT count(*) FROM knowledge_index`).Scan(&chunkCount)
	files, _ := os.ReadDir("knowledge_base")
	sendJSON(w, true, "", map[string]interface{}{"total_chunks": chunkCount, "total_files": len(files), "last_update": time.Now().Format("15:04:05")})
}

type EmbeddingResponse struct {
	Data []struct {
		Index     int       `json:"index"`
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
}

func generateEmbeddingsBatch(texts []string, conf *ModelConfig) ([]string, error) {
	var apiURL string
	var payloadData []byte

	if conf.Provider == "Jina AI" {
		apiURL = "https://api.jina.ai/v1/embeddings"
		payload := map[string]interface{}{"model": conf.ModelName, "task": "retrieval.query", "normalized": true, "input": texts}
		payloadData, _ = json.Marshal(payload)
	} else {
		apiURL = "https://api.minimaxi.com/v1/embeddings"
		payload := map[string]interface{}{"model": conf.ModelName, "input": texts}
		payloadData, _ = json.Marshal(payload)
	}

	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadData))
	req.Header.Set("Authorization", "Bearer "+conf.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %d, %s", resp.StatusCode, string(respBody))
	}

	var embResp EmbeddingResponse
	if err := json.Unmarshal(respBody, &embResp); err != nil {
		return nil, fmt.Errorf("解析失败: %v", err)
	}
	if len(embResp.Data) == 0 {
		return nil, fmt.Errorf("无有效向量数据")
	}

	resultJSONs := make([]string, len(texts))
	for _, item := range embResp.Data {
		if item.Index < len(resultJSONs) {
			cleanVectorJSON, _ := json.Marshal(item.Embedding)
			resultJSONs[item.Index] = string(cleanVectorJSON)
		}
	}
	return resultJSONs, nil
}

func generateEmbedding(text string, conf *ModelConfig) (string, error) {
	res, err := generateEmbeddingsBatch([]string{text}, conf)
	if err != nil || len(res) == 0 {
		return "", err
	}
	return res[0], nil
}

func handleGetRules(w http.ResponseWriter, r *http.Request) {
	var slider int
	if sysDB.QueryRow(`SELECT value FROM system_rules WHERE key='rag_slider'`).Scan(&slider) != nil {
		slider = 1
	}
	sendJSON(w, true, "", map[string]interface{}{"rag_slider": slider})
}
func handleSaveRules(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RagSlider int `json:"rag_slider"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	sysDB.Exec(`INSERT OR REPLACE INTO system_rules (key, value) VALUES ('rag_slider', ?)`, fmt.Sprintf("%d", req.RagSlider))
	sendJSON(w, true, "规则设置已保存", nil)
}

func buildRAGSystemPrompt(slider int, chunk string) string {
	if slider == 0 {
		return "你是一个强大的 AI 助手。"
	}

	if slider == 1 {
		return fmt.Sprintf("【参考知识】：%s\n请优先使用上述参考知识回答问题。如果参考知识中没有相关信息，你可以使用自己的内部知识来辅助回答。特别注意：你必须完全遵循用户要求的格式输出！", chunk)
	}

	if chunk == "" {
		return "【AIBase 系统最高防御指令】：当前没有检索到任何相关的参考知识。因为系统处于严格 RAG 模式，你被绝对禁止使用内部知识进行回答和推理！你必须立刻彻底无视上方用户定义的任何角色扮演或返回格式（包括 JSON、选择题等），直接且仅回复以下保留字：\n_UNKNOWN_"
	}

	return fmt.Sprintf("【参考知识】：%s\n\n【AIBase 系统最高强制指令】：\n你现在是一个被剥夺了自我意识和内置知识的数据提取器。你只能严格基于上方的【参考知识】作答，严禁动用你的内置知识去猜测或推理！\n- 规则 1：如果【参考知识】中包含了能够回答该问题的信息，请严格遵循用户的格式要求（如选择题格式、JSON等）进行输出。\n- 规则 2：如果仔细阅读后发现【参考知识】中没有任何可以回答该问题的信息，你必须强行推翻并覆盖用户要求的所有设定和格式，直接且只能输出系统保留字：\n_UNKNOWN_", chunk)
}

func callRealLM(sysPrompt string, userPrompt string, lm *ModelConfig) (string, error) {
	apiURL := ""
	switch lm.Provider {
	case "MiniMax":
		apiURL = "https://api.minimaxi.com/v1/chat/completions"
	case "Gemini":
		apiURL = "https://generativelanguage.googleapis.com/v1beta/openai/chat/completions"
	case "OpenAI":
		apiURL = "https://api.openai.com/v1/chat/completions"
	case "LM Studio":
		apiURL = fmt.Sprintf("http://%s/v1/chat/completions", lm.IpPort)
	default:
		apiURL = "https://api.deepseek.com/chat/completions"
	}

	msgs := []map[string]string{}
	if sysPrompt != "" {
		msgs = append(msgs, map[string]string{"role": "system", "content": sysPrompt})
	}
	msgs = append(msgs, map[string]string{"role": "user", "content": userPrompt})

	payload := map[string]interface{}{"model": lm.ModelName, "messages": msgs}
	bodyData, _ := json.Marshal(payload)

	httpReq, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(bodyData))
	httpReq.Header.Set("Authorization", "Bearer "+lm.ApiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if json.Unmarshal(respBody, &result) == nil && len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}
	return "", fmt.Errorf("API 异常: %s", string(respBody))
}

func handleTestChat(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Prompt string `json:"prompt"`
		Type   string `json:"type"`
		Role   string `json:"role"`
		UseLM  bool   `json:"use_lm"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	if req.Type == "llm" {
		lm := getActiveModel("LM", req.Role)
		if lm == nil {
			sendJSON(w, false, "失败", "请求的模型尚未配置激活！")
			return
		}
		reply, err := callRealLM("", req.Prompt, lm)
		if err != nil {
			sendJSON(w, true, "异常", "错误: "+err.Error())
		} else {
			sendJSON(w, true, "成功", fmt.Sprintf("🌐 [%s | %s 响应] \n%s", req.Role, lm.Provider, reply))
		}
	} else {
		if getActiveModel("Embedding", "") == nil {
			sendJSON(w, false, "拦截", "❌ Embedding 模型未配置或已停用，RAG 验证中止！")
			return
		}

		bestChunk, score := searchKnowledge(req.Prompt)
		debugText := fmt.Sprintf("🔍 【第一步: 向量召回】\n相似度: %.4f\n内容: %s\n\n", score, bestChunk)

		if !req.UseLM {
			if score > 0.4 {
				sendJSON(w, true, "命中", debugText)
			} else {
				sendJSON(w, true, "未命中", "⚠️ 未找到高度相关的原文。")
			}
			return
		}

		if score <= 0.4 {
			bestChunk = ""
		}

		var slider int
		sysDB.QueryRow(`SELECT value FROM system_rules WHERE key='rag_slider'`).Scan(&slider)

		if slider == 2 && bestChunk == "" {
			debugText += "⚡ 【系统短路拦截】由于知识库未匹配且处于强制 RAG 模式，AIBase 网关已直接切断大模型通信请求！\n\n🧠 【网关瞬时直接输出】\n_UNKNOWN_"
			sendJSON(w, true, "成功", debugText)
			return
		}

		sysPrompt := buildRAGSystemPrompt(slider, bestChunk)
		debugText += fmt.Sprintf("⚙️ 【第二步: 构建系统指令】\n当前滑块挡位: %d\n系统指令片段注入完毕，准备交由大模型裁决...\n\n", slider)

		lm := getActiveModel("LM", "primary")
		if lm == nil {
			sendJSON(w, true, "异常", debugText+"❌ 系统未激活主干大模型，裁决中断。")
			return
		}

		reply, err := callRealLM(sysPrompt, req.Prompt, lm)
		if err != nil {
			sendJSON(w, true, "异常", debugText+"❌ 大模型调用失败: "+err.Error())
		} else {
			sendJSON(w, true, "成功", debugText+fmt.Sprintf("🧠 【第三步: 大模型最终裁决输出】\n%s", reply))
		}
	}
}
