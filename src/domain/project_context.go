package domain

import (
	"fmt"
	"strings"
)

type ProjectContext struct {
	Target        TargetPlatform
	Title         string
	Description   string
	Documentation []string
}

func (c ProjectContext) Validate() error {
	var problems []string

	if _, err := ParseTargetPlatform(string(c.Target)); err != nil {
		problems = append(problems, "target is required")
	}

	if strings.TrimSpace(c.Title) == "" {
		problems = append(problems, "title is required")
	}

	if strings.TrimSpace(c.Description) == "" {
		problems = append(problems, "description is required")
	}

	for i, item := range c.Documentation {
		if strings.TrimSpace(item) == "" {
			problems = append(problems, fmt.Sprintf("documentation[%d] is required", i))
		}
	}

	if len(problems) > 0 {
		return fmt.Errorf("invalid project context: %s", strings.Join(problems, "; "))
	}

	return nil
}

func (c ProjectContext) Normalized() ProjectContext {
	return ProjectContext{
		Target:        TargetPlatform(strings.TrimSpace(string(c.Target))),
		Title:         strings.TrimSpace(c.Title),
		Description:   strings.TrimSpace(c.Description),
		Documentation: dedupeAndTrim(c.Documentation),
	}
}
