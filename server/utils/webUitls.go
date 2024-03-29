package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrBadRequest = errors.New("infoplus: Bad request")

func JsonDecode(url string) (map[string]interface{}, error) {
	response, err := http.Get(url)
	if response.StatusCode >= 400 {
		return nil, ErrBadRequest
	}
	if err != nil {
		return nil, err
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	resultJson := make(map[string]interface{})
	json.Unmarshal(responseData, &resultJson)

	if err != nil {
		return nil, err
	}
	return resultJson, nil
}

func GetQueryOrDefault(c *gin.Context, urlParam string, defaultValue string) string {
	queryList, ok := c.Request.URL.Query()[urlParam]
	query := defaultValue
	if ok && len(queryList) > 0 {
		query = queryList[0]
	}
	return query
}
