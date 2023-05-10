package httpService

import (
	"bytes"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

func SetupClient() (client *http.Client) {
	client = &http.Client{
		Timeout: 30 * time.Second,
	}
	return
}

func SetupHeader() (head http.Header, err error) {
	head = http.Header{}
	for key, value := range viper.GetStringMapString("httpService.header") {
		head.Add(key, value)
	}
	return
}

func SetupRequestByUrl(URL string) (req http.Request, err error) {
	header, err := SetupHeader()
	if err != nil {
		return
	}
	rawUrl, err := url.Parse(URL)
	if err != nil {
		return
	}
	req = http.Request{URL: rawUrl, Header: header}
	return
}

func ReadResponse(resp *http.Response) (s string, err error) {
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return
	}
	s = buf.String()
	return
}
