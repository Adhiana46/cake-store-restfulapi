package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStringToSqlOrder_ValidFormat(t *testing.T) {
	rawOrder := "rating.DESC,title.ASC"

	expected := []SqlOrder{
		SqlOrder{
			Field: "rating",
			Dir:   "DESC",
		},
		SqlOrder{
			Field: "title",
			Dir:   "ASC",
		},
	}

	result, err := ParseStringToSqlOrder(rawOrder)

	assert.Nil(t, err)

	assert.Equal(t, expected, result)
}

func TestParseStringToSqlOrder_InvalidFormat(t *testing.T) {
	rawOrder := "rating,title.ASC"

	result, err := ParseStringToSqlOrder(rawOrder)

	assert.Nil(t, result)

	assert.NotNil(t, err)
}

func TestParseEmptyStringToSqlOrder(t *testing.T) {
	rawOrder := ""

	result, err := ParseStringToSqlOrder(rawOrder)

	assert.NotNil(t, err)

	assert.Nil(t, result)
}
