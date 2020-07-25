package main

import (
	"jindouyunERP/app/model/users"
	_ "jindouyunERP/boot"
	_ "jindouyunERP/router"
	"log"

	"github.com/gogf/gf/frame/g"

	"github.com/BurntSushi/toml"
)

type configToml struct {
	DB database `toml:"database"`
}
type database struct {
	Link string
}

func main() {
	var config configToml
	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		log.Fatal(err)
		return
	}
	link := config.DB.Link
	users.Init(link)
	g.Server().Run()
}
