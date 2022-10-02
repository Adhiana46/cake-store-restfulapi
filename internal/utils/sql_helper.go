package utils

import (
	"errors"
	"fmt"
	"strings"
)

const (
	SqlOrderAsc  = "ASC"
	SqlOrderDesc = "DESC"
)

type SqlWhere struct {
	Field    string // Nama field
	Operator string // Operator: =, >=, <=, LIKE
	Value    any    // Value
}

type SqlOrder struct {
	Field string
	Dir   string
}

// rating.desc,title.asc
func ParseStringToSqlOrder(rawOrder string) ([]SqlOrder, error) {
	if rawOrder == "" {
		return nil, errors.New("ParseStringToSqlOrder `rawOrder` cannot empty")
	}

	sqlOrders := []SqlOrder{}

	orders := strings.Split(rawOrder, ",")
	for _, order := range orders {
		orderArr := strings.Split(order, ".")

		if len(orderArr) != 2 {
			return nil, errors.New(fmt.Sprintf("ParseStringToSqlOrder invalid format for `%s`", order))
		}

		sqlOrders = append(sqlOrders, SqlOrder{
			Field: orderArr[0],
			Dir:   orderArr[1],
		})
	}

	return sqlOrders, nil
}

// mysql escape string
func MysqlRealEscapeString(value string) string {
	replace := map[string]string{"\\": "\\\\", "'": `\'`, "\\0": "\\\\0", "\n": "\\n", "\r": "\\r", `"`: `\"`, "\x1a": "\\Z"}

	for b, a := range replace {
		value = strings.Replace(value, b, a, -1)
	}

	return value
}
