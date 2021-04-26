package service

import (
	"authentication/config"
	"authentication/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)


func Register(dto model.UserRegisterDto) (interface{}, string){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	findErr := config.UserCollection.FindOne(ctx, bson.M{ "username": dto.Username}); findErr == nil {
		msg := fmt.Println("Username has already existed!")	
		return nil, msg
	}
	user := model.UserRegisterDtoToEntity(dto)
	result, insertErr := config.UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Println("Can not create user!")		
			return nil, msg
		}
	return result, "Regisger successful!"
}