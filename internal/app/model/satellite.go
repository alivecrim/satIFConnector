package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Satellite struct {
	Id          primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	ProjectId   *primitive.ObjectID `bson:"project_id" json:"projectId,omitempty"`
	Name        string              `bson:"name" json:"name,omitempty"`
	Description string              `bson:"description" json:"description"`
}

type SatelliteConnector struct {
	Id           primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	SatelliteId  *primitive.ObjectID `bson:"project_id" json:"projectId"`
	Name         string              `bson:"name" json:"name"`
	Description  string              `bson:"description" json:"description"`
	Antenna      string              `json:"antenna,omitempty" bson:"antenna"`
	Polarization string              `json:"polarization,omitempty" bson:"polarization"`
	Direction    string              `json:"direction,omitempty" bson:"direction"`
	Band         string              `json:"band,omitempty" bson:"band"`
}
