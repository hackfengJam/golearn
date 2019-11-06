package sdk

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

var (
	host        = "api.1sapp.com"
	accesstoken = "d941hNTEJ4Nmn5yg2xHn-rk_TzOEFudmh3an7islrN72DV5nN65xsSufKSyVcz4ibpZaeHKl1nFPFr3zWA"
	appId       = "3b3Z1v2CTbrGCDPSLLcy"
	appSecret   = []byte("Q&uHbSg7FIkOP!HVf6Oz")
)

func TestMac_SignRequest(t *testing.T) {
	mac := &Mac{
		AccessKey: appId,
		SecretKey: appSecret,
		Host:      host,
	}
	targetUrl := fmt.Sprintf("%v/member/getCustomInfo?token=%v", host, url.QueryEscape(accesstoken))

	httpReq, err := http.NewRequest(http.MethodGet, targetUrl, nil)
	if err != nil {
		t.Fatal(err)
	}

	sa := &SignArgs{
		http.MethodGet,
		httpReq.URL.Path,
		httpReq.URL.RawQuery,
		"",
		GetBody(httpReq),
	}

	sign, err := mac.SignRequest(sa)
	if err != nil {
		t.Fatal(err)
	}

	// mt  3b3Z1v2CTbrGCDPSLLcy:VGlk2cbaxb4bXmwDWuo__gLSmJ4=
	t.Log("mt ", sign)
}
