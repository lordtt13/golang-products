package main

import (
	"fmt"
	"log"
	"lordtt13/GO_Prods/TODO-API/controller"
	"lordtt13/GO_Prods/TODO-API/model"
	"net/http"
	"github.com/go-sql-driver/mysql"
)

/*
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getgoing/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello world"))
	})
	mux.HandleFunc("/getgoing/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello world go"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
*/

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}