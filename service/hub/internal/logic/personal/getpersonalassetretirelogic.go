package personal

import (
	"context"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonalAssetRetireLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalAssetRetireLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalAssetRetireLogic {
	return &GetPersonalAssetRetireLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalAssetRetireLogic) GetPersonalAssetRetire(uid string) (resp *types.GetPersonalAssetRetireResp, err error) {
	// todo: add your logic here and delete this line
	var all *[]model.Retire
	if all, err = l.svcCtx.MysqlServiceContext.Retire.FindALLByUserId(l.ctx, uid); err != nil {
		return nil, err
	}

	var list []types.RetireOnce
	for _, v := range *all {
		list = append(list, types.RetireOnce{
			RId:             v.RId,
			AssId:           v.AssId,
			Number:          v.Number,
			CertificateLink: v.CertificateLink,
			Status:          v.Status,
			CreateTime:      v.CreateTime.Unix(),
		})
	}
	return &types.GetPersonalAssetRetireResp{
		RetireList: list,
	}, nil
}
