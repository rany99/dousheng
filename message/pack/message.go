package pack

import (
	"douyin/code_gen/kitex_gen/messageproto"
	"douyin/message/infra/dal/model"
	redisModel "douyin/message/infra/redis/model"
)

func Message(message *model.Message) *messageproto.MessageInfo {
	return &messageproto.MessageInfo{
		MessageId:  message.MessageUUID,
		FromUserId: message.FromUserId,
		ToUserId:   message.ToUserId,
		Content:    message.Contents,
		CreateTime: message.CreateTime,
	}
}

func Messages(messages []*model.Message) []*messageproto.MessageInfo {
	messageInfos := make([]*messageproto.MessageInfo, len(messages))
	for i, message := range messages {
		messageInfos[i] = Message(message)
	}
	return messageInfos
}

func MessageFromRedisModel(message *redisModel.MessageRedis) *model.Message {
	return &model.Message{
		FromUserId:  message.FromUserId,
		ToUserId:    message.ToUserId,
		Contents:    message.Content,
		MessageUUID: message.MessageId,
		CreateTime:  message.CreateTime,
	}
}
