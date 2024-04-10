package webapp

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed frontend-resources/*
var frontendStatic embed.FS

type WebAppHandler struct {
	webAppDir fs.FS
	logger    *log.Logger
}

func NewWebAppHandler(logger *log.Logger) (*WebAppHandler, error) {
	webAppDir, err := getWebAppAssets()
	if err != nil {
		return nil, fmt.Errorf("getting web app assets: %w", err)
	}

	return &WebAppHandler{
		webAppDir: webAppDir,
		logger:    logger,
	}, nil
}

func (h *WebAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(h.webAppDir)).ServeHTTP(w, r)
}

func getWebAppAssets() (fs.FS, error) {
	if err := isFrontendFolderValid(frontendStatic); err != nil {
		return embed.FS{}, fmt.Errorf("checking frontend static folder: %w", err)
	}

	webDir, err := fs.Sub(frontendStatic, "frontend-resources")
	if err != nil {
		return embed.FS{}, fmt.Errorf("getting frontend-resources folder from embeded front app: %w", err)
	}

	return webDir, nil
}

func isFrontendFolderValid(frontendFs embed.FS) error {
	if _, err := fs.Stat(frontendFs, "frontend-resources/index.html"); err != nil {
		return fmt.Errorf("frontendStatic does not contain index.html: %w", err)
	}

	return nil
}
