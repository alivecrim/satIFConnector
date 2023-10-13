package satellite

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"pasha.com/satifconnector/internal/app/store"
	"pasha.com/satifconnector/internal/app/store/entity"
	"testing"
)

func TestMain(m *testing.M) {
	log.Printf("Start all tests")
	store.InitTestStore()
	m.Run()
	many, err := store.Db().Collection(CollectionName).DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("After all test was delete %v items", many)

}
func TestCreateFindAllDelete(t *testing.T) {
	p := entity.Satellite{
		Id:          primitive.ObjectID{},
		Name:        "Satellite 1",
		Description: "Description for project 1",
	}
	create1, err := Create(p)
	create2, err := Create(p)
	if err != nil {
		return
	}
	all, err := FindAll()
	if err != nil {
		return
	}
	if len(all) != 2 {
		t.Fail()
	}
	deleteCount, err := Delete(create1)
	if deleteCount != 1 {
		log.Printf("Первый делит не сработал")
		t.Fail()
	}
	deleteCount, err = Delete(create2)

	if deleteCount != 1 {
		log.Printf("Второй  делит не сработал")
		t.Fail()
	}
}

func handleErr(err error, t *testing.T) {
	if err != nil {
		t.Fail()
	}
}
