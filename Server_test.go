// ChatServer project main.go
//package main

//import (
//	"database/sql"
//	"fmt"
//)

//func main() {
//	db := sql.Open("mysql", "root/arvin@tcp(localhost:3306)/test?cha-utf-8")
//	fmt.Println("Hello World!  123123123")
//	fmt.Println("This's ok!!!")
//}
package FDServer

import "testing"

func TestServer(t *testing.T) {
	var s ChatServer
	s.ReadConf("./ChatServer.conf")

	s.InitServer()
}
