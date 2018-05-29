package noxus_sdk

import (
	"time"
	"context"
	pb "github.com/zm-dev/noxus-go-sdk/pb"
	"google.golang.org/grpc"
)

type AppClient struct {
	asc     pb.AppServiceClient
	timeout time.Duration
}

func (ac *AppClient) ValidateApp(appID int32, appSecret string) (isValid bool, err error) {
	ctx, _ := context.WithTimeout(context.Background(), ac.timeout)
	if res, err := ac.asc.ValidateApp(ctx, &pb.AppCredential{Id: appID, Secret: appSecret}); err != nil {
		return false, err
	} else {
		return res.GetIsValid(), nil
	}
}

func (ac *AppClient) FindApp(appID int32) (*pb.Application, error) {
	ctx, _ := context.WithTimeout(context.Background(), ac.timeout)
	return ac.asc.FindApp(ctx, &pb.AppID{Id: appID})
}

func (ac *AppClient) ListApp(perPage, page int32) ([]*pb.Application, error) {
	ctx, _ := context.WithTimeout(context.Background(), ac.timeout)
	appList, err := ac.asc.ListApp(ctx, &pb.AppListReq{PerPage: perPage, Page: page})
	if err != nil {
		return nil, err
	}
	return appList.GetApps(), nil
}

func NewAppClient(noxusHost string, timeout time.Duration) (*AppClient, error) {
	conn, err := grpc.Dial(noxusHost, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &AppClient{asc: pb.NewAppServiceClient(conn), timeout: timeout}, nil
}

func ValidateApp(ctx context.Context, appID int32, appSecret string) (isValid bool, err error) {
	return FromContext(ctx).ValidateApp(appID, appSecret)
}

func FindApp(ctx context.Context, appID int32) (*pb.Application, error) {
	return FromContext(ctx).FindApp(appID)
}

func ListApp(ctx context.Context, perPage, page int32) ([]*pb.Application, error) {
	return FromContext(ctx).ListApp(perPage, page)
}
