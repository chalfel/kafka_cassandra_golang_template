package helpers

import (
	"encoding/base64"
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Metadata PaginationMeta `json:"metadata"`
	Data     interface{}    `json:"data"`
}

type PaginationQueryData struct {
	ActualLimit int32
	Limit       int32
	Cursor      string
}

type PaginationMeta struct {
	NextCursor *string `json:"nextCursor"`
	PastCursor *string `json:"pastCursor"`
	HasMore    bool    `json:"hasMore"`
}

var (
	ErrInvalidCursorField = errors.New("invalid cursor field")
)

func DecodeString(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

func EncodeTimeBasedString(timeBased time.Time) string {
	return base64.StdEncoding.EncodeToString([]byte(timeBased.Format(time.RFC3339)))
}

func CreatePagination[T any](s []T, query PaginationQueryData, cursorField string) (Pagination, error) {
	pagination := Pagination{}
	metadata := PaginationMeta{}

	hasMore := query.ActualLimit == int32(len(s))

	if hasMore {
		lastRow := s[query.Limit-1]
		lastRowValue := reflect.ValueOf(lastRow)
		lastCursor := reflect.Indirect(lastRowValue).FieldByName(cursorField)

		if lastCursor.Kind() == 0 || !lastCursor.CanConvert(reflect.TypeOf(time.Time{})) {
			return pagination, ErrInvalidCursorField
		}

		lastCursorTimestamp := lastCursor.Interface().(time.Time)
		nextCursor := EncodeTimeBasedString(lastCursorTimestamp)
		metadata.NextCursor = &nextCursor
		s = s[:query.Limit]
	}

	if len(s) > 0 {
		firstRow := s[0]
		firstRowValue := reflect.ValueOf(firstRow)
		firstCursor := reflect.Indirect(firstRowValue).FieldByName(cursorField)

		if firstCursor.Kind() == 0 || !firstCursor.CanConvert(reflect.TypeOf(time.Time{})) {
			return pagination, ErrInvalidCursorField
		}

		pagination.Data = s

		firstCursorTimestamp := firstCursor.Interface().(time.Time)
		pastCursor := EncodeTimeBasedString(firstCursorTimestamp)
		metadata.PastCursor = &pastCursor
	} else {
		//This is a temporary code to avoid the error when the cursor is empty
		pagination.Data = interface{}(nil)
	}

	metadata.HasMore = hasMore
	pagination.Metadata = metadata

	return pagination, nil
}

func DecodeQueryParams(c *gin.Context) (PaginationQueryData, error) {
	var data = PaginationQueryData{}
	var err error

	queryCursor := c.Query("cursor")
	queryLimit := c.Query("limit")

	if queryLimit == "" {
		data.Limit = int32(25)
	} else {
		limit32, err := strconv.Atoi(queryLimit)

		if err != nil {
			return data, err
		}

		data.Limit = int32(limit32)
	}

	data.ActualLimit = data.Limit + 1

	if queryCursor == "" {
		data.Cursor = time.Now().Format(time.RFC3339)
	} else {
		data.Cursor, err = DecodeString(queryCursor)

		if err != nil {
			return data, err
		}
	}

	return data, nil
}
