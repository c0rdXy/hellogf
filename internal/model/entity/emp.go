// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Emp is the golang structure for table emp.
type Emp struct {
	Id     uint   `json:"id"      ` // ID
	DeptId uint   `json:"dept_id" ` // 所属部门
	Name   string `json:"name"    ` // 姓名
	Gender int    `json:"gender"  ` // 性别: 0=男 1=女
	Phone  string `json:"phone"   ` // 联系电话
	Email  string `json:"email"   ` // 邮箱
	Avatar string `json:"avatar"  ` // 照片

	Dept  *Dept  `json:"dept" orm:"with:id=dept_id"` // 所属部门
	Hobby *Hobby `json:"hobby" orm:"with:emp_id=id"` // 爱好
}
