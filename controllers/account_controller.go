package controllers

import (
	"net/http"

	"github.com/Nezent/GoCrud/config"
	"github.com/Nezent/GoCrud/models"
	"github.com/gin-gonic/gin"
)

func GetAccount(ctx *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	ctx.JSON(http.StatusOK,&users)
}

func CreateAccount(ctx *gin.Context){
	user := models.User{}
	ctx.Bind(&user)
	config.DB.Create(&user)
	ctx.JSON(http.StatusCreated,&user)
}

func DeleteAccount(ctx *gin.Context){
	user := models.User{}
	config.DB.Where("id = ?",ctx.Param("id")).Delete(&user)
	ctx.JSON(http.StatusOK,&user)
}

func UpdateAccount(ctx *gin.Context) {
	user := models.User{}
	config.DB.Where("id = ?",ctx.Param("id")).First(&user)
	ctx.Bind(&user)
	config.DB.Save(&user)
	ctx.JSON(http.StatusOK,&user)
}