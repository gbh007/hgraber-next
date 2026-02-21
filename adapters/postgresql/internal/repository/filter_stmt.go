package repository

import (
	"regexp"
	"slices"
	"strings"
)

const (
	insertKey = "insert"
	selectKey = "select"
	updateKey = "update"
	deleteKey = "delete"
)

var (
	selectRgx = regexp.MustCompile(`(?mi)select.+?from\s+([^\s,()]+)`)
	joinRgx   = regexp.MustCompile(`(?mi)join\s*([^\s,()]+)`)
	insertRgx = regexp.MustCompile(`(?mi)insert\s+into\s+([^\s,()]+)`)
	updateRgx = regexp.MustCompile(`(?mi)update\s+([^\s,()]+)\s+set`)
	deleteRgx = regexp.MustCompile(`(?mi)delete\s+from\s+([^\s,()]+)`)
)

func filterStmt(s string) (string, bool) {
	tmp := map[string][]string{}

	for _, match := range selectRgx.FindAllStringSubmatch(s, -1) {
		if len(match) > 1 {
			tmp[selectKey] = append(tmp[selectKey], match[1:]...)
		}
	}

	for _, match := range joinRgx.FindAllStringSubmatch(s, -1) {
		if len(match) > 1 {
			tmp[selectKey] = append(tmp[selectKey], match[1:]...)
		}
	}

	for _, match := range insertRgx.FindAllStringSubmatch(s, -1) {
		if len(match) > 1 {
			tmp[insertKey] = append(tmp[insertKey], match[1:]...)
		}
	}

	for _, match := range updateRgx.FindAllStringSubmatch(s, -1) {
		if len(match) > 1 {
			tmp[updateKey] = append(tmp[updateKey], match[1:]...)
		}
	}

	for _, match := range deleteRgx.FindAllStringSubmatch(s, -1) {
		if len(match) > 1 {
			tmp[deleteKey] = append(tmp[deleteKey], match[1:]...)
		}
	}

	result := &strings.Builder{}

	for _, k := range []string{
		insertKey,
		selectKey,
		updateKey,
		deleteKey,
	} {
		if values := tmp[k]; len(values) > 0 {
			if result.Len() > 0 {
				_, _ = result.WriteString(" ")
			}

			_, _ = result.WriteString(k + ":")

			for i, v := range values {
				values[i] = strings.ToLower(v)
			}

			slices.Sort(values)
			values = slices.Compact(values)

			_, _ = result.WriteString(strings.Join(values, ","))
		}
	}

	return result.String(), result.Len() > 0
}
