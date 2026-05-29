import requests
import urllib3
import json
import sys

# certifiacate warning
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

# ==========================================
#                  Config
# ==========================================
# 基础网关地址 gateway addr
BASE_URL = "http://127.0.0.1:8444"
API_KEY = "sk-deepnow-18acaabc63392260ecdeb6e2cb05c8a2"

DEBUG_MODE = False

# 核心模式控制：
# True  -> Demo for  SSE 流式握手 (stream=True)，实现打字机跑字效果
# False -> SSE disabled as a standard way to make a chat (v1/chat/completions)
DEBUG_STREAM_MODE = True
# ==========================================

def main():
    # 动态拼接端点
    API_URL = f"{BASE_URL}/v1/chat/completions"
    print("==================================================")
    print("🚀 Deepnow Hybrid Stateless & SSE Demo Client")
    print(f"🔗 Target Endpoint: {API_URL}")
    print(f"🔧 Transfer Mode: {'🌊 SSE Stream Mode' if DEBUG_STREAM_MODE else '📦 Standard Non-Stream Mode (Block)'}")
    print(f"🔧 Raw Capture: {'🟢 Enabled' if (DEBUG_MODE and DEBUG_STREAM_MODE) else '🔴 Disabled'}")
    print("💡 Type your message to start, type '!!exit' to quit")
    print("==================================================\n")    


    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {API_KEY}"
    }

    messages = []

    while True:
        try:
            user_input = input("\n🧑 Question: ")
            
            if not user_input.strip():
                continue

            if user_input.strip() == "!!exit":
                print("👋 quit demo as normally.")
                break

            messages.append({"role": "user", "content": user_input})

            # 根据全局配置动态设置 stream 载荷
            payload = {
                "model": "ignored_by_gateway", 
                "messages": messages,
                "stream": DEBUG_STREAM_MODE  
            }

            # 发送请求：非流式模式下不需要启用 requests 的 stream=True 保持长连接
            response = requests.post(
                API_URL, 
                headers=headers, 
                json=payload, 
                verify=False, 
                stream=DEBUG_STREAM_MODE
            )

            if response.status_code == 200:
                ai_reply = ""

                # ===========================================================
                # 分支一：🌊 开启流式传输 (DEBUG_STREAM_MODE == True)
                # ===========================================================
                if DEBUG_STREAM_MODE:
                    print("\n🤖 Answer: ", end="")
                    sys.stdout.flush()
                    
                    for line in response.iter_lines():
                        if not line:
                            continue
                            
                        decoded_line = line.decode('utf-8').strip()
                        if not decoded_line.startswith("data: "):
                            continue
                            
                        content_data = decoded_line[6:].strip()
                        if content_data == "[DONE]":
                            break
                        
                        if DEBUG_MODE:
                            print(f"\n[RAW DATA]: {content_data}")
                            continue

                        try:
                            result = json.loads(content_data)
                            
                            # 轨道 A 适配：OpenAI 标准流式结构提取
                            if "choices" in result and len(result["choices"]) > 0:
                                delta = result["choices"][0].get("delta", {})
                                delta_content = delta.get("content", "")
                                if delta_content:
                                    print(delta_content, end="")
                                    sys.stdout.flush()
                                    ai_reply += delta_content
                                    
                            # 轨道 B 适配：自研 Responses 高级协议
                            elif "type" in result:
                                event_type = result.get("type")
                                if event_type == "response.output_text.delta":
                                    delta_content = result.get("delta", "")
                                    print(delta_content, end="")
                                    sys.stdout.flush()
                                    ai_reply += delta_content
                                elif event_type == "response.reasoning.delta":
                                    delta_reason = result.get("delta", "")
                                    print(f"\033[90m{delta_reason}\033[0m", end="")
                                    sys.stdout.flush()
                            
                        except json.JSONDecodeError:
                            if content_data:
                                print(content_data, end="")
                                sys.stdout.flush()
                                ai_reply += content_data
                    print() # 换行

                # ===========================================================
                # 分支二：📦 采用标准的非 SSE 模式握手 (DEBUG_STREAM_MODE == False)
                # ===========================================================
                else:
                    try:
                        result = response.json()
                        
                        # 解析标准 OpenAI 非流式响应包：choices[0].message.content
                        if "choices" in result and len(result["choices"]) > 0:
                            message_node = result["choices"][0].get("message", {})
                            ai_reply = message_node.get("content", "").strip()
                            
                            print(f"\n🤖 回答: {ai_reply}")
                        else:
                            print(f"\n⚠️ [异常]: 未能识别的非流式返回结构: {result}")
                    except json.JSONDecodeError:
                        # 兜底纯文本返回
                        ai_reply = response.text.strip()
                        print(f"\n🤖 回答: {ai_reply}")

                # 统一将有效的 AI 回复追加到多轮历史上下文中
                if ai_reply:
                    messages.append({"role": "assistant", "content": ai_reply})
                else:
                    if DEBUG_STREAM_MODE and not DEBUG_MODE:
                        print("\n⚠️ [警告]: 接收流已结束，但未解析到任何有效文本。")

            else:
                print(f"\n❌ 请求失败: HTTP {response.status_code}")
                if messages and messages[-1]["role"] == "user":
                    messages.pop()
                try:
                    print(f"错误详情: {response.text}")
                except Exception:
                    pass

        except KeyboardInterrupt:
            print("\n\n👋 catch quit signal.")
            break
        except requests.exceptions.RequestException as e:
            print(f"\n❌ 网络连接异常: {e}")
            if messages and messages[-1]["role"] == "user":
                messages.pop()

if __name__ == "__main__":
    main()