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
func changeMenuStruct(list []*models.ManagerMenu) (map[int]interface{}, error) {
	res := make(map[int]interface{})
	for k, v := range list {
		tmp := make(map[string]interface{})
		tmp["id"] = v.ID
		tmp["parentId"] = v.ParentID
		tmp["type"] = v.Type
		tmp["code"] = v.Code
		tmp["name"] = v.Name
		tmp["path"] = v.Path
		tmp["method"] = v.Method
		tmp["apiPath"] = v.APIPath
		tmp["sequence"] = v.Sequence
		tmp["status"] = v.Status
		tmp["create_time"] = v.CreateTime

		items, err := models.ListManagerMenuByParentId(v.ID)
		if err != nil {
			return nil, err
		}
		if len(items) > 0 {
			tmp["children"], err = changeMenuStruct(items)
			if err != nil {
				return nil, err
			}
		}

		res[k] = tmp
		continue
	}

	return res, nil
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
