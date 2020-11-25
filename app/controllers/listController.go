package controllers

import (
	"github.com/gin-gonic/gin"
	"go-proj/app/models"
	"net/http"
	"reflect"
)

type ListController struct {
	User UserController
}

func (l *ListController) Index(c *gin.Context) {
	var lists []models.List

	l.User.Auth.GetAuth(c)

	models.DB.Find(&lists).Where("user_id = ?", l.User.Auth.UserId)

	c.JSON(http.StatusOK, gin.H{
		"data": lists,
	})
}

func (l ListController) Store(c *gin.Context) {
	var input models.CreateListInput
	l.User.Auth.GetAuth(c)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list := models.List{
		UserId: l.User.Auth.UserId,
		Title:  input.Title,
		Status: input.Status,
	}

	models.DB.Create(&list)

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (l ListController) Show(c *gin.Context) {
	var list models.List
	l.User.Auth.GetAuth(c)

	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), l.User.Auth.UserId).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (l ListController) Update(c *gin.Context) {
	var list models.List
	l.User.Auth.GetAuth(c)

	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), l.User.Auth.UserId).First(&list).Error; err != nil {
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

func (l ListController) Delete(c *gin.Context) {
	var list models.List
	l.User.Auth.GetAuth(c)

	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), l.User.Auth.UserId).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&list)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
