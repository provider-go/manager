package api

import (
	"github.com/provider-go/manager/models"
	_ "github.com/provider-go/manager/models"
	"github.com/provider-go/pkg/types"
	"testing"
)

func TestName(t *testing.T) {
	a := &models.ManagerMenu{
		ID:         1,
		ParentID:   1,
		Type:       "1",
		Code:       "1",
		Name:       "1",
		Path:       "1",
		Method:     "1",
		APIPath:    "1",
		Sequence:   0,
		Status:     "1",
		CreateTime: types.Time{},
		UpdateTime: types.Time{},
	}

	var list []*models.ManagerMenu
	list = append(list, a)
	res, err := changeMenuStruct(list)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}
