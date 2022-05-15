package verify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
)

type res struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`

	Data struct {
		BizCode string `json:"bizCode"`
	} `json:"data"`
}

func TestVerify(t *testing.T) {
	host := "https://id2meta.market.alicloudapi.com"
	path := "/id2meta"
	appcode := "1efaadcf0eba402396a4df709f27a6df"

	reqURL, err := url.Parse(host + path)
	if err != nil {
		fmt.Println(err)
		return
	}
	params := url.Values{}
	params.Add("identifyNum", "321324199909155076")
	params.Add("userName", "韩拓")
	reqURL.RawQuery = params.Encode()

	client := http.DefaultClient

	request, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		fmt.Println("request", err)
		return
	}

	request.Header.Add("Authorization", "APPCODE "+appcode)

	response, err := client.Do(request)
	if err != nil {

		fmt.Println("response", err)
		return
	}


	all, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	res:=res{}
	json.Unmarshal(all,&res)

	fmt.Println(res)

}
