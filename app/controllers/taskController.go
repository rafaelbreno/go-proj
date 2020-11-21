package controllers

import (
	"github.com/gin-gonic/gin"
	"go-proj/app/models"
	"net/http"
	"reflect"
)

type TaskController struct{}

func (_ TaskController) Index(c *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (_ TaskController) Store(c *gin.Context) {
	var input models.CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		ListId: input.ListId,
		Title:  input.Title,
		Status: input.Status,
	}

	models.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (_ TaskController) Show(c *gin.Context) {
	var task models.Task

	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (_ TaskController) Update(c *gin.Context) {
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input models.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	/* This is a workaround
	 * TODO:
	 *    - Investigate this, should've been working without
	 * 		having to convert to a map[string]interface{}
	**/
	v := reflect.ValueOf(input)
	typeOfV := v.Type()

	inputData := map[string]interface{}{}

	for i := 0; i < v.NumField(); i++ {
		inputData[typeOfV.Field(i).Name] = v.Field(i).Interface()
	}
	/*
		if err := models.DB.Model(&task).Updates(inputData).Error; err != nil {
											.....^^^^^^^^^.....
											should've been just input
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	*/
	// ---- END OF WORKAROUND ----

	if err := models.DB.Model(&task).Updates(inputData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (_ TaskController) Delete(c *gin.Context) {
	var task models.Task

	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
