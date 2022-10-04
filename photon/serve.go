package photon

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//serve embedded static frontend files
func Serve(content embed.FS, dir string, fileServerPort string) {
	router := gin.Default()

	static, err := fs.Sub(content, dir)
	if err != nil {
		fmt.Println(err)
	}
	router.StaticFS("/", http.FS(static))

	router.Run(fileServerPort)
}
