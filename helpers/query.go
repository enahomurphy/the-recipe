package helpers

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

// Query params for pagination
type Query struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Q      string `json:"q"`
}

// Pagination  metadata
type Pagination struct {
	PageCount   int `json:"page_count"`
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
	TotalCount  int `json:"total_count"`
}

// UpdateBuilder generates update query dynamically
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

// GetCount get user count
func GetCount(db *sql.DB, table string) (int, error) {
	var count int

	row, countErr := db.Query(`SELECT COUNT(*) as count FROM ` + table)
	if countErr != nil {
		errMsg := fmt.Errorf("an unknown error occurred %s", countErr.Error())
		return 0, errMsg
	}

	for row.Next() {
		scanErr := row.Scan(&count)
		if scanErr != nil {
			errMsg := fmt.Errorf("an unknown error occurred %s", scanErr.Error())
			return 0, errMsg
		}
	}
	return count, nil
}

// MetaData creates meta_data for pagination
func MetaData(count int, rowsLength int, query *Query) interface{} {
	next := math.Ceil(float64(count) / float64(query.Limit))
	fmt.Println(strconv.Itoa(count))
	fmt.Println(strconv.Itoa(query.Limit))

	currentPage := math.Floor(float64(query.Offset/query.Limit) + 1)

	pagination := Pagination{
		PageCount:   int(next),
		CurrentPage: int(currentPage),
		PageSize:    rowsLength,
		TotalCount:  count,
	}
	return pagination
}

// GetQuery gets user query
func GetQuery(r *http.Request) (Query, error) {

	query := Query{}

	formLimit := r.FormValue("limit")
	formOffset := r.FormValue("offset")
	formQ := r.FormValue("q")

	if formLimit != "" {
		limit, err := strconv.Atoi(formLimit)
		if err != nil {
			return query, errors.New("Invalid limit")
		}
		query.Limit = limit
	} else {
		query.Limit = 10
	}

	if formOffset != "" {
		offset, err := strconv.Atoi(formOffset)
		if err != nil {
			return query, errors.New("Invalid offset")
		}
		query.Offset = offset
	} else {
		query.Offset = 0
	}

	if formQ != "" {
		query.Q = formQ
	}

	return query, nil

}
