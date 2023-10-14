export interface CalibrationWithSign {
  calId: string
  sign: string
}

export interface GroundConnector {
  antenna: string//       string                `json:"antenna,omitempty" bson:"antenna"`
  polarization: string//     string                `json:"polarization,omitempty" bson:"polarization"`
  direction: string//       string                `json:"direction,omitempty" bson:"direction"`
  band: string//      string                `json:"band,omitempty" bson:"band"`
  name: string   //              string                `json:"name,omitempty" bson:"name"`
  description: string        //        `json:"description,omitempty" bson:"description"`
  calibrationsWithSign: CalibrationWithSign[] //// []CalibrationWithSign `json:"calibrations,omitempty" bson:"calibrations"`
  connectTo?: string//            primitive.ObjectID    `json:"connectTo,omitempty" bson:"connectTo"`
  stationConfig?: string//        string                `json:"stationConfig,omitempty" bson:"stationConfig"`
}

export interface Lo {
  name: string
  value: number
}

export interface GroundStation {
  id?: string
  projectId?: string
  satelliteId?: string
  name: string
  description: string
  calibrationIds: string[]
  los: Lo[]
  configTemplate?: string
  groundConnectors: GroundConnector[]
}
