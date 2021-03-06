// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package staffs

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Fill with you ideas below.

// Staff 员工表
type Staff struct {
	ID      string `gorm:"primary_key"`
	Name    string
	Age     string
	Address string
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
	db.AutoMigrate(&Staff{})
}

//查询员工信息
func SelectStaff() (staffs []Staff, err error) {
	err = db.Table("staffs").Select("id,name,age,address").Find(&staffs).Error
	if err != nil {
		return
	}
	return
}

//添加或修改员工信息
func UpdateStaff(ID string, Name string, Age string, Address string) error {
	return db.Table("staffs").Save(Staff{ID: ID, Name: Name, Age: Age, Address: Address}).Error
}

//删除员工信息
func DeleteStaff(ID string) error {
	//删除单条记录，批量则使用Where
	return db.Table("staffs").Delete(Staff{ID: ID}).Error
}

//保存全部员工信息
func SaveStaff(staffs []Staff) error {
	for _, v := range staffs {
		err := db.Table("staffs").Save(&v).Error
		if err != nil {
			return err
		}
	}
	return nil
}
