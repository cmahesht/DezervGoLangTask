package login

import (
	"DezervGoLangTask/api/model"
	"unicode"
)

/*SaveUserService : Author - Mahesh Chinvar
Purpose :Save user details
 Input  :
	 UserModel : user specifed user model
return  :
	error : service error
*/
func SaveUserService(user *model.User) error {
	return saveUser(user)
}

// Password validates plain password against the rules defined below.

/*
upper	: at least one upper case letter.
lower 	: at least one lower case letter.
number 	: at least one digit.
symbol	: at least one special character.
lenght	: at least 8 characters long.
No empty string or whitespace.
*/
func PasswordRulesValidationNewUser(user *model.User) (bool, []string) {
	errList := BasicPasswordRulesValidation(user)

	if len(errList) > 0 {
		return false, errList
	}

	isUserAlreadyExist, _ := CheckIfUserNameExist(user.UserName)
	if isUserAlreadyExist {
		errList = append(errList, "USER WITH SAME USERNAME ALREADY EXIST")
	}
	if len(errList) > 0 {
		return false, errList
	}

	return true, errList
}

func BasicPasswordRulesValidation(user *model.User) []string {
	var (
		upp, low, num, sym bool
		tot uint8
	)
	errList := []string{}
	if user.UserName == user.Password {
		errList = append(errList, "PASSWORD MUST NOT BE SAME AS USERNAME")
		return errList
	}
	for _, char := range user.Password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		}
	}

	if !upp {
		errList = append(errList, "PASSWORD MUST CONTAINS ATLEAST ONE UPPER CASE LETTER")
	}
	if !low {
		errList = append(errList, "PASSWORD MUST CONTAINS ATLEAST ONE LOWER CASE LETTER")
	}
	if !num {
		errList = append(errList, "PASSWORD MUST CONTAINS ATLEAST ONE NUMBER CASE LETTER")
	}
	if !sym {
		errList = append(errList, "PASSWORD MUST CONTAINS ATLEAST ONE SPECIAL CHARACTER")
	}
	if tot < 8 || tot > 16 {
		errList = append(errList, "PASSWORD LENGTH SHOULD BE MINIMUM 8 TO MAXIMUM 16 CHARACTER LONG")
	}

	return errList
}

func PasswordRulesValidationUpdateUser(user *model.User) (bool, []string) {
	errList := []string{}
	isUserNameExit, _ := CheckIfUserNameExist(user.UserName)
	if !isUserNameExit {
		errList = append(errList, "USERNAME IS NOT EXIST")
		return false, errList
	}
	errList = append(errList, BasicPasswordRulesValidation(user)...)

	isPreviousPasswordSame, _ := CheckIfUserPasswordIsPrevious(user)
	if isPreviousPasswordSame {
		errList = append(errList, "NEW PASSWORD SHOULDN’T BE THE SAME AS LAST PASSWORD")
	}

	if len(errList) > 0 {
		return false, errList
	}

	return true, errList
}

func UpdateUserService(user *model.User) error {
	return updateUser(user)
}
