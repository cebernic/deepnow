# 🚀 使用 Deepnow 转发 OpenAI Codex 请求指南

本教程将引导你如何配置 Codex，将其请求通过 **Deepnow** 转发到任意模型。

---

## 🛠️ 配置步骤

### 1. 寻找配置文件目录

首先，你需要定位到 Codex 的主目录，并在其中找到 `config.toml` 和 `auth.json` 两个配置文件。
> 💡 默认安装codex workdir 一般都在 Windows 的 ["C:\\Users\\你名字\\.codex\\"] 目录下。
---

### 2. 配置 `auth.json` (鉴权文件)

在 Codex 主目录下找到 `auth.json` 文件（如果没有该文件，请手动创建一个新文本文件并重命名为 `auth.json`）。

(最新版直接跳到第三步 step 3 即可.for Latest codex:To step 3 directly,copy config.toml to codex's workdir is the step only one.)
1. **清空** 该文件内的所有既有内容。
2. 复制并**粘贴**以下代码：

```json

{
  "OPENAI_API_KEY": "sk-deepnow-xxxx"
}

```
**注意**：请将代码中的 "sk-deepnow-xxxx" 替换为你在 Deepnow 平台上实际生成的 API Key。

### 3. 配置 `config.toml` (核心配置文件)

接下来，在同一目录下找到 config.toml 文件，然后清空该文件内的所有内容，复制并粘贴以下配置代码：

**2026-05-31 Updated:(最新版codex可用,记住把 sk-deepnow-xxx 换成你自己通过deepnow 后台生成的任意一个 key )**

```ini

model = "deepnow"
model_provider = "deepnow"
model_reasoning_effort = "high"

[model_providers.deepnow]
name = "deepnow"
base_url = "http://127.0.0.1:8444/v1"
http_headers = { "Authorization" = "sk-deepnow-xxx" }

[windows]
sandbox = "elevated"

[mcp_servers.deepnow]
enabled = true
url = "http://127.0.0.1:8444/mcp"
http_headers = { "Authorization" = "sk-deepnow-xxx" }

[features]
js_repl = false


```
**注意**：如果开启Deepnow HTTPS模式的Endpoint，必须使用真实证书，否则codex会拒绝握手。

### 4. 重新启动 Codex 即可，记得要关闭图标才算真正关闭codex