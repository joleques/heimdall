package templates

import "embed"

// DefaultTemplates embeds the canonical template library used by the CLI at runtime.
//
//go:embed default/**
var DefaultTemplates embed.FS
