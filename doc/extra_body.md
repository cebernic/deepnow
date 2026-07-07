==Extra_body definitions for major reasoning models.
<hr>

Deepseek Family :

reasoning_effort = "high"
thinking = {"type":"enabled"}


Gemini Family :
Gemini 2.5 series:
extra_body = {"google":{"thinking_config":{"thinking_budget":8192,"include_thoughts":true}}}

Gemini 3 series:
extra_body = {"google":{"thinking_config":{"thinking_level":"high","include_thoughts":true}}}

