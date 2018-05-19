package filter

import (
	"github.com/astaxie/beego/context"
	"github.com/shiruitao/go-one/application/bbs/common"
)

func LoginFilter(ctx *context.Context) {
	if _, ok := MapFilter[ctx.Request.RequestURI]; !ok {
		userID := ctx.Input.CruSession.Get(common.SessionUser)
		if userID == nil {
			ctx.Output.JSON(map[string]interface{}{common.RespKeyStatus: common.ErrLoginRequired}, false, false)
		}
		isAdmin := ctx.Input.CruSession.Get(common.SessionIsAdmin)
		if isAdmin == nil {
			ctx.Output.JSON(map[string]interface{}{common.RespKeyStatus: common.ErrPermission}, false, false)
		}
	}
}