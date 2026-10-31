package srvconfig

import (
	"time"
)

type ServerConfig struct {
	Log            Log                  `toml:"log"`
	Application    ServerApplication    `toml:"application"`
	Parsing        ServerParsing        `toml:"parsing"`
	Workers        []ServerWorker       `toml:"workers"`
	Storage        ServerStorage        `toml:"storage"`
	FileStorage    ServerFileStorage    `toml:"file_storage"`
	API            ServerAPI            `toml:"api"`
	AgentServer    ServerAgentServer    `toml:"agent_server"`
	AttributeRemap ServerAttributeRemap `toml:"attribute_remap"`
	MCPServer      ServerMCPServer      `toml:"mcp_server"`
}

type ServerParsing struct {
	ParseBookTimeout time.Duration `toml:"parse_book_timeout"`
	AgentTimeout     time.Duration `toml:"agent_timeout"`
}

type ServerAttributeRemap struct {
	Auto     bool `toml:"auto"`
	AllLower bool `toml:"all_lower"`
}

func ServerConfigDefault() ServerConfig {
	return ServerConfig{
		Log:         LogDefault(),
		Application: ServerApplicationDefault(),
		Parsing:     ServerParsingDefault(),
		API:         ServerAPIDefault(),
		Storage:     ServerStorageDefault(),
		FileStorage: ServerFileStorageDefault(),
		AgentServer: ServerAgentServerDefault(),
	}
}

func ServerParsingDefault() ServerParsing {
	return ServerParsing{
		ParseBookTimeout: DefaultParsingBookTimeout,
		AgentTimeout:     DefaultParsingAgentTimeout,
	}
}

func ServerAttributeRemapDefault() ServerAttributeRemap {
	return ServerAttributeRemap{}
}
