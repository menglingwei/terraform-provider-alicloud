package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryInfoByToken invokes the emr.QueryInfoByToken API synchronously
// api document: https://help.aliyun.com/api/emr/queryinfobytoken.html
func (client *Client) QueryInfoByToken(request *QueryInfoByTokenRequest) (response *QueryInfoByTokenResponse, err error) {
	response = CreateQueryInfoByTokenResponse()
	err = client.DoAction(request, response)
	return
}

// QueryInfoByTokenWithChan invokes the emr.QueryInfoByToken API asynchronously
// api document: https://help.aliyun.com/api/emr/queryinfobytoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryInfoByTokenWithChan(request *QueryInfoByTokenRequest) (<-chan *QueryInfoByTokenResponse, <-chan error) {
	responseChan := make(chan *QueryInfoByTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryInfoByToken(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QueryInfoByTokenWithCallback invokes the emr.QueryInfoByToken API asynchronously
// api document: https://help.aliyun.com/api/emr/queryinfobytoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryInfoByTokenWithCallback(request *QueryInfoByTokenRequest, callback func(response *QueryInfoByTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryInfoByTokenResponse
		var err error
		defer close(result)
		response, err = client.QueryInfoByToken(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// QueryInfoByTokenRequest is the request struct for api QueryInfoByToken
type QueryInfoByTokenRequest struct {
	*requests.RpcRequest
}

// QueryInfoByTokenResponse is the response struct for api QueryInfoByToken
type QueryInfoByTokenResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TokenUid  string `json:"TokenUid" xml:"TokenUid"`
	TokenBid  string `json:"TokenBid" xml:"TokenBid"`
}

// CreateQueryInfoByTokenRequest creates a request to invoke QueryInfoByToken API
func CreateQueryInfoByTokenRequest() (request *QueryInfoByTokenRequest) {
	request = &QueryInfoByTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "QueryInfoByToken", "emr", "openAPI")
	return
}

// CreateQueryInfoByTokenResponse creates a response to parse from QueryInfoByToken response
func CreateQueryInfoByTokenResponse() (response *QueryInfoByTokenResponse) {
	response = &QueryInfoByTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
