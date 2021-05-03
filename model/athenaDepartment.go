// 自动生成模板AthenaDepartment
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type AthenaDepartment struct {
	global.GVA_MODEL
	DepartmentName string `json:"departmentName" form:"departmentName" gorm:"column:department_name;comment:部门名称;type:varchar(255);size:255;"`
	Creator        string `json:"creator" form:"creator" gorm:"column:creator;comment:创建者;type:varchar(255);size:255;"`
	//CreateTime  time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:timestamp;"`
	Updater string `json:"updater" form:"updater" gorm:"column:updater;comment:更新者;type:varchar(255);size:255;"`
	//UpdateTime  time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;type:timestamp;"`
	DelFlag bool `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:删除标志;type:tinyint;"`
}

func (AthenaDepartment) TableName() string {
	return "athena_department"
}
