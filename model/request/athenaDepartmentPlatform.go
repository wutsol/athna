package request

import "gin-vue-admin/model"

type AthenaDepartmentPlatformSearch struct{
    model.AthenaDepartmentPlatform
    PageInfo
}