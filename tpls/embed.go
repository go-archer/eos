package tpls

import "embed"

//go:embed creator/*.tpl
var CreateTemplateFS embed.FS
