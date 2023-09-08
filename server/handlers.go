package server

import (
	"errors"
	"fmt"
	bm "infoplus/server/services/badgemaker"
	"infoplus/server/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrEmptyRequest = errors.New("items is empty")

func StackoverflowReputationhandler(c *gin.Context) {
	badge, badeTheme, err := stackExchangeBaseHandler(c, "stackoverflow", "8522463", "555555", "97ca00")
	badgeHandler(c, "stackoverflow", badge, badeTheme, err)
}

func stackExchangeBaseHandler(c *gin.Context, site string, userid string, defaultlabelColor string, defaultMessageColor string) (*bm.Badge, *bm.BadgeTheme, error) {
	userId := utils.GetQueryOrDefault(c, "userid", "8522463")

	resp, err := utils.JsonDecode(fmt.Sprintf("https://api.stackexchange.com/2.3/users/%s?&site=%s", userId, site))

	if err != nil {
		return nil, nil, err
	}
	if len(resp["items"].([]interface{})) == 0 {
		return nil, nil, ErrEmptyRequest
	}
	val := resp["items"].([]interface{})[0].(map[string]interface{})["reputation"].(float64)
	valStr := utils.HumanReadable(val)

	badgeTheme := bm.CreateBadgeTheme(c)
	badge := badgeTheme.CreateBadgeWithTheme(site, valStr)
	return badge, badgeTheme, nil
}

func setBadgeTemplate(badge *bm.Badge, badgeTheme *bm.BadgeTheme) *gin.H {
	return &gin.H{
		"Label":        badge.Label,
		"Message":      badge.Message,
		"TotalWidth":   badge.TotalWidth,
		"LabelWidth":   badge.LabelWidth,
		"MessageWidth": badge.MessageWidth,
		"LabelX":       badge.LabelX,
		"MessageX":     badge.MessageX,

		"LabelColor":   badgeTheme.LabelColor,
		"MessageColor": badgeTheme.MessageColor,
	}
}

func onFailedToFetch(c *gin.Context, site string) {
	badgeTheme := bm.CreateBadgeTheme(c)
	err500 := badgeTheme.CreateBadgeWithTheme(site, "error")
	c.HTML(http.StatusOK, badgeTheme.Style, setBadgeTemplate(err500, badgeTheme))
}
func badgeHandler(c *gin.Context, site string, badge *bm.Badge, badgeTheme *bm.BadgeTheme, err error) {
	if err != nil {
		log.Println(err)
		onFailedToFetch(c, site)
		return
	}
	c.HTML(
		http.StatusOK,
		badgeTheme.Style,
		setBadgeTemplate(badge, badgeTheme))
}
