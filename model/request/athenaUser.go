package request

import "gin-vue-admin/model"

type AthenaUserSearch struct{
    model.AthenaUser
    PageInfo
}