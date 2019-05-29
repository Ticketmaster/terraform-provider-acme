package alidns

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

// DnsProduct is a nested struct in alidns response
type DnsProduct struct {
	InstanceId            string `json:"InstanceId" xml:"InstanceId"`
	VersionCode           string `json:"VersionCode" xml:"VersionCode"`
	VersionName           string `json:"VersionName" xml:"VersionName"`
	StartTime             string `json:"StartTime" xml:"StartTime"`
	EndTime               string `json:"EndTime" xml:"EndTime"`
	StartTimestamp        int    `json:"StartTimestamp" xml:"StartTimestamp"`
	EndTimestamp          int    `json:"EndTimestamp" xml:"EndTimestamp"`
	Domain                string `json:"Domain" xml:"Domain"`
	BindCount             int    `json:"BindCount" xml:"BindCount"`
	BindUsedCount         int    `json:"BindUsedCount" xml:"BindUsedCount"`
	TTLMinValue           int    `json:"TTLMinValue" xml:"TTLMinValue"`
	SubDomainLevel        int    `json:"SubDomainLevel" xml:"SubDomainLevel"`
	DnsSLBCount           int    `json:"DnsSLBCount" xml:"DnsSLBCount"`
	URLForwardCount       int    `json:"URLForwardCount" xml:"URLForwardCount"`
	DDosDefendFlow        int    `json:"DDosDefendFlow" xml:"DDosDefendFlow"`
	DDosDefendQuery       int    `json:"DDosDefendQuery" xml:"DDosDefendQuery"`
	OverseaDDosDefendFlow int    `json:"OverseaDDosDefendFlow" xml:"OverseaDDosDefendFlow"`
	SearchEngineLines     string `json:"SearchEngineLines" xml:"SearchEngineLines"`
	ISPLines              string `json:"ISPLines" xml:"ISPLines"`
	ISPRegionLines        string `json:"ISPRegionLines" xml:"ISPRegionLines"`
	OverseaLine           string `json:"OverseaLine" xml:"OverseaLine"`
	MonitorNodeCount      int    `json:"MonitorNodeCount" xml:"MonitorNodeCount"`
	MonitorFrequency      int    `json:"MonitorFrequency" xml:"MonitorFrequency"`
	MonitorTaskCount      int    `json:"MonitorTaskCount" xml:"MonitorTaskCount"`
	RegionLines           bool   `json:"RegionLines" xml:"RegionLines"`
	Gslb                  bool   `json:"Gslb" xml:"Gslb"`
	InClean               bool   `json:"InClean" xml:"InClean"`
	InBlackHole           bool   `json:"InBlackHole" xml:"InBlackHole"`
	BindDomainCount       int    `json:"BindDomainCount" xml:"BindDomainCount"`
	BindDomainUsedCount   int    `json:"BindDomainUsedCount" xml:"BindDomainUsedCount"`
	DnsSecurity           string `json:"DnsSecurity" xml:"DnsSecurity"`
}
