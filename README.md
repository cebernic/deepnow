# ⚡ DeepNow(深脑) : The Ultimate AI Compute Gateway ( Router )

<div align="center">

<img src="./assets/logo.png" alt="DeepNow Logo" width="200" />

**全场景聚合 AI 路由(网关) & 智能 RAG 融合基座**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)
[![OpenAI Compatible](https://img.shields.io/badge/OpenAI-100%25_Compatible-412991?style=for-the-badge&logo=openai)](#)
[![SQLite-Vec](https://img.shields.io/badge/Vector_DB-SQLite--Vec-success?style=for-the-badge)](#)
[![Zig CC](https://img.shields.io/badge/CGO_Cross--Compiled-Zig_CC-F7A41D?style=for-the-badge&logo=zig)](https://ziglang.org/)

*打造极致稳定、无缝接入的私有化 AI 算力中枢。*

</div>

## 🌌 什么是 DeepNow

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 在 __Token__ 为王的当下，**Deepnow** 是一个专为(个人/企业)打造，面向高可用、高并发，全场景使用的 AI 模型网关(路由)与知识融合底座，它不仅能将各种孤立的大语言模型（LLM/VLM）和向量模型（Embedding）统一管理、调度，还可以通过绑定算法使他们聚合使用，实现算力整合与容灾。通过把多家模型运营商的资源整合利用，最大化为你的前端应用或开发场景提供最强劲的Token动力。通过聚合，你可以轻松突破各家模型运营商的各类TPM、并发等限制，不管是面向多人团队 vibe coding 或是长文本融合应用，亦或者是高密集度突发调用，Deepnow 都可以实际解决你的问题，基本上 Deepnow 就像是在AI时代你的私人软路由器(Soft Router)，只不过这个 Router 管理的流量不再是 TCP/IP 网路包，而是 Token 流量，所以你可以把 Deepnow 看成是一种全新的 TR 或 TG (Token Router / Token Gateway)。

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 使用 Deepnow 后，你可以把前端所有的app应用或者开发工具的 Token 端点(Endpoint)指向 Deepnow ，以实现对所有应用、特定应用的流量走向、负载均衡等统一控制，还可以随时在线热切换模型，你不必为每一个应用都去单独配置某家Token运营商的API key，所有应用都统一配置为 Deepnow 生成的 key 和 Endpoint 地址即可。当你希望切换运营商时又要去修改每一个应用的 Endpoint URL 和 API Key，如果是面向用户的应用，那这种切换会变得成本极其高昂。你甚至还可以让前端应用自由选择 Deepnow 系统中已配置的所有模型，实现如 OpenRouter 网关应用一样的自定义多模型功能，但 Openrouter 本身也是一个远端多模型网关，所有的端点都在远程，你完全无法掌控，一个在本地运行的多模型AI网关才是令人安心的，更不会有网络抖动或安全带来的问题。

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 但，Deepnow 的能力远不止于这些。

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 如果你是一个团队(企业)管理员，你肯定不希望给每一个员工都单独购买一个私人 Key，因为模型提供商的 Key 如果被员工随意分享可能会给企业带来巨大损失；同样，你肯定还希望能够实时了解每一个员工或者某个应用的流量使用情况，或者在某些特定场景希望多人分摊算力和成本但有苦于没有工具实现。如今， Deepnow 都可以帮到你，因为 Deepnow 可以生成和分发自己的 Key 给使用者，并且可以基于 Key 来绑定特定的并发能力、模型选择和设定流量阈值等，还可以随时收回。

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 如果你希望搭建一个业务机器人，你也不需要额外的工具。不管是面向企业内部知识，还是让AI承袭某些业务经验，你都能够以近乎傻瓜式的操作方式直接喂给 Deepnow ，你完全不用有任何模型训练能力，你就拥有了一个拿世界顶级模型大脑来读懂和学会你私有知识和经验的解决方案，瞬间就成为一个懂你的私有化AI和数字员工，在多个应用调用你的私有化AI时都能承袭你的私人知识和对这种知识的逻辑化推理和响应的能力。在培训、医疗、科研、产品功能展示，后端维护等等诸多严禁知识幻觉的场景，Deepnow 都可以随时为你打造。

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 市面上眼花缭乱的智能体和天女散花般的模型，导致前端、后端可能都有各自的协议，而随着时代的发展，以往全面被兼容的 OpenAI v1/completions 协议正在被行业抛弃，随之取代的是更加高级的 Responses 和 Messages 协议，但这些协议之间往往都有各家自己的特色和“私货”在兼容性方面难以互相兼顾，而 Deepnow 却可以轻松帮你解决这些问题。你可以不用购买额外的远端服务，在本地就可以轻松实现 Openclaw / Codex / Hermes 主流智能体挂载任何大模型的能力，且你可以为这些智能体开启更加高级的并发模式，成倍缩短智能体解决综合事物的时间。

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 未来，Deepnow 还将具备整合mcp服务、Plugin、ServerSide task、模型无关的自有上下文维护、Severside skills、多模态聚合（即：你可以用不同能力的模型拼接在一起，实现图片声音识别、图形生成的总响应能力，因为我们都知道世上没有一个模型在所有领域都可以同时顶尖，我们需要按需分配AI算力路由）等等这些能力。

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 总之，Deepnow 将会是一种AI基础设施，一种强大的算力底座，一种可挂载各种先进能力的 AI 流量路由器。

## ⚡ DeepNow 解决了哪些痛点？


<details>
  <summary><b>🚀 高可用、极致算力融合，突破并发限制</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;你可能正在苦恼：买了一堆 Token 算力，却没有一款能支撑起你的应用实现极速、高并发的响应。普通的非企业级模型资源通常有严格的速率限制（如 `RPM` 限制每分钟请求数、`TPM` 限制每分钟 Token 数、`RPD` 限制每日请求数等），无论应用端如何优化，都无法避开官方的接口限流。<b>DeepNow 彻底打破了这一枷锁</b>： 通过独创的算力融合技术，你只需组合多个廉价的个人 Token 算力（或极具性价比的开发者套餐），就能使其综合吞吐量达到甚至超越企业级算力标准。此外，你还可以将不同物理设备上部署的本地开源模型聚合起来集中调度，实现一个接口同时服务多个应用，无需购买昂贵的专业级算力设备，**彻底实现 Token 自由**。同时，在需求高可用的业务场景，前端应用都不想收到任何 ""...high demand, please try again later..." 此类的尴尬提示。此类非正常应答都会被 deepnow 捕获，当意外应答出现，将会全自动的使用备用模型重发请求，前端是无感的。
</details>
<details>
  <summary><b>🧠 实现优答 (Token-Aware Optimal Routing)</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Deepnow 除了有同模型聚合、多模型混合、灾备等调度算法外，还有一套自主独有的“优答”算法，该算法的核心是对2种业务场景实现最优解。<br><br>
  1）推理速度优先：用户可以绑定>=2个极速且廉价模型，deepnow 会把一个推理请求同时分发给N个极速且轻量化模型(可以是本地)，任何一个有应答即刻把应答内容返回请求者，丢弃其它应答；<br><br>
  2）推理层级排序 Re-rank：用户可以把多个模型按照其推理能力排列应答次序。首先挂载分析模型，由一个轻量级大模型先对问题的意图（Intent）和关键词特征信息，并并将其归类（Classification）后推给分类不同层级的大模型来按需实现推理难度的解答，从而做到精细化的Token付费和更准确的推理结果。
</details>
<details>
  <summary><b>🛡️ 算力共享</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;无论面向个人AI geeker（如重度 Hermes/ OpenClaw 玩家），还是企业级多人AI应用，Deepnow 都可以实现同一算力资源分享给不同应用和不同使用者。因为 Deepnow 自己可以生成和管理一套自己的鉴权系统 (API Key) 你可以让一个Key共用给你的多个应用，还可以让多个Key分享给多个人。以往，多人异地 IP 使用同模型Key调用同一模型商极易触发风控，导致封号或隐私泄露，责任难以界定。而使用Deepnow 后真实调用者将被隐藏，对模型商来说所有调用均来自 Deepnow，这就让共享变得安全可控。 Deepnow 还可以基于自身分发的Key统计Token进出流量，并写入内、外部数据库（sqlite/mysql) 你可以自己开发基于数据库的Token管理UI实现流量二次计费、管理等。
</details>
<details>
  <summary><b>📦 敏捷部署与强大性能（零依赖，开箱即用）</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;当前各种开源AI应用大部分使用Python开发，除了要安装各种错综复杂的依赖库以外(甚至依赖库版本错误也可能无法启动)，还有承受python天生的性能不佳问题。而 Deepnow 一开始的目标就是不对性能妥协的对标企业级承载应用，使用编译型Go和C语言混合开发，同时内置所有组件而无需额外安装，就连UI界面也是打包进本体而无需额外挂载，真正做到“下载和执行”2步开箱即尝。没有细碎的组件和依赖库，更没有沉重的 Docker ，你只要选好针对目标平台的分发二进制文件，一个文件你就可以在N个设备上1秒部署并运行。无需手动编写复杂的配置文件，所有设置均可通过极其直观的可视化 H5 管理面板完成。先让服务跑起来，再进行精细化调优，这是 DeepNow 针对敏捷部署设计的极致体验。
</details>
<details>
  <summary><b>🧠 无感的知识库挂载（插拔式大脑）</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;无论是个人还是企业，都有将其私有知识转化为 AI 记忆的需求。<b>DeepNow 可以无缝调度任意大模型来处理这些知识</b>。前端应用无需再开发复杂的 RAG（检索增强生成）召回逻辑，DeepNow 会在执行推理服务的同时，直接在底层召回相关知识并无缝织入上下文中。系统支持自定义检索兜底策略，如：在未命中任何知识的 **情况下**，直接拦截或降级使用大模型自身知识作答。更巧妙的是，算力轮询架构同样作用于知识检索过程，这使得单一模型提供商无法获取你完整的上下文，从物理层面间接降低了数据链整体泄密的风险。DeepNow还能针对0幻想为前提(temp 0)在专业垂直场景，使用Inject Prompt技术将不同能力模型的推理输出结果拉齐一致性返回，让你的私有专业知识召回时不管用任何模型都按请求的精确格式。Deepnow 知识系统将会支持图文、音频的多模态召回，且每一种即可使用单独专业模型交叉混合输出不同模态的内容，也可以独立使用多模态模型直接实现，但在最终推理结果上却有实现融合。无论对文本知识，还是结构化的知识图谱都能完成学习并可以合并用于推理任务，从而让智能体 Agent 等前端应用实现完全定制化的精确任务流。<br><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**DeepNow** 的核心理念是让“私有逻辑”与“推理模型”彻底解耦。在这里，大模型退化为纯粹的计算单元、属于本能认知层，底层知识才是你专属AI的大脑。 通过 DeepNow 统一的网关调度，你不仅可以随时热拔插、切换远端底层模型或本地模型，更可以切换不同种类的底层知识挂载，甚至对知识片段单独实现移除和重恢复以适应各种场景。其核心技术是对知识语义(包含图像)的筛略、通过向量空间比对、重索引和排级(Reindex/ReRank)、语义扩张命中等多个维度实现知识的定位和实时知识流的推理。达到类似预训练私有模型的效果，却无需付出训练私有模型的成本。<br><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;此外，Deepnow 还将具备自我知识学习和存储的能力，可以根据权限设置把所有请求内容转换成记忆，且无长度限制。
</details>
<details>
  <summary><b>🌐 树状分布算力网（无限裂变与级联）</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**DeepNow** 节点之间可以相互级联！除了纯粹的算力共享外，其状态记忆与私有知识系统也能被其他 DeepNow 实例无缝调用，因为上下文和知识召回已经高度融合在返回的流式响应中。你可以构建多个专精于不同领域的 DeepNow 节点（例如：A节点挂载的是企业员工的知识图谱大脑，B节点挂载的是企业产品大脑，C节点为公司经营分析大脑）以一个上层节点将它们聚合，从而打造出一个具备全领域综合能力的超级私有化 AI 中枢。同时，前端应用可以分权限、需求连接A或B节点或者为最顶层超级节点。
</details>

## ⚡ 部分UI截图
<details>
  <summary>::Screenshots ::</summary>
  <div align="center">
    <img src="./assets/ss1.jpg" alt="DeepNow Dashboard Overview" width="400" />
    <p><i>网络配置</i></p>
  </div>
  <div align="center">
    <img src="./assets/ss2.jpg" alt="DeepNow Dashboard RAG Settings" width="400" />
    <p><i>算力配置</i></p>
  </div>
  <div align="center">
    <img src="./assets/ss4.jpg" alt="DeepNow Dashboard Stateful Memory" width="400" />
    <p><i>记忆策略配置</i></p>
  </div>
  <div align="center">
    <img src="./assets/ss5.jpg" alt="DeepNow Dashboard Stats & Authentication" width="400" />
    <p><i>Key全景</i></p>
  </div>
  <div align="center">
    <img src="./assets/ss3.jpg" alt="DeepNow Dashboard Stats & Authentication" width="400" />
    <p><i>Key全景</i></p>
  </div>
</details>

---

<details>
  <summary><h2 style="display:inline-block; cursor:pointer;">🔥 核心特性 (Core Features) <i><small>- 点击展开详情</small></i></h2></summary>

### 🛡️ 1. 全域算力汇聚 (Compute Combined)
* **主备冗余 (Primary/Slave Failover):** 为生产环境量身打造的高可用防线。不仅支持常规单体大模型的无缝毫秒级切换，保障 API 调用 99.99% 的可用性，更支持将一整套“聚合模型”挂载为灾备节点。
* **同构聚合 (Token Aggregation):** 组合 ≥2 个相同的模型实例。在绝对保证原始推理智商与效果 100% 一致的前提下，完美分摊请求负载，彻底突破官方 API 的 TPM (Tokens Per Minute) 与并发 (Concurrency) 硬限制。
* **混合轮询 (Hybrid Round-Robin):** 支持将多个异构大模型组合为“混合调度池”。它能将你手中零散、免费的模型额度全部“榨干”，自动进行负载均衡，大幅缓解业务高峰期的单一 API 压力。
* **超混编排 (Super Hybrid):** 终极算力形态！支持将“异构单体模型”与“同构聚合模型池”进行二次组合，形成深度的“超混调度网”，满足极其复杂的企业级并发分流策略。

### 🧠 2. 动态探针 RAG 向量引擎 (Dynamic RAG Engine)
* **极致轻量底座:** 彻底摒弃臃肿的独立向量数据库（如 Milvus 等）。底层采用 CGO 直接绑定的 `sqlite-vec`，实现零依赖、高性能的单机亿级向量检索。
* **维度自适应探针 (Auto-Probe):** 无论你接入的是 768 维的经典模型，还是 3072 维的最前沿模型（如 `gemini-embedding-001`），系统在首次摄入知识时，会自动发射探针测定 Embedding 模型的向量维度，并动态重构底层张量表结构，真正做到“即插即用”。
* **时空溯源追踪:** 每一条被 RAG 引擎召回的上下文，不仅提供精准的知识切片，还精确携带录入时间戳与来源文件属性。这不仅彻底消灭了 AI 的幻觉，更提供了极简的知识库维护体验，让大模型在特定业务下直接拥有类似 LoRA 级微调的专有知识表现。

### ⏳ 3. 滑动窗口记忆增强 (Stateful Memory Injection)
突破标准 `/v1/chat/completions` 接口的无状态限制。DeepNow 在网关层内置了高性能关系型记忆存储引擎。外部客户端无需自行维护庞大的历史上下文，DeepNow 能基于 `API Key + Session ID` 自动实施滑动窗口拦截，将历史对话无缝拼接到当前请求中，并推送给承接算力的底层大模型。同时还支持最新的SSE或流式响应协议 `v1/Responses` 这使得Deepnow 不仅可用于计算推理，还可以应用于Agent 前端。 

### 🔐 4. 细粒度资源管控 (Access Control & Stats)
支持无限量生成以 `sk-deepnow-` 开头的专属消费令牌。并可对每一个令牌实施物理级管控：
* 总 Token 消耗绝对额度限制。
* 最高并发请求数 (Concurrency) 硬拦截。
* 提供细化到每日的流量走势图表与高精度 Token 使用明细，便于管理员进行跨团队成本分摊与模型提供商对账。

### 🤖 5. 自驱型 Agent 演进 (Self-Learning Agent - *TODO*)
未来，DeepNow 将在后台引入自驱型 Agent 机制。它能够在空闲算力期间，主动调用外部 Search API 获取全网实时信息，并结合大模型推理进行“自我总结与消化”。配合用户日常被动喂入的 RAG 知识碎块，系统将实现专有领域知识库的“主动生长与进化”，甚至可以主动完成事务性连续工作任务的响应，还可以通过custom协议来实现沙箱、搜索工具等的主要能力的整体集成化，无需recall本地工具实现，极大的减少交互次数和提高智能体效能。

</details>

<div align="center">
<img src="./assets/arch.png" alt="DeepNow Architecture" width="800" />
<p><i>Basic Architecture</i></p>
</div>

---

## 🚀 极速接入 (Quick Start)

DeepNow 采用全静态编译与资源内嵌技术，开箱即用，无需复杂的部署流程，除需配置大模型和Embedding模型的接口地址外，无需再配置或安装任何第三方组件；release 版无需docker或npm等相关环境安装、无任何三方依赖且跨平台，干净免维护一个二进制文件走天下（系统自带GUI dashboard/向量数据库/关系存储系统等）

```bash
# 1. 运行 DeepNow 服务端
Linux 下:
./deepnow

Windows 下:
直接执行deepnow.exe 

看到控制台输出文字并监听成功后可打开浏览器直达控制面板，deepnow 默认绑定设备的所有IP。

# GUI Dashboard 默认绑定 8084 端口，使用 http 访问

浏览器打开 http://127.0.0.1:8084/ 即可，端口可以在后台自己重新配置。

# 注意，默认情况下系统使用https提供endpoint 端点服务，第一次运行系统会下发一个为期10年的自签名证书。
# 但是遇到需要强制验证签名的客户端可能无法通过，所以最好配置真实证书并使用域名访问。
