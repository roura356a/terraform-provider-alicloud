package cdn

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

// DeleteFCTrigger invokes the cdn.DeleteFCTrigger API synchronously
// api document: https://help.aliyun.com/api/cdn/deletefctrigger.html
func (client *Client) DeleteFCTrigger(request *DeleteFCTriggerRequest) (response *DeleteFCTriggerResponse, err error) {
	response = CreateDeleteFCTriggerResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFCTriggerWithChan invokes the cdn.DeleteFCTrigger API asynchronously
// api document: https://help.aliyun.com/api/cdn/deletefctrigger.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFCTriggerWithChan(request *DeleteFCTriggerRequest) (<-chan *DeleteFCTriggerResponse, <-chan error) {
	responseChan := make(chan *DeleteFCTriggerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFCTrigger(request)
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

// DeleteFCTriggerWithCallback invokes the cdn.DeleteFCTrigger API asynchronously
// api document: https://help.aliyun.com/api/cdn/deletefctrigger.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFCTriggerWithCallback(request *DeleteFCTriggerRequest, callback func(response *DeleteFCTriggerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFCTriggerResponse
		var err error
		defer close(result)
		response, err = client.DeleteFCTrigger(request)
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

// DeleteFCTriggerRequest is the request struct for api DeleteFCTrigger
type DeleteFCTriggerRequest struct {
	*requests.RpcRequest
	TriggerARN string           `position:"Query" name:"TriggerARN"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteFCTriggerResponse is the response struct for api DeleteFCTrigger
type DeleteFCTriggerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteFCTriggerRequest creates a request to invoke DeleteFCTrigger API
func CreateDeleteFCTriggerRequest() (request *DeleteFCTriggerRequest) {
	request = &DeleteFCTriggerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DeleteFCTrigger", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteFCTriggerResponse creates a response to parse from DeleteFCTrigger response
func CreateDeleteFCTriggerResponse() (response *DeleteFCTriggerResponse) {
	response = &DeleteFCTriggerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
