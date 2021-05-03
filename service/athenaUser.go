package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAthenaUser
//@description: 创建AthenaUser记录
//@param: athenaUser model.AthenaUser
//@return: err error

func CreateAthenaUser(athenaUser model.AthenaUser) (err error) {
	athenaUser.ID = utils.UniqueId()
	err = global.GVA_DB.Debug().Create(&athenaUser).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAthenaUser
//@description: 删除AthenaUser记录
//@param: athenaUser model.AthenaUser
//@return: err error

func DeleteAthenaUser(athenaUser model.AthenaUser) (err error) {
	err = global.GVA_DB.Debug().Model(&model.AthenaUser{}).Where("id = ?", athenaUser.ID).Update("del_flag", 1).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAthenaUserByIds
//@description: 批量删除AthenaUser记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAthenaUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Debug().Model(&model.AthenaUser{}).Where("id in ?", ids.Ids).Update("del_flag", 1).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAthenaUser
//@description: 更新AthenaUser记录
//@param: athenaUser *model.AthenaUser
//@return: err error

func UpdateAthenaUser(athenaUser model.AthenaUser) (err error) {
	err = global.GVA_DB.Debug().Save(&athenaUser).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAthenaUser
//@description: 根据id获取AthenaUser记录
//@param: id uint
//@return: err error, athenaUser model.AthenaUser

func GetAthenaUser(id int64) (err error, athenaUser model.AthenaUser) {
	err = global.GVA_DB.Debug().Where("id = ? and del_flag = ?", id, 0).First(&athenaUser).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAthenaUserInfoList
//@description: 分页获取AthenaUser记录
//@param: info request.AthenaUserSearch
//@return: err error, list interface{}, total int64

func GetAthenaUserInfoList(info request.AthenaUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Debug().Model(&model.AthenaUser{}).Where("del_flag = ?", 0)
    var athenaUsers []model.AthenaUser
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&athenaUsers).Error
	return err, athenaUsers, total
}