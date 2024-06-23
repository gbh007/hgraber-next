package static

import "embed"

//go:embed *.html assets/* css/* js/*
var StaticDir embed.FS
