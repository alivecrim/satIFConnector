package groundStation

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"pasha.com/satifconnector/internal/app/store"
	"pasha.com/satifconnector/internal/app/store/entity"
)

const CollectionName = "groundstation"

func FindAll() ([]entity.GroundStation, error) {
	ctx := context.TODO()
	filter := bson.D{}
	cursor, err := store.Db().Collection(CollectionName).Find(ctx, filter)
	if err != nil {
		log.Printf("error %v, failed to find etities in collection %v", err, CollectionName)
		return nil, err
	}

	var groundStationResult []entity.GroundStation
	err = cursor.All(ctx, &groundStationResult)
	if err != nil {
		return nil, err
	}

	return groundStationResult, nil
}

func Create(gs entity.GroundStation) (interface{}, error) {
	ctx := context.TODO()
	insertedResult, err := store.Db().Collection(CollectionName).InsertOne(ctx, gs)
	if err != nil {
		return nil, err
	}
	return insertedResult.InsertedID, err
}

func Delete(id interface{}) (int64, error) {
	ctx := context.TODO()
	one, err := store.Db().Collection(CollectionName).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return 0, nil
	}
	return one.DeletedCount, nil
}

func Update(id interface{}, station entity.GroundStation) error {
	ctx := context.TODO()
	_, err := store.Db().Collection(CollectionName).UpdateByID(ctx, id,
		bson.M{"$set": bson.M{
			"project_id":        station.ProjectId,
			"satellite_id":      station.SatelliteId,
			"name":              station.Name,
			"description":       station.Description,
			"calibration_ids":   station.CalibrationIds,
			"los":               station.Los,
			"config_template":   station.ConfigTemplate,
			"ground_interfaces": station.GroundConnectors,
		}},
	)
	if err != nil {
		return err
	}
	return nil
}
