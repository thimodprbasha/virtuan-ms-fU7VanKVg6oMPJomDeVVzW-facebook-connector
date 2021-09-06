package api

import (
	"context"
	arfsidecar "facebook-connector/pkg/rule/ARF_Sidecar"
	util "facebook-connector/pkg/util"
	"reflect"
)

import (
	"facebook-connector/pkg/dto"
	"facebook-connector/sidecar"
	"github.com/jinzhu/copier"
)

func Sidecar(request *sidecar.RequestDto, ctx context.Context) (*sidecar.ResponseDto, error) {

	mRequest := dto.Request{}
	copier.Copy(&mRequest, &request)
	cntxt := util.CustomContext{}
	cntxt.GoContext = ctx
	cntxt.AppGoContext = context.Background()
	config := make(map[string]interface{})
	result := arfsidecar.ARF_Sidecar(&mRequest, &cntxt, config)
	if reflect.TypeOf(result) == reflect.TypeOf(dto.Response{}) || reflect.TypeOf(result) == reflect.TypeOf(&dto.Response{}) {
		var data *sidecar.ResponseDto
		copier.Copy(&data, result.(*dto.Response))
		return data, nil
	} else {
		var data sidecar.ResponseDto
		return &data, nil
	}
}
