package satellite

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"pasha.com/satifconnector/internal/app/store"
	"pasha.com/satifconnector/internal/app/store/entity"
)

const CollectionName = "satellite"

func FindAll() ([]entity.Satellite, error) {
	ctx := context.TODO()
	filter := bson.D{}
	cursor, err := store.Db().Collection(CollectionName).Find(ctx, filter)
	if err != nil {
		log.Printf("error %v, failed to find etities in collection %v", err, CollectionName)
		return nil, err
	}
	var result []entity.Satellite
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Create(p entity.Satellite) (interface{}, error) {
	ctx := context.TODO()
	insertedResult, err := store.Db().Collection(CollectionName).InsertOne(ctx, p)
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
