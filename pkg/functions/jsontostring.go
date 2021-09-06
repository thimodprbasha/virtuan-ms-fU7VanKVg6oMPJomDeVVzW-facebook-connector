package functions

import (
	"facebook-connector/pkg/model"
)

func JsonToString(responsestring *model.Responsestring) (string, error) {
	return responsestring.ReturnData, nil
}
