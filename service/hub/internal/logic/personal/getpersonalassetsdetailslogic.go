package personal

import (
	"context"
	"fmt"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonalAssetDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalAssetDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalAssetDetailsLogic {
	return &GetPersonalAssetDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalAssetDetailsLogic) GetPersonalAssetDetails(req *types.PersonalAssetDetailsReq) (resp *types.PersonalAssetOnceResp, err error) {
	var (
		result *model.Assets
	)

	result, err = l.svcCtx.MysqlServiceContext.Assets.FindAssIdOne(l.ctx, req.AssId)
	if err != nil {
		return nil, err
	}

	return &types.PersonalAssetOnceResp{
		Code:         types.SuccessCode,
		AssId:        result.AssId,
		GS:           result.GsId,
		Serial:       fmt.Sprintf("%v-%v", result.VersHead, result.VersTail),
		Project:      result.Project,
		SerialNumber: result.SerialNumber + fmt.Sprintf("-%v-%v", result.VersHead, result.VersTail),
		Country:      result.Country,
		Product:      result.Product,
		ProjectType:  result.ProjectType,
		Source:       result.Source,
		Day:          result.Day,
	}, nil
}
