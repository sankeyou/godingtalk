package godingtalk

import (
	"fmt"
	"net/url"
)

type User struct {
	OAPIResponse
	Userid     string
	Name       string
	Mobile     string
	Tel        string
	Remark     string
	Order      int
	IsAdmin    bool
	IsBoss     bool
	DingId     string `json:"dingId"`
	Unionid    string `json:"unionid"`
	IsLeader   bool
	IsSys      bool `json:"is_sys"`
	SysLevel   int  `json:"sys_level"`
	Active     bool
	Department []int
	Position   string
	Email      string
	Avatar     string
	Extattr    interface{}
}

type UserList struct {
	OAPIResponse
	HasMore  bool
	Userlist []User
}

type Department struct {
	OAPIResponse
	Id                    int
	Name                  string
	ParentId              int
	Order                 int
	DeptPerimits          string
	UserPerimits          string
	OuterDept             bool
	OuterPermitDepts      string
	OuterPermitUsers      string
	OrgDeptOwner          string
	DeptManagerUseridList string
}

type DepartmentList struct {
	OAPIResponse
	Departments []Department `json:"department"`
}

// DepartmentList is 获取部门列表
func (c *DingTalkClient) DepartmentList(parentId string) (DepartmentList, error) {
	var data DepartmentList
	params := url.Values{}
	params.Add("lang", "zh_CN")
	params.Add("fetch_child", "true")
	if parentId == "" {
		parentId = "1"
	}
	params.Add("id", parentId)
	err := c.httpRPC("department/list", params, nil, &data)
	return data, err
}

//DepartmentDetail is 获取部门详情
func (c *DingTalkClient) DepartmentDetail(id int) (Department, error) {
	var data Department
	params := url.Values{}
	params.Add("id", fmt.Sprintf("%d", id))
	err := c.httpRPC("department/get", params, nil, &data)
	return data, err
}

//UserList is 获取部门成员
func (c *DingTalkClient) UserList(departmentID int) (UserList, error) {
	var data UserList
	params := url.Values{}
	params.Add("department_id", fmt.Sprintf("%d", departmentID))
	err := c.httpRPC("user/list", params, nil, &data)
	return data, err
}

//CreateChat is
func (c *DingTalkClient) CreateChat(name string, owner string, useridlist []string) (string, error) {
	var data struct {
		OAPIResponse
		Chatid string
	}
	request := map[string]interface{}{
		"name":       name,
		"owner":      owner,
		"useridlist": useridlist,
	}
	err := c.httpRPC("chat/create", nil, request, &data)
	return data.Chatid, err
}

//UserInfoByCode 校验免登录码并换取用户身份
func (c *DingTalkClient) UserInfoByCode(code string) (User, error) {
	var data User
	params := url.Values{}
	params.Add("code", code)
	err := c.httpRPC("user/getuserinfo", params, nil, &data)
	params.Add("userid", data.Userid)
	err = c.httpRPC("user/get", params, nil, &data)
	return data, err
}

//UserInfoByUserId userid获取用户分身
func (c *DingTalkClient) UserInfoByUserId(userId string) (User, error) {
	params := url.Values{}
	params.Add("userid", userId)
	user := User{}
	err := c.httpRPC("user/get", params, nil, &user)
	return user, err
}

//UseridByUnionId 通过UnionId获取玩家Userid
func (c *DingTalkClient) UseridByUnionId(unionid string) (string, error) {
	var data struct {
		OAPIResponse
		UserID string `json:"userid"`
	}

	params := url.Values{}
	params.Add("unionid", unionid)
	err := c.httpRPC("user/getUseridByUnionid", params, nil, &data)
	if err != nil {
		return "", err
	}

	return data.UserID, err
}
