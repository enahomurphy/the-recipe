package helpers

import (
	"bytes"
	"fmt"
)

// UpdateBuilder
func UpdateBuilder(u map[string]string, table string) string {
	var query bytes.Buffer
	query.WriteString("UPDATE " + table + " SET ")
	var result string
	length := len(u)

	for value, key := range u {
		comma := ","
		if length == 1 {
			comma = ""
		}
		query.WriteString(value + " = " + "'" + key + "'" + comma + " ")
		length--
	}

	query.WriteString("WHERE id = ? ")
	fmt.Println(query.String())
	result = query.String()
	return result
}
