package domain

import "fmt"

type TargetPlatform string

const (
	TargetCodex       TargetPlatform = "codex"
	TargetAntigravity TargetPlatform = "antigravity"
	TargetClaude      TargetPlatform = "claude"
	TargetCursor      TargetPlatform = "cursor"
)

func ParseTargetPlatform(value string) (TargetPlatform, error) {
	t := TargetPlatform(value)
	switch t {
	case TargetCodex, TargetAntigravity, TargetClaude, TargetCursor:
		return t, nil
	default:
		return "", fmt.Errorf("invalid target platform: %q", value)
	}
}
