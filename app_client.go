package noxus_sdk

import (
	pb "code.aliyun.com/digital_campus/noxus/src/rpc/pb"
	"time"
	"context"
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

func NewAppClient(timeout time.Duration) *AppClient {
	return &AppClient{timeout: timeout}
}
