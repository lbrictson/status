package server

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/lbrictson/status/pkg"

	"github.com/lbrictson/status/web"

	"github.com/labstack/echo/v4"
)

// Server holds all the deps of our server instance
type Server struct {
	store pkg.StorageBackend
}

type NewServerConfig struct {
	Port  int
	Store pkg.StorageBackend
}

// Run starts a blocking webserver
func Run(config NewServerConfig) {
	server := Server{
		store: config.Store,
	}
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
	s.GET("/login", server.loginView)
	s.POST("/login", server.loginForm)
	s.Logger.Fatal(s.Start(fmt.Sprintf(":%v", config.Port)))
}
