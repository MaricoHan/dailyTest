package signTest

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSign(t *testing.T) {
	apiKey := "y680WS4FZCywcRB4irMM3xxxxxx"
	apiSecret := "g1V2VKTnlWlUzZ2DQd0kxxxxx"
	url := "https://stage.apis.avata.bianjie.ai"
	path := "/v1beta1/nft/classes"

	body:= map[string]string{
		"name": "MBOX-ttt",
		"symbol": "MBOX-ttt",
		"owner": "iaa1xd44h5cktcv2f7k96quwpmx2cczkzr48re4aje",
		"operation_id": "123333333",
	}

	bodyBZ, err := json.Marshal(body)
	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodGet,url+path,bytes.NewBuffer(bodyBZ))
	if err != nil {
		return
	}

	signRequest := SignRequest(request, apiKey, apiSecret)

	fmt.Println("-----------------",signRequest.Header.Get("X-Signature"))

}


// SignRequest 对请求进行签名
func SignRequest(r *http.Request, apiKey, apiSecret string) *http.Request {
	timestamp := "123123"
	// 获取 path params
	params := map[string]interface{}{}
	params["path_url"] = r.URL.Path
	// 获取 query params
	for k, v := range r.URL.Query() {
		k = "query_" + k
		params[k] = v[0]
	}
	// 获取 body params
	// 把request的内容读取出来
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// 把刚刚读出来的再写进去
	if bodyBytes != nil {
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	paramsBody := map[string]interface{}{}
	_ = json.Unmarshal(bodyBytes, &paramsBody)
	hexHash := hash(timestamp + apiSecret)
	for k, v := range paramsBody {
		k = "body_" + k
		params[k] = v
	}
	sortParams := params
	if sortParams != nil {
		sortParamsBytes, _ := json.Marshal(sortParams)
		hexHash = hash(string(sortParamsBytes) + timestamp + apiSecret)
	}
	r.Header.Set("X-Api-Key", apiKey)
	r.Header.Set("X-Signature", hexHash)
	r.Header.Set("X-Timestamp", timestamp)
	return r
}
func hash(oriText string) string {
	oriTextHashBytes := sha256.Sum256([]byte(oriText))
	return hex.EncodeToString(oriTextHashBytes[:])
}