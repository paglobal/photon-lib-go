package photon

import (
	"embed"
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize(content embed.FS, dir string, start func(), fileServerPort string, ipcPort string) {
	//serve files only when not in dev mode
	args := os.Args[1:]
	if argsCount := len(args); argsCount <= 0 {
		gin.SetMode(gin.ReleaseMode)
		go Serve(content, dir, fileServerPort)
	}

	go start()
}
