package service

import (
	"dousheng/dao"
	"dousheng/data"
	tool "dousheng/middleware/rabbitMQ"
	"log"
	"strconv"
	"strings"
	"time"
)

func SendMessage(res *data.DouyinMessageActionRequest) error {
	sb := strings.Builder{}
	_, err := sb.WriteString(strconv.Itoa(int(res.UserId)))
	if err != nil {
		return err
	}
	sb.WriteString(" ")
	_, err1 := sb.WriteString(strconv.Itoa(int(res.ToUserId)))
	if err1 != nil {
		return err1
	}
	sb.WriteString(" ")
	_, err2 := sb.WriteString(res.Content)
	if err2 != nil {
		return err2
	}
	tool.RmqMeaasge.Publish(sb.String())
	log.Println("消息打入成功。")
	return nil
}

func HistoryMessage(res *data.DouyinMessageHistoryRequest) (history []*data.MessageHistory, err error) {
	messages, err := dao.MessageHistory(res.UserId, res.ToUserId, res.LastTime)
	if err != nil {
		return nil, err
	}
	history = make([]*data.MessageHistory, len(messages))
	loc, _ := time.LoadLocation("Local")
	for i, value := range messages {
		theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", value.CreateTime, loc)
		history[i] = &data.MessageHistory{
			Id:         value.Id,
			UserID:     value.UserID,
			ToUserID:   value.ToUserID,
			Content:    value.Content,
			CreateTime: theTime.Unix(),
		}
	}
	return
}
