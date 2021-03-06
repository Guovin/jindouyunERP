// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package purchase_orders

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Fill with you ideas below.

// PurchaseOrder 采购订单表
type PurchaseOrder struct {
	gorm.Model
	Operator string `gorm:"size:255"`
	Remarks  string
	Amount   float32
	Freight  float32
	State    string `gorm:"default:'未完成'"`
}

var db *gorm.DB

// Init 初始化
func Init(link string) {
	var err error
	db, err = gorm.Open("mysql", link+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("failed to connect database, ", err)
	}
	db.DB().SetConnMaxLifetime(60 * time.Second)
	// db.LogMode(true)
	db.AutoMigrate(&PurchaseOrder{})
}
