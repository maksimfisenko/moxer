package generator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

var varPattern = regexp.MustCompile(`{{\s*([a-zA-Z0-9_]+)\s*}}`)

var generators = map[string]func() any{
	"uuid": func() any { return uuid.New().String() },
	"name": func() any { return "random_name" },
}

func GenerateData(content map[string]any, count int) []map[string]any {
	data := make([]map[string]any, count)
	for i := range count {
		data[i] = resolveMap(content)
	}
	return data
}

func resolveMap(m map[string]any) map[string]any {
	result := make(map[string]any)
	for k, v := range m {
		result[k] = resolve(v)
	}
	return result
}

func resolve(value any) any {
	switch v := value.(type) {
	case string:
		return resolveString(v)
	case map[string]any:
		return resolveMap(v)
	default:
		return v
	}
}

func resolveString(s string) string {
	return varPattern.ReplaceAllStringFunc(s, func(match string) string {
		varName := strings.TrimSpace(match[2 : len(match)-2])
		if genFunc, ok := generators[varName]; ok {
			return fmt.Sprintf("%v", genFunc())
		}
		return match
	})
}
