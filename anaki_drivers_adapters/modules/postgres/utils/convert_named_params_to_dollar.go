package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ConvertNamedParams(query string, params map[string]interface{}) (string, []interface{}, error) {
	re := regexp.MustCompile(`:\w+`)
	matches := re.FindAllString(query, -1)
	if len(matches) == 0 {
		return query, nil, fmt.Errorf("no parameters found in query")
	}

	uniqueParams := make(map[string]struct{})
	var orderedParams []string

	for _, m := range matches {
		param := strings.TrimPrefix(m, ":")
		if _, exists := uniqueParams[param]; !exists {
			uniqueParams[param] = struct{}{}
			orderedParams = append(orderedParams, param)
		}
	}

	// sort.Strings(orderedParams)

	args := make([]interface{}, len(orderedParams))
	for i, param := range orderedParams {
		value, ok := params[param]
		if !ok {
			return "", nil, fmt.Errorf("missing value for parameter: %s", param)
		}
		args[i] = value
	}

	finalQuery := query
	for i, param := range orderedParams {
		finalQuery = strings.ReplaceAll(finalQuery, ":"+param, fmt.Sprintf("$%d", i+1))
	}

	return finalQuery, args, nil
}
