package ARF_PostingSRV

import util "facebook-connector/pkg/util"

import (
	"facebook-connector/pkg/functions"
	"facebook-connector/pkg/model"
)

type CFGContext struct {
	Requeststring  *model.Requeststring
	token          string
	message        string
	responseString *model.Responsestring
	_context       *util.CustomContext
	_ruleConfig    map[string]interface{}
	_returnValue   interface{}
	_errorValue    interface{}
}

func NewCFGContext(pRequeststring *model.Requeststring, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Requeststring: pRequeststring,
		_context:      pContext,
		_ruleConfig:   pRuleConfig,
	}
}
func ARF_PostingSRV(pRequeststring *model.Requeststring, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequeststring, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiHybridFunctionNodeM01()
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.responseString, err = functions.CreatingPost(cfg.message, cfg.token)
	cfg._returnValue = cfg.responseString
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.token = cfg.Requeststring.Token
	cfg.message = cfg.Requeststring.Message

	return nil
}
