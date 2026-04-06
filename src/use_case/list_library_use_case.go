package usecase

import (
	"context"
	"sort"
	"strings"
)

type ListLibraryUseCase struct {
	catalog CatalogGateway
}

func NewListLibraryUseCase(catalog CatalogGateway) ListLibraryUseCase {
	return ListLibraryUseCase{catalog: catalog}
}

func (uc ListLibraryUseCase) Execute(ctx context.Context, request ListLibraryRequest) (ListLibraryResult, error) {
	catalog, err := uc.catalog.Load(ctx, request.OutputDir)
	if err != nil {
		return ListLibraryResult{}, err
	}

	filterCategory := normalizeListCategory(request.Category)
	includeSkills := request.IncludeSkills || filterCategory != ""

	items := make([]AssistantLibraryItem, 0, len(catalog.Assistants))
	for _, assistant := range catalog.Assistants {
		if !matchesCategory(assistant.Categories, filterCategory) {
			continue
		}

		items = append(items, AssistantLibraryItem{
			ID:          assistant.ID,
			Name:        assistant.Name,
			Description: assistant.Description,
			Skills:      append([]string(nil), assistant.Skills...),
			Categories:  append([]string(nil), assistant.Categories...),
		})
	}

	sort.Slice(items, func(i, j int) bool { return items[i].ID < items[j].ID })

	result := ListLibraryResult{Assistants: items}
	if includeSkills {
		skills := make([]SkillLibraryItem, 0, len(catalog.Skills))
		for _, skill := range catalog.Skills {
			if !matchesCategory(skill.Categories, filterCategory) {
				continue
			}

			description := ""
			if skill.Contract != nil {
				description = skill.Contract.Description
			}

			skills = append(skills, SkillLibraryItem{
				ID:          skill.Name,
				Description: description,
				Categories:  append([]string(nil), skill.Categories...),
			})
		}

		sort.Slice(skills, func(i, j int) bool { return skills[i].ID < skills[j].ID })
		result.Skills = skills
	}

	return result, nil
}

func normalizeListCategory(value string) string {
	normalized := strings.ToLower(strings.TrimSpace(value))
	normalized = strings.ReplaceAll(normalized, "_", "-")
	normalized = strings.ReplaceAll(normalized, " ", "-")
	return normalized
}

func matchesCategory(categories []string, desired string) bool {
	if desired == "" {
		return true
	}

	for _, category := range categories {
		if normalizeListCategory(category) == desired {
			return true
		}
	}

	return false
}

var _ ListLibrary = ListLibraryUseCase{}
