package filter

import (
	"github.com/astaxie/beego/context"
	"github.com/shiruitao/go-one/application/health/common"
)

func LoginFilter(ctx *context.Context) {
	if _, ok := MapFilter[ctx.Request.RequestURI]; !ok {
		userID := ctx.Input.CruSession.Get(common.SessionUserID)
		if userID == nil {
			ctx.Output.JSON(map[string]interface{}{common.RespKeyStatus: common.ErrLoginRequired}, false, false)
		}
		power := ctx.Input.CruSession.Get(common.SessionPower)
		if power == nil {
			ctx.Output.JSON(map[string]interface{}{common.RespKeyStatus: common.ErrPermission}, false, false)
		}
		userName := ctx.Input.CruSession.Get(common.SessionUserName)
		if userName == nil {
			ctx.Output.JSON(map[string]interface{}{common.RespKeyStatus: common.ErrPermission}, false, false)
		}
		userNum := ctx.Input.CruSession.Get(common.SessionUserNum)
		if userNum == nil {
			ctx.Output.JSON(map[string]interface{}{common.RespKeyStatus: common.ErrPermission}, false, false)
		}
	}
}
