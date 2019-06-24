package vod

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

// BatchStartVodDomain invokes the vod.BatchStartVodDomain API synchronously
// api document: https://help.aliyun.com/api/vod/batchstartvoddomain.html
func (client *Client) BatchStartVodDomain(request *BatchStartVodDomainRequest) (response *BatchStartVodDomainResponse, err error) {
	response = CreateBatchStartVodDomainResponse()
	err = client.DoAction(request, response)
	return
}

// BatchStartVodDomainWithChan invokes the vod.BatchStartVodDomain API asynchronously
// api document: https://help.aliyun.com/api/vod/batchstartvoddomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchStartVodDomainWithChan(request *BatchStartVodDomainRequest) (<-chan *BatchStartVodDomainResponse, <-chan error) {
	responseChan := make(chan *BatchStartVodDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchStartVodDomain(request)
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

// BatchStartVodDomainWithCallback invokes the vod.BatchStartVodDomain API asynchronously
// api document: https://help.aliyun.com/api/vod/batchstartvoddomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchStartVodDomainWithCallback(request *BatchStartVodDomainRequest, callback func(response *BatchStartVodDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchStartVodDomainResponse
		var err error
		defer close(result)
		response, err = client.BatchStartVodDomain(request)
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

// BatchStartVodDomainRequest is the request struct for api BatchStartVodDomain
type BatchStartVodDomainRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainNames   string           `position:"Query" name:"DomainNames"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// BatchStartVodDomainResponse is the response struct for api BatchStartVodDomain
type BatchStartVodDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchStartVodDomainRequest creates a request to invoke BatchStartVodDomain API
func CreateBatchStartVodDomainRequest() (request *BatchStartVodDomainRequest) {
	request = &BatchStartVodDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "BatchStartVodDomain", "vod", "openAPI")
	return
}

// CreateBatchStartVodDomainResponse creates a response to parse from BatchStartVodDomain response
func CreateBatchStartVodDomainResponse() (response *BatchStartVodDomainResponse) {
	response = &BatchStartVodDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
