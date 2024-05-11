package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/models"
	"github.com/provider-go/pkg/output"
)

func CreateMenu(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)

	code := output.ParamToString(json["code"])
	name := output.ParamToString(json["name"])
	description := output.ParamToString(json["description"])
	menuType := output.ParamToString(json["menuType"])
	path := output.ParamToString(json["path"])
	properties := output.ParamToString(json["properties"])
	parentId := output.ParamToString(json["parentId"])
	parentPath := output.ParamToString(json["parentPath"])
	err := models.CreateManagerMenu(code, name, description, menuType, path, properties, parentId, parentPath)
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
	code := output.ParamToString(json["code"])
	name := output.ParamToString(json["name"])
	description := output.ParamToString(json["description"])
	sequence := output.ParamToInt(json["sequence"])
	menuType := output.ParamToString(json["menuType"])
	path := output.ParamToString(json["path"])
	properties := output.ParamToString(json["properties"])
	status := output.ParamToString(json["status"])
	parentId := output.ParamToString(json["parentId"])
	parentPath := output.ParamToString(json["parentPath"])

	err := models.UpdateManagerMenu(id, code, name, description, sequence, menuType, path, properties, status, parentId, parentPath)
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
