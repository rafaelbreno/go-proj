package controllers

import (
	"github.com/gin-gonic/gin"
	"go-proj/app/models"
	"net/http"
	"reflect"
)

type ListController struct{}

func (_ ListController) Index(c *gin.Context) {
	var lists []models.List
	models.DB.Find(&lists)

	c.JSON(http.StatusOK, gin.H{"data": lists})
}

func (_ ListController) Store(c *gin.Context) {
	var input models.CreateListInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list := models.List{
		UserId: input.UserId,
		Title:  input.Title,
		Status: input.Status,
	}

	models.DB.Create(&list)

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (_ ListController) Show(c *gin.Context) {
	var list models.List

	if err := models.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (_ ListController) Update(c *gin.Context) {
	var list models.List
	if err := models.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input models.UpdateListInput
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
		if err := models.DB.Model(&list).Updates(inputData).Error; err != nil {
											.....^^^^^^^^^.....
											should've been just input
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	*/
	// ---- END OF WORKAROUND ----

	if err := models.DB.Model(&list).Updates(inputData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (_ ListController) Delete(c *gin.Context) {
	var list models.List

	if err := models.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&list)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
