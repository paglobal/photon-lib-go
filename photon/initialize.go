package photon

import (
	"embed"
	"os"
)

func Initialize(content embed.FS, dir string, start func()) {
	//serve files only when not in dev mode
	args := os.Args[1:]
	if argsCount := len(args); argsCount <= 0 {
		go Serve(content, dir)
	}

	go start()
}
