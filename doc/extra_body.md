## Extra_body definitions for major reasoning models. ##
<hr>
These are only a partial example of configurations; you can refer to each model's official API manual to add more switches and parameters.

<hr>

> Copy the full code into the "extra body" on Dashboard GUI  -> Menu "Model Configure" in Deepnow with model name matching.

### Deepseek 4 Family (think mode enable, for clients which connected via deepnow ): ###
```json
reasoning_effort = "high"
thinking = {"type":"enabled"}
```

### Gemini Family : ###

- Gemini 2.5 series (think mode & thought msg for clients which connected via deepnow ):
```json
extra_body = {"google":{"thinking_config":{"thinking_budget":1024,"include_thoughts":true}}}
```

- Gemini 3 series ( think mode & thought msg for clients which connected via deepnow ):
```json
extra_body = {"google":{"thinking_config":{"thinking_level":"high","include_thoughts":true}}}
```
### MiniMax Family (think mode): ###

- M3 or M2.7
```json
reasoning = {"effort":"high"}
reasoning_split = true
service_tier = "priority"
thinking = {"type":"adaptive"}
```
