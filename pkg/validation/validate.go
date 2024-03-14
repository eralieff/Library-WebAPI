package validation

import (
	"github.com/hashicorp/go-hclog"
)

type Validate struct {
	logger hclog.Logger
}

func NewValidation(logger hclog.Logger) *Validate {
	return &Validate{
		logger: logger,
	}
}

// fix
/*
func (v *Validate) ValidateAuthorUpdateFields(updatedAuthor *model.Author) error {
	validFields := map[string]bool{
		"id":         true,
		"full_name":  true,
		"nickname":   true,
		"speciality": true,
	}

	val := reflect.ValueOf(updatedAuthor).Elem()

	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Tag.Get("json")
		fieldName = strings.ToLower(fieldName)
		if !validFields[fieldName] {
			v.logger.Error("Unknown field in update request: ", fieldName)
			return fmt.Errorf("Unknown field in update request: %s", fieldName)
		}
	}

	return nil
}
*/
