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
	menuType := output.ParamToString(json["menuType"])
	code := output.ParamToString(json["code"])
	name := output.ParamToString(json["name"])
	path := output.ParamToString(json["path"])
	method := output.ParamToString(json["method"])
	apiPath := output.ParamToString(json["apiPath"])
	err := models.CreateManagerMenu(parentId, menuType, code, name, path, method, apiPath)
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
	menuType := output.ParamToString(json["menuType"])
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
		res["records"] = list
		res["total"] = total
		output.ReturnSuccessResponse(ctx, res)
	}

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
