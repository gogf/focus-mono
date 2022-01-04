package handler

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"focus-single/apiv1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

var (
	// Register 注册控制器
	Register = handlerRegister{}
)

type handlerRegister struct{}

func (a *handlerRegister) Index(ctx context.Context, req *apiv1.RegisterIndexReq) (res *apiv1.RegisterIndexRes, err error) {
	service.View.Render(ctx, model.View{})
	return
}

func (a *handlerRegister) Register(ctx context.Context, req *apiv1.RegisterDoReq) (res *apiv1.RegisterDoRes, err error) {
	if !service.Captcha.VerifyAndClear(g.RequestFromCtx(ctx), consts.CaptchaDefaultName, req.Captcha) {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}
	if err = service.User.Register(ctx, model.UserRegisterInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	}); err != nil {
		return
	}

	// 自动登录
	err = service.User.Login(ctx, model.UserLoginInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}
