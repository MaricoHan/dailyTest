package http

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	request, err := http.NewRequest(http.MethodPut, "https://stage.apis.avata.bianjie.ai/v1beta1/nft/classes", bytes.NewBuffer([]byte("132456789")))
	fmt.Println("===========", request.URL.Path)

	if err != nil {
		return
	}
	//	body := request.Body
	//	all, err := ioutil.ReadAll(body)
	//	if err != nil {
	//		return
	//	}
	//	fmt.Println(string(all))
	//
	//
	//
	//
	//	body = request.Body
	//	body.Close()
	//	all, err = ioutil.ReadAll(body)
	//	fmt.Println(string(all))

}
