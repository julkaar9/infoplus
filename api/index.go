package handler

import (
	"html/template"
	"infoplus/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

var parsedTemplates *template.Template

func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	var err error
	parsedTemplates, err = server.GetTemplates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	g.Use(server.SVGMiddleware())
	g.SetHTMLTemplate(parsedTemplates)

	// Routes
	g.GET("/api/stackoverflow", server.StackoverflowReputationhandler)

	g.ServeHTTP(w, r)
}
