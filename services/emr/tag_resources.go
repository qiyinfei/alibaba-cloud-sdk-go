package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// TagResources invokes the emr.TagResources API synchronously
// api document: https://help.aliyun.com/api/emr/tagresources.html
func (client *Client) TagResources(request *TagResourcesRequest) (response *TagResourcesResponse, err error) {
	response = CreateTagResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// TagResourcesWithChan invokes the emr.TagResources API asynchronously
// api document: https://help.aliyun.com/api/emr/tagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TagResourcesWithChan(request *TagResourcesRequest) (<-chan *TagResourcesResponse, <-chan error) {
	responseChan := make(chan *TagResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TagResources(request)
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

// TagResourcesWithCallback invokes the emr.TagResources API asynchronously
// api document: https://help.aliyun.com/api/emr/tagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TagResourcesWithCallback(request *TagResourcesRequest, callback func(response *TagResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TagResourcesResponse
		var err error
		defer close(result)
		response, err = client.TagResources(request)
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

// TagResourcesRequest is the request struct for api TagResources
type TagResourcesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer   `position:"Query" name:"ResourceOwnerId"`
	ResourceId           *[]string          `position:"Query" name:"ResourceId"  type:"Repeated"`
	ResourceOwnerAccount string             `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string             `position:"Query" name:"OwnerAccount"`
	Tag                  *[]TagResourcesTag `position:"Query" name:"Tag"  type:"Repeated"`
	OwnerId              requests.Integer   `position:"Query" name:"OwnerId"`
	ResourceType         string             `position:"Query" name:"ResourceType"`
}

// TagResourcesTag is a repeated param struct in TagResourcesRequest
type TagResourcesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// TagResourcesResponse is the response struct for api TagResources
type TagResourcesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateTagResourcesRequest creates a request to invoke TagResources API
func CreateTagResourcesRequest() (request *TagResourcesRequest) {
	request = &TagResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2019-10-29", "TagResources", "emr", "openAPI")
	return
}

// CreateTagResourcesResponse creates a response to parse from TagResources response
func CreateTagResourcesResponse() (response *TagResourcesResponse) {
	response = &TagResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
