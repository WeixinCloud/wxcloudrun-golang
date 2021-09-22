package main

import (
	"fmt"
	"log"
	"net/http"

	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/adduser", service.AddUser)
	http.HandleFunc("/deleteuser", service.DeleteUser)
	http.HandleFunc("/updateuser", service.UpdateUser)
	http.HandleFunc("/queryuser", service.QueryUser)

	log.Fatal(http.ListenAndServe(":80", nil))
}
