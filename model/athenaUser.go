// 自动生成模板AthenaUser
package model

import (
      "gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type AthenaUser struct {
      global.GVA_MODEL
      UserName  string `json:"userName" form:"userName" gorm:"column:user_name;comment:用户名;type:varchar(255);size:255;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:用户密码;type:varchar(255);size:255;"`
      UserType  *bool `json:"userType" form:"userType" gorm:"column:user_type;comment:管理员类型，0普通用户，1超级管理员，2部门管理员，3平台管理员;type:tinyint;"`
      DepartmentName  string `json:"departmentName" form:"departmentName" gorm:"column:department_name;comment:部门;type:varchar(255);size:255;"`
      PlatformName  int `json:"platformName" form:"platformName" gorm:"column:platform_name;comment:平台;type:bigint;size:19;"`
      Creator  string `json:"creator" form:"creator" gorm:"column:creator;comment:创建者;type:varchar(255);size:255;"`
      //CreateTime  time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:timestamp;"`
      Updater  string `json:"updater" form:"updater" gorm:"column:updater;comment:更新者;type:varchar(255);size:255;"`
      //UpdateTime  time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;type:timestamp;"`
      DelFlag  *bool `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:删除标志;type:tinyint;"`
}


func (AthenaUser) TableName() string {
  return "athena_user"
}

