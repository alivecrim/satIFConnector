package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lo struct {
	Name  string  `bson:"name" json:"name"`
	Value float64 `bson:"value" json:"value"`
}

type CalibrationSign struct {
	CalId primitive.ObjectID `bson:"calId" json:"calId,omitempty" `
	Sign  string             `bson:"sign" json:"sign,omitempty"`
}

type GroundConnector struct {
	Antenna       string             `json:"antenna,omitempty" bson:"antenna"`
	Polarization  string             `json:"polarization,omitempty" bson:"polarization"`
	Direction     string             `json:"direction,omitempty" bson:"direction"`
	Band          string             `json:"band,omitempty" bson:"band"`
	Name          string             `json:"name,omitempty" bson:"name"`
	Description   string             `json:"description,omitempty" bson:"description"`
	Calibrations  []CalibrationSign  `json:"calibrations,omitempty" bson:"calibrations"`
	ConnectTo     primitive.ObjectID `json:"connectTo,omitempty" bson:"connectTo"`
	StationConfig string             `json:"stationConfig,omitempty" bson:"stationConfig"`
}

type GroundStation struct {
	Id               primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	ProjectId        primitive.ObjectID   `bson:"project_id" json:"projectId"`
	SatelliteId      primitive.ObjectID   `bson:"satellite_id" json:"satelliteId"`
	Name             string               `bson:"name" json:"name"`
	Description      string               `bson:"description" json:"description"`
	CalibrationIds   []primitive.ObjectID `bson:"calibration_ids" json:"calibrationIds"`
	Los              []Lo                 `bson:"los" json:"los"`
	ConfigTemplate   string               `bson:"config_template" json:"configTemplate"`
	GroundInterfaces []GroundConnector    `bson:"ground_interfaces" json:"groundInterfaces"`
}

type Calibration struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	GroundStationId primitive.ObjectID `bson:"project_id" json:"projectId"`
	Name            string             `bson:"name" json:"name"`
	Data            []DataPoint        `bson:"data" json:"data"`
}

type DataPoint struct {
	X float64 `bson:"x" json:"x"`
	Y float64 `bson:"y" json:"y"`
}
