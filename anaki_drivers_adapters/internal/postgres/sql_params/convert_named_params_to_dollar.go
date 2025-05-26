package sqlparams

import (
	"fmt"
	"regexp"
	"strings"
)

// ConvertNamedParams converts a query with named parameters (e.g., :param) into a
// query with dollar-sign parameters (e.g., $1, $2) and extracts the corresponding
// arguments from a map based on the order of appearance of named parameters in the query.
//
// If the query does not contain any named parameters, it returns the original query
// and an empty slice of arguments.
//
// It returns an error if a named parameter in the query does not have a corresponding
// value in the provided params map.
//
// Example:
//
//   Input:
//     query = "SELECT * from users where id = :id AND name = :name AND id = :id"
//     params = { "id": 2, "name": "Marcos" }
//
//   Output:
//     finalQuery = "SELECT * from users where id = $1 AND name = $2 AND id = $1"
//     args = []interface{}{2, "Marcos"}
//     err = nil
//
func ConvertNamedParamsToDollar(query string, params map[string]interface{}) (string, []interface{}, error) {
	if !strings.Contains(query, ":") {
		return query, []interface{}{}, nil
	}

	re := regexp.MustCompile(`:\w+`)
	matches := re.FindAllString(query, -1)
	if len(matches) == 0 {
		return query, nil, nil
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
