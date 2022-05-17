package repository

import (
	"event/delivery/helpers/response"
	"event/entity"
)

type EventModel interface {
	Insert(task *entity.Event) response.InsertEvent
	GetAll() []response.GetEvent
}