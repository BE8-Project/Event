package repository

import (
	"event/delivery/helpers/response"
	"event/entity"
)

type EventModel interface {
	Insert(task *entity.Event) response.InsertEvent
	GetAll(name, location string, limit, page int) []response.GetEvent
	Get(id uint) (response.GetEvent, error)
	Update(id, user_id uint, task *entity.Event) (response.UpdateEvent, error)
	Delete(id, user_id uint) (response.DeleteEvent, error)
	GetByUser(user_id uint) []response.GetEvent
}