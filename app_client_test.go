package noxus_sdk

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/zm-dev/noxus-go-sdk/test/mock"
	pb "github.com/zm-dev/noxus-go-sdk/pb"
	"time"
)

func TestAppClient_Validate(t *testing.T) {
	
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	masc := mock_pb.NewMockAppServiceClient(ctrl)
	masc.EXPECT().Validate(gomock.Any(), gomock.Any(), gomock.Any()).Return(&pb.AppValidateRes{
		IsValid: true,
	}, nil)
	ac := NewAppClient(masc, 10*time.Second)
	isValid, err := ac.Validate(123, "456")
	if err != nil {
		t.Errorf("unexcept error, error:%+v", err)
	}
	if !isValid {
		t.Error("isValid except true")
	}
}

func TestAppClient_Find(t *testing.T) {
	var appID int32 = 123
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	masc := mock_pb.NewMockAppServiceClient(ctrl)
	masc.EXPECT().Find(gomock.Any(), gomock.Any()).Return(&pb.Application{
		Id: appID,
	}, nil)
	ac := NewAppClient(masc, 10*time.Second)
	app, err := ac.Find(appID)
	if err != nil {
		t.Errorf("unexcept error, error:%+v", err)
	}
	if app.Id != appID {
		t.Errorf("app.Id=%d, except %d", app.Id, appID)
	}
}

func TestAppClient_List(t *testing.T) {

	tests := []struct {
		perPage, page int32
	}{
		{1, 1},
		{1, 2},
		{20, 2},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	masc := mock_pb.NewMockAppServiceClient(ctrl)
	masc.EXPECT().List(gomock.Any(), gomock.Any()).Return(&pb.AppList{}, nil).Times(len(tests))
	ac := NewAppClient(masc, 10*time.Second)
	for _, test := range tests {
		_, err := ac.List(test.perPage, test.page)
		if err != nil {
			t.Errorf("unexcept error, error:%+v", err)
		}
	}

}
