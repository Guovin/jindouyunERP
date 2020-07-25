// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package orders

import (
	"database/sql"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table custormer_orders.
type Entity struct {
	Id              uint        `orm:"id,primary"       json:"id"`               //
	CreatedAt       *gtime.Time `orm:"created_at"       json:"created_at"`       //
	UpdatedAt       *gtime.Time `orm:"updated_at"       json:"updated_at"`       //
	DeletedAt       *gtime.Time `orm:"deleted_at"       json:"deleted_at"`       //
	Operator        string      `orm:"operator"         json:"operator"`         //
	Name            string      `orm:"name"             json:"name"`             //
	Tel             string      `orm:"tel"              json:"tel"`              //
	DeliveryAddress string      `orm:"delivery_address" json:"delivery_address"` //
	DeliveryTime    string      `orm:"delivery_time"    json:"delivery_time"`    //
	Amount          float64     `orm:"amount"           json:"amount"`           //
	Deposit         float64     `orm:"deposit"          json:"deposit"`          //
	Remarks         string      `orm:"remarks"          json:"remarks"`          //
	State           string      `orm:"state"            json:"state"`            //
	Freight         float64     `orm:"freight"          json:"freight"`          //
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}
