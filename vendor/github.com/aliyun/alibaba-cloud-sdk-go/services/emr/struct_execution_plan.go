package emr

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

// ExecutionPlan is a nested struct in emr response
type ExecutionPlan struct {
	BizId                string `json:"BizId" xml:"BizId"`
	Name                 string `json:"Name" xml:"Name"`
	Status               int    `json:"Status" xml:"Status"`
	LastExeStatus        int    `json:"LastExeStatus" xml:"LastExeStatus"`
	IsCreateCluster      bool   `json:"IsCreateCluster" xml:"IsCreateCluster"`
	IsInterruptWhenError bool   `json:"IsInterruptWhenError" xml:"IsInterruptWhenError"`
	IsCycle              bool   `json:"IsCycle" xml:"IsCycle"`
	ScheduleCycle        string `json:"ScheduleCycle" xml:"ScheduleCycle"`
	RegionId             string `json:"RegionId" xml:"RegionId"`
	CronExpr             string `json:"CronExpr" xml:"CronExpr"`
	UtcCreateTimestamp   int64  `json:"UtcCreateTimestamp" xml:"UtcCreateTimestamp"`
	UtcModifiedTimestamp int64  `json:"UtcModifiedTimestamp" xml:"UtcModifiedTimestamp"`
	Version              string `json:"Version" xml:"Version"`
	ClusterTemplateId    string `json:"ClusterTemplateId" xml:"ClusterTemplateId"`
}
