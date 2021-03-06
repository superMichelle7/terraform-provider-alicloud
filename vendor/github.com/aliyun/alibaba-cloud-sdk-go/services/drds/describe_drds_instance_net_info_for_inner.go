package drds

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

// DescribeDrdsInstanceNetInfoForInner invokes the drds.DescribeDrdsInstanceNetInfoForInner API synchronously
// api document: https://help.aliyun.com/api/drds/describedrdsinstancenetinfoforinner.html
func (client *Client) DescribeDrdsInstanceNetInfoForInner(request *DescribeDrdsInstanceNetInfoForInnerRequest) (response *DescribeDrdsInstanceNetInfoForInnerResponse, err error) {
	response = CreateDescribeDrdsInstanceNetInfoForInnerResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsInstanceNetInfoForInnerWithChan invokes the drds.DescribeDrdsInstanceNetInfoForInner API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsinstancenetinfoforinner.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsInstanceNetInfoForInnerWithChan(request *DescribeDrdsInstanceNetInfoForInnerRequest) (<-chan *DescribeDrdsInstanceNetInfoForInnerResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsInstanceNetInfoForInnerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsInstanceNetInfoForInner(request)
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

// DescribeDrdsInstanceNetInfoForInnerWithCallback invokes the drds.DescribeDrdsInstanceNetInfoForInner API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsinstancenetinfoforinner.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsInstanceNetInfoForInnerWithCallback(request *DescribeDrdsInstanceNetInfoForInnerRequest, callback func(response *DescribeDrdsInstanceNetInfoForInnerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsInstanceNetInfoForInnerResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsInstanceNetInfoForInner(request)
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

// DescribeDrdsInstanceNetInfoForInnerRequest is the request struct for api DescribeDrdsInstanceNetInfoForInner
type DescribeDrdsInstanceNetInfoForInnerRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// DescribeDrdsInstanceNetInfoForInnerResponse is the response struct for api DescribeDrdsInstanceNetInfoForInner
type DescribeDrdsInstanceNetInfoForInnerResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	DrdsInstanceId string   `json:"DrdsInstanceId" xml:"DrdsInstanceId"`
	NetworkType    string   `json:"NetworkType" xml:"NetworkType"`
	NetInfos       NetInfos `json:"NetInfos" xml:"NetInfos"`
}

// CreateDescribeDrdsInstanceNetInfoForInnerRequest creates a request to invoke DescribeDrdsInstanceNetInfoForInner API
func CreateDescribeDrdsInstanceNetInfoForInnerRequest() (request *DescribeDrdsInstanceNetInfoForInnerRequest) {
	request = &DescribeDrdsInstanceNetInfoForInnerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsInstanceNetInfoForInner", "drds", "openAPI")
	return
}

// CreateDescribeDrdsInstanceNetInfoForInnerResponse creates a response to parse from DescribeDrdsInstanceNetInfoForInner response
func CreateDescribeDrdsInstanceNetInfoForInnerResponse() (response *DescribeDrdsInstanceNetInfoForInnerResponse) {
	response = &DescribeDrdsInstanceNetInfoForInnerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
