package smartag

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

// DescribeSagRouteList invokes the smartag.DescribeSagRouteList API synchronously
// api document: https://help.aliyun.com/api/smartag/describesagroutelist.html
func (client *Client) DescribeSagRouteList(request *DescribeSagRouteListRequest) (response *DescribeSagRouteListResponse, err error) {
	response = CreateDescribeSagRouteListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSagRouteListWithChan invokes the smartag.DescribeSagRouteList API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesagroutelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSagRouteListWithChan(request *DescribeSagRouteListRequest) (<-chan *DescribeSagRouteListResponse, <-chan error) {
	responseChan := make(chan *DescribeSagRouteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSagRouteList(request)
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

// DescribeSagRouteListWithCallback invokes the smartag.DescribeSagRouteList API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesagroutelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSagRouteListWithCallback(request *DescribeSagRouteListRequest, callback func(response *DescribeSagRouteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSagRouteListResponse
		var err error
		defer close(result)
		response, err = client.DescribeSagRouteList(request)
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

// DescribeSagRouteListRequest is the request struct for api DescribeSagRouteList
type DescribeSagRouteListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
}

// DescribeSagRouteListResponse is the response struct for api DescribeSagRouteList
type DescribeSagRouteListResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Routes    []Route `json:"Routes" xml:"Routes"`
}

// CreateDescribeSagRouteListRequest creates a request to invoke DescribeSagRouteList API
func CreateDescribeSagRouteListRequest() (request *DescribeSagRouteListRequest) {
	request = &DescribeSagRouteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSagRouteList", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSagRouteListResponse creates a response to parse from DescribeSagRouteList response
func CreateDescribeSagRouteListResponse() (response *DescribeSagRouteListResponse) {
	response = &DescribeSagRouteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
