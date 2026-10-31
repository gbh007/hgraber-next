package agnconfig

import "time"

type AgentFSBase struct {
	ImportPath          string `toml:"import_path"`
	FilePath            string `toml:"file_path"`
	EnableDeduplication bool   `toml:"enable_deduplication"`
	ImportLimitOnFolder int    `toml:"import_limit_on_folder"`
}

func AgentDefaultFSBase() AgentFSBase {
	return AgentFSBase{}
}

type AgentSqlite struct {
	FilePath string `toml:"file_path"`
}

func AgentDefaultSqlite() AgentSqlite {
	return AgentSqlite{}
}

type AgentZipScanner struct {
	MasterAddr  string `toml:"master_addr"`
	MasterToken string `toml:"master_token"`
}

func AgentDefaultZipScanner() AgentZipScanner {
	return AgentZipScanner{}
}

type AgentHighway struct {
	PrivateKey    string        `toml:"private_key"`
	TokenLifetime time.Duration `toml:"token_lifetime"`
}

func AgentDefaultHighway() AgentHighway {
	return AgentHighway{}
}
