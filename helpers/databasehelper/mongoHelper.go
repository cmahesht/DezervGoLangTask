package databasehelper

import (
	"DezervGoLangTask/helpers/confighelper"
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

/*Logger : Author - Mahesh Chinvar
Purpose : Mongo Connection initilisation once
*/

var dbInstance *mongo.Client
var dbName string
var sessionError error
var once sync.Once

/*InitMongo : Author - Mahesh Chinvar
Purpose :Mongo Connection initilisation
 Return :
	sessionError : session initilisation error
*/
func InitMongo() error {
	var sessionError error
	once.Do(func() {
		sessionError = NewMongoConnection(confighelper.GetConfig("MONGODSN")+confighelper.GetConfig("DBNAME"), confighelper.GetConfig("DBNAME"))
	})
	return sessionError
}

/*NewMongoConnection : Author - Mahesh Chinvar
Purpose : create new connection to mongo uri and return the db instance
Input :
	mongoURI : mongo DNS
	databaseName : mongo Database Name
Return :
	error : database error
*/
func NewMongoConnection(mongoURI, databaseName string) error {
	db, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	err = db.Connect(ctx)
	if err != nil {
		return err
	}
	err = db.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	dbName = databaseName
	dbInstance = db
	return nil
}

/*GetMongoInstance : Author - Mahesh Chinvar
Purpose : GetMongoInstance returns connected mongo db connection instance
Return :
	mongoInstnce : mongo.Client
	error : database error
*/
func GetMongoInstance() (*mongo.Client, error) {
	err := dbInstance.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return dbInstance, nil
}

/*Find : Author - Mahesh Chinvar
Purpose : Find will return data for selector or query
Input :
	collectionName : collection name
	selector : selector map for query
Return :
	Result : resultset
	error : database error
*/
func Find(collectionName string, selector map[string]interface{}) (*gjson.Result, error) {
	db, err := GetMongoInstance()
	if err != nil {
		return nil, err
	}

	collection := db.Database(dbName).Collection(collectionName)

	if id, ok := selector["_id"]; ok {
		objID, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}
		selector["_id"] = objID
	}

	findOptions := options.Find()                                // build a `findOptions`
	findOptions.SetSort(map[string]int{"creationTimestamp": -1}) // reverse order by `when`

	cur, err := collection.Find(context.Background(), selector, findOptions)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var results []interface{}
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	ba, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}
	rs := gjson.ParseBytes(ba)
	return &rs, nil
}

/*FindByID : Author - Mahesh Chinvar
Purpose : FindByID return results by Object ID
Input :
	collectionName : collection name
	ID : selector ID
Return :
	Result : resultset
	error : database error
*/
func FindByID(collectionName, ID string) (*gjson.Result, error) {
	db, err := GetMongoInstance()
	if err != nil {
		return nil, err
	}

	collection := db.Database(dbName).Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	res := collection.FindOne(context.Background(), bson.M{"_id": objID})

	var result bson.M
	err = res.Decode(&result)
	if err != nil {
		return nil, err
	}

	ba, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	rs := gjson.ParseBytes(ba)
	return &rs, nil
}

/*Update : Author - Mahesh Chinvar
Purpose : Update updates the document by selector query
Input :
	collectionName : collection name
	ID : selector ID
	dataQuery : update query
Return :
	error : database error
*/
func Update(collectionName string, selector map[string]interface{}, dataQuery map[string]interface{}) error {
	db, err := GetMongoInstance()
	if err != nil {
		return err
	}

	/* 	if id, ok := selector["_id"]; ok {
	   		objID, err := primitive.ObjectIDFromHex(id.(string))
	   		if err != nil {
	   			return err
	   		}
	   		selector["_id"] = objID
	   	}
	*/
	collection := db.Database(dbName).Collection(collectionName)
	_, updateError := collection.UpdateOne(context.Background(), selector, dataQuery)
	if updateError != nil {
		return updateError
	}
	return nil
}

/*UpdateByID : Author - Mahesh Chinvar
Purpose : updates the document by Object ID
Input :
	collectionName : collection name
	ID : selector ID
	dataQuery : update query
Return :
	error : database error
*/
func UpdateByID(collectionName string, ID string, dataQuery map[string]interface{}) error {
	db, err := GetMongoInstance()
	if err != nil {
		return err
	}

	collection := db.Database(dbName).Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	_, updateError := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, dataQuery)
	if updateError != nil {
		return updateError
	}
	return nil
}

/*Delete : Author - Mahesh Chinvar
Purpose : Delete will delete data given for selector
Input :
	collectionName : collection name
	selector : update selector query
Return :
	error : database error
*/
func Delete(collectionName string, selector map[string]interface{}) error {

	db, err := GetMongoInstance()
	if err != nil {
		return err
	}

	if id, ok := selector["_id"]; ok {
		objID, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return err
		}
		selector["_id"] = objID
	}

	collection := db.Database(dbName).Collection(collectionName)
	_, deleteError := collection.DeleteOne(context.Background(), selector)
	if deleteError != nil {
		return deleteError
	}

	return nil
}

/*InsertOne : Author - Mahesh Chinvar
Purpose : InsertOne Save data in mongo db
Input :
	collectionName : collection name
	data : insert object
Return :
	error : database error
*/
func InsertOne(collectionName string, data interface{}) (string, error) {
	db, err := GetMongoInstance()
	if err != nil {
		return "", err
	}

	collection := db.Database(dbName).Collection(collectionName)
	opts, insertError := collection.InsertOne(context.Background(), data)
	if insertError != nil {
		return "", insertError
	}
	return getInsertedId(opts.InsertedID), nil
}

func getInsertedId(id interface{}) string {
	switch v := id.(type) {
	case string:
		return v
	case primitive.ObjectID:
		return v.Hex()
	default:
		return ""
	}
}
