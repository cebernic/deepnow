# ⚡ DeepNow(深脑) : The Ultimate AI Compute Gateway 

<div align="center">

<img src="[此处填入你的 Logo 图片 URL]" alt="DeepNow Logo" width="200" />


**全场景聚合算力网关 & 智能 RAG 融合基座**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)
[![OpenAI Compatible](https://img.shields.io/badge/OpenAI-100%25_Compatible-412991?style=for-the-badge&logo=openai)](#)
[![SQLite-Vec](https://img.shields.io/badge/Vector_DB-SQLite--Vec-success?style=for-the-badge)](#)

*打造极致稳定、无缝接入的私有化 AI 算力中枢。*

</div>

---

## 🌌 什么是 DeepNow？

**DeepNow** 是一个专为企业级高可用场景打造的 AI 模型网关与知识融合底座。它不仅能将各种孤立的大语言模型（LLM）和向量模型（Embedding）统一管理，还能为那些原本无状态、无长期记忆的 API 调用，**注入原生级别的 RAG 知识外脑与 SQL 级滑动窗口记忆**。

对于外部客户端（如 Chatbox、NextChat 等）而言，**DeepNow 是完全透明的**。你只需将 Endpoint 替换为 DeepNow 的地址，即可瞬间让所有客户端拥有主备容灾、算力轮询以及深度的企业知识库支持，**100% 兼容 OpenAI 接口标准**。

<div align="center">
<img src="[此处填入后台 Dashboard 截图 URL]" alt="DeepNow Dashboard" width="800" />
<p><i>DeepNow 控制台 - 全局算力编排与监控</i></p>
</div>

---

## 🔥 核心特性 (Core Features)

### 🛡️ 1. 战区级算力编排 (Compute Orchestration)
* **主干备援 (Primary/Slave Failover):** 狙击枪级的单点探测与战区级的火力覆盖。当主模型（Primary）因网络或官方 API 故障宕机时，系统将在毫秒级静默切换至备用模型（Slave），客户端完全无感，确保业务 99.99% 可用。
* **混合轮询 (Hybrid Round-Robin):** 支持将多个异构大模型组合为“混合调度池”，系统会自动进行负载均衡，分散单一 API 的并发压力。

### 🧠 2. 动态探针 RAG 向量引擎 (Dynamic RAG Engine)
* **极致轻量底座:** 抛弃臃肿的独立向量数据库，底层采用 CGO 绑定的 `sqlite-vec`，实现单文件、高性能的亿级向量检索。
* **维度自适应探针 (Auto-Probe):** 无论你接入的是 768 维的旧模型，还是 3072 维的最前沿模型（如 `gemini-embedding-001`），系统在首次摄入知识时会自动发射探针探明维度，并动态重构底层张量表结构。
* **时空溯源追踪:** 每一条 RAG 召回不仅提供知识切片，还精确携带 UNIX 级上传时间戳与来源文件名称，彻底消灭 AI 幻觉。

### ⏳ 3. 滑动窗口记忆增强 (Stateful Memory Injection)
突破标准 `/v1/chat/completions` 接口的无状态限制。DeepNow 在网关层内置 SQL 关系型明细存储。无论客户端是否携带历史上下文，DeepNow 都能基于 `API Key + Session ID` 自动实施滑动窗口拦截，将历史对话无缝推给承接算力的底层大模型。

### 🔐 4. 细粒度资源管控 (Access Control & Stats)
支持无限生成以 `sk-deepnow-` 开头的专属令牌。可对每个令牌实施：
* 总 Token 消耗额度限制。
* 最高并发请求（Concurrency）硬拦截。
* 每日流量走势统计。

<div align="center">
<img src="[此处填入架构流程图 URL]" alt="DeepNow Architecture" width="800" />
<p><i>请求流转引擎底座架构图</i></p>
</div>

---

## 🚀 极速接入 (Quick Start)

DeepNow 采用全静态编译与资源内嵌技术，前端 HTML/JS/CSS 资产已通过 `//go:embed` 彻底打包进单一可执行文件中。**只需一个二进制文件，开箱即用。**

```bash
# 1. 运行 DeepNow 服务端
./deepnow

# 服务将默认在 8443 端口启动自愈式 HTTPS 服务