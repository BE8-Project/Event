package repository

import (
	"event/delivery/helpers/response"
	"event/entity"
)

type ParticipantModel interface {
	Insert(participant *entity.Participant) (response.InsertParticipant, error)
	Update(order_id string, user_id uint, participant *entity.Participant) (response.UpdateParticipat, error)
	GetByUser(user_id uint) []response.GetParticipant
}