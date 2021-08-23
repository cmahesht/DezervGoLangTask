package login

import (
	"DezervGoLangTask/api/constants"
	"DezervGoLangTask/api/model"
	"DezervGoLangTask/helpers/logginghelper"
	"DezervGoLangTask/helpers/validationhelper"
	"net/http"

	"github.com/labstack/echo"
)

/*Route : Author - Mahesh Chinvar
Purpose : All User api endpoints
*/

/*Init : Author - Mahesh Chinvar
Purpose :registration of routes
*/
func Init(o *echo.Group, r *echo.Group) {

	//	URL should be ‘/createUser
	o.POST("/createUser", createUser)

	//	URL should be ‘/updateUserDetails
	o.POST("/updateUserDetails", updateUserDetails)

	// post login request
	r.POST("/userdetail", updateUserDetails)

}

/* createUser
Author - Mahesh Chinvar
Purpose :  /*  Create User: user create name, email, phone number.
					a registered user should have unique password followed by password rules.
					a registered user can update exisiting password followed by password rules.
			Should be a POST request
			Use JSON request body
			URL should be ‘/createUser
	CommentViewModel :
			'userName' validate:"required"` - user name
			'email' validate:"required"` - user email id
			'mobile' validate:"required"` - user mobile/phone number
return: Must return success or error message
*/
func createUser(r echo.Context) error {

	userModel := new(model.User)
	//Bind incoming request with model
	bindErr := r.Bind(userModel)
	if bindErr != nil {
		logginghelper.Error("ERROR occurred in createUser : Parameter binding")
		return r.JSON(http.StatusExpectationFailed, constants.DezervErrorCodeParameterBindError)
	}

	//Validate model with validator
	validationErr := validationhelper.Validate(userModel)
	if validationErr != nil {
		logginghelper.Error("ERROR occurred in createUser : Validate ")
		return r.JSON(http.StatusExpectationFailed, constants.DezervErrorCodeRequiredFieldValidationFailed)
	}

	isValid, errList := PasswordRulesValidationNewUser(userModel)
	if !isValid {
		logginghelper.Error("ERROR occurred in PasswordRulesValidation ", errList)
		return r.JSON(http.StatusExpectationFailed, errList)
	}

	// SaveUserService service call
	serviceErr := SaveUserService(userModel)
	if serviceErr != nil {
		logginghelper.Error("ERROR occurred in saveUser : SaveUserService")
		return r.JSON(http.StatusInternalServerError, constants.DezervErrorCodeServiceNotAvailable)
	}

	return r.JSON(http.StatusOK, constants.DezervUserCreatedSavedSuccessfully)
}

func updateUserDetails(r echo.Context) error {

	userModel := new(model.User)

	//Bind incoming request with model
	bindErr := r.Bind(userModel)
	if bindErr != nil {
		logginghelper.Error("ERROR occurred in UpdateUserDetails : Parameter binding")
		return r.JSON(http.StatusExpectationFailed, constants.DezervErrorCodeParameterBindError)
	}

	//Validate model with validator
	validationErr := validationhelper.Validate(userModel)
	if validationErr != nil {
		logginghelper.Error("ERROR occurred in UpdateUserDetails : Validate ")
		return r.JSON(http.StatusExpectationFailed, constants.DezervErrorCodeRequiredFieldValidationFailed)
	}

	isValid, errList := PasswordRulesValidationUpdateUser(userModel)
	if !isValid {
		logginghelper.Error("ERROR occurred in PasswordRulesValidation ", errList)
		return r.JSON(http.StatusExpectationFailed, errList)
	}

	// SaveUserService service call
	serviceErr := UpdateUserService(userModel)
	if serviceErr != nil {
		logginghelper.Error("ERROR occurred in UpdateUserDetails : UpdateUserService")
		return r.JSON(http.StatusInternalServerError, constants.DezervErrorCodeServiceNotAvailable)
	}

	return r.JSON(http.StatusOK, constants.DezervUserUpdatedSuccessfully)
}
