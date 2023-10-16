package exchange

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExchangeAssetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type Stand struct {
	All int
	JCM int
	GS  int
	VCS int
}

var (
	stand = &Stand{}
)

func NewGetExchangeAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExchangeAssetListLogic {
	return &GetExchangeAssetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExchangeAssetListLogic) GetExchangeAssetList(req *types.ExchangeAssetListReq) (resp *types.ExchangeAssetListResp, err error) {

	var (
		all       *[]model.AssetsSell
		standBool = false
	)

	if req.LocationTabulation != "" {
		all, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindAllLocationTabulation(l.ctx, req.LocationTabulation)
		standBool = false
	} else if req.SourceTabulation != "" {
		all, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindAllSourceTabulation(l.ctx, req.SourceTabulation)
		standBool = false
	} else {
		all, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindAll(l.ctx)
		standBool = true
		stand = &Stand{}
	}

	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}

	var list []types.ExchangeAssetOnceResp
	for _, v := range *all {
		var (
			one *model.Assets
			//ec  model.AssetsSell
		)
		one, err = l.svcCtx.MysqlServiceContext.Assets.FindAssIdOne(l.ctx, v.AssId)
		if err != nil {
			break
		}

		if standBool {
			switch one.Source {
			case "jcm":
				stand.JCM++
			case "gs":
				stand.GS++
			case "vcs":
				stand.VCS++
			}
			stand.All++
		}

		if req.Search != "" {
			if strings.Contains(one.Product, req.Search) || strings.Contains(one.ProjectType, req.Search) || strings.Contains(strconv.FormatInt(one.Day, 10), req.Search) || strings.Contains(one.SerialNumber, req.Search) || strings.Contains(one.Country, req.Search) || strings.Contains(one.Project, req.Search) || strings.Contains(one.GsId, req.Search) {
				list = append(list, types.ExchangeAssetOnceResp{
					ExId:         v.ExId,
					Assid:        v.AssId,
					UserId:       v.UserId,
					Number:       v.Number,
					GS:           one.GsId,
					Serial:       fmt.Sprintf("%v-%v", one.VersHead, one.VersTail),
					Project:      one.Project,
					SerialNumber: one.SerialNumber + fmt.Sprintf("-%v-%v", one.VersHead, one.VersTail),
					Source:       one.Source,
					Day:          one.Day,
					Amount:       v.Amount,
					Country:      one.Country,
					Product:      one.Product,
					ProjectType:  one.ProjectType,
					CreateTime:   v.CreateTime.Unix(),
					EndTime:      v.EndTime.Unix(),
				})
			}
		} else {
			list = append(list, types.ExchangeAssetOnceResp{
				ExId:         v.ExId,
				Assid:        v.AssId,
				UserId:       v.UserId,
				Number:       v.Number,
				GS:           one.GsId,
				Serial:       fmt.Sprintf("%v-%v", one.VersHead, one.VersTail),
				Project:      one.Project,
				SerialNumber: one.SerialNumber + fmt.Sprintf("-%v-%v", one.VersHead, one.VersTail),
				Source:       one.Source,
				Day:          one.Day,
				Amount:       v.Amount,
				Country:      one.Country,
				Product:      one.Product,
				ProjectType:  one.ProjectType,
				CreateTime:   v.CreateTime.Unix(),
				EndTime:      v.EndTime.Unix(),
			})
		}
	}

	return &types.ExchangeAssetListResp{
		ExchangeAssetList: list,
		ExchangeAssetTabulationResp: types.ExchangeAssetTabulationResp{
			SourceTabulation: []string{
				"JCM",
				"GS",
				"VCS",
			},
			Stand: []types.Stand{
				{
					Name: "ALL",
					Val:  stand.All,
				},
				{
					Name: "JCM",
					Val:  stand.JCM,
				},
				{
					Name: "GS",
					Val:  stand.GS,
				},
				{
					Name: "VCS",
					Val:  stand.VCS,
				},
			},
			LocationTabulation: []string{
				"cn",
				"tr",
				"rw",
				"er",
				"ru",
				"bd",
				"vn",
				"jp",
				"th",
				"sa",
				"pw",
				"mv",
				"mg",
				"la",
				"kh",
				"ke",
				"id",
			},
		},
	}, nil
}
