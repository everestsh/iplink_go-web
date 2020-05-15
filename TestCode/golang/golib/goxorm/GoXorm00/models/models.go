package models

import (
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"`
}

var x *xorm.Engine

func init() {
	var err error
	x, err = xorm.NewEngine("sqlite3", "./bank.db")
	if err != nil {
		log.Fatalf("Fail to create engine %s", err)
	}

	if err = x.Sync(new(Account)); err != nil {
		log.Fatalf("Fail to sync database: %v", err)
	}
}

func NewAccount(name string, balance float64) error {
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err
}
