package controller

import (
	"encoding/json"
	"fmt"
	"lordtt13/GO_Prods/TODO-API/model"
	"lordtt13/GO_Prods/TODO-API/views"
	"net/http"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}
