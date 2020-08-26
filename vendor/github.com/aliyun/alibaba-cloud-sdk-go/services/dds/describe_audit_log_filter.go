package dds

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

// DescribeAuditLogFilter invokes the dds.DescribeAuditLogFilter API synchronously
// api document: https://help.aliyun.com/api/dds/describeauditlogfilter.html
func (client *Client) DescribeAuditLogFilter(request *DescribeAuditLogFilterRequest) (response *DescribeAuditLogFilterResponse, err error) {
	response = CreateDescribeAuditLogFilterResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAuditLogFilterWithChan invokes the dds.DescribeAuditLogFilter API asynchronously
// api document: https://help.aliyun.com/api/dds/describeauditlogfilter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAuditLogFilterWithChan(request *DescribeAuditLogFilterRequest) (<-chan *DescribeAuditLogFilterResponse, <-chan error) {
	responseChan := make(chan *DescribeAuditLogFilterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAuditLogFilter(request)
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

// DescribeAuditLogFilterWithCallback invokes the dds.DescribeAuditLogFilter API asynchronously
// api document: https://help.aliyun.com/api/dds/describeauditlogfilter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAuditLogFilterWithCallback(request *DescribeAuditLogFilterRequest, callback func(response *DescribeAuditLogFilterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAuditLogFilterResponse
		var err error
		defer close(result)
		response, err = client.DescribeAuditLogFilter(request)
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

// DescribeAuditLogFilterRequest is the request struct for api DescribeAuditLogFilter
type DescribeAuditLogFilterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	RoleType             string           `position:"Query" name:"RoleType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeAuditLogFilterResponse is the response struct for api DescribeAuditLogFilter
type DescribeAuditLogFilterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Filter    string `json:"Filter" xml:"Filter"`
	RoleType  string `json:"RoleType" xml:"RoleType"`
}

// CreateDescribeAuditLogFilterRequest creates a request to invoke DescribeAuditLogFilter API
func CreateDescribeAuditLogFilterRequest() (request *DescribeAuditLogFilterRequest) {
	request = &DescribeAuditLogFilterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeAuditLogFilter", "Dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAuditLogFilterResponse creates a response to parse from DescribeAuditLogFilter response
func CreateDescribeAuditLogFilterResponse() (response *DescribeAuditLogFilterResponse) {
	response = &DescribeAuditLogFilterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
