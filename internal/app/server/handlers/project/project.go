package project_http

import (
	"encoding/json"
	"log"
	"net/http"
	"pasha.com/satifconnector/internal/app/store/entity"
	"pasha.com/satifconnector/internal/app/store/project"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
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
