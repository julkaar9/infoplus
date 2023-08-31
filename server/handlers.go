package server

import (
	"fmt"
	"infoplus/server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func onFailedToFetch(c *gin.Context) {
	err500 := Badge{"stackoverflow", "error"}
	c.HTML(http.StatusOK, "badge.html", gin.H{"Title": err500.Title, "Value": err500.Value})
}

func StackoverflowReputationhandler(c *gin.Context) {
	userIds, ok := c.Request.URL.Query()["userid"]
	userId := "8522463"
	if ok && len(userIds) > 0 {
		userId = userIds[0]
	}
	resp, err := utils.JsonDecode(fmt.Sprintf("https://api.stackexchange.com/2.3/users/%s?&site=stackoverflow", userId))

	if err != nil || len(resp["items"].([]interface{})) == 0 {
		onFailedToFetch(c)
		return
	}
	val := resp["items"].([]interface{})[0].(map[string]interface{})["reputation"].(float64)
	valStr := utils.HumanReadable(val)

	badge := Badge{"stackoverflow", valStr}
	c.HTML(http.StatusOK, "badge.html", gin.H{"Title": badge.Title, "Value": badge.Value})
	//c.IndentedJSON(http.StatusOK, badge_obj)
}
