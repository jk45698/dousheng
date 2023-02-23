package service

import (
	"dousheng/data"
	"dousheng/middleware/rabbitMQ"
	"dousheng/middleware/redis"
	"log"
	"strconv"
	"strings"
)

func AddFollowRelation(userId int64, targetId int64) (bool, error) {
	key := strconv.Itoa(int(userId)) + strconv.Itoa(int(targetId)) + "follow"
	exist := redis.IsExistsCache(key, redis.GetRdbRelationClient())
	if exist == 1 {
		return true, nil
	}
	// 加信息打入消息队列。
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(userId)))
	sb.WriteString(" ")
	sb.WriteString(strconv.Itoa(int(targetId)))
	tool.RmqFollowAdd.Publish(sb.String())
	// 记录日志
	log.Println("消息打入成功。")
	return true, nil
}

//生产者弄丢消息：没有把消息打入消息队列
//RabbitMQ 本身丢失的可能性就非常低，其次如果这里需要落库再用定时任务扫描重发还要开发一堆代码，分布式定时任务…
//再其次定时任务扫描肯定会增加消息延迟，不是很有必要。真实业务场景是记录一下日志就行了，方便问题回溯，顺便发个邮件给相关人员，
//如果真的极其罕见的是生产者弄丢消息，那么开发往数据库补数据就行了。

// 消费者弄丢消息：可以设置ack，手动确认机制，但是代码报错并不能因为重试而解决，可能会造成死循环。
// 2，MQ收到消息，暂存内存中，还没消费，自己挂掉，数据会都丢失
// 解决方式：MQ设置为持久化。将内存数据持久化到磁盘中
func DeleteFollowRelation(userId int64, targetId int64) (bool, error) {
	// 加信息打入消息队列。
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(userId)))
	sb.WriteString(" ")
	sb.WriteString(strconv.Itoa(int(targetId)))
	err := tool.RmqFollowDel.Publish(sb.String())
	// 记录日志
	if err != nil {
		log.Printf("消息%v打入失败！err:%v", sb, err.Error())
	}
	log.Println("消息打入成功。")
	// 更新redis信息。
	return true, nil
}
func RelationAction(req *data.DouyinRelationActionRequest) error {
	// 1-关注
	if req.ActionType == 1 {
		go AddFollowRelation(req.UserId, req.ToUserId)
	}
	// 2-取消关注
	if req.ActionType == 2 {
		go DeleteFollowRelation(req.UserId, req.ToUserId)
	}
	return nil
}
