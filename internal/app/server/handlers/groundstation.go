package handlers

import "net/http"

const CollectionName = "groundstation"

//type DevType struct {
//	Id primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
//	V  string             `bson:"v" json:"v"`
//}

//func GroundStationUpdate() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		ctx := context.Background()
//		opts := options.Find().SetSort(bson.D{{"v", 1}})
//
//		cursor, err := store.Db().Collection("devType").Find(context.TODO(), bson.D{}, opts)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fullResult := make([]DevType, 0)
//		for cursor.Next(ctx) {
//			cur := DevType{}
//			err = cursor.Decode(&cur)
//			if err != nil {
//				return
//			}
//			fullResult = append(fullResult, cur)
//		}
//		marshal, err := json.Marshal(fullResult)
//		if err != nil {
//			return
//		}
//		_, err = w.Write(marshal)
//		if err != nil {
//			return
//		}
//	}
//}
//func DeviceEntityHandlerGet() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		ctx := context.Background()
//		opts := options.Find().SetSort(bson.D{{"name", 1}})
//		cursor, err := store.Db().Collection(CollectionName).Find(context.TODO(), bson.D{}, opts)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fullResult := make([]entity.Device, 0)
//		for cursor.Next(ctx) {
//			cur := entity.Device{}
//			err = cursor.Decode(&cur)
//			if err != nil {
//				return
//			}
//			fullResult = append(fullResult, cur)
//		}
//		marshal, err := json.Marshal(fullResult)
//		if err != nil {
//			return
//		}
//		_, err = w.Write(marshal)
//		if err != nil {
//			return
//		}
//	}
//}

//func DeviceEntityHandlerPost() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		p := entity.Device{}
//		err := json.NewDecoder(r.Body).Decode(&p)
//		new := bson.D{{"name", p.Name}}
//		result, err := store.Db().Collection(CollectionName).InsertOne(context.TODO(), new)
//		if err != nil {
//			panic(err)
//		}
//		marshal, err := json.Marshal(result)
//		if err != nil {
//			return
//		}
//		_, err = w.Write(marshal)
//		if err != nil {
//			return
//		}
//
//	}
//}
//
//type OperationResult struct {
//	Result string `json:"result"`
//	Err    int    `json:"err"`
//}

/*
 История
*/
func GroundStationGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
