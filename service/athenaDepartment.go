package service

import (
	athenaModel "athena/model"
	"context"
	"gin-vue-admin/model/request"
	"gin-vue-admin/rpc/athena"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAthenaDepartment
//@description: 创建AthenaDepartment记录
//@param: athenaDepartment model.AthenaDepartment
//@return: err error

func CreateAthenaDepartment(ctx context.Context, athenaDepartment athenaModel.Department) (err error) {
	//athenaDepartment.ID = utils.UniqueId()
	//err = global.GVA_DB.Debug().Create(&athenaDepartment).Error
	//return err
	err = athena.AthenaRpc.CreateDepartment(ctx, athenaDepartment)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAthenaDepartment
//@description: 删除AthenaDepartment记录
//@param: athenaDepartment model.AthenaDepartment
//@return: err error

func DeleteAthenaDepartment(ctx context.Context, athenaDepartment athenaModel.Department) (err error) {
	//err = global.GVA_DB.Debug().Model(&model.AthenaDepartment{}).Where("id = ?", athenaDepartment.ID).Update("del_flag", 1).Error
	//return err
	req := athenaModel.IdsReq{
		Ids: []int{int(athenaDepartment.ID)},
	}
	err = athena.AthenaRpc.BatchDeleteDepartmentByIds(ctx, req)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAthenaDepartmentByIds
//@description: 批量删除AthenaDepartment记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAthenaDepartmentByIds(ctx context.Context, req athenaModel.IdsReq) (err error) {
	//err = global.GVA_DB.Debug().Model(&model.AthenaDepartment{}).Where("id in ?", ids.Ids).Update("del_flag", 1).Error
	//return err
	//req := athenaModel.IdsReq{
	//	Ids: ids.Ids,
	//}
	err = athena.AthenaRpc.BatchDeleteDepartmentByIds(ctx, req)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAthenaDepartment
//@description: 更新AthenaDepartment记录
//@param: athenaDepartment *model.AthenaDepartment
//@return: err error

func UpdateAthenaDepartment(ctx context.Context, athenaDepartment athenaModel.Department) (err error) {
	//err = global.GVA_DB.Debug().Save(&athenaDepartment).Error
	//return err
	err = athena.AthenaRpc.UpdateDepartment(ctx, athenaDepartment)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAthenaDepartment
//@description: 根据id获取AthenaDepartment记录
//@param: id uint
//@return: err error, athenaDepartment model.AthenaDepartment

func GetAthenaDepartment(ctx context.Context, id int) (err error, athenaDepartment *athenaModel.Department) {
	//err = global.GVA_DB.Debug().Where("id = ? and del_flag = ?", id, 0).First(&athenaDepartment).Error
	//return
	req := athenaModel.IdReq{
		Id: int(id),
	}
	athenaDepartment, err = athena.AthenaRpc.GetDepartmentById(ctx, req)
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAthenaDepartmentInfoList
//@description: 分页获取AthenaDepartment记录
//@param: info request.AthenaDepartmentSearch
//@return: err error, list interface{}, total int64

func GetAthenaDepartmentInfoList(ctx context.Context, info request.AthenaDepartmentSearch) (err error, list interface{}, total int64) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
	//// 创建db
	//db := global.GVA_DB.Debug().Model(&model.AthenaDepartment{}).Where("del_flag = ?", 0)
	//var athenaDepartments []model.AthenaDepartment
	//// 如果有条件搜索 下方会自动创建搜索语句
	//err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&athenaDepartments).Error
	//return err, athenaDepartments, total

	req := athenaModel.GetDepartmentListReq{
		PageInfo: athenaModel.PageInfo{
			PageNo:   info.Page,
			PageSize: info.PageSize,
		},
	}
	if len(info.DepartmentName) != 0 {
		req.DepartmentName = &info.DepartmentName
	}
	total, list, err = athena.AthenaRpc.GetDepartmentList(ctx, req)
	return
}
