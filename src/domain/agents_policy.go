package domain

import "fmt"

type AgentsPolicy string

const (
	AgentsPolicySkip      AgentsPolicy = "skip"
	AgentsPolicyIfMissing AgentsPolicy = "if-missing"
	AgentsPolicyOverwrite AgentsPolicy = "overwrite"
)

const DefaultAgentsPolicy = AgentsPolicyIfMissing

func (p AgentsPolicy) Validate() error {
	switch p {
	case AgentsPolicySkip, AgentsPolicyIfMissing, AgentsPolicyOverwrite:
		return nil
	default:
		return fmt.Errorf("invalid agents policy: %q", p)
	}
}
