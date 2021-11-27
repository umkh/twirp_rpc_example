package http_server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/twitchtv/twirp"
	pb_auth "github.com/umkh/twirp_rpc_example/pkg/genproto/services/auth"
	"net/http"
)

type auth struct{}

func NewAuth() (string, http.HandlerFunc) {
	srv := pb_auth.NewAuthServer(
		&auth{},
		twirp.WithServerPathPrefix(urlPrefix),
	)

	return srv.PathPrefix(), middleware(srv).ServeHTTP
}

func (a *auth) Login(ctx context.Context, req *pb_auth.LoginReq) (*pb_auth.LoginRes, error) {
	if err := req.Validate(); err != nil {
		return nil, twirp.NewError(twirp.InvalidArgument, err.Error())
	}

	return &pb_auth.LoginRes{
		AccessToken: "sdkfkdsjksfjgkfdmgkmdflkgma;lkflasdlasd",
		UpdateToken: "ldsfsdkfl;dksflksdf;ldk;kfdgjkkdfjgk;fakf;",
	}, nil
}

func (a *auth) Register(ctx context.Context, req *pb_auth.RegisterReq) (*empty.Empty, error) {
	return nil, nil
}

func (a *auth) Forgot(ctx context.Context, req *pb_auth.ForgotReq) (*empty.Empty, error) {
	return nil, nil
}

func (a *auth) Activate(ctx context.Context, req *pb_auth.ActivateReq) (*empty.Empty, error) {
	return nil, nil
}
