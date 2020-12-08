package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gapi/models"
	"github.com/lvxin0315/gapi/services"
)

func Index(c *gin.Context) {

	demoService := services.NewDemoService()

	demoModel := new(models.DemoModel)
	demoModel.Age = 12
	demoModel.Name = "lvxin"

	err := demoService.Save(demoModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("demoModel.ID: ", demoModel.ID)

	err = demoService.UpdateWithId(demoModel.ID, map[string]interface{}{
		"name": "lvxin_new",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("update after name: ", demoModel.Name)

	newDemoModel := new(models.DemoModel)
	newDemoModel.ID = demoModel.ID
	err = demoService.First(newDemoModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("first name: ", newDemoModel.Name)

	var demoModelList []*models.DemoModel
	err = demoService.Find(&demoModelList, map[string]interface{}{
		"name": newDemoModel.Name,
	})
	if err != nil {
		panic(err)
	}
	for _, m := range demoModelList {
		fmt.Println("m name: ", m.Name)
	}
}
