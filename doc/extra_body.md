## Extra_body definitions for major reasoning models. ##
<hr>
These are only a partial example of configurations; you can refer to each model's official API manual to add more switches and parameters.

<hr>

> Copy the full code into the extra body on Menu Model Configure in Deepnow.

### Deepseek 4 Family : ###
```json
reasoning_effort = "high"
thinking = {"type":"enabled"}
```

### Gemini Family : ###

- Gemini 2.5 series:
```json
extra_body = {"google":{"thinking_config":{"thinking_budget":8192,"include_thoughts":true}}}
```

- Gemini 3 series:
```json
extra_body = {"google":{"thinking_config":{"thinking_level":"high","include_thoughts":true}}}
```
### MiniMax Family : ###

- M3
```json
thinking: {"type": "adaptive"}
reasoning_split": true
```
