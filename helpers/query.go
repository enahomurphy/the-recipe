package helpers

import (
	"bytes"
)

// UpdateBuilder
func UpdateBuilder(u map[string]string) string {
	var query bytes.Buffer
	query.WriteString("UPDATE users SET ")

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

	return query.String()
}

func B(id int) int {
	return id
}

// MapBuilder is
// func MapBuilder(arr []string, vals ) {
// 	for value, key := range arr {

// 	}
// }
