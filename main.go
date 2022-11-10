package main

import (
	"golang.org/x/tools/go/packages"
)

func main() {
	conf := packages.Config{
		Dir:  "/home/omer/repos/assisted-installer-agent/src/container_image_availability",
		Mode: packages.LoadSyntax,
	}

	_, err := packages.Load(&conf, "file=./container_image_availability.go")

	if err != nil {
		panic(err)
	}
}
