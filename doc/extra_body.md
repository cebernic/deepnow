# Deepnow `extra_body` 使用指南

## 1. 功能说明

为了更加灵活的配置模型，Deepnow 的“大模型配置“页面里支持 **额外参数（`extra_body`）** 配置，用于向上游大模型请求体中追加厂商自定义参数。

它适合配置 OpenAI 标准字段之外的能力，例如：

- 思考模式开关
- 推理强度
- 厂商私有采样参数
- 自定义布尔开关
- 嵌套 JSON 对象
- JSON 数组
- 数字或字符串枚举值

Deepnow 会将这里填写的参数合并到发送给上游模型的 JSON 请求体中。

---

## 2. 基本格式

每行填写一个参数，支持以下两种分隔方式：

```text
参数名 = 参数值
```

或：

```text
参数名: 参数值
```

例如：

```text
reasoning_effort = "high"
thinking = {"type":"disabled"}
```

也可以写成：

```text
reasoning_effort: "high"
thinking: {"type":"disabled"}
```

建议一个参数占一行。

---

## 3. 支持的数据类型

Deepnow 会自动识别参数值的 JSON 类型。

### 3.1 字符串

带引号字符串：

```text
reasoning_effort = "high"
```

最终请求体：

```json
{
  "reasoning_effort": "high"
}
```

普通未加引号字符串也支持：

```text
reasoning_effort = high
```

最终同样转换为：

```json
{
  "reasoning_effort": "high"
}
```

建议：

- 枚举值可以加引号，也可以不加；
- 为了与标准 JSON 表达一致，推荐加双引号。

---

### 3.2 布尔值

```text
thinking_mode = true
```

或：

```text
thinking_mode = false
```

最终请求体：

```json
{
  "thinking_mode": true
}
```

注意：

```text
thinking_mode = "true"
```

表示字符串 `"true"`，不是布尔值 `true`。

---

### 3.3 数字

整数：

```text
max_tokens = 4096
```

小数：

```text
temperature = 0.7
```

最终请求体：

```json
{
  "max_tokens": 4096,
  "temperature": 0.7
}
```

数字不需要加引号。

---

### 3.4 JSON 对象

```text
thinking = {"type":"disabled"}
```

最终请求体：

```json
{
  "thinking": {
    "type": "disabled"
  }
}
```

对象值必须使用合法 JSON：

- 属性名使用双引号；
- 字符串值使用双引号；
- 不允许使用 JavaScript 单引号对象；
- 不允许省略属性名引号。

正确：

```text
thinking = {"type":"enabled"}
```

错误：

```text
thinking = {type: enabled}
```

错误：

```text
thinking = {'type':'enabled'}
```

---

### 3.5 JSON 数组

```text
stop = ["END","STOP"]
```

最终请求体：

```json
{
  "stop": [
    "END",
    "STOP"
  ]
}
```

也可以填写数字数组：

```text
custom_ids = [1,2,3]
```

---

### 3.6 `null`

```text
custom_value = null
```

最终请求体：

```json
{
  "custom_value": null
}
```

---

## 4. DeepSeek 思考模式配置

DeepSeek 支持通过 `thinking` 参数控制思考模式。

### 4.1 关闭思考模式

```text
thinking = {"type":"disabled"}
```

最终请求体：

```json
{
  "thinking": {
    "type": "disabled"
  }
}
```

对于通过 Deepnow 的 Responses → Completions 转换使用 Codex 等 Agent 客户端时，关闭思考模式可以避免部分工具调用续写阶段出现：

```text
The `reasoning_content` in the thinking mode must be passed back to the API.
```

推荐的稳定配置：

```text
thinking = {"type":"disabled"}
```

---

### 4.2 开启思考模式

```text
thinking = {"type":"enabled"}
```

最终请求体：

```json
{
  "thinking": {
    "type": "enabled"
  }
}
```

注意：

如果模型在工具调用轮次返回了 `reasoning_content`，下一轮请求可能要求完整回传该字段。

在 Deepnow 尚未完整保存和恢复该字段前，Codex 等工具型客户端可能偶发出现：

```text
The `reasoning_content` in the thinking mode must be passed back to the API.
```

---

### 4.3 设置推理强度

```text
reasoning_effort = "high"
```

也支持：

```text
reasoning_effort = high
```

最终请求体：

```json
{
  "reasoning_effort": "high"
}
```

常见值可能包括：

```text
low
medium
high
max
xhigh
```

具体支持范围以对应模型和厂商文档为准。

---

### 4.4 同时配置思考模式与推理强度

开启思考：

```text
thinking = {"type":"enabled"}
reasoning_effort = "high"
```

关闭思考：

```text
thinking = {"type":"disabled"}
reasoning_effort = "high"
```

当 `thinking.type` 为 `disabled` 时，`reasoning_effort` 通常不会实际发挥作用，但上游是否忽略该字段取决于模型实现。

为了保持配置简洁，关闭思考模式时建议只填写：

```text
thinking = {"type":"disabled"}
```

---

## 5. 常用示例

### 5.1 字符串、布尔值和数字

```text
reasoning_effort = "high"
thinking_mode = true
temperature = 0.7
max_tokens = 4096
```

对应请求体：

```json
{
  "reasoning_effort": "high",
  "thinking_mode": true,
  "temperature": 0.7,
  "max_tokens": 4096
}
```

---

### 5.2 嵌套对象

```text
thinking = {"type":"enabled"}
stream_options = {"include_usage":true}
```

对应请求体：

```json
{
  "thinking": {
    "type": "enabled"
  },
  "stream_options": {
    "include_usage": true
  }
}
```

---

### 5.3 数组参数

```text
stop = ["END","STOP"]
```

对应请求体：

