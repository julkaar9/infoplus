package infoplus_tests

import (
	"html/template"
	"infoplus/server"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

var parsedTemplates *template.Template

func getRouter(withTemplates bool) (*gin.Engine, error) {
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

func TestStackoverflowReputationhandler(t *testing.T) {
	r, err := getRouter(true)
	if err != nil {
		t.Fail()
	}

	r.GET("/api/stackoverflow", server.StackoverflowReputationhandler)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/stackoverflow", nil)
	q := req.URL.Query()
	q.Add("userid", "1111111")
	req.URL.RawQuery = q.Encode()

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	want := `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="110" height="20" id="svg1">
	<linearGradient id="b" x2="0" y2="100%">
	  <stop offset="0" stop-color="#bbb" stop-opacity=".1" />
	  <stop offset="1" stop-opacity=".1" />
	</linearGradient>
	<clipPath id="a">
	  <rect width="110" height="20" rx="3" fill="#fff" />
	</clipPath>
	<g clip-path="url(#a)">
	  <path fill="#555" d="M0 0h75v20H0z" />
	  <path fill="#5eba7d" d="M75 0h35v20H75z" />
	  <path fill="url(#b)" d="M0 0h110v20H0z" />
	</g>
	<g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="110">
	  <text id="t1" x="385" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="650">
		stackoverflow
	  </text>
	  <text id="t1" x="385" y="140" transform="scale(.1)" textLength="650">
	  stackoverflow
	  </text>
	  <text id="t2" x="915" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)">
		error
	  </text>
	  <text id="t2" x="915" y="140" transform="scale(.1)">
		error
	  </text>
	</g>
  </svg>`
	assert.Equal(t, standardizeSpaces(want), standardizeSpaces(w.Body.String()))
}
