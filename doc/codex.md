# 🚀 使用 Deepnow 转发 OpenAI Codex 请求指南

本教程将引导你如何配置 Codex，将其请求通过 **Deepnow** 转发到任意模型。

---

## 🛠️ 配置步骤

### 1. 寻找配置文件目录

首先，你需要定位到 Codex 的主目录，并在其中找到 `config.toml` 和 `auth.json` 两个配置文件。
> 💡 默认安装一般都在 Windows 的 "C:\Users\你名字\.codex\" 目录下。
---

### 2. 配置 `auth.json` (鉴权文件)

在 Codex 主目录下找到 `auth.json` 文件（如果没有该文件，请手动创建一个新文本文件并重命名为 `auth.json`）。

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

```ini

profile = "deepnow"

model_provider = "deepnow"

[profiles.deepnow]
model = "deepnow"
model_provider = "deepnow"
model_reasoning_effort = "low"

[profiles.deepnow.windows]
sandbox = "elevated"

[model_providers.deepnow]
name = "deepnow"
requires_openai_auth = true
base_url = "http://127.0.0.1:8444/v1"
wire_api = "responses"

[desktop]
conversationDetailMode = "STEPS_COMMANDS"
ambient-suggestions-enabled = false

```
**注意**：如果开启Deepnow HTTPS模式的Endpoint，必须使用真实证书，否则codex会拒绝握手。

### 4. 重新启动 Codex 即可，记得要关闭图标才算真正关闭codex