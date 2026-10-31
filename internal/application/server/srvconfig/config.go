package srvconfig

import (
	"time"
)

type ServerConfig struct {
	Log            Log                  `toml:"log" yaml:"log" envconfig:"LOG"`
	Application    ServerApplication    `toml:"application" yaml:"application" envconfig:"APPLICATION"`
	Parsing        ServerParsing        `toml:"parsing" yaml:"parsing" envconfig:"PARSING"`
	Workers        []ServerWorker       `toml:"workers" yaml:"workers" envconfig:"WORKERS"`
	Storage        ServerStorage        `toml:"storage" yaml:"storage" envconfig:"STORAGE"`
	FileStorage    ServerFileStorage    `toml:"file_storage" yaml:"file_storage" envconfig:"FILE_STORAGE"`
	API            ServerAPI            `toml:"api" yaml:"api" envconfig:"API"`
	AgentServer    ServerAgentServer    `toml:"agent_server" yaml:"agent_server" envconfig:"AGENT_SERVER"`
	AttributeRemap ServerAttributeRemap `toml:"attribute_remap" yaml:"attribute_remap" envconfig:"ATTRIBUTE_REMAP"`
	MCPServer      ServerMCPServer      `toml:"mcp_server" yaml:"mcp_server" envconfig:"MCP_SERVER"`
}

type ServerParsing struct {
	ParseBookTimeout time.Duration `toml:"parse_book_timeout" yaml:"parse_book_timeout" envconfig:"PARSE_BOOK_TIMEOUT"`
	AgentTimeout     time.Duration `toml:"agent_timeout" yaml:"agent_timeout" envconfig:"AGENT_TIMEOUT"`
}

type ServerAttributeRemap struct {
	Auto     bool `toml:"auto" yaml:"auto" envconfig:"AUTO"`
	AllLower bool `toml:"all_lower" yaml:"all_lower" envconfig:"ALL_LOWER"`
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
