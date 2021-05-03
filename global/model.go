package global

import (
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	ID        int64          `gorm:"column:id"`
	CreatedAt time.Time      `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdatedAt time.Time      `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	//Creator   string         `gorm:"column:creator" json:"creator"`
	//Updater   string         `gorm:"column:updater" json:"updater"`
	//DelFlag   int            `gorm:"column:del_flag" json:"delFlag"`
}
