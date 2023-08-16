package personal

import (
	"context"
	"errors"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"
	public "hub/utils"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PersonalAssetRetireLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPersonalAssetRetireLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersonalAssetRetireLogic {
	return &PersonalAssetRetireLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PersonalAssetRetireLogic) PersonalAssetRetire(uid string, req *types.PersonalAssetRetireReq) (resp *types.PersonalAssetRetireResp, err error) {
	// todo: add your logic here and delete this line

	var one *model.Assets
	if one, err = l.svcCtx.MysqlServiceContext.Assets.FindAssIdOne(l.ctx, req.AssId); err == nil {
		if one.Number < req.Number {
			return nil, errors.New("exceeding maximum value")
		}
		if err = l.svcCtx.MysqlServiceContext.Assets.UpdateNumber(l.ctx, &model.Assets{AssId: req.AssId, RetireNumber: one.RetireNumber + req.Number, Number: one.Number - req.Number}); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	if _, err = l.svcCtx.MysqlServiceContext.Retire.Insert(l.ctx, &model.Retire{
		Number:          req.Number,
		UserId:          uid,
		AssId:           req.AssId,
		RId:             public.RID(),
		Status:          1001,
		CertificateLink: "",
		CreateTime:      time.Now(),
	}); err != nil {
		return nil, err
	}
	return &types.PersonalAssetRetireResp{
		Code:    types.SuccessCode,
		Message: "do is success",
	}, nil
}
