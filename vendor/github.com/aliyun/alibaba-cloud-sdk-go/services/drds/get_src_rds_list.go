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

// GetSrcRdsList invokes the drds.GetSrcRdsList API synchronously
// api document: https://help.aliyun.com/api/drds/getsrcrdslist.html
func (client *Client) GetSrcRdsList(request *GetSrcRdsListRequest) (response *GetSrcRdsListResponse, err error) {
	response = CreateGetSrcRdsListResponse()
	err = client.DoAction(request, response)
	return
}

// GetSrcRdsListWithChan invokes the drds.GetSrcRdsList API asynchronously
// api document: https://help.aliyun.com/api/drds/getsrcrdslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSrcRdsListWithChan(request *GetSrcRdsListRequest) (<-chan *GetSrcRdsListResponse, <-chan error) {
	responseChan := make(chan *GetSrcRdsListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSrcRdsList(request)
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

// GetSrcRdsListWithCallback invokes the drds.GetSrcRdsList API asynchronously
// api document: https://help.aliyun.com/api/drds/getsrcrdslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSrcRdsListWithCallback(request *GetSrcRdsListRequest, callback func(response *GetSrcRdsListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSrcRdsListResponse
		var err error
		defer close(result)
		response, err = client.GetSrcRdsList(request)
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

// GetSrcRdsListRequest is the request struct for api GetSrcRdsList
type GetSrcRdsListRequest struct {
	*requests.RpcRequest
	DbName           string                           `position:"Query" name:"DbName"`
	PartitionMapping *[]GetSrcRdsListPartitionMapping `position:"Query" name:"PartitionMapping"  type:"Repeated"`
	DrdsInstanceId   string                           `position:"Query" name:"DrdsInstanceId"`
}

// GetSrcRdsListPartitionMapping is a repeated param struct in GetSrcRdsListRequest
type GetSrcRdsListPartitionMapping struct {
	DbShardValue string `name:"DbShardValue"`
	HotDbName    string `name:"HotDbName"`
	HotTableName string `name:"HotTableName"`
	TbShardValue string `name:"TbShardValue"`
	LogicTable   string `name:"LogicTable"`
}

// GetSrcRdsListResponse is the response struct for api GetSrcRdsList
type GetSrcRdsListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateGetSrcRdsListRequest creates a request to invoke GetSrcRdsList API
func CreateGetSrcRdsListRequest() (request *GetSrcRdsListRequest) {
	request = &GetSrcRdsListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "GetSrcRdsList", "drds", "openAPI")
	return
}

// CreateGetSrcRdsListResponse creates a response to parse from GetSrcRdsList response
func CreateGetSrcRdsListResponse() (response *GetSrcRdsListResponse) {
	response = &GetSrcRdsListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
