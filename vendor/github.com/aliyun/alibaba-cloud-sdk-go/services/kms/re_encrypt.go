package kms

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

// ReEncrypt invokes the kms.ReEncrypt API synchronously
// api document: https://help.aliyun.com/api/kms/reencrypt.html
func (client *Client) ReEncrypt(request *ReEncryptRequest) (response *ReEncryptResponse, err error) {
	response = CreateReEncryptResponse()
	err = client.DoAction(request, response)
	return
}

// ReEncryptWithChan invokes the kms.ReEncrypt API asynchronously
// api document: https://help.aliyun.com/api/kms/reencrypt.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReEncryptWithChan(request *ReEncryptRequest) (<-chan *ReEncryptResponse, <-chan error) {
	responseChan := make(chan *ReEncryptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReEncrypt(request)
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

// ReEncryptWithCallback invokes the kms.ReEncrypt API asynchronously
// api document: https://help.aliyun.com/api/kms/reencrypt.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReEncryptWithCallback(request *ReEncryptRequest, callback func(response *ReEncryptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReEncryptResponse
		var err error
		defer close(result)
		response, err = client.ReEncrypt(request)
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

// ReEncryptRequest is the request struct for api ReEncrypt
type ReEncryptRequest struct {
	*requests.RpcRequest
	DestinationEncryptionContext string `position:"Query" name:"DestinationEncryptionContext"`
	SourceEncryptionAlgorithm    string `position:"Query" name:"SourceEncryptionAlgorithm"`
	SourceKeyVersionId           string `position:"Query" name:"SourceKeyVersionId"`
	DestinationKeyId             string `position:"Query" name:"DestinationKeyId"`
	SourceKeyId                  string `position:"Query" name:"SourceKeyId"`
	SourceEncryptionContext      string `position:"Query" name:"SourceEncryptionContext"`
	CiphertextBlob               string `position:"Query" name:"CiphertextBlob"`
}

// ReEncryptResponse is the response struct for api ReEncrypt
type ReEncryptResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	KeyId          string `json:"KeyId" xml:"KeyId"`
	KeyVersionId   string `json:"KeyVersionId" xml:"KeyVersionId"`
	CiphertextBlob string `json:"CiphertextBlob" xml:"CiphertextBlob"`
}

// CreateReEncryptRequest creates a request to invoke ReEncrypt API
func CreateReEncryptRequest() (request *ReEncryptRequest) {
	request = &ReEncryptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "ReEncrypt", "kms-service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReEncryptResponse creates a response to parse from ReEncrypt response
func CreateReEncryptResponse() (response *ReEncryptResponse) {
	response = &ReEncryptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
