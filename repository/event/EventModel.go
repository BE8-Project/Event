package repository

import (
	"errors"
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

func (m *eventModel) GetAll(name, location string) []response.GetEvent {
	var tasks []entity.Event
	m.DB.Where("name LIKE ? AND location LIKE ?", "%"+name+"%", "%"+location+"%").Find(&tasks)

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

func (m *eventModel) Get(id uint) (response.GetEvent, error) {
	var task entity.Event
	record := m.DB.Where("id = ?", id).First(&task)

	if record.RowsAffected == 0 {
		return response.GetEvent{}, errors.New("event not found")
	} else {
		return response.GetEvent {
			Name : task.Name,
			HostedBy: task.HostedBy,
			DateStart: task.DateStart,
			DateEnd: task.DateEnd,
			Location: task.Location,
			Details: task.Details,
			Ticket: task.Ticket,
		}, nil
	}
}

func (m *eventModel) Update(id, user_id uint, task *entity.Event) (response.UpdateEvent, error) {
	if task.Name == "" && task.HostedBy == "" && task.Location == "" && task.Details == "" && task.Ticket == 0 {
		return response.UpdateEvent{}, errors.New("required")
	}

	update := m.DB.Model(&entity.Event{}).Where("id = ? AND user_id = ?", id, user_id).Updates(&task)

	if task.Name == "" {
		m.DB.Where("id = ?", id).First(&task)
	}

	if update.RowsAffected == 0 {
		return response.UpdateEvent{}, errors.New("you are not allowed to access this resource")
	} else {
		return response.UpdateEvent{
			Name: 	task.Name,
			UpdatedAt: task.UpdatedAt,
		}, nil
	}
}

func (m *eventModel) Delete(id, user_id uint) (response.DeleteEvent, error) {
	var task entity.Event
	record := m.DB.Where("id = ? AND user_id = ?", id, user_id).First(&task)

	if record.RowsAffected == 0 {
		return response.DeleteEvent{}, errors.New("you are not allowed to access this resource")
	} else {
		m.DB.Delete(&task)
		return response.DeleteEvent{
			Name: 	task.Name,
			DeletedAt: task.DeletedAt,
		}, nil
	}
}