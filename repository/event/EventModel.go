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

func (m *eventModel) GetAll() []response.GetEvent {
	var tasks []entity.Event
	m.DB.Find(&tasks)

	var results []response.GetEvent
	for _, result := range tasks {
		results = append(results, response.GetEvent{
			Name : result.Name,
			HostedBy: result.HostedBy,
			DateStart: result.DateStart,
			DateEnd: result.DateEnd,
			Location: result.Location,
			Details: result.Details,
			Ticket: result.Ticket,
		})
	}

	return results
}