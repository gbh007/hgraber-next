package config

import "time"

type AgentFSBase struct {
	ImportPath          string `toml:"import_path" yaml:"import_path" envconfig:"IMPORT_PATH"`
	FilePath            string `toml:"file_path" yaml:"file_path" envconfig:"FILE_PATH"`
	EnableDeduplication bool   `toml:"enable_deduplication" yaml:"enable_deduplication" envconfig:"ENABLE_DEDUPLICATION"`
	ImportLimitOnFolder int    `toml:"import_limit_on_folder" yaml:"import_limit_on_folder" envconfig:"IMPORT_LIMIT_ON_FOLDER"`
}

func AgentDefaultFSBase() AgentFSBase {
	return AgentFSBase{}
}

type AgentSqlite struct {
	FilePath string `toml:"file_path" yaml:"file_path" envconfig:"FILE_PATH"`
}

func AgentDefaultSqlite() AgentSqlite {
	return AgentSqlite{}
}

type AgentZipScanner struct {
	MasterAddr  string `toml:"master_addr" yaml:"master_addr" envconfig:"MASTER_ADDR"`
	MasterToken string `toml:"master_token" yaml:"master_token" envconfig:"MASTER_TOKEN"`
}

func AgentDefaultZipScanner() AgentZipScanner {
	return AgentZipScanner{}
}

type AgentHighway struct {
	PrivateKey    string        `toml:"private_key" yaml:"private_key" envconfig:"PRIVATE_KEY"`
	TokenLifetime time.Duration `toml:"token_lifetime" yaml:"token_lifetime" envconfig:"TOKEN_LIFETIME"`
}

func AgentDefaultHighway() AgentHighway {
	return AgentHighway{}
}
