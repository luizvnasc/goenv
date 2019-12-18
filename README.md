[![GitHub license](https://img.shields.io/github/license/luizvnasc/goenv)](https://github.com/luizvnasc/goenv/blob/master/LICENSE)
[![Build Status](https://travis-ci.com/luizvnasc/goenv.svg?branch=master)](https://travis-ci.com/luizvnasc/goenv)
[![Go Report Card](https://goreportcard.com/badge/github.com/luizvnasc/goenv)](https://goreportcard.com/report/github.com/luizvnasc/goenv)
[![Coverage Status](https://coveralls.io/repos/github/luizvnasc/goenv/badge.svg?branch=master)](https://coveralls.io/github/luizvnasc/goenv?branch=master)

# Gonfig

A simple module to unmarshal environment variables to struct;

## Getting Started

To install this package run:

```sh
go get -u github.com/luizvnasc/goenv
```

### Example

```go
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

```

## Authors
* Luiz Augusto Volpi Nascimento - Initial work - [@luizvnasc](https://github.com/luizvnasc)

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/luizvnasc/goenv/blob/master/LICENSE) file for details
