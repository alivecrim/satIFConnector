package groundStation

import (
	"context"
	"fmt"
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

func TestCreateAndDelete(t *testing.T) {
	station := entity.GroundStation{
		Id:               primitive.ObjectID{},
		ProjectId:        primitive.ObjectID{},
		SatelliteId:      primitive.ObjectID{},
		Name:             "Somename",
		Description:      "some description",
		CalibrationIds:   nil,
		Los:              nil,
		ConfigTemplate:   "some template",
		GroundConnectors: nil,
	}
	id, err := Create(station)
	if err != nil {
		return
	}
	count, err := Delete(id)
	if err != nil {
		return
	}
	fmt.Println(count)
}

func TestFindAll(t *testing.T) {
	station := entity.GroundStation{
		Id:               primitive.ObjectID{},
		ProjectId:        primitive.ObjectID{},
		SatelliteId:      primitive.ObjectID{},
		Name:             "Somename",
		Description:      "some description",
		CalibrationIds:   nil,
		Los:              nil,
		ConfigTemplate:   "some template",
		GroundConnectors: nil,
	}
	for i := 0; i < 20; i++ {
		_, err := Create(station)
		if err != nil {
			log.Println(err)
			t.Fail()
		}
	}

	all, err := FindAll()
	if err != nil {
		return
	}
	if len(all) != 20 {
		log.Printf("all_items_len!=20")
		t.Fail()
	}
}
func TestUpdate(t *testing.T) {
	station := entity.GroundStation{
		Id:          primitive.ObjectID{},
		ProjectId:   primitive.ObjectID{},
		SatelliteId: primitive.ObjectID{},
		Name:        "Somename",
		Description: "some description",
		CalibrationIds: []primitive.ObjectID{
			primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID(),
		},
		Los: []entity.Lo{
			{
				Name:  "Some Lo1",
				Value: 1,
			},
			{
				Name:  "Some Lo2",
				Value: 2,
			},
			{
				Name:  "Some Lo3",
				Value: 3,
			},
		},
		ConfigTemplate: "some template",
		GroundConnectors: []entity.GroundConnector{
			{
				Antenna:      "An1",
				Polarization: "RHCP",
				Direction:    "UP",
				Band:         "C",
				Name:         "An1-RHCP-UP-C",
				Description:  "Some description",
				CalibrationsWithSign: []entity.CalibrationWithSign{
					{
						CalId: primitive.NewObjectID(),
						Sign:  "+",
					},
					{
						CalId: primitive.NewObjectID(),
						Sign:  "-",
					},
					{
						CalId: primitive.NewObjectID(),
						Sign:  "+",
					},
				},
				ConnectTo:     primitive.NewObjectID(),
				StationConfig: "вsadsaвыadasdsaвыфdsdsadsaвыфв",
			},
		},
	}
	Create(station)

}
