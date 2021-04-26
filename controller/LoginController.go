package controller

import (
	"authentication/model"
	"authentication/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c * gin.Context){
	var input model.UserLoginDto;
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resutl, msg := service.Login(input)
	if(!resutl){
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}
	fmt.Println(msg)
	c.JSON(http.StatusAccepted, gin.H{"token":msg})
}