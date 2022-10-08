package main

import (
	"embed"
)

//go:embed dist/static/*
var content embed.FS

var fileServerPort string = ":53172"
var ipcPort string = ":53174"
