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