package controller

import (
	"authentication/model"
	"authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)



func Register(c * gin.Context){
	var input model.UserRegisterDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	  }
	result, msg := service.Register(input);
	if(result == nil){
		c.JSON(http.StatusInternalServerError, gin.H{"messageError": msg})
	}
	c.JSON(http.StatusOK, gin.H{"result" : result, "message": msg})
}