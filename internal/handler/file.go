package handler

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"focus-single/apiv1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

var (
	// File 文件管理
	File = handlerFile{}
)

type handlerFile struct{}

func (a *handlerFile) Upload(ctx context.Context, req *apiv1.FileUploadReq) (res *apiv1.FileUploadRes, err error) {
	var (
		request = g.RequestFromCtx(ctx)
		file    = request.GetUploadFile("file")
	)
	if file == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	result, err := service.File.Upload(ctx, model.FileUploadInput{
		File:       file,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	res = &apiv1.FileUploadRes{
		Name: result.Name,
		Url:  result.Url,
	}
	return
}
