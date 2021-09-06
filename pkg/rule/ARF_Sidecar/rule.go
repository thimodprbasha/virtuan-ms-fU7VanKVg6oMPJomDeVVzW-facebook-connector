package ARF_Sidecar

import util "facebook-connector/pkg/util"

import (
	"facebook-connector/pkg/dto"
	"facebook-connector/pkg/functions"
	"facebook-connector/pkg/model"
	"facebook-connector/pkg/rule/ARF_PostingSRV"
)

type CFGContext struct {
	Request       *dto.Request
	responseData  *model.Responsestring
	returnDTO     string
	response      *dto.Response
	requestString string
	jsonModel     *model.Requeststring
	_context      *util.CustomContext
	_ruleConfig   map[string]interface{}
	_returnValue  interface{}
	_errorValue   interface{}
}

func NewCFGContext(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Request:     pRequest,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_Sidecar(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequest, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiHybridFunctionNodeM01()

	cfg.xiSubRuleNodeM12()

	cfg.xiModelPropertyNodeM23()

	cfg.xiModelPropertyNodeM34()

	cfg.xiHybridFunctionNodeM45()
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM45() error {
	var err error
	cfg.returnDTO, err = functions.JsonToString(cfg.responseData)
	cfg._returnValue = cfg.returnDTO
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM34() error {
	cfg.returnDTO = cfg.response.Data

	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM23() error {
	cfg.response = &dto.Response{}

	return nil
}

func (cfg *CFGContext) xiSubRuleNodeM12() error {
	cfg.responseData = ARF_PostingSRV.ARF_PostingSRV(cfg.jsonModel, cfg._context, cfg._ruleConfig).(*model.Responsestring)
	cfg._returnValue = cfg.responseData
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.jsonModel, err = functions.StringToJson(cfg.requestString)
	cfg._returnValue = cfg.jsonModel
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.requestString = cfg.Request.Data
	cfg.jsonModel = &model.Requeststring{}
	cfg.responseData = &model.Responsestring{}

	return nil
}
