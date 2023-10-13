package glob

import (
	"encoding/json"
	"net/http"
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
	HasProjects           bool                   `json:"hasProjects"`
	HasActiveProject      bool                   `json:"hasActiveProject"`
	Projects              []entity.Project       `json:"projects"`
	RelatedGroundStations []entity.GroundStation `json:"relatedGroundStations"`
	RelatedSatellite      entity.Satellite       `json:"relatedSatellite"`
	Satellites            []entity.Satellite     `json:"satellites"`
	ActiveProject         entity.Project         `json:"activeProject"`
	GroundStations        []entity.GroundStation `json:"groundStations"`
}

func RequestForCurrentRelationState() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := InitRequestResult{
			HasProjects:           false,
			Projects:              []entity.Project{},
			HasActiveProject:      false,
			ActiveProject:         entity.Project{},
			GroundStations:        []entity.GroundStation{},
			RelatedGroundStations: []entity.GroundStation{},
			Satellites:            []entity.Satellite{},
			RelatedSatellite:      entity.Satellite{},
		}

		projectResult, err := project.FindAll()
		if err != nil {
			return
		}
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
				resp.HasActiveProject = true
				resp.ActiveProject = p
				break
			}
		}
		if !resp.HasActiveProject {
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
		}
		resp.GroundStations = stations
		for _, station := range stations {
			if station.ProjectId == resp.ActiveProject.Id {
				resp.RelatedGroundStations = append(resp.RelatedGroundStations)
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
		}
		resp.Satellites = satellites
		for _, sat := range satellites {
			if sat.ProjectId == resp.ActiveProject.Id {
				resp.RelatedSatellite = append(resp.RelatedGroundStations)
			}
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
	if err != nil {
		return err
	}
	return err
}
