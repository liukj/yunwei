package logic

import (
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xerr"
	"context"
	"github.com/pkg/errors"
	"fmt"
	"ywadmin-v3/service/identity/rpc/internal/svc"
	"ywadmin-v3/service/identity/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)


var ValidateTokenError = xerr.NewErrCode(xerr.TOKEN_EXPIRE_ERROR)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  validateToken ，只很对用户服务、授权服务api开放
func (l *ValidateTokenLogic) ValidateToken(in *pb.ValidateTokenReq) (*pb.ValidateTokenResp, error) {
	userTokenKey := fmt.Sprintf(globalkey.CacheUserTokenKey, in.UserId)
	dbToken, err := l.svcCtx.RedisClient.Get(userTokenKey)
	if err != nil {
		return nil, errors.Wrapf(ValidateTokenError, "ValidateToken RedisClient Get userId:%d ,err:%v", in.UserId, err)
	}

	if dbToken != in.Token {
		return nil, errors.Wrapf(ValidateTokenError, "ValidateToken is invalid  userId:%d , token:%s , dbToken:%s", in.UserId, in.Token, dbToken)
	}

	return &pb.ValidateTokenResp{
		Ok: true,
	}, nil
}
