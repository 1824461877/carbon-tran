package gsf_back

import (
	"context"
	"gsf/internal/svc"
	"gsf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChineseTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChineseTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChineseTemplateLogic {
	return &ChineseTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChineseTemplateLogic) ChineseTemplate(req *types.ChineseTemplateRequest) (resp *types.ChineseTemplateResponse, err error) {
	// todo: add your logic here and delete this line
	return
}
