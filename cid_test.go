package jpush

import (
	"context"
	"os"
	"testing"
)

var cidClient *CIDClient

func TestMain(m *testing.M) {
	println("TestMain setup.")

	cidClient = NewCIDClient(20, SetAppKey(appKey), SetMasterSecret(masterSecret), SetCIDCount(2))
	retCode := m.Run() // 执行测试，包括单元测试、性能测试和示例测试

	println("TestMain tear-down.")

	os.Exit(retCode)
}

func TestNewCIDClient(t *testing.T) {
	NewCIDClient(20,
		SetAppKey(appKey),
		SetMasterSecret(masterSecret),
		SetCIDCount(2), )
}

func TestGetPushID(t *testing.T) {
	res, err := cidClient.GetPushID(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestGetScheduleID(t *testing.T) {
	res, err := cidClient.GetScheduleID(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestNewCIDItem(t *testing.T) {
	options := &options{
		host:     "https://api.jpush.cn",
		cidCount: 1000,
		appKey:appKey,
		masterSecret:masterSecret,
	}

	res, err := newCIDItem(options, "push", 20).Get(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
