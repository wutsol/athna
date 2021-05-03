package request

import "gin-vue-admin/model"

type AthenaDepartmentSearch struct{
    model.AthenaDepartment
    PageInfo
}