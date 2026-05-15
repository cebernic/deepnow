# ⚡ DeepNow(深脑) : The Ultimate AI Compute Gateway 

<div align="center">

<img src="./assets/logo.png" alt="DeepNow Logo" width="200" />

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

## 🌌 Deepnow 在Token为王的时代，解决了哪些痛点？

1.[算力融合] -> 你可能正在苦恼买了一堆Token算力，但发现没有一款可以让你的应用能够支撑急速且高并发的响应能力。因为普通的模型资源（非企业级）都是面向个人用户进行轻量推理使用的，一般都有很多限制（如RPM/TPM/RPD:Request Per minute/Tokens per minute/Peak requests per day等）你的应用不管如何优化都避不开模型提供商的这些限制！但 Deepnow 的出现可使你立刻摆脱这些烦恼。通过算力融合的技术，让你只需要购买廉价的个人Token算力（或code plan此类极便宜套餐）来实现综合能力达到或者超过企业算力，甚至你还可以共享这些算力给其它应用或任意第三方。当然，你也可以把本地多个不同小设备上部署的本地模型聚合起来集中使用，达到一个接口同时服务多个应用的目的，而不必为了并发和token输出性能购买价格昂贵的专业算力设备，彻底实现token自由。

2.[高可用、可共享] -> 不管你是个人龙虾openclaw，还是企业级的多人推理应用，谁都不希望在调用模型时总是遭遇网络问题或遇到模型商发来的诸如“当前算力高峰，请稍后再试...”等等恼人的错误信息，这不仅可能导致一系列后续任务失败，还可能造成应用处理任务时的极大延迟；同时，当你买了一个按次结算的套餐后在自己不使用时为了避免浪费又希望共享给朋友使用，但多个IP调用同一个KEY可能会被模型商认为key泄露或者是触犯其规则的行为带来封号等严重后果。最重要的是，你的key可能会被朋友泄露而造成责任界定不清晰，只能频繁更换Key来保障安全，而 Deepnow 则可轻松解决这个问题，甚至还可以计算token共用人的用量，使多人平摊token使用费用成为可能。

3.[极简部署、高强壮、高性能] -> 所有人都知道任何虚拟化技术（即便轻虚技术docker）都会或多或少带来不必要的开销，且诸如 Docker 等打包封装环境部署本身安装过程繁杂且易环境问题报错，为了部署和使用一款软件，翻阅各种资料折腾几小时甚至几天的这种经历相信大家都有。另外，开源社区的软件往往为了便于迭代，基本服务型系统都会有非常多的依赖环境，即便部署其最终编译亦或是docker版可能还需要部署各类三方组件和编辑各种各样的烦人config才能顺利运行，但deepnow的目标是一个可执行文件足以，配置也是全图形化h5面板操作，先运行起来再精细化配置是deepnow的核心特点。

4.[无感的知识库挂载能力] -> 不管是个人还是企业可能都有一些私有知识需要记忆和检索。deepnow 可以无缝利用任一大模型来检索这些知识，前端无需再开发复杂的召回应用，可以在推理服务的同时直接召回并插入到推理结果中。还可以根据检索等级来决定当检索无命中的清空下直接用模型来推理回答，而且聚合和轮询算法同样可以作用于知识检索，使每种模型都不可能去摄取太多的内部知识，还间接降低了泄密的可能性。

5.[模型无关性] -> Deepnow 的核心理念就是要让前端应用彻底与模型剥离，不依赖具体某个模型，甚至还可以热切换模型。模型成为计算单元（类似CPU的存在），deepnow 可以通过主动学习和被动喂给知识在专业场景中拉平任何模型的推理结果一致性。

6.[分布式无中心化算力供给] -> Deepnow 之间也可以相互聚合，除了算力共享外其记忆和知识系统也可以被任意deepnow系统使用，因为上下文记忆和知识的召回是高度融合在推理结果中的。你可以搭建不同用途的deepnow使其学习不同的专业性知识，并用一个deepnow聚合不同的专业deepnow 使一个或多个deepnow 具备综合知识能力，也可以分别使用他们。


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
* **极致轻量底座:** 无需搭建和维护臃肿的独立向量数据库，底层采用 CGO 绑定的 `sqlite-vec`，实现高性能的亿级向量检索。
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