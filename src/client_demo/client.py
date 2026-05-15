import requests
import urllib3
import json

# 忽略自签名证书的安全警告
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

# ==========================================
#                  配置区
# ==========================================
# 目标 deepnow Server 地址
API_URL = "https://127.0.0.1:8444/v1/chat/completions"

# 授权 Key
API_KEY = "sk-deepnow-18ad38cd8fcb608c62c68aeed5090f3a"

# 设定系统人设与强制返回格式 (可留空)
SYSTEM_PROMPT = "你是一个AI助手，协助用户解答问题。请确保回答简洁明了，直接切入主题，不要添加无关信息。"
#"你是一个单选题回答助手，每次回答只帮我返回一个选择.比如A,B,C,D这种，A,B,C这种。注意：无论如何都要做一个选择,并且需要附带选择后面的选项内容而不是一个单纯的字母且不能有其它废话。"
# ==========================================

def main():
    print("==================================================")
    print("🚀 Deepnow 终端控制台客户端已启动")
    print(f"🔗 目标端点: {API_URL}")
    if SYSTEM_PROMPT:
        print(f"🧠 系统指令: {SYSTEM_PROMPT}")
    print("💡 输入内容开始对话，输入 '!!exit' 或按 Ctrl+C 退出")
    print("==================================================\n")

    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {API_KEY}"
    }

    # 维护对话上下文历史
    messages = []
    
    # 如果配置了系统提示词，将其作为第一条消息注入
    if SYSTEM_PROMPT.strip():
        messages.append({"role": "system", "content": SYSTEM_PROMPT})

    while True:
        try:
            user_input = input("\n🧑 提问: ")
            
            if not user_input.strip():
                continue

            if user_input.strip() == "!!exit":
                print("👋 退出测试。")
                break

            messages.append({"role": "user", "content": user_input})

            payload = {
                "model": "ignored_by_gateway", 
                "messages": messages
            }

            response = requests.post(API_URL, headers=headers, json=payload, verify=False)

            if response.status_code == 200:
                result = response.json()
                ai_reply = result.get("choices", [{}])[0].get("message", {}).get("content", "")
                
                print(f"\n🤖 回答: {ai_reply}")

                # 将 AI 的回复加入上下文，实现持续对话
                messages.append({"role": "assistant", "content": ai_reply})
            else:
                print(f"\n❌ 请求失败: HTTP {response.status_code}")
                try:
                    err_msg = response.json().get("error", response.text)
                    print(f"详细信息: {err_msg}")
                except json.JSONDecodeError:
                    print(f"详细信息: {response.text}")
                
                # 请求失败时弹出刚加入的用户消息，防污染
                messages.pop()

        except KeyboardInterrupt:
            print("\n\n👋 退出测试。")
            break
        except requests.exceptions.RequestException as e:
            print(f"\n❌ 网络连接异常: {e}")
            if len(messages) > 0 and messages[-1]["role"] == "user":
                messages.pop()

if __name__ == "__main__":
    main()