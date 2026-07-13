# ⚡ DeepNow : The Ultimate AI Compute Gateway ( Router )

<div align="center">

<img src="./assets/logo.png" alt="DeepNow Logo" width="200" />

**All-in-One Aggregated AI Gateway & Intelligent RAG Fusion Engine**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)
[![OpenAI Compatible](https://img.shields.io/badge/OpenAI-100%25_Compatible-412991?style=for-the-badge&logo=openai)](#)
[![SQLite-Vec](https://img.shields.io/badge/Vector_DB-SQLite--Vec-success?style=for-the-badge)](#)
[![Zig CC](https://img.shields.io/badge/CGO_Cross--Compiled-Zig_CC-F7A41D?style=for-the-badge&logo=zig)](https://ziglang.org/)


* Core Philosophy: Achieving model-agnostic continuation for batch tasks, seamless multimodal drifting, context migration, provider switching, tool task distribution, multi-key load balancing, and multi-model management. It is inherently engineered to serve diverse concurrent task agents while negotiating maximum compatibility across all models.

</div>

<div align="right">
  <a href="README-eng.md">English</a> | <a href="README.md">简体中文</a>
</div>

<hr>

**Deepnow Protocol List:**

| Client Protocol ↓ \ Endpoint Protocol → | R Protocol Model | C Protocol Model | M Protocol Model |
| :--- | :--- | :--- | :--- |
| **R Protocol Client** | **Transparent Pass-Through**<br>(R → R) | **Wrapper Translation**<br>(R → C) | **Wrapper Translation**<br>(R → M) |
| **C Protocol Client** | **Wrapper Translation**<br>(C → R) | **Transparent Pass-Through**<br>(C → C) | **Wrapper Translation**<br>(C → M) |
| **M Protocol Client** | **Wrapper Translation**<br>(M → R) | **Wrapper Translation**<br>(M → C) | **Transparent Pass-Through**<br>(M → M) |

<hr>




## ⚡ What is DeepNow

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; In the contemporary __Token__-centric paradigm of artificial intelligence, **Deepnow** serves as an enterprise-grade and individual-centric infrastructure engineered for high availability, high concurrency, and comprehensive scenario coverage. It functions as an intelligent AI model gateway (routing system) and a knowledge integration foundation. By unifying the management and scheduling of disparate Large Language Models (LLMs/VLMs) and Vector Embedding models, Deepnow aggregates these distinct entities via specialized binding algorithms to achieve compute consolidation and robust fault tolerance. By orchestrating resources across heterogeneous model providers, Deepnow maximizes token throughput and efficiency for downstream client applications and development environments. Through its aggregation mechanism, you can easily bypass the variable restrictions—such as TPM and concurrency ceilings—imposed by individual model operators. Whether deployed for multi-user collaborative development ("vibe coding"), long-context integration, or high-density impulse invocations, Deepnow definitively resolves your operational bottlenecks. Conceptually, Deepnow introduces the paradigm of a software-defined **Soft Router** for the AI era. However, rather than managing standard TCP/IP network packets, this router arbitrates Token traffic, establishing a novel product classification: a **Token Router / Token Gateway (TR/TG)**.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; You may have previously utilized multi-model bridging solutions such as CC Switch or OpenRouter to seamlessly mount various models across autonomous agents like Codex, Hermes, Claude, or Openclaw without protocol friction. However, while the former is fundamentally a front-end application designed for model switching, the latter operates essentially as a commercial Token reselling service, multiplexing multiple models via a single endpoint. It is highly improbable that you would desire to permanently anchor your proprietary AI applications to a third-party Token broker. Deepnow resolves this by offering a novel, open-source backend architecture that synthesizes both capabilities, providing both robust protocol translation and stateful multi-model persistence.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Deepnow is architected as an autonomous, dependency-free, cross-platform microservice capable of continuous 24/7 deployment in both local and remote environments, engineered using Go and C—languages natively optimized for highly concurrent network services. The gateway integrates sophisticated architectural paradigms, including cross-modal load sharing, disaster recovery, dynamic multi-path Retrieval-Augmented Generation (RAG) loading, and high-concurrency multi-agent task load balancing. Furthermore, it incorporates dedicated upstream adapters to unlock the proprietary, high-tier capabilities of specific advanced models rather than relying on lowest-common-denominator compatibility.


&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; From an administrative perspective, Deepnow addresses the fiscal and security risks associated with widespread API key distribution within an enterprise or team. Provisioning individual upstream credentials across an organization escalates operational expenditures and heightens the risk of credential leakage. Similarly, in collaborative individual environments where users seek to pool resources for high-tier compute and share expenses proportional to Token consumption, multi-IP concurrency frequently triggers upstream fraud-detection algorithms, resulting in account suspension. Deepnow mitigates these risks comprehensively by empowering administrators to generate and distribute custom downstream API keys. These keys can be programmatically bound to strict concurrency quotas, specific model access controls, and rate-limiting thresholds, with the capacity for instantaneous revocation. Furthermore, because all external mutations originate deterministically from the static deployment IP of the Deepnow instance, downstream clients are entirely insulated from upstream account suspension risks.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Beyond routing matrix features, Deepnow integrates a native knowledge-base layer directly into the telemetry request paths. This architecture allows individual or enterprise-proprietary corpora to be seamlessly ingested into the gateway via an intuitive operational workflow. Consequently, underlying base models are augmented into domain-specific reasoning agents, formalizing heuristic inference into structured knowledge graphs. This system is optimally suited for deterministic programming, corporate training, automated customer support, and medical informatics—domains with strict epistemic-fidelity requirements where hallucinations must be strictly prohibited. This architecture effectively provides zero-configuration persistent memory, enabling all default requests to maintain cross-application, cross-session long-term memory while completely circumventing context-window constraints without requiring any front-end or back-end source code modifications.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; The current AI landscape is characterized by a fragmentation of customized front-end and back-end protocols. Providers frequently offer "OpenAI-compatible" API variants that include proprietary modifications designed to expose advanced features. These non-standard variations present significant barriers to universal interoperability, which underscores the existential value of Deepnow. Acting as an intermediary "protocol translator," Deepnow decouples applications and upstream models from mutual protocol dependencies. It negotiates session contexts by prioritizing transparent data pass-through; when direct pass-through is structurally unviable, the engine executes a near-lossless semantic translation. For example, a client restricted to the newer Responses schema can successfully interface with a legacy-only Completions model, and vice versa. 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; In the near future, Deepnow's developmental trajectory encompasses integration with the Model Context Protocol (MCP), external plugin ecosystems, server-side tasks, cross-client context persistence, server-side skills, and advanced multi-modal matrices. (Note: A multi-modal matrix allows you to synthesize specialized, disparate models into a unified composite architecture to handle auditory, visual, and text modalities concurrently. This acknowledges the empirical reality that no single monolithic model achieves state-of-the-art performance across all modalities simultaneously, necessitating the dynamic allocation of AI compute paths).

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; In synthesis, Deepnow constitutes a fundamental layer of AI infrastructure—a high-performance compute foundation and an endpoint gateway routing system engineered to support advanced cognitive capabilities.


## ⚡ Basic Architecture

  <div align="center">
    <img src="./assets/deepnow-arch-eng.png" alt="DeepNow Architecture" width="800" />
    <p><i>Basic Architecture</i></p>
  </div>

## ⚡ What Pain Points Does DeepNow Address?

<details>
  <summary><b>🎯 Arbitrary Model Delivery</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Probably the only AI Routing solution that enables self cached cross-model context delivery.<br><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;That is to say, after you have completed part of your work, as long as your agent or chat client maintains the context, you can have the first part of the same session completed by one model, and when an unexpected situation requires switching to another model to continue, you do not need to start a new conversation at all. deepnow will reload your context and redeliver it in a format acceptable to the target model, allowing you to resume your unfinished work. For example: Deepseek writes a piece of code but runs out of tokens, and you happen to have Gemini on hand – you can immediately mount Gemini to continue your work. Under normal circumstances, the two models have many differences in their payloads; without deepnow, simply changing the endpoint in your agent or client would not allow you to continue.
</details>
<details>
  <summary><b>🚀 Protocol Wrapper & Transpilation</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The model landscape is mutating rapidly, with fragmented protocol formats: some are native single-model multimodal, others expose unified multi-model endpoints, while some still lean on legacy completions or strict SSE streams. If a frontend tool attempts to switch to an arbitrary upstream model, payload structural mismatches frequently yield HTTP 400 Bad Request errors. What if an advanced IDE plugin expects the latest stateful Responses protocol, but your target model only interprets legacy Completions parameters? Conventional AI gateways simply act as transparent pass-through proxies, failing to provide downstream schema consistency. This forces you to either abandon your favorite frontend client or change your model.<br><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;DeepNow's protocol wrapper smoothly irons out these frontend-backend discrepancies. We support dual-mode routing: 1) Protocol Hijacking & Transpilation, and 2) Native Pass-through. Impressively, whether utilizing transpilation or raw proxying, the system fully maintains its cascading load-balancing matrix and local RAG fusion pipelines. Frontend apps remain decoupled from upstream variations, requiring compatibility with only a single generic endpoint standard (Completions/Responses/Messages).
</details>
<details>
  <summary><b>🧠 Token-Aware Optimal Routing ("Optimal Choice")</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Beyond model pooling, multi-channel hybrid routing, and failover mechanics, DeepNow embeds a proprietary "Optimal Choice" scheduling algorithm, delivering optimal solutions for two critical runtime patterns:<br><br>
  1) Concurrency Racing (Latency-First): The gateway broadcasts an incoming inference request simultaneously to $\ge 2$ configured fast, lightweight channels (including local models). The first channel to return an indexable First Token (TTFT) instantly locks onto the client stream, while the redundant worker threads are gracefully aborted, squeezing down latency under heavy load.<br><br>
  2) Layered Cascade Routing (Re-rank): Queries are automatically evaluated against model reasoning tiers. A lightweight model first intercepts the payload to classify intent and extract semantic features, subsequently escalating the query to an optimally rated reasoning node. This ensures granular token-billing optimization and highly accurate inference outcomes.
</details>
<details>
  <summary><b>🚀 High Availability & Compute Pooling: Breaking Concurrency Ceilings</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;You might be struggling with fragmented token pools that fail to deliver high-concurrency, low-latency responsiveness. Standard non-enterprise API keys come bound to strict rate limiting frameworks (`RPM` for requests per minute, `TPM` for tokens per minute, `RPD` for daily limits). No matter how much you optimize application-side code, upstream rate limits cannot be bypassed. <b>DeepNow shatters this bottleneck.</b> Through our innovative compute pooling architecture, you can bundle multiple cost-effective developer or individual keys into a unified virtual cluster whose combined throughput mirrors or surpasses premium enterprise tiers. Furthermore, you can aggregate open-source models running across separate physical hardware networks into a single, highly available endpoint array. In mission-critical environments, frontend apps will never display embarrassing errors like "...high demand, please try again later...". DeepNow captures upstream anomalies in real-time, executing an atomic failover retry to a backup model in milliseconds, keeping the entire recovery lifecycle completely transparent to the client.
</details>
<details>
  <summary><b>🛡️ Secure Multi-Tenant Compute Sharing</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Whether serving hardcore AI geeks running intense Hermes/OpenClaw setups or orchestrating multi-user corporate spaces, DeepNow facilitates safe resource sharing across multiple clients. DeepNow ships with a robust self-managed API Key provisioning framework. A single master key can be bound across all your tools, or individual sub-keys can be distributed to separate users. Historically, multi-regional concurrent calls utilizing identical upstream keys frequently triggered account bans or leaked credential hashes. By routing traffic through DeepNow, true upstream keys are securely abstracted behind the gateway layer; to the provider, all incoming streams originate from DeepNow, neutralizing telemetry risks. DeepNow records complete token ingestion and emission logs inside in-memory caches and persists them to local/external database layers (SQLite/MySQL), enabling downstream tracking or custom secondary billing UIs.
</details>
<details>
  <summary><b>📦 Agile Deployment & High Performance (Zero Dependencies, Out-of-the-Box)</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The majority of open-source AI utilities depend heavily on Python, which introduces convoluted dependency dependency chains, version mismatches, and structural runtime performance bottlenecks. DeepNow is engineered from day one to match enterprise-grade high-throughput demands, written in compiled Go fused with highly optimized C components. The system features a fully self-contained architecture—even the static UI administration dashboard is compiled directly into the single binary asset. No messy component tracking, no external database dependencies, and no heavy Docker footprints. Download and execute; that’s it. Simply grab the target architecture binary and achieve instant deployment. All runtime parameters are fully manageable via an intuitive HTML5 control center, eliminating manual configuration editing. Get the service up in milliseconds, then fine-tune parameters—this is the peak agile experience designed by DeepNow.
</details>
<details>
  <summary><b>🧠 Transparent RAG Injection (Hot-Swappable Semantic Brain)</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Both individual power users and large organizations face the challenge of converting private knowledge vectors into robust AI memory. <b>DeepNow seamlessly orchestrates any arbitrary LLM to compute over this knowledge</b>. Frontend applications are completely liberated from compiling complex RAG chunking and retrieval flows; DeepNow handles vector space lookups at the gateway layer, dynamically weaving relevant context directly into the prompt payload before it hits the model. The gateway supports strict fallback constraints: if a query fails to trip semantic thresholds, it can be short-circuited or downgraded to the LLM's baseline weights. More cleverly, our multi-channel routing pipeline applies to the retrieval phase as well; no single upstream provider ever receives your full raw context stack, physically neutralizing data chain leakage risks. In hyper-specialized domains requiring zero hallucination (Temperature = 0), DeepNow applies target prompt injection techniques to align structural consistency across divergent models. The semantic framework supports cross-modal retrieval (text, diagrams, audio charts), allowing cross-model processing to combine outputs while returning a single, unified inference response. Regardless of whether textual records or structured knowledge graphs are supplied, DeepNow seamlessly indexes and orchestrates them for accurate downstream reasoning task enforcement.<br><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The foundational core of **DeepNow** is the absolute decoupling of private data logic from raw inference engines. The LLM scales down to a raw, stateless compute unit—an intellectual commodity—while the underlying knowledge base serves as the actual brain of your AI system. DeepNow provides an integrated scheduling ring where you can hot-swap, plug in, or disconnect cloud models, local nodes, or knowledge segments on the fly, tailoring your deployment state instantly. Technologically, this is achieved through multi-dimensional semantic search, cosine distance filtering, atomic re-indexing, Re-rank algorithms, and semantic query expansion. It mirrors the exact behavioral profile of a custom pre-trained private model without the prohibitive overhead of compute training.<br><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Furthermore, DeepNow possesses continuous self-directed learning capabilities, transforming authorized transaction logs into persistent vector memories with zero context window boundaries.
</details>
<details>
  <summary><b>🌐 Hierarchical Distributed Compute Network (Infinite Cascading)</b></summary>
  <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**DeepNow** nodes can be recursively cascaded! Beyond raw compute sharing, session memory states and private localized RAG indices can be seamlessly invoked by sibling or parent DeepNow instances, as the semantic context is fully serialized within the streaming server-sent events (SSE). You can deploy multiple specialized DeepNow instances across separate operational zones (e.g., Node A managing corporate HR knowledge graphs, Node B tracking product engineering specifications, Node C processing executive operational data) and aggregate them under a master root node. This lets you synthesize a multi-domain super-intelligent private AI hub, where frontend clients connect to restricted sub-nodes or the overarching root gateway based on explicit authorization profiles.
</details>

##  ⚡ Competitive Landscape & Matrix Comparison


| Dimension / Feature | **DeepNow (Edge-Local Gateway)** | **Cherry Studio (CC Switch)** | **OpenRouter (Cloud SaaS)** | **OneAPI / NewAPI (Proxy System)** |
| :--- | :--- | :--- | :--- | :--- |
| **Local Gateway Capability** | ✅**Core or Edge Gateway**<br>· Ultra-agile deployment, lightweight runtime<br>· Embedded high-concurrency middleware, runs as a local background daemon | ❌**Desktop Client Only**<br>· View-level interface switching<br>· Restricted to single-user local application context | ❌**Cloud Commercial SaaS**<br>· Centralized remote hosting<br>· Mandatory external internet connection | ❌**Cloud-Native Middleware**<br>· Dependent on heavy Docker containers<br>· Strict runtime reliance on multiple external components, complex deployment |
| **High-Concurrency Scheduling & Cache** | ✅**Microsecond Memory State Machine**<br>· Atomic in-memory cache and dictionary tries for $O(1)$ lookup performance<br>· Concurrency lock mechanisms prevent I/O starvation under burst traffic peaks | ❌**No Concurrency Handling**<br>· Limited to static client-side proxying and single-point routing switching | ✅**Cloud Elastic Scaling**<br>· High-throughput remote handling capacity<br>· Enterprise metering and tenant traffic controls | ❌**Database-Driven Polling**<br>· High-frequency relational DB reads per request to deduct quota and check channels<br>· DB connection limits act as severe performance bottlenecks under high QPS |
| **Credential Security** | ✅**Device-Fingerprint AES Encryption**<br>· Derived cryptographic MasterID locked to local physical hardware<br>· Local config files heavily obfuscated to prevent credential theft; admin-exportable | ❌**Plaintext Storage**<br>· API credentials stored raw inside local client storage directories without device-level protection | ❌**Remote Managed Vault**<br>· API keys are fully entrusted to and held on OpenRouter's cloud databases | ❌**Standard Relational Encryption**<br>· Relies on baseline MySQL/PgSQL database protections and typical hashing metrics |
| **Vector Engine** | ✅**Embedded Native Semantic Base**<br>· Low-overhead C-level `sqlite-vec` engine compiled into core binary<br>· Supports concurrent high-frequency text chunking and immediate vectorization | ❌**External Hook / Missing Core**<br>· Depends on external app integrations or third-party cloud-hosted vector endpoints | ❌**No Semantic Layer**<br>· Strictly restricted to raw LLM protocol handling and token stream distribution | ❌**No In-Memory Vector Pool**<br>· Acts strictly as a forwarding pipe; some forks patch this via clunky third-party plugins |
| **Knowledge Retrieval** | ✅**RRF Dual-Track Reciprocal Rank Fusion + Hard Kill**<br>· Normalized hybrid evaluation blending **Vector + Tag + Semantic Re-rank**<br>· Anti-representation drift safeguards built-in<br>· **2-Stage Slider Short-Circuiting**: Queries below threshold completely cut off external model routing and return custom markers, stopping hallucinations at zero cost | ❌**No Retrieval Engine**<br>· RAG processes must be handled manually by client-side application prompt assemblers | ❌**No Filtering Layer**<br>· Incapable of parsing tenant business data; all blind-spot queries are pushed upstream, racking up billable tokens | ❌**No Semantic Guardrails**<br>· Purely transparent network pipeline; unable to intercept or short-circuit queries based on knowledge hit-rates |
| **Protocol Compatibility** | ✅**Full Responses/Completions Wrapper**<br>· Advanced parsing of stateful Responses and client Runs topology state machines<br>· Auto-reconstructs multi-turn Agent trace histories and multimodal payloads; bidirectional transpilation | ❌**Standard Single-Turn Protocol**<br>· Primitive structural conversion profiles | ✅**Generic Cloud Translation**<br>· Focuses on cloud vendor Chat protocol flattening; lacks deep stateful Responses tracking | ❌**Legacy Protocol Focused**<br>· Designed for basic `v1/chat/completions` flattening; completely lacks stateful agent thread orchestration |
| **Edge Microservices & Expansion** | ✅**Built-In Secured MCP Server**<br>· Evaluates and executes model tools locally<br>· Directly exposes and tracks local compute microservices to clients like Codex Desktop | ❌**MCP Client Only**<br>· Imports external MCP endpoints; incapable of operating as a secured local microservice provider | ❌**Cloud Endpoint Only**<br>· Geographically isolated from interacting with or managing resources in a local network cluster | ❌**No Edge MCP Protocol Layer**<br>· Lacks mechanisms to intercept, evaluate, and restructure local hardware toolchains |
| **High Availability & Load Balancing** | ✅**Primary/Slave Dual Track + Auto-Self-Healing + Compute Pooling + Hyper-Hybrid Configuration**<br>· Instant gateway-layer orchestration with native TLS certificate attachment capabilities | ❌**Basic Failover Only**<br>· Elementary error-switching optimized solely for local desktop clients | ✅**Cloud Failover, Rigid Balancing**<br>· Remote multi-channel fallback switching; however, long network paths can impact TTFT stability | ✅**Weighted Relational Load Balancing**<br>· Weighted round-robin routing driven by database status tables; lacks end-of-life transactional reconstruction |
| **Self-Provisioned Key System** | ✅**Comprehensive Downstream Key Management**<br>· Limitless key provisioning with fine-grained tracking metrics for models, time windows, and token ceilings | ❌**Single-User Desktop Profile**<br>· Missing tenant provisioning system | ❌**Commercial Consumer SaaS**<br>· Sells aggregated upstream compute; cannot be deployed to build your own private key provisioning infrastructure | ✅**Commercial Quota Key Generation**<br>· Rich multi-tier key structure with quota configurations |
| **Efficiency & Usability** | ✅**Zero Dependency Runtime**<br>· No npm, no node.js required; packages high-speed WebUI out-of-the-box for edge deployment<br>· Native Go runtime thread concurrency | ❌**Heavily Dependent Runtime**<br>· Electron-based node.js desktop interface talking to headless processes via IPC; bounded to client machines | ❌**Zero Deployment Overhead**<br>· Instant API consumption but subject to external premium billing structures | ❌**Multi-Layer Component Matrix**<br>· Resource-intensive deployment; high-concurrency capability hampered by frequent database read validation loops |
| **Protocol Wrapper** | ✅**Bi-Directional Wrapper Architecture**<br>· Dynamically transforms Completions queries to Responses and vice versa, rendering all models universal | ❌**Transparent Proxying**<br>· Standard protocol pass-through mapping | ❌**Transparent Proxying**<br>· Flat standard cloud payload mapping | ❌**Transparent Proxying**<br>· Linear string/field rewriting |



## ⚡ Pictures
<details>
  <summary>Click to Expand</summary>
  <div align="center">
    <img src="./assets/ss1.jpg" alt="DeepNow Dashboard Overview" width="400" />
    <p><i>Network Topologies</i></p>
  </div>
  <div align="center">
    <img src="./assets/ss2.jpg" alt="DeepNow Dashboard RAG Settings" width="400" />
    <p><i>Compute Configuration</i></p>
  </div>
  <div align="center">
    <img src="./assets/ss4.jpg" alt="DeepNow Dashboard Stateful Memory" width="400" />
    <p><i>Session Memory Policies</i></p>
  </div>
  <div align="center">
    <img src="./assets/ss5.jpg" alt="DeepNow Dashboard Stats & Authentication" width="400" />
    <p><i>Downstream API Token Overview</i></p>
  </div>
  <div align="center">
    <img src="./assets/ss3.jpg" alt="DeepNow Model Selection" width="400" />
    <p><i>Compute Core Allocation</i></p>
  </div>
</details>

## ⚡ Capability Building

  <div align="center">
    <img src="./assets/arch.png" alt="Capability Building" width="800" />
    <p><i>Capability Building</i></p>
  </div>

---

## ⚡ Quick Start (Out-of-the-Box)

DeepNow leverages full static compilation and asset inlining technologies to provide an out-of-the-box runtime environment with zero deployment friction. Outside of configuring your base model and embedding API addresses, no external databases or runtime dependencies are required. The production release asset drops all dependencies on Docker or Node.js/npm environments—a clean, self-healing, cross-platform single binary managing your complete suite (built-in GUI control console, embedded vector database, relational system tracker).

```bash
# 1. Spin up the DeepNow daemon
With Linux:
./deepnow       (console)
or
nohup ./deepnow >& /dev/null &       (background daemon)

Windows:
double click deepnow.exe (It is not recommended to run this on your local machine.)

# When the initialization logs and listening address appear in the terminal, the service is fully operational. DeepNow binds natively across all available interface IPs.

# Once initialized, the administrative configuration UI becomes accessible.

# The GUI Dashboard binds defaults to port 8084 via HTTP (restricted strictly to 127.0.0.1 for baseline security).

Simply navigate your browser to [http://targetip:8084] ,or running under local just visit http://127.0.0.1:8084/ to start. Ports are fully customizable via the control center.

default user and password : admin / 12345 . You can modify password and UI language anytime. 

# Note: By default, the gateway exposes AI endpoint routes over HTTP.
# You can easily toggle TLS (HTTPS) mode inside the 8084 GUI dashboard. Upon the first toggle, the system automatically cuts a 10-year self-signed certificate asset. However, strict downstream enterprise clients (e.g., Codex Desktop) might require validated certificate authorities.
# For standalone sandbox development or local engineering clusters, keeping the endpoint in HTTP mode guarantees maximum cross-compatibility across all client tools.

Running it for the first time will generate a vector database and multiple configuration files in the DeepNow installation directory.


<div align="left">
  <a href="/doc/codex.md">Codex Config HOWTO</a> | <a href="/doc/hermes.md">Hermes Config HOWTO</a>
</div>

