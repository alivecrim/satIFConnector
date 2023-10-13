package project

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"pasha.com/satifconnector/internal/app/state"
	"pasha.com/satifconnector/internal/app/store"
	"pasha.com/satifconnector/internal/app/store/entity"
)

const CollectionName = "project"

func FindAll() ([]entity.Project, error) {
	ctx := context.TODO()
	filter := bson.D{}
	cursor, err := store.Db().Collection(CollectionName).Find(ctx, filter)
	if err != nil {
		log.Printf("error %v, failed to find etities in collection %v", err, CollectionName)
		return nil, err
	}
	var result []entity.Project
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Create(p entity.Project) (interface{}, error) {
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

func SetActive(id interface{}) error {
	ctx := context.TODO()
	_, err := store.Db().Collection(CollectionName).UpdateMany(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"active": true}})
	if err != nil {
		return err
	}
	_, err = store.Db().Collection(CollectionName).UpdateMany(ctx, bson.M{"_id": bson.M{"$ne": id}}, bson.M{"$set": bson.M{"active": false}})
	if err != nil {
		return err
	}
	relationState := state.State()
	objectID := id.(primitive.ObjectID)
	relationState.ActiveProjectId = &objectID
	return nil
}
