package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type badge struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

func getJson(url string) (map[string]interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	resultJson := make(map[string]interface{})
	json.Unmarshal(responseData, &resultJson)
	if err != nil {
		log.Fatal(err)
	}
	return resultJson, nil
}
func onError(c *gin.Context) {
	err500 := map[string]string{"message": "failed to fetch", "status": "error"}
	c.IndentedJSON(http.StatusOK, err500)
}
func getStackovReps(c *gin.Context) {
	user_ids, ok := c.Request.URL.Query()["userid"]
	user_id := "8522463"
	if !ok && len(user_ids) > 0 {
		user_id = user_ids[0]
	}
	resp, err := getJson(fmt.Sprintf("https://api.stackexchange.com/2.3/users/%s?&site=stackoverflow", user_id))
	if err != nil {
		onError(c)
	}
	val := int64(resp["items"].([]interface{})[0].(map[string]interface{})["reputation"].(float64))
	val_str := strconv.FormatInt(int64(val), 10)

	badge_obj := badge{"stackoverflow", val_str}
	c.HTML(http.StatusOK, "badge.html", gin.H{"Title": badge_obj.Title, "Value": badge_obj.Value})
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/stackoverflow", getStackovReps)

	router.Run("localhost:8080")
}