```json
{
  "stop": [
    "END",
    "STOP"
  ]
}
```

---

## 6. 带引号和不带引号的区别

以下两种写法都会得到字符串 `"high"`：

```text
reasoning_effort = high
```

```text
reasoning_effort = "high"
```

但布尔值和数字不同。

布尔值：

```text
thinking_mode = true
```

表示：

```json
{
  "thinking_mode": true
}
```

而：

```text
thinking_mode = "true"
```

表示：

```json
{
  "thinking_mode": "true"
}
```

数字：

```text
temperature = 0.7
```

表示数字。

而：

```text
temperature = "0.7"
```

表示字符串。

因此建议：

| 类型 | 推荐写法 |
|---|---|
| 字符串 | `"high"` |
| 布尔值 | `true` / `false` |
| 整数 | `4096` |
| 小数 | `0.7` |
| 对象 | `{"type":"enabled"}` |
| 数组 | `["END","STOP"]` |
| 空值 | `null` |

---

## 7. 常见错误

### 7.1 字符串重复包含引号

旧版本 Deepnow 中填写：

```text
reasoning_effort = "high"
```

可能被错误转换成内容包含引号的字符串：

```json
{
  "reasoning_effort": "\"high\""
}
```

上游可能报错：

```text
unknown variant `"high"`
```

新版 Deepnow 已支持正确解析带引号字符串。

---

### 7.2 JSON 对象被当成字符串

错误请求体：

```json
{
  "thinking": "{\"type\":\"disabled\"}"
}
```

上游会认为 `thinking` 是字符串，而不是对象，并可能报错：

```text
invalid type: string, expected struct ThinkingOptions
```

正确填写：

```text
thinking = {"type":"disabled"}
```

新版 Deepnow 会将其解析为真正的 JSON 对象。

---

### 7.3 使用单引号编写 JSON

错误：

```text
thinking = {'type':'disabled'}
```

JSON 标准要求属性名和字符串使用双引号。

正确：

```text
thinking = {"type":"disabled"}
```

---

### 7.4 布尔值错误加引号

错误：

```text
thinking_mode = "true"
```

这会发送字符串。

正确：

```text
thinking_mode = true
```

---

### 7.5 数字错误加引号

错误：

```text
temperature = "0.7"
```

这会发送字符串。

正确：

```text
temperature = 0.7
```

---

## 8. 参数冲突说明

`extra_body` 中的参数会合并到最终请求体。

如果填写的参数与 Deepnow 已生成的标准参数同名，可能发生覆盖或冲突。例如：

```text
model = "another-model"
stream = false
messages = []
```

不建议在 `extra_body` 中填写以下核心字段：

```text
model
messages
input
tools
stream
previous_response_id
```

这些字段应由 Deepnow 的协议转换和路由逻辑管理。

`extra_body` 更适合填写厂商扩展字段，例如：

```text
thinking
reasoning_effort
thinking_mode
stream_options
```

---

## 9. 厂商兼容性

并不是所有模型都支持相同的自定义字段。

例如：

```text
thinking = {"type":"disabled"}
```

可能只对特定 DeepSeek 模型有效。

其他模型可能：

- 忽略未知字段；
- 返回 HTTP 400；
- 使用不同字段名；
- 要求不同的数据类型。

因此配置前应确认：

1. 当前 Provider 是否支持该字段；
2. 当前模型是否支持该能力；
3. 字段值类型是否正确；
4. 字段是否适用于 Chat Completions 或 Responses；
5. 是否仅适用于流式或非流式请求。

---

## 10. 推荐配置

### DeepSeek + Codex，优先稳定工具调用

```text
thinking = {"type":"disabled"}
```

### DeepSeek 开启思考模式

```text
thinking = {"type":"enabled"}
reasoning_effort = "high"
```

### 普通布尔开关示例

```text
thinking_mode = true
```

### 使用统计信息

```text
stream_options = {"include_usage":true}
```

是否生效取决于上游接口支持情况。

---

## 11. 调试方法

开启 Deepnow Debug 日志后，应检查最终请求体中的参数类型。

正确的对象：

```json
"thinking": {
  "type": "disabled"
}
```

错误的字符串对象：

```json
"thinking": "{\"type\":\"disabled\"}"
```

正确的字符串：

```json
"reasoning_effort": "high"
```

错误的包含引号字符串：

```json
"reasoning_effort": "\"high\""
```

正确的布尔值：

```json
"thinking_mode": true
```

错误的布尔字符串：

```json
"thinking_mode": "true"
```

---

## 12. 完整示例

Deepnow 输入：

```text
thinking = {"type":"disabled"}
reasoning_effort = "high"
thinking_mode = true
temperature = 0.7
max_tokens = 4096
stop = ["END","STOP"]
stream_options = {"include_usage":true}
```

最终 JSON 类型应等价于：

```json
{
  "thinking": {
    "type": "disabled"
  },
  "reasoning_effort": "high",
  "thinking_mode": true,
  "temperature": 0.7,
  "max_tokens": 4096,
  "stop": [
    "END",
    "STOP"
  ],
  "stream_options": {
    "include_usage": true
  }
}
```

---

## 13. 使用建议

- 一个参数填写一行；
- 推荐使用 `参数名 = 参数值` 格式；
- 字符串推荐使用双引号；
- 布尔值和数字不要加引号；
- 嵌套对象和数组必须使用合法 JSON；
- 不要通过 `extra_body` 覆盖核心协议字段；
- 不同模型使用不同的 `extra_body` 配置；
- 遇到 HTTP 400 时，优先检查参数类型和上游文档；
- DeepSeek 与 Codex 工具链追求稳定时，建议关闭 thinking；
- 开启 thinking 后，应留意 `reasoning_content` 回传兼容问题。
