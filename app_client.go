package noxus_sdk

import (
	"time"
	"context"
	pb "github.com/zm-dev/noxus-go-sdk/pb"
)

type AppClient struct {
	asc     pb.AppServiceClient
	timeout time.Duration
}

func (ac *AppClient) Validate(appID int32, appSecret string) (isValid bool, err error) {
	ctx, _ := context.WithTimeout(context.Background(), ac.timeout)
	if res, err := ac.asc.Validate(ctx, &pb.AppCredential{Id: appID, Secret: appSecret}); err != nil {
		return false, err
	} else {
		return res.GetIsValid(), nil
	}
}

func (ac *AppClient) Find(appID int32) (*pb.Application, error) {
	ctx, _ := context.WithTimeout(context.Background(), ac.timeout)
	return ac.asc.Find(ctx, &pb.AppID{Id: appID})
}

func (ac *AppClient) List(perPage, page int32) ([]*pb.Application, error) {
	ctx, _ := context.WithTimeout(context.Background(), ac.timeout)
	appList, err := ac.asc.List(ctx, &pb.AppListReq{PerPage: perPage, Page: page})
	if err != nil {
		return nil, err
	}
	return appList.GetApps(), nil
}

func NewAppClient(asc pb.AppServiceClient, timeout time.Duration) *AppClient {
	return &AppClient{asc: asc, timeout: timeout}
}
