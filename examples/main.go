package main

import (
	"fmt"
	"os"

	"github.com/luizvnasc/goenv"
)

type Project struct {
	Name    string `env:"NAME"`
	Version string `env:"VERSION"`
}

func main() {
	os.Setenv("NAME", "goenv")
	os.Setenv("VERSION", "1.0.0")

	var project Project
	goenv.Unmarshal(&project)

	fmt.Printf("Project: %v\n", project)

}
