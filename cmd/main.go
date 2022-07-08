package main

import (
	"log"
	server "social/pkg/server"
)

func main() {
	// store, err := sqlite.CreateDB()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println("Database successfully inited")

	serv := server.NewServer()
	log.Fatal(serv.Start(":8080"))
}
