package ons

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

// OnsTopicUpdate invokes the ons.OnsTopicUpdate API synchronously
// api document: https://help.aliyun.com/api/ons/onstopicupdate.html
func (client *Client) OnsTopicUpdate(request *OnsTopicUpdateRequest) (response *OnsTopicUpdateResponse, err error) {
	response = CreateOnsTopicUpdateResponse()
	err = client.DoAction(request, response)
	return
}

// OnsTopicUpdateWithChan invokes the ons.OnsTopicUpdate API asynchronously
// api document: https://help.aliyun.com/api/ons/onstopicupdate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsTopicUpdateWithChan(request *OnsTopicUpdateRequest) (<-chan *OnsTopicUpdateResponse, <-chan error) {
	responseChan := make(chan *OnsTopicUpdateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsTopicUpdate(request)
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

// OnsTopicUpdateWithCallback invokes the ons.OnsTopicUpdate API asynchronously
// api document: https://help.aliyun.com/api/ons/onstopicupdate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsTopicUpdateWithCallback(request *OnsTopicUpdateRequest, callback func(response *OnsTopicUpdateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsTopicUpdateResponse
		var err error
		defer close(result)
		response, err = client.OnsTopicUpdate(request)
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

// OnsTopicUpdateRequest is the request struct for api OnsTopicUpdate
type OnsTopicUpdateRequest struct {
	*requests.RpcRequest
	PreventCache requests.Integer `position:"Query" name:"PreventCache"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	Perm         requests.Integer `position:"Query" name:"Perm"`
	Topic        string           `position:"Query" name:"Topic"`
}

// OnsTopicUpdateResponse is the response struct for api OnsTopicUpdate
type OnsTopicUpdateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
}

// CreateOnsTopicUpdateRequest creates a request to invoke OnsTopicUpdate API
func CreateOnsTopicUpdateRequest() (request *OnsTopicUpdateRequest) {
	request = &OnsTopicUpdateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsTopicUpdate", "ons", "openAPI")
	return
}

// CreateOnsTopicUpdateResponse creates a response to parse from OnsTopicUpdate response
func CreateOnsTopicUpdateResponse() (response *OnsTopicUpdateResponse) {
	response = &OnsTopicUpdateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}