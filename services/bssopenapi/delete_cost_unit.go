package bssopenapi

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

// DeleteCostUnit invokes the bssopenapi.DeleteCostUnit API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/deletecostunit.html
func (client *Client) DeleteCostUnit(request *DeleteCostUnitRequest) (response *DeleteCostUnitResponse, err error) {
	response = CreateDeleteCostUnitResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCostUnitWithChan invokes the bssopenapi.DeleteCostUnit API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/deletecostunit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCostUnitWithChan(request *DeleteCostUnitRequest) (<-chan *DeleteCostUnitResponse, <-chan error) {
	responseChan := make(chan *DeleteCostUnitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCostUnit(request)
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

// DeleteCostUnitWithCallback invokes the bssopenapi.DeleteCostUnit API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/deletecostunit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCostUnitWithCallback(request *DeleteCostUnitRequest, callback func(response *DeleteCostUnitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCostUnitResponse
		var err error
		defer close(result)
		response, err = client.DeleteCostUnit(request)
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

// DeleteCostUnitRequest is the request struct for api DeleteCostUnit
type DeleteCostUnitRequest struct {
	*requests.RpcRequest
	UnitId   requests.Integer `position:"Query" name:"UnitId"`
	OwnerUid requests.Integer `position:"Query" name:"OwnerUid"`
}

// DeleteCostUnitResponse is the response struct for api DeleteCostUnit
type DeleteCostUnitResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDeleteCostUnitRequest creates a request to invoke DeleteCostUnit API
func CreateDeleteCostUnitRequest() (request *DeleteCostUnitRequest) {
	request = &DeleteCostUnitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "DeleteCostUnit", "bssopenapi", "openAPI")
	return
}

// CreateDeleteCostUnitResponse creates a response to parse from DeleteCostUnit response
func CreateDeleteCostUnitResponse() (response *DeleteCostUnitResponse) {
	response = &DeleteCostUnitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
