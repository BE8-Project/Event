package repository

import (
	"errors"
	"event/delivery/helpers/response"
	"event/entity"

	"gorm.io/gorm"
)

type participantModel struct {
	DB *gorm.DB
}

func NewParticipantModel(db *gorm.DB) *participantModel {
	return &participantModel{
		DB: db,
	}
}

func (m *participantModel) Insert(participant *entity.Participant) (response.InsertParticipant, error) {
	record := m.DB.Create(participant)

	if record.RowsAffected == 0 {
		return response.InsertParticipant{}, record.Error
	} else {
		return response.InsertParticipant{
			OrderID: participant.OrderID,
			Total: participant.Total,
			CreatedAt: participant.CreatedAt,
		}, nil
	}
}

func (m *participantModel) Update(order_id string, user_id uint, participant *entity.Participant) (response.UpdateParticipat, error) {
	record := m.DB.Where("order_id = ? AND user_id = ?", order_id, user_id).Updates(&participant)

	if record.RowsAffected == 0 {
		return response.UpdateParticipat{}, errors.New("you are not allowed to access this resource")
	} else {
		return response.UpdateParticipat{
			OrderID: participant.OrderID,
			Status: 	   participant.Status,
			UpdatedAt: 	participant.UpdatedAt,
		}, nil
	}
}

func (m *participantModel) GetByUser(user_id uint) []response.GetParticipant {
	var tasks []entity.Participant

	m.DB.Where("user_id = ?", user_id).Find(&tasks)

	var results []response.GetParticipant
	for _, task := range tasks {
		var event entity.Event
		m.DB.Where("id = ?", task.EventID).Find(&event)
		results = append(results, response.GetParticipant{
			ID: task.ID,
			OrderID: task.OrderID,
			EventID: task.EventID,
			Status: task.Status,
			Total: task.Total,
			Event: response.GetEvent{
				ID: event.ID,
				Name: event.Name,
				HostedBy: event.HostedBy,
				DateStart: event.DateStart,
				DateEnd: event.DateEnd,
				Location: event.Location,
				Details: event.Details,
				Ticket: event.Ticket,
				Price: event.Price,
				Image: event.Image,
			},

		})

	}

	return results
}