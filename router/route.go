package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	studentController "go-restapi/controller"
	"go-restapi/helper"
	"net/http"
)

type Respons struct {
	Message string `json:"message"`
}

func Routes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			result, err := studentController.Index()
			helper.BasicHandler(err)

			jsonData, err := json.Marshal(result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(jsonData)
			return
		}

		http.Error(w, "", http.StatusBadRequest)
	})

	http.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			id := r.FormValue("id")
			result, err := studentController.Show(id)

			if err != nil {
				if err == sql.ErrNoRows {
					message := fmt.Sprintf("Data not found for id %s", id)
					responMessage := Respons{message}

					jsonData, err := json.Marshal(responMessage)
					helper.BasicHandler(err)

					w.Write(jsonData)
					http.Error(w, "", http.StatusNotFound)
				} else {
					http.Error(w, "", http.StatusInternalServerError)
				}

				return
			}

			jsonData, err := json.Marshal(result)
			helper.BasicHandler(err)

			w.Write(jsonData)
			return
		}

		http.Error(w, "", http.StatusBadRequest)
	})
}
