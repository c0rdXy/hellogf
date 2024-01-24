// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Book is the golang structure for table book.
type Book struct {
	Id          uint        `json:"id"           ` // ID
	Name        string      `json:"name"         ` // 书名
	Author      string      `json:"author"       ` // 作者
	Price       float64     `json:"price"        ` // 价格
	PublishTime *gtime.Time `json:"publish_time" ` // 出版时间
	CreatedAt   *gtime.Time `json:"created_at"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"   ` // 更新时间
	DeletedAt   *gtime.Time `json:"deleted_at"   ` // 删除时间
}
