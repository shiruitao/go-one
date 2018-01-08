package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/shiruitao/go-one/application/cheng/common"
)

func LoginFilter(ctx *context.Context) {
	if _, ok := MapFilter[ctx.Request.RequestURI]; !ok {
		userID := ctx.Input.CruSession.Get(common.SessionAdminID)

		if userID == nil {
			ctx.Output.JSON(map[string]interface{}{common.RespKeyStatus: common.ErrLoginRequired}, false, false)
		}
	}
}