package validationhelper

import (
	"github.com/go-playground/locales/en"

	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

/*ValidationHelper : Author - Mahesh Chinvar
Purpose : Default validations
*/

//Validate method
func Validate(s interface{}) map[string]string {
	var validate *validator.Validate
	var uni *ut.UniversalTranslator

	validate = validator.New()

	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	//For custom error message
	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	err := validate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		customErrs := make(map[string]string, len(errs))

		for _, e := range errs {
			// can translate each error one at a time.
			customErrs[e.Namespace()] = e.Translate(trans)
		}
		return customErrs
	}
	return nil
}
