package vpc

import (
	"time"

	"github.com/hdksky/ksyungo/common"
)

type VpcType struct {
	CreateTime time.Time
	VpcName    string
	VpcId      string
	CidrBlock  string
	IsDefault  bool
}

type DescribeVpcsResponse struct {
	common.Response
	VpcSet struct {
		Item []VpcType
	}
}

// DescribeVpcs describe vpc info
// You can read doc at https://docs.ksyun.com/read/latest/56/_book/Action/Vpcs/DescribeVpcs.html
func (c *Client) DescribeVpcs(vpcIds []string) (*DescribeVpcsResponse, error) {
	response := DescribeVpcsResponse{}
	err := c.Invoke("DescribeVpcs", vpcIds, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

type Kv struct {
	Name  string
	Value string
}

type DescribeSubnetsArgs struct {
	SubnetId []string
	Filter   []Kv
}

type SubnetType struct {
	CreateTime   string
	VpcId        string
	SubnetId     string
	SubnetType   string
	SubnetName   string
	CidrBlock    string
	DhcpIpFrom   string
	DhcpIpTo     string
	GatewayIp    string
	Dns1         string
	Dns2         string
	NetworkAclId string
	NatId        string
}

type DescribeSubnetsResonse struct {
	common.Response
	SubnetSet struct {
		Item []SubnetType
	}
}

// DescribeSubnets describe subnets info
// You can read doc at https://docs.ksyun.com/read/latest/56/_book/Action/Subnets/DescribeSubnets.html
func (c *Client) DescribeSubnets(args *DescribeSubnetsArgs) ([]SubnetType, error) {
	response := DescribeSubnetsResonse{}
	err := c.Invoke("DescribeSubnets", args, &response)
	if err == nil {
		return response.SubnetSet.Item, nil
	}
	return nil, err
}

type DescribeSecurityGroupsArgs struct {
	SecurityGroupId []string
	Filter          []Kv
}

type SecurityGroupType struct {
	CreateTime            string
	VpcId                 string
	SecurityGroupName     string
	SecurityGroupId       string
	SecurityGroupType     string
	SecurityGroupEntrySet string
}

type DescribeSecurityGroupsResponse struct {
	common.Response
	SecurityGroupSet struct {
		Item []SecurityGroupType
	}
}

// DescribeSecurityGroups describe security groups
// You can read doc at https://docs.ksyun.com/read/latest/56/_book/Action/SecurityGroups/DescribeSecurityGroups.html
func (c *Client) DescribeSecurityGroups(args *DescribeSecurityGroupsArgs) ([]SecurityGroupType, error) {
	response := DescribeSecurityGroupsResponse{}
	err := c.Invoke("DescribeSecurityGroups", args, &response)
	if err == nil {
		return response.SecurityGroupSet.Item, nil
	}
	return nil, err
}