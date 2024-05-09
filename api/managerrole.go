package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/models"
	"github.com/provider-go/pkg/output"
)

func CreateRole(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)

	code := output.ParamToString(json["code"])
	name := output.ParamToString(json["name"])
	description := output.ParamToString(json["description"])
	sequence := output.ParamToInt(json["sequence"])
	err := models.CreateManagerRole(code, name, description, sequence)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}
func UpdateRole(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	code := output.ParamToString(json["code"])
	name := output.ParamToString(json["name"])
	description := output.ParamToString(json["description"])
	sequence := output.ParamToInt(json["sequence"])
	status := output.ParamToInt(json["status"])
	err := models.UpdateManagerRole(id, code, name, description, sequence, status)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}
func DeleteRole(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	err := models.DeleteManagerRole(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}
func ListRole(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	pageSize := output.ParamToInt(json["pageSize"])
	pageNum := output.ParamToInt(json["pageNum"])
	list, total, err := models.ListManagerRole(pageSize, pageNum)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		output.ReturnSuccessResponse(ctx, res)
	}

}
func ViewRole(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	row, err := models.ViewManagerRole(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, row)
	}

}
