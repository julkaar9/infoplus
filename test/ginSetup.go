package infoplus_tests

import (
	"html/template"
	"infoplus/server"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

var parsedTemplates *template.Template

func GetGinRouter() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	var err error
	parsedTemplates, err = server.GetTemplates()
	if err != nil {
		return nil, err
	}

	g.Use(server.SVGMiddleware())
	g.SetHTMLTemplate(parsedTemplates)
	return g, nil
}

func GetHttpRecorder(url string, urlParams map[string]string) (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	for key, value := range urlParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	return w, req

}
