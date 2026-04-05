package domain

import (
	"fmt"
	"regexp"
	"strings"
)

var assistantIDPattern = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

type AssistantSpec struct {
	ID           string
	Name         string
	Description  string
	Instructions string
	Version      string
	Skills       []string
	Inputs       []InputSpec
	Tools        []string
	Tags         []string
	Metadata     map[string]string
}

type InputSpec struct {
	Name        string
	Description string
	Required    bool
	Default     string
}

func (a AssistantSpec) Validate() error {
	var problems []string

	if a.ID == "" {
		problems = append(problems, "id is required")
	} else if !assistantIDPattern.MatchString(a.ID) {
		problems = append(problems, "id must match ^[a-z0-9]+(-[a-z0-9]+)*$")
	}

	if strings.TrimSpace(a.Name) == "" {
		problems = append(problems, "name is required")
	}

	if strings.TrimSpace(a.Description) == "" {
		problems = append(problems, "description is required")
	}

	instructions := strings.TrimSpace(a.Instructions)
	if instructions == "" {
		problems = append(problems, "instructions is required")
	} else if len([]rune(instructions)) < 30 {
		problems = append(problems, "instructions must have at least 30 characters")
	}

	inputNames := make(map[string]struct{}, len(a.Inputs))
	for _, in := range a.Inputs {
		name := strings.TrimSpace(in.Name)
		if name == "" {
			problems = append(problems, "inputs.name is required")
			continue
		}

		if strings.TrimSpace(in.Description) == "" {
			problems = append(problems, fmt.Sprintf("input %q description is required", name))
		}

		if _, exists := inputNames[name]; exists {
			problems = append(problems, fmt.Sprintf("input %q is duplicated", name))
			continue
		}
		inputNames[name] = struct{}{}
	}

	if hasDuplicate(a.Skills) {
		problems = append(problems, "skills must not contain duplicates")
	}

	if hasDuplicate(a.Tools) {
		problems = append(problems, "tools must not contain duplicates")
	}

	if hasDuplicate(a.Tags) {
		problems = append(problems, "tags must not contain duplicates")
	}

	if len(problems) > 0 {
		return fmt.Errorf("invalid assistant spec: %s", strings.Join(problems, "; "))
	}
	return nil
}

func (a AssistantSpec) Normalized() AssistantSpec {
	normalized := a
	normalized.ID = strings.TrimSpace(a.ID)
	normalized.Name = strings.TrimSpace(a.Name)
	normalized.Description = strings.TrimSpace(a.Description)
	normalized.Instructions = strings.TrimSpace(a.Instructions)

	if strings.TrimSpace(normalized.Version) == "" {
		normalized.Version = "0.1.0"
	} else {
		normalized.Version = strings.TrimSpace(normalized.Version)
	}

	normalized.Skills = dedupeAndTrim(a.Skills)

	normalized.Inputs = make([]InputSpec, 0, len(a.Inputs))
	for _, in := range a.Inputs {
		normalized.Inputs = append(normalized.Inputs, InputSpec{
			Name:        strings.TrimSpace(in.Name),
			Description: strings.TrimSpace(in.Description),
			Required:    in.Required,
			Default:     strings.TrimSpace(in.Default),
		})
	}

	normalized.Tools = dedupeAndTrim(a.Tools)
	normalized.Tags = dedupeAndTrim(a.Tags)

	if normalized.Metadata == nil {
		normalized.Metadata = map[string]string{}
	}

	return normalized
}

func hasDuplicate(values []string) bool {
	seen := make(map[string]struct{}, len(values))
	for _, v := range values {
		normalized := strings.TrimSpace(v)
		if normalized == "" {
			continue
		}

		if _, exists := seen[normalized]; exists {
			return true
		}
		seen[normalized] = struct{}{}
	}
	return false
}

func dedupeAndTrim(values []string) []string {
	if len(values) == 0 {
		return []string{}
	}

	result := make([]string, 0, len(values))
	seen := make(map[string]struct{}, len(values))
	for _, v := range values {
		normalized := strings.TrimSpace(v)
		if normalized == "" {
			continue
		}

		if _, exists := seen[normalized]; exists {
			continue
		}
		seen[normalized] = struct{}{}
		result = append(result, normalized)
	}

	return result
}
