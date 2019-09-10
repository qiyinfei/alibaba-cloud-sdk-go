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

// Event is a nested struct in ros20190910 response
type Event struct {
	CreateTime         string `json:"CreateTime" xml:"CreateTime"`
	EventId            string `json:"EventId" xml:"EventId"`
	LogicalResourceId  string `json:"LogicalResourceId" xml:"LogicalResourceId"`
	PhysicalResourceId string `json:"PhysicalResourceId" xml:"PhysicalResourceId"`
	ResourceType       string `json:"ResourceType" xml:"ResourceType"`
	StackId            string `json:"StackId" xml:"StackId"`
	StackName          string `json:"StackName" xml:"StackName"`
	Status             string `json:"Status" xml:"Status"`
	StatusReason       string `json:"StatusReason" xml:"StatusReason"`
}
