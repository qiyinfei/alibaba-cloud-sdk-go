package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UntagResources invokes the emr.UntagResources API synchronously
// api document: https://help.aliyun.com/api/emr/untagresources.html
func (client *Client) UntagResources(request *UntagResourcesRequest) (response *UntagResourcesResponse, err error) {
	response = CreateUntagResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// UntagResourcesWithChan invokes the emr.UntagResources API asynchronously
// api document: https://help.aliyun.com/api/emr/untagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UntagResourcesWithChan(request *UntagResourcesRequest) (<-chan *UntagResourcesResponse, <-chan error) {
	responseChan := make(chan *UntagResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UntagResources(request)
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

// UntagResourcesWithCallback invokes the emr.UntagResources API asynchronously
// api document: https://help.aliyun.com/api/emr/untagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UntagResourcesWithCallback(request *UntagResourcesRequest, callback func(response *UntagResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UntagResourcesResponse
		var err error
		defer close(result)
		response, err = client.UntagResources(request)
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

// UntagResourcesRequest is the request struct for api UntagResources
type UntagResourcesRequest struct {
	*requests.RpcRequest
	All                  requests.Boolean `position:"Query" name:"All"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceId           *[]string        `position:"Query" name:"ResourceId"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TagKey               *[]string        `position:"Query" name:"TagKey"  type:"Repeated"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
}

// UntagResourcesResponse is the response struct for api UntagResources
type UntagResourcesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUntagResourcesRequest creates a request to invoke UntagResources API
func CreateUntagResourcesRequest() (request *UntagResourcesRequest) {
	request = &UntagResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2019-10-29", "UntagResources", "emr", "openAPI")
	return
}

// CreateUntagResourcesResponse creates a response to parse from UntagResources response
func CreateUntagResourcesResponse() (response *UntagResourcesResponse) {
	response = &UntagResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
