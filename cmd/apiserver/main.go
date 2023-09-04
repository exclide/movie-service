package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/exclide/movie-service/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}

	serv := apiserver.NewServer(config)

	if err := serv.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello world")

}
