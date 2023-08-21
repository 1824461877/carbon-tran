package personal

import (
	"context"
	"errors"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"
	public "hub/utils"
	"retire_cert"
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

	var (
		filename string
		onTime   = time.Now()
	)

	if filename, _, err = l.svcCtx.RetireCert.RetireCertInter.Create(&retire_cert.RetireCertificate{
		ByFrom:         "hniee 碳资产中心",
		ByTo:           one.UserId,
		OnTime:         onTime.Format("2006-01-02"),
		VerifiedNumber: req.Number,
		ProjectName:    one.Project,
		RetiredBy:      one.UserId,
	}); err != nil {
		return nil, err
	}

	if _, err = l.svcCtx.MysqlServiceContext.Retire.Insert(l.ctx, &model.Retire{
		Number:          req.Number,
		UserId:          uid,
		AssId:           req.AssId,
		RId:             public.RID(),
		Status:          1001,
		CertificateLink: filename,
		CreateTime:      onTime,
	}); err != nil {
		return nil, err
	}
	return &types.PersonalAssetRetireResp{
		Code:    types.SuccessCode,
		Message: "do is success",
	}, nil
}
