package validation

import (
	"encoding/json"
	"errors"
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

func (v *Validate) ValidateAuthorUpdateFields(requestBody []byte) error {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(requestBody, &jsonData); err != nil {
		v.logger.Error("Error parsing request body: ", err.Error())
		return errors.New("Error parsing request body: " + err.Error())
	}

	validFields := map[string]bool{
		"id":         true,
		"full_name":  true,
		"nickname":   true,
		"speciality": true,
	}

	for fieldName := range jsonData {
		if !validFields[fieldName] {
			v.logger.Error("Unknown field in update request: ", fieldName)
			return errors.New("Unknown field in update request: " + fieldName)
		}
	}

	return nil
}

func (v *Validate) ValidateBookUpdateFields(requestBody []byte) error {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(requestBody, &jsonData); err != nil {
		v.logger.Error("Error parsing request body: ", err.Error())
		return errors.New("Error parsing request body: " + err.Error())
	}

	validFields := map[string]bool{
		"id":        true,
		"title":     true,
		"genre":     true,
		"isbn":      true,
		"author_id": true,
	}

	for fieldName := range jsonData {
		if !validFields[fieldName] {
			v.logger.Error("Unknown field in update request: ", fieldName)
			return errors.New("Unknown field in update request: " + fieldName)
		}
	}

	return nil
}
