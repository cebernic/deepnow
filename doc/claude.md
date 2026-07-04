## 🛠️ Claude 如何通过开发者模式配置自定义端点

### 1. 启用开发者模式
* **打开应用**：启动 Claude Desktop 客户端。
* **开启模式**：导航至顶部菜单，选择 **Help（帮助） > Troubleshooting（故障排除） > Enable Developer Mode（启用开发者模式）**。
* **重启**：重新启动该应用程序。

---

### 2. 访问配置界面
* **查找入口**：重启后，应用程序的菜单栏中会出现一个全新的 **Developer（开发者）** 标签页。
* **进入配置**：依次点击 **Developer（开发者） > Configure Third-Party Inference（配置第三方推理）**。

---

### 3. 设置您的网关/服务商
在配置界面中，按以下要求填写参数：
* **Connection（连接）**：将 Inference Provider（推理服务商）设置为 **Gateway（网关）**。
* **Gateway Base URL（网关基础 URL）**：（这里填写你的 Deepnow endpoint 地址,注意不是管理后台的端口，一般是 http://127.0.0.1:8444/ ，这里务必不要加v1 ）。
* **Credential Kind（凭据类型）**：选择 **Static API Key（静态 API 密钥）** 并输入您在 Deepnow 后台生成的任意 1 个 API Key ，类似: sk-deepnow-18acaabc6xxxxxxx 这样的ID。
* **Gateway auth scheme（认证方案）**：将其设置为 Bearer 或 x-api-key 皆可。

---

### 4. 配置模型
* **MODELS（模型）**：拉到下方 **Model List** 位置，点击旁边 **+ Add Model** ，看到 **Model ID** 输入框这里输入 **claude-3-5-sonnet** 即可。

---

### 5. 保存并生效
* **应用配置**：点击设置界面底部的 **Apply Changes** 按钮，等待几秒以后就可以开始使用 Claude Desktop 了。

---

### 6. 后续你可以在 Deepnow 里挂载任何模型给 Claude 使用。