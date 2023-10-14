package project_http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"pasha.com/satifconnector/internal/app/store/entity"
	"pasha.com/satifconnector/internal/app/store/project"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p = entity.Project{}
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = project.Create(p)
		if err != nil {
			log.Println(err)
		}
	}
}

func SetActive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			fmt.Println("error parse id")
			return
		}
		hex, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
		err = project.SetActive(hex)
		if err != nil {
			return
		}
	}
}
