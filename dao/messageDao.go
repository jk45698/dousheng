package dao

import (
	"dousheng/data"
	"time"
)

func CreateMes(UserID int64, To_UserId int64, Content string) error {
	err := DB.Create(&data.MessageRaw{UserID: UserID, ToUserID: To_UserId, Content: Content, CreateTime: time.Now().Format("2006-01-02 15:04:05")}).Error
	return err
}

func MessageHistory(UserID int64, To_UserId int64, LastTime time.Time) ([]*data.Message, error) {
	var Message []*data.Message
	err := DB.Debug().Where("create_time > ?", LastTime).Where("user_id = ? and to_user_id = ?  or user_id = ? and to_user_id = ?",
		UserID, To_UserId, To_UserId, UserID).Order("create_time").Find(&Message).Error
	if err != nil {
		return nil, err
	}
	return Message, nil
}
