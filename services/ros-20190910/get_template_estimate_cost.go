package ros20190910

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

// GetTemplateEstimateCost invokes the ros20190910.GetTemplateEstimateCost API synchronously
// api document: https://help.aliyun.com/api/ros20190910/gettemplateestimatecost.html
func (client *Client) GetTemplateEstimateCost(request *GetTemplateEstimateCostRequest) (response *GetTemplateEstimateCostResponse, err error) {
	response = CreateGetTemplateEstimateCostResponse()
	err = client.DoAction(request, response)
	return
}

// GetTemplateEstimateCostWithChan invokes the ros20190910.GetTemplateEstimateCost API asynchronously
// api document: https://help.aliyun.com/api/ros20190910/gettemplateestimatecost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTemplateEstimateCostWithChan(request *GetTemplateEstimateCostRequest) (<-chan *GetTemplateEstimateCostResponse, <-chan error) {
	responseChan := make(chan *GetTemplateEstimateCostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTemplateEstimateCost(request)
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

// GetTemplateEstimateCostWithCallback invokes the ros20190910.GetTemplateEstimateCost API asynchronously
// api document: https://help.aliyun.com/api/ros20190910/gettemplateestimatecost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTemplateEstimateCostWithCallback(request *GetTemplateEstimateCostRequest, callback func(response *GetTemplateEstimateCostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTemplateEstimateCostResponse
		var err error
		defer close(result)
		response, err = client.GetTemplateEstimateCost(request)
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

// GetTemplateEstimateCostRequest is the request struct for api GetTemplateEstimateCost
type GetTemplateEstimateCostRequest struct {
	*requests.RpcRequest
	ClientToken  string                               `position:"Query" name:"ClientToken"`
	TemplateBody string                               `position:"Query" name:"TemplateBody"`
	TemplateURL  string                               `position:"Query" name:"TemplateURL"`
	Parameters   *[]GetTemplateEstimateCostParameters `position:"Query" name:"Parameters"  type:"Repeated"`
}

// GetTemplateEstimateCostParameters is a repeated param struct in GetTemplateEstimateCostRequest
type GetTemplateEstimateCostParameters struct {
	ParameterValue string `name:"ParameterValue"`
	ParameterKey   string `name:"ParameterKey"`
}

// GetTemplateEstimateCostResponse is the response struct for api GetTemplateEstimateCost
type GetTemplateEstimateCostResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Resources map[string]interface{} `json:"Resources" xml:"Resources"`
}

// CreateGetTemplateEstimateCostRequest creates a request to invoke GetTemplateEstimateCost API
func CreateGetTemplateEstimateCostRequest() (request *GetTemplateEstimateCostRequest) {
	request = &GetTemplateEstimateCostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ros20190910", "2019-09-10", "GetTemplateEstimateCost", "ros20190910", "openAPI")
	return
}

// CreateGetTemplateEstimateCostResponse creates a response to parse from GetTemplateEstimateCost response
func CreateGetTemplateEstimateCostResponse() (response *GetTemplateEstimateCostResponse) {
	response = &GetTemplateEstimateCostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
