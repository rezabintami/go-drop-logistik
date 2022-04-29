package helpers

import (
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

var (
	vd    *validator.Validate
	trans ut.Translator
)

// Validate checking data
func Validate(data interface{}) (map[string]string, bool, error) {
	err := vd.Struct(data)
	if err != nil {
		errorMap := map[string]string{}
		errorType := ""
		for _, e := range err.(validator.ValidationErrors) {
			if errorType == "" {
				errorType = e.Tag()
				errorMap[e.StructField()] = formatMessage(e, trans)
			} else {
				if e.Tag() == errorType {
					errorMap[e.StructField()] = formatMessage(e, trans)
				}
			}
		}
		return errorMap, true, nil
	}
	return map[string]string{}, false, nil
}

// Validation for initial function validation
func InitValidation() {
	vd = validator.New()
}

func formatMessage(err validator.FieldError, trans ut.Translator) string {
	message := ""

	switch err.Tag() {
	case "phone_number":
		message = "nomor handphone tidak sesuai (081234567890)"
	default:
		message = err.Translate(trans)
	}
	return message
}
