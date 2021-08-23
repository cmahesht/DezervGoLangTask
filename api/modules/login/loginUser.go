package login

import (
	"DezervGoLangTask/api/model"
	"DezervGoLangTask/helpers/confighelper"
	"DezervGoLangTask/helpers/databasehelper"
	"DezervGoLangTask/helpers/logginghelper"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

/*commentDAO : Author - Mahesh Chinvar
Purpose : Comment data access layer
*/

/*saveCommentDAO : Author - Mahesh Chinvar
Purpose :Save user details to database
 Input  :
	 commentParentIDs : Comment parent id list
	 comment : user specifed comment data
 Return :
	error : database error
*/
func saveUser(user *model.User) error {
	// generate ObjectID for user object
	objectId, objectIDErr := primitive.ObjectIDFromHex(bson.NewObjectId().Hex())
	if objectIDErr != nil {
		logginghelper.Error("ERROR : Error occurred in saveUser : primitive.ObjectIDFromHex ")
		return objectIDErr
	}
	user.Id = objectId
	user.CreationTimestamp = time.Now()

	// databasehelper.InsertOne insert new comment object
	_, updateError := databasehelper.InsertOne(confighelper.GetConfig("UserCollectionName"), user)
	if updateError != nil {
		logginghelper.Error("ERROR : Error occurred in saveUser : databasehelper.InsertOne ")
		return updateError
	}
	return nil
}

func CheckIfUserNameExist(userName string) (bool, error) {

	user, findError := databasehelper.Find(confighelper.GetConfig("UserCollectionName"), bson.M{"username": userName})
	if findError != nil {
		logginghelper.Error("ERROR : Error occurred in CheckIfUserNameExist : databasehelper.Find ")
		return false, findError
	}
	fmt.Println("user.Value() :", user.Value())
	if user.Value() != nil {
		return true, nil
	}

	return false, nil

}

func CheckIfUserPasswordIsPrevious(newUserDetails *model.User) (bool, error) {
	existingUserDetails := []model.User{}

	userData, findError := databasehelper.Find(confighelper.GetConfig("UserCollectionName"), bson.M{"username": newUserDetails.UserName})
	if findError != nil {
		logginghelper.Error("ERROR : Error occurred in CheckIfUserPasswordIsPrevious : databasehelper.Find ")
		return false, findError
	}
	
	// userdata - is a interface returned by the database find method
	// now we are converting userdata to json using Marshal
	data, marshallErr := json.Marshal(userData.Value())

	if marshallErr != nil {
		logginghelper.Error("ERROR : Error occurred in userdata : json.Marshal ")
		return false, marshallErr
	}

	// converting json data to go struct which is go struct
	marshallErr = json.Unmarshal(data, &existingUserDetails)
	if marshallErr != nil {
		logginghelper.Error("ERROR : Error occurred in userdata : json.Unmarshal ", marshallErr)
		fmt.Println("marshallErr :", marshallErr)
		return false, marshallErr
	}

	if newUserDetails.Password == existingUserDetails[0].Password {
		return true, nil
	}

	return false, nil

}

func updateUser(user *model.User) error {
	filter := bson.M{"name": bson.M{"$eq": user.UserName}}
	update := bson.M{"$set": bson.M{"password": user.Password}}

	
	updateError := databasehelper.Update(confighelper.GetConfig("UserCollectionName"), filter, update)
	if updateError != nil {
		logginghelper.Error("ERROR : Error occurred in updateUserDAO : databasehelper.updateUserDAO ")
		return updateError
	}

	return nil
}
