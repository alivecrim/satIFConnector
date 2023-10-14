package glob

import (
	"encoding/json"
	"log"
	"net/http"
	"pasha.com/satifconnector/internal/app/state"
	"pasha.com/satifconnector/internal/app/store/entity"
	"pasha.com/satifconnector/internal/app/store/groundStation"
	"pasha.com/satifconnector/internal/app/store/project"
	"pasha.com/satifconnector/internal/app/store/satellite"
)

/*
	История
	открываем страницу с настройкой связи между спутником и земной станцией
	делаем запрос в базу на предмет наличия текущего активного проекта

	если проект не найден - предлагаем создать
	если проект найден, но активных не назначено - предлагаем назначить
	если проект найден и есть активный - говорим софту, как его звать

	далее для активного проекта ищем земные станции - найдена одна - отображаем ее
	если не найдена - предлагаем создать - отображаем
	если найдно несколько - предлагаем выбор - отображаем

	после отображаения станции ищем связаный спутник
	если спутник найден - отображаем спутник
	если не найден - предлагаем создать - связываем и отображаем

	далее настраиваем отдельные компоненты
*/

type InitRequestResult struct {
	ActiveProject         *entity.Project        `json:"activeProject,omitempty"`
	Projects              []entity.Project       `json:"projects"`
	RelatedGroundStations []entity.GroundStation `json:"relatedGroundStations"`
	RelatedSatellite      *entity.Satellite      `json:"relatedSatellite,omitempty"`
	Satellites            []entity.Satellite     `json:"satellites"`
	GroundStations        []entity.GroundStation `json:"groundStations"`
}

func RequestForCurrentRelationState() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := InitRequestResult{
			Projects:              []entity.Project{},
			ActiveProject:         nil,
			GroundStations:        []entity.GroundStation{},
			RelatedGroundStations: []entity.GroundStation{},
			Satellites:            []entity.Satellite{},
			RelatedSatellite:      nil,
		}

		projectResult, err := project.FindAll()
		if err != nil {
			return
		}
		resp.Projects = projectResult
		// Нет проектов вообще
		if len(projectResult) == 0 {
			err = writeResponse(w, &resp)
			if err != nil {
				panic(err)
			}
			return
		}
		// проверка на наличие активных проектов
		for _, p := range projectResult {
			if p.Active {
				resp.ActiveProject = &p
				state.State().ActiveProjectId = &p.Id
				break
			}
		}
		if resp.ActiveProject == nil {
			err = writeResponse(w, &resp)
			if err != nil {
				panic(err)
			}
			return
		}

		// проверка наличия земных станций в принципе
		stations, err := groundStation.FindAll()
		if err != nil {
			return
		}
		if len(stations) == 0 {
			err = writeResponse(w, &resp)
			if err != nil {
				return
			}
			return
		}
		resp.GroundStations = stations
		for _, station := range stations {
			if station.ProjectId == resp.ActiveProject.Id {
				resp.RelatedGroundStations = append(resp.RelatedGroundStations, station)
			}
		}

		// проверка наличия спутников
		satellites, err := satellite.FindAll()
		if err != nil {
			return
		}
		if len(satellites) == 0 {
			err = writeResponse(w, &resp)
			if err != nil {
				return
			}
			return
		}
		resp.Satellites = satellites
		for _, sat := range satellites {
			if sat.ProjectId == resp.ActiveProject.Id {
				resp.RelatedSatellite = &sat
				break
			}
		}
		err = writeResponse(w, &resp)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func writeResponse(writer http.ResponseWriter, resp *InitRequestResult) error {
	marshal, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = writer.Write(marshal)
	if err != nil {
		return err
	}
	return nil
}
