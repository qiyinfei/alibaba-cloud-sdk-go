package aliyuncvc

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

// CreateMeeting invokes the aliyuncvc.CreateMeeting API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/createmeeting.html
func (client *Client) CreateMeeting(request *CreateMeetingRequest) (response *CreateMeetingResponse, err error) {
	response = CreateCreateMeetingResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMeetingWithChan invokes the aliyuncvc.CreateMeeting API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/createmeeting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMeetingWithChan(request *CreateMeetingRequest) (<-chan *CreateMeetingResponse, <-chan error) {
	responseChan := make(chan *CreateMeetingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMeeting(request)
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

// CreateMeetingWithCallback invokes the aliyuncvc.CreateMeeting API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/createmeeting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMeetingWithCallback(request *CreateMeetingRequest, callback func(response *CreateMeetingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMeetingResponse
		var err error
		defer close(result)
		response, err = client.CreateMeeting(request)
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

// CreateMeetingRequest is the request struct for api CreateMeeting
type CreateMeetingRequest struct {
	*requests.RpcRequest
	MeetingName string `position:"Body" name:"MeetingName"`
	UserId      string `position:"Body" name:"UserId"`
}

// CreateMeetingResponse is the response struct for api CreateMeeting
type CreateMeetingResponse struct {
	*responses.BaseResponse
	ErrorCode   int         `json:"ErrorCode" xml:"ErrorCode"`
	Message     string      `json:"Message" xml:"Message"`
	Success     bool        `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	MeetingInfo MeetingInfo `json:"MeetingInfo" xml:"MeetingInfo"`
}

// CreateCreateMeetingRequest creates a request to invoke CreateMeeting API
func CreateCreateMeetingRequest() (request *CreateMeetingRequest) {
	request = &CreateMeetingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-09-19", "CreateMeeting", "aliyuncvc", "openAPI")
	return
}

// CreateCreateMeetingResponse creates a response to parse from CreateMeeting response
func CreateCreateMeetingResponse() (response *CreateMeetingResponse) {
	response = &CreateMeetingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
