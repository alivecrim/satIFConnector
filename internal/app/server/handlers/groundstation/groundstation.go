package groundstation_http

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"pasha.com/satifconnector/internal/app/state"
	"pasha.com/satifconnector/internal/app/store/entity"
	"pasha.com/satifconnector/internal/app/store/groundStation"
)

func FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		all, err := groundStation.FindAll()
		if err != nil {
			return
		}
		marshal, err := json.Marshal(all)
		if err != nil {
			return
		}
		_, err = w.Write(marshal)
		if err != nil {
			return
		}
	}
}

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var gs entity.GroundStation
		err := json.NewDecoder(r.Body).Decode(&gs)
		gs.ProjectId = *(state.State().ActiveProjectId)
		if err != nil {
			return
		}
		create, err := groundStation.Create(gs)
		if err != nil {
			return
		}
		gs.Id = create.(primitive.ObjectID)
		marshal, err := json.Marshal(gs)
		if err != nil {
			return
		}
		_, err = w.Write(marshal)
		if err != nil {
			return
		}

	}
}
