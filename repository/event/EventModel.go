package repository

import (
	"event/delivery/helpers/response"
	"event/entity"

	"gorm.io/gorm"
)

type eventModel struct {
	DB *gorm.DB
}

func NewEventModel(db *gorm.DB) *eventModel {
	return &eventModel{
		DB: db,
	}
}

func (m *eventModel) Insert(task *entity.Event) response.InsertEvent {
	m.DB.Create(&task)

	return response.InsertEvent{
		Name: 	task.Name,
		CreatedAt: task.CreatedAt,
	}
}