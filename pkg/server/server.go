package server

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/lbrictson/status/web"

	"github.com/labstack/echo/v4"
	"github.com/lbrictson/status/pkg/configuration"
)

// Server holds all the deps of our server instance
type Server struct {
}

// Run starts a blocking webserver
func Run(config *configuration.Config) {
	server := Server{}
	embeddedFiles := web.Assets
	fSys, err := fs.Sub(embeddedFiles, "static")
	if err != nil {
		panic(err)
	}
	// Serve static files from our virtual file system 'static' directory
	assetHandler := http.FileServer(http.FS(fSys))
	s := echo.New()
	t := Template{templates: template.Must(template.ParseFS(web.Assets, "templates/*.tmpl"))}
	s.Renderer = &t
	s.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))
	s.GET("/", server.indexView)
	s.Logger.Fatal(s.Start(fmt.Sprintf(":%v", config.WebPort)))
}
