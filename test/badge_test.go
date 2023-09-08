package infoplus_tests

import (
	bm "infoplus/server/services/badgemaker"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateBadgeTheme(t *testing.T) {
	badgeThemeTests := []struct {
		urlParams map[string]string
		want      *bm.BadgeTheme
	}{
		{
			map[string]string{"style": "plastic", "labelcolor": bm.DefaultGrey, "messagecolor": bm.DefaultGreen},
			&bm.BadgeTheme{Style: "plastic.html", LabelColor: bm.DefaultGrey, MessageColor: bm.DefaultGreen},
		},
		{
			map[string]string{"style": "for_the_badge", "labelcolor": bm.DefaultGrey},
			&bm.BadgeTheme{Style: "for_the_badge.html", LabelColor: bm.DefaultGrey, MessageColor: bm.DefaultGreen},
		},
		{
			map[string]string{},
			&bm.BadgeTheme{Style: "flat.html", LabelColor: bm.DefaultGrey, MessageColor: bm.DefaultGreen},
		},
	}

	const url = "/test/createBadgeTheme"
	for _, tt := range badgeThemeTests {
		r, err := GetGinRouter()
		if err != nil {
			t.Fail()
		}
		r.GET(url, func(c *gin.Context) {
			got := bm.CreateBadgeTheme(c)
			assert.Equal(t, tt.want, got)

		})
		w, req := GetHttpRecorder(url, tt.urlParams)

		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	}
}

func TestCreateBadgeWithTheme(t *testing.T) {
	badgeTests := []struct {
		label      string
		message    string
		badgeTheme *bm.BadgeTheme
		want       *bm.Badge
	}{
		{"len4", "length7",
			&bm.BadgeTheme{Style: "flat.html", LabelColor: bm.DefaultGrey, MessageColor: bm.DefaultGreen},
			&bm.Badge{Label: "len4", Message: "length7", TotalWidth: 92, LabelWidth: 37, MessageWidth: 55, LabelX: 19.5, MessageX: 63.5},
		},
		{"len_5", "length___11",
			&bm.BadgeTheme{Style: "for_the_badge.html", LabelColor: bm.DefaultGrey, MessageColor: bm.DefaultGreen},
			&bm.Badge{Label: "LEN_5", Message: "LENGTH___11", TotalWidth: 141, LabelWidth: 48, MessageWidth: 93, LabelX: 25, MessageX: 93.5},
		},
	}

	for _, tt := range badgeTests {
		got := tt.badgeTheme.CreateBadgeWithTheme(tt.label, tt.message)
		assert.Equal(t, tt.want, got)
	}
}
