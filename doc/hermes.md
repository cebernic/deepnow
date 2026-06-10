# 🚀 使用 Deepnow 转发 Hermes 请求指南

本教程将引导你如何配置 Hermes CLI/Desktop，将其请求通过 **Deepnow** 转发到任意模型。

---

## 🛠️ 配置步骤

### 1. 寻找配置文件目录

首先，你需要定位到 Codex 的主目录，并在其中找到 `config.yaml` 这个配置文件。
> 💡 默认安装hermes workdir 一般都在 Windows 的 ["C:\\Users\\yourname\\AppData\\Local\\hermes\"] 目录下 ，可以输入 ["%LOCALAPPDATA%\hermes"] 直达。其它的系统应该可以输入： ['~/.hermes/config.yaml'] 到达。
---

### 2. 配置 `config.yaml` 

1. 安装好hermes后在此文件顶部粘贴如下代码，auxiliary 部分内容如果原来的配置中已有，请用auxiliary的部分覆盖原来的。

```json

model:
  default: deepnow-auto-model
  provider: deepnow
  base_url: http://127.0.0.1:8444/v1
  api_mode: codex_responses
  api_key: sk-deepnow-xxx
providers:
  deepnow:
    base_url: http://localhost:8444/v1
    

auxiliary:
  vision:
    provider: auto
    model: 'deepnow-auto-model'
    base_url: 'http://127.0.0.1:8444/v1'
    api_key: 'sk-deepnow-xxx
    timeout: 120
    extra_body: {}
    download_timeout: 30

```

**注意**：请将代码中的 "sk-deepnow-xxxx" 替换为你在 Deepnow 平台上实际生成的 API Key。由于hermes 针对本地搭建的Endpoint写入配置文件有BUG，请直接使用以上方法配置即可。同时，如果你在deepnow里配置的模型是多模态模型，或者次模型是多模态，auxiliary里配置的视觉能力也是可以生效的。


### 3. 重启hermes 即可。另外， hermes desktop 版也可用。