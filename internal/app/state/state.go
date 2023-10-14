package state

import "go.mongodb.org/mongo-driver/bson/primitive"

type CurrentRelationState struct {
	ActiveProjectId *primitive.ObjectID
}

var state *CurrentRelationState

func init() {
	state = &CurrentRelationState{ActiveProjectId: nil}
}

func State() *CurrentRelationState {
	return state
}
