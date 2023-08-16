package personal

import (
	"context"
	"fmt"
	"hub/internal/svc"
	"hub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonalAssetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalAssetListLogic {
	return &GetPersonalAssetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalAssetListLogic) GetPersonalAssetList() (resp *types.PersonalAssetListResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.MysqlServiceContext.Assets.FindUidAll(context.Background(), l.ctx.Value("uid").(string))
	if err != nil {
		return nil, err
	}

	var respList []types.PersonalAssetOnceResp
	for _, v := range *result {
		respList = append(respList, types.PersonalAssetOnceResp{
			UserID:       v.UserId,
			Code:         types.SuccessCode,
			AssId:        v.AssId,
			GS:           v.GsId,
			Serial:       fmt.Sprintf("%v-%v", v.VersHead, v.VersTail),
			Project:      v.Project,
			SerialNumber: v.SerialNumber + fmt.Sprintf("-%v-%v", v.VersHead, v.VersTail),
			Country:      v.Country,
			HID:          v.Hid,
			Product:      v.Product,
			ProjectType:  v.ProjectType,
			Source:       v.Source,
			Number:       v.Number,
			RetireNumber: v.RetireNumber,
			Day:          v.Day,
			Listing:      v.Listing,
		})
	}
	return &types.PersonalAssetListResp{
		PersonalAssetList: respList,
	}, nil
}
