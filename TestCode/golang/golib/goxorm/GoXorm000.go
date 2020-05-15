package main

import (
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		log.Printf("fail: %s", err)
	}
}
