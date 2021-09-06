package functions

import (
	"encoding/json"
	"facebook-connector/pkg/model"
)

func StringToJson(requeststring string) (*model.Requeststring, error) {

	var jsonModel *model.Requeststring
	err := json.Unmarshal([]byte(requeststring), jsonModel)

	// Code

	return jsonModel, err
}
