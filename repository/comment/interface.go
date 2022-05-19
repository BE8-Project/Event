package comment

import "event/entity"

type CommentRepo interface {
	Insert(comment *entity.Comment) (string, error)
	Delete(id uint, idUser uint) (string, error)
}
