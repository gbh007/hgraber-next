package generatorcore

import (
	"strings"

	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func StrToPtr(s string) *string {
	return &s
}

func BoolToPtr(b bool) *bool {
	return &b
}

func NameToVarDS(s string) string {
	return "${" + s + "}"
}

func NameToVar(s string) string {
	return "$" + s
}

func ValuesFromString(s string) dashboard.StringOrMap {
	return dashboard.StringOrMap{
		String: &s,
	}
}

func ValuesFromArray(a []string) dashboard.StringOrMap {
	s := strings.Join(a, ",")

	return dashboard.StringOrMap{
		String: &s,
	}
}
