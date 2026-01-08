# Фича с MCP

Пример конфигурации сервера

```toml
[mcp_server]
addr = ":8888"
token = "mcp-hg"
debug = true
```

Пример конфигурации LM-studio

```json
{
  "mcpServers": {
    "hg-next": {
      "url": "http://localhost:8888/mcp",
      "headers": {
        "X-HG-Token": "mcp-hg"
      }
    }
  }
}
```
