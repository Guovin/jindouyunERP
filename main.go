package main

import (
	"jindouyunERP/app/model/commodities"
	"jindouyunERP/app/model/custormer_goods"
	"jindouyunERP/app/model/custormer_orders"
	"jindouyunERP/app/model/purchase_goods"
	"jindouyunERP/app/model/purchase_orders"
	"jindouyunERP/app/model/roles"
	"jindouyunERP/app/model/staffs"
	"jindouyunERP/app/model/users"
	_ "jindouyunERP/boot"
	_ "jindouyunERP/router"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/gogf/gf/frame/g"

	"github.com/BurntSushi/toml"
)

type configToml struct {
	DB database `toml:"database"`
}
type database struct {
	Link string
}

var db *gorm.DB

func main() {
	var config configToml
	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		log.Fatal(err)
		return
	}
	link := config.DB.Link
	users.Init(link)
	staffs.Init(link)
	roles.Init(link)
	purchase_orders.Init(link)
	purchase_goods.Init(link)
	custormer_orders.Init(link)
	custormer_goods.Init(link)
	commodities.Init(link)
	g.Server().Run()
}
