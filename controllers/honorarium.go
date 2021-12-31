package controllers

import (
	"net/http"

	"rest-api-go/structs"

	"github.com/gin-gonic/gin"
)

//get data by id
func (idb *InDB) GetHonorarium(c *gin.Context) {
	var (
		honorarium structs.Honorarium
		result     gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&honorarium).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": honorarium,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}
