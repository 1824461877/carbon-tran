package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hub/internal/svc"
	"io/ioutil"
	"os"
)

type RetireFile struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetireFile(ctx context.Context, svcCtx *svc.ServiceContext) *RetireFile {
	return &RetireFile{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetireFile) RetireFile(fid string) (all []byte, err error) {
	// todo: add your logic here and delete this line
	var (
		file *os.File
	)
	file, err = os.OpenFile(fid, os.O_RDONLY, 0666)
	all, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return
}
