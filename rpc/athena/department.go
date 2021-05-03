package athena

import (
	"context"
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/rpc"

	athenaModel "athena/model"
	"go.uber.org/zap"
)

type athenaRpc struct{}

var AthenaRpc athenaRpc

func (r athenaRpc) CreateDepartment(ctx context.Context, req athenaModel.Department) error {
	resp := new(athenaModel.BaseResp)
	err := rpc.AthenaCli.Call(ctx, "CreateDepartment", req, resp)
	if err != nil {
		global.GVA_LOG.Error("call CreateDepartment fail", zap.Any("err", err))
		return err
	}
	if resp.Code != 0 {
		global.GVA_LOG.Error("call CreateDepartment status invalid", zap.Any("code", resp.Code), zap.Any("msg", resp.Message))
		return errors.New("CreateDepartment resp code invalid")
	}
	return nil
}

func (r athenaRpc) BatchDeleteDepartmentByIds(ctx context.Context, req athenaModel.IdsReq) error {
	resp := new(athenaModel.BaseResp)
	err := rpc.AthenaCli.Call(ctx, "BatchDeleteDepartmentByIds", req, resp)
	if err != nil {
		global.GVA_LOG.Error("call BatchDeleteDepartmentByIds fail", zap.Any("err", err))
		return err
	}
	if resp.Code != 0 {
		global.GVA_LOG.Error("call BatchDeleteDepartmentByIds status invalid", zap.Any("code", resp.Code), zap.Any("msg", resp.Message))
		return errors.New("BatchDeleteDepartmentByIds resp code invalid")
	}
	return nil
}

func (r athenaRpc) UpdateDepartment(ctx context.Context, req athenaModel.Department) error {
	resp := new(athenaModel.BaseResp)
	err := rpc.AthenaCli.Call(ctx, "UpdateDepartment", req, resp)
	if err != nil {
		global.GVA_LOG.Error("call UpdateDepartment fail", zap.Any("err", err))
		return err
	}
	if resp.Code != 0 {
		global.GVA_LOG.Error("call UpdateDepartment status invalid", zap.Any("code", resp.Code), zap.Any("msg", resp.Message))
		return errors.New("UpdateDepartment resp code invalid")
	}
	return nil
}

func (r athenaRpc) GetDepartmentById(ctx context.Context, req athenaModel.IdReq) (*athenaModel.Department, error) {
	resp := new(athenaModel.GetDepartmentByIdResp)
	err := rpc.AthenaCli.Call(ctx, "GetDepartmentById", req, resp)
	if err != nil {
		global.GVA_LOG.Error("call GetDepartmentById fail", zap.Any("err", err))
		return nil, err
	}
	if resp.BaseResp == nil {
		global.GVA_LOG.Error("GetDepartmentById baseResp nil", zap.Any("resp", resp))
		return nil, errors.New("GetDepartmentById resp code invalid")
	}
	if resp.BaseResp.Code != 0 {
		global.GVA_LOG.Error("call GetDepartmentById status invalid", zap.Any("code", resp.BaseResp.Code), zap.Any("msg", resp.BaseResp.Message))
		return nil, errors.New("GetDepartmentById resp code invalid")
	}
	return &resp.Department, nil
}

func (r athenaRpc) GetDepartmentList(ctx context.Context, req athenaModel.GetDepartmentListReq) (int64, []athenaModel.Department, error) {
	resp :=new(athenaModel.GetDepartmentListResp)
	err := rpc.AthenaCli.Call(ctx, "GetDepartmentList", req, resp)
	if err != nil {
		global.GVA_LOG.Error("call GetDepartmentList fail", zap.Any("err", err))
		return 0, nil, err
	}
	if resp.BaseResp == nil {
		global.GVA_LOG.Error("GetDepartmentList baseResp nil", zap.Any("resp", resp))
		return 0, nil, errors.New("GetDepartmentList resp code invalid")
	}
	if resp.BaseResp.Code != 0 {
		global.GVA_LOG.Error("call GetDepartmentList status invalid", zap.Any("code", resp.BaseResp.Code), zap.Any("msg", resp.BaseResp.Message))
		return 0, nil, errors.New("GetDepartmentList resp code invalid")
	}
	return resp.Total, resp.List, nil
}
