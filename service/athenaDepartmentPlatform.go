package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAthenaDepartmentPlatform
//@description: 创建AthenaDepartmentPlatform记录
//@param: athenaDepartmentPlatform model.AthenaDepartmentPlatform
//@return: err error

func CreateAthenaDepartmentPlatform(athenaDepartmentPlatform model.AthenaDepartmentPlatform) (err error) {
	athenaDepartmentPlatform.ID = utils.UniqueId()
	err = global.GVA_DB.Create(&athenaDepartmentPlatform).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAthenaDepartmentPlatform
//@description: 删除AthenaDepartmentPlatform记录
//@param: athenaDepartmentPlatform model.AthenaDepartmentPlatform
//@return: err error

func DeleteAthenaDepartmentPlatform(athenaDepartmentPlatform model.AthenaDepartmentPlatform) (err error) {
	err = global.GVA_DB.Model(&model.AthenaDepartmentPlatform{}).Where("id = ?", athenaDepartmentPlatform.ID).Update("del_flag", 1).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAthenaDepartmentPlatformByIds
//@description: 批量删除AthenaDepartmentPlatform记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAthenaDepartmentPlatformByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Debug().Model(&model.AthenaDepartmentPlatform{}).Where("id in ?", ids.Ids).Update("del_flag", 1).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAthenaDepartmentPlatform
//@description: 更新AthenaDepartmentPlatform记录
//@param: athenaDepartmentPlatform *model.AthenaDepartmentPlatform
//@return: err error

func UpdateAthenaDepartmentPlatform(athenaDepartmentPlatform model.AthenaDepartmentPlatform) (err error) {
	err = global.GVA_DB.Save(&athenaDepartmentPlatform).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAthenaDepartmentPlatform
//@description: 根据id获取AthenaDepartmentPlatform记录
//@param: id uint
//@return: err error, athenaDepartmentPlatform model.AthenaDepartmentPlatform

func GetAthenaDepartmentPlatform(id int64) (err error, athenaDepartmentPlatform model.AthenaDepartmentPlatform) {
	err = global.GVA_DB.Debug().Where("id = ? and del_flag = ?", id, 0).First(&athenaDepartmentPlatform).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAthenaDepartmentPlatformInfoList
//@description: 分页获取AthenaDepartmentPlatform记录
//@param: info request.AthenaDepartmentPlatformSearch
//@return: err error, list interface{}, total int64

func GetAthenaDepartmentPlatformInfoList(info request.AthenaDepartmentPlatformSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.AthenaDepartmentPlatform{}).Where("del_flag = ?", 0)
	var athenaDepartmentPlatforms []model.AthenaDepartmentPlatform
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&athenaDepartmentPlatforms).Error
	return err, athenaDepartmentPlatforms, total
}
