package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListTagResources invokes the emr.ListTagResources API synchronously
// api document: https://help.aliyun.com/api/emr/listtagresources.html
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (response *ListTagResourcesResponse, err error) {
	response = CreateListTagResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// ListTagResourcesWithChan invokes the emr.ListTagResources API asynchronously
// api document: https://help.aliyun.com/api/emr/listtagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagResourcesWithChan(request *ListTagResourcesRequest) (<-chan *ListTagResourcesResponse, <-chan error) {
	responseChan := make(chan *ListTagResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTagResources(request)
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

// ListTagResourcesWithCallback invokes the emr.ListTagResources API asynchronously
// api document: https://help.aliyun.com/api/emr/listtagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagResourcesWithCallback(request *ListTagResourcesRequest, callback func(response *ListTagResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTagResourcesResponse
		var err error
		defer close(result)
		response, err = client.ListTagResources(request)
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

// ListTagResourcesRequest is the request struct for api ListTagResources
type ListTagResourcesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer       `position:"Query" name:"ResourceOwnerId"`
	ResourceId           *[]string              `position:"Query" name:"ResourceId"  type:"Repeated"`
	ResourceOwnerAccount string                 `position:"Query" name:"ResourceOwnerAccount"`
	NextToken            string                 `position:"Query" name:"NextToken"`
	OwnerAccount         string                 `position:"Query" name:"OwnerAccount"`
	Tag                  *[]ListTagResourcesTag `position:"Query" name:"Tag"  type:"Repeated"`
	OwnerId              requests.Integer       `position:"Query" name:"OwnerId"`
	ResourceType         string                 `position:"Query" name:"ResourceType"`
}

// ListTagResourcesTag is a repeated param struct in ListTagResourcesRequest
type ListTagResourcesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListTagResourcesResponse is the response struct for api ListTagResources
type ListTagResourcesResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	NextToken    string       `json:"NextToken" xml:"NextToken"`
	TagResources TagResources `json:"TagResources" xml:"TagResources"`
}

// CreateListTagResourcesRequest creates a request to invoke ListTagResources API
func CreateListTagResourcesRequest() (request *ListTagResourcesRequest) {
	request = &ListTagResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2019-10-29", "ListTagResources", "emr", "openAPI")
	return
}

// CreateListTagResourcesResponse creates a response to parse from ListTagResources response
func CreateListTagResourcesResponse() (response *ListTagResourcesResponse) {
	response = &ListTagResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
