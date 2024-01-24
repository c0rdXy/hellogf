// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"         ` // ID
	Username  string      `json:"username"   ` // 用户名
	Nickname  string      `json:"nickname"   ` // 昵称
	Password  string      `json:"password"   ` // 密码
	Avatar    string      `json:"avatar"     ` // 头像
	CreatedAt *gtime.Time `json:"created_at" ` // 创建时间
}
