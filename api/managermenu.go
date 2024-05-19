package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/models"
	"github.com/provider-go/pkg/output"
)

func CreateMenu(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)

	parentId := output.ParamToInt32(json["parentId"])
	menuType := output.ParamToString(json["type"])
	code := output.ParamToString(json["code"])
	name := output.ParamToString(json["name"])
	path := output.ParamToString(json["path"])
	method := output.ParamToString(json["method"])
	apiPath := output.ParamToString(json["apiPath"])
	status := output.ParamToString(json["status"])
	err := models.CreateManagerMenu(parentId, menuType, code, name, path, method, apiPath, status)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func UpdateMenu(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	parentId := output.ParamToInt32(json["parentId"])
	menuType := output.ParamToString(json["type"])
	code := output.ParamToString(json["code"])
	name := output.ParamToString(json["name"])
	path := output.ParamToString(json["path"])
	method := output.ParamToString(json["method"])
	apiPath := output.ParamToString(json["apiPath"])
	sequence := output.ParamToString(json["sequence"])
	status := output.ParamToString(json["status"])

	err := models.UpdateManagerMenu(id, parentId, menuType, code, name, path, method, apiPath, sequence, status)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func DeleteMenu(ctx *gin.Context) {
	id := output.ParamToInt32(ctx.Query("id"))
	err := models.DeleteManagerMenu(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func ListMenu(ctx *gin.Context) {
	pageSize := output.ParamToInt(ctx.Query("pageSize"))
	pageNum := output.ParamToInt(ctx.Query("pageNum"))

	list, total, err := models.ListManagerMenu(pageSize, pageNum)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["list"] = list
		res["total"] = total
		output.ReturnSuccessResponse(ctx, res)
	}

}

type AllManagerMenu struct {
	models.ManagerMenu
	Children []AllManagerMenu `json:"children"`
}

func ListAllMenu(ctx *gin.Context) {
	list, err := models.ListManagerMenuByParentId(0)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
		return
	}

	items, err := changeMenuStruct(list)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
		return
	}
	res := make(map[string]interface{})
	res["list"] = items
	output.ReturnSuccessResponse(ctx, res)

}

// changeMenuStruct 递归格式化菜单结构
func changeMenuStruct(list []*models.ManagerMenu) ([]AllManagerMenu, error) {
	var rows []AllManagerMenu
	for k, v := range list {
		rows[k].ParentID = v.ID
		rows[k].ParentID = v.ParentID
		rows[k].Type = v.Type
		rows[k].Code = v.Code
		rows[k].Name = v.Name
		rows[k].Path = v.Path
		rows[k].Method = v.Method
		rows[k].APIPath = v.APIPath
		rows[k].Sequence = v.Sequence
		rows[k].Status = v.Status
		rows[k].CreateTime = v.CreateTime
		rows[k].UpdateTime = v.UpdateTime

		items, err := models.ListManagerMenuByParentId(v.ID)
		if err != nil {
			return nil, err
		}
		if len(items) > 0 {
			rows[k].Children, err = changeMenuStruct(items)
			if err != nil {
				return nil, err
			}
		}
	}

	return rows, nil
}

func ViewMenu(ctx *gin.Context) {
	id := output.ParamToInt32(ctx.Query("id"))
	row, err := models.ViewManagerMenu(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, row)
	}

}
