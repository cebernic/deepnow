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

**DeepNow** 是一个专为企业级高可用场景打造的 AI 模型网关与知识融合底座。它不仅能将各种孤立的大语言模型（LLM/VLM) 和向量模型（Embedding）统一管理，还能为那些原本无状态、无长期记忆的 API 调用，**注入原生级别的 RAG 知识外脑与 SQL 级滑动窗口记忆**。

对于外部客户端（如 Chatbox、NextChat 、Openclaw 等）而言，**DeepNow 是完全透明的**。你只需将 Endpoint 替换为 DeepNow 的地址，即可瞬间让所有客户端拥有主备容灾、算力轮询以及深度的企业知识库支持，**100% 兼容 OpenAI API调用标准**。 

## 🌌 解决了目前Token时代的什么痛点？

1.[算力融合] -> 你可能正在苦恼我买了一堆Token算力，但发现没有一款可以让你的应用能够支撑急速且高并发的响应能力，因为普通的模型资源（非企业级）都是面向个人用户进行轻量推理使用的，一般都有很多限制（如RPM/TPM/RPD:Request Per minute/Tokens per minute/Peak requests per day等）

或者说购买了比较廉价的

但是由于供应商

任何技术新手皆可轻松的把手头所有Token资源融合为一体，通过编排融合把廉价的Token资源转外为可提供一体化的超级Token算力的服务底座。
或者模型商的并发限制、调用过程的各种规则限制（如RPM 

2.[高可用和规则突破] -> 不管你是个人龙虾openclaw，还是企业级的多人推理应用，谁都不希望在调用第三方模型时总是遭遇网络问题、Request per min 限制、频繁多人不同的IP地址限制导致的封号）或模型商突然由于各种原因导致的不可用等等，这些都可能极大的影响前端应用的稳定性特别是重创业务场景的口碑；同时频繁重试

前端应用调用大模型对失败的无感化

3.部署 -> Docker 等虚拟化满天飞，更新以及部署麻烦，且不能使用原生计算性能;安装和维护一堆组件头大

4.无感知识召回

5.模型无关性


<div align="center">
<img src="[此处填入后台 Dashboard 截图 URL]" alt="DeepNow Dashboard" width="800" />
<p><i>DeepNow 控制台 - 全局算力编排与监控</i></p>
</div>

---

## 🔥 核心特性 (Core Features)

### 🛡️ 1. 算力汇聚 (Compute Combined)
* **主备模型 (Primary/Slave Failover):** 支持生产环境的高可用模式。使大模型调用实现99.99%高可用，还可以使用聚合模型来实现主备。
* **模型聚合 (Token Aggregation):** 用>=2个相同的模型聚合，达到完全一致推理效果，同时还能分担负载突破TPM(TPS)、并发等相关限制。
* **混合轮询 (Hybrid Round-Robin):** 支持将多个异构大模型组合为“混合调度池”，虽然不同模型推理效果不同，但可把零散的模型资源全部利用，自动进行负载均衡，分散业务压力。
* **超混汇聚 (Super Hybrid):** 支持将多个异构大模型+聚合模型组合为“超混调度池”，系统会自动进行负载均衡，分散单一API的并发压力。


### 🧠 2. 动态探针 RAG 向量引擎 (Dynamic RAG Engine)
* **极致轻量底座:** 抛弃臃肿的独立向量数据库，底层采用 CGO 绑定的 `sqlite-vec`，实现单文件、高性能的亿级向量检索。
* **维度自适应探针 (Auto-Probe):** 无论你接入的是 768 维的旧模型，还是 3072 维的最前沿模型（如 `gemini-embedding-001`），系统在首次摄入知识时会自动探明embedding model的维度，并动态重构底层张量表结构。
* **时空溯源追踪:** 每一条 RAG 召回不仅提供知识切片，还精确携带属性信息，彻底消灭 AI 幻觉，且便于知识内容的维护，彻底实现类似Lora的能力。

### ⏳ 3. 滑动窗口记忆增强 (Stateful Memory Injection)
突破标准 `/v1/chat/completions` 接口的无状态限制。DeepNow 在网关层内置记忆存储能力，无需客户端维护历史上下文，DeepNow 都能基于 `Key + Session ID` 自动实施滑动窗口拦截，将历史对话无缝推给承接算力的底层大模型。

### 🔐 4. 细粒度资源管控 (Access Control & Stats)
支持无限生成以 `sk-deepnow-` 开头的专属令牌。可对每个令牌实施：
* 总 Token 消耗额度限制。
* 最高并发请求（Concurrency）硬拦截。
* 每日流量走势统计，以及精细的Token使用量记录，便于与大模型提供商对账

<div align="center">
<img src="[此处填入架构流程图 URL]" alt="DeepNow Architecture" width="800" />
<p><i>请求流转引擎底座架构图</i></p>
</div>

---

## 🚀 极速接入 (Quick Start)

DeepNow 采用全静态编译与资源内嵌技术，开箱即用，无需复杂的部署流程，除需配置大模型和Embedding模型的接口地址外，无需再配置或安装任何第三方组件；release 版无需docker或npm等相关环境安装、无任何三方依赖且跨平台，干净免维护一个二进制文件走天下（系统自带GUI dashboard/向量数据库/关系存储系统等）

```bash
# 1. 运行 DeepNow 服务端
./deepnow

# 服务将默认在 8443 端口启动自愈式 HTTPS 服务