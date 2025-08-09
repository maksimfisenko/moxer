package generator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
)

var varPattern = regexp.MustCompile(`{{\s*([a-zA-Z0-9_]+)\s*}}`)

var generators = map[string]func() any{
	// Misc
	"uuid": func() any { return gofakeit.UUID() },
	"bool": func() any { return strconv.FormatBool(gofakeit.Bool()) },
	// Person
	"name":        func() any { return gofakeit.Name() },
	"first_name":  func() any { return gofakeit.FirstName() },
	"middle_name": func() any { return gofakeit.MiddleName() },
	"last_name":   func() any { return gofakeit.LastName() },
	"phone":       func() any { return gofakeit.Phone() },
	"email":       func() any { return gofakeit.Email() },
	"username":    func() any { return gofakeit.Username() },
	"password":    func() any { return gofakeit.Password(true, true, true, true, false, 8) },
	"gender":      func() any { return gofakeit.Gender() },
	// Address
	"country":    func() any { return gofakeit.Country() },
	"city":       func() any { return gofakeit.City() },
	"street":     func() any { return gofakeit.Street() },
	"zip":        func() any { return gofakeit.Zip() },
	"latitude":   func() any { return fmt.Sprintf("%f", gofakeit.Latitude()) },
	"longtitude": func() any { return fmt.Sprintf("%f", gofakeit.Longitude()) },
	// Words
	"word":        func() any { return gofakeit.Word() },
	"noun":        func() any { return gofakeit.Noun() },
	"verb":        func() any { return gofakeit.Verb() },
	"adverb":      func() any { return gofakeit.Adverb() },
	"preposition": func() any { return gofakeit.Preposition() },
	"adjective":   func() any { return gofakeit.Adjective() },
	"pronoun":     func() any { return gofakeit.Pronoun() },
	"phrase":      func() any { return gofakeit.Phrase() },
	"question":    func() any { return gofakeit.Question() },
	// Colors
	"color": func() any { return gofakeit.Color() },
	"hex":   func() any { return gofakeit.HexColor() },
	// Internet
	"url":           func() any { return gofakeit.URL() },
	"domain_name":   func() any { return gofakeit.DomainName() },
	"domain_suffix": func() any { return gofakeit.DomainSuffix() },
	"ipv4":          func() any { return gofakeit.IPv4Address() },
	"ipv6":          func() any { return gofakeit.IPv6Address() },
	// Date/time
	"date":        func() any { return gofakeit.Date().String() },
	"past_date":   func() any { return gofakeit.PastDate().String() },
	"future_date": func() any { return gofakeit.FutureDate().String() },
	"year":        func() any { return strconv.Itoa(gofakeit.Year()) },
	"month":       func() any { return strconv.Itoa(gofakeit.Month()) },
	"weekday":     func() any { return gofakeit.WeekDay() },
	"hour":        func() any { return strconv.Itoa(gofakeit.Hour()) },
	"minute":      func() any { return strconv.Itoa(gofakeit.Minute()) },
	"second":      func() any { return strconv.Itoa(gofakeit.Second()) },
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
