package comment

import (
	"event/entity"

	"gorm.io/gorm"
)

type CommentModel struct {
	DB *gorm.DB
}

func NewCommenttModel(db *gorm.DB) *CommentModel {
	return &CommentModel{
		DB: db,
	}
}

func (m *CommentModel) Insert(comment *entity.Comment) (string, error) {
	record := m.DB.Create(comment)

	if record.RowsAffected == 0 {
		return "comment failed", record.Error
	} else {
		return "Comment success", nil
	}
}
func (m *CommentModel) Delete(id uint, idUser uint) (string, error) {
	var comment entity.Comment
	if err := m.DB.Where("id = ? AND user_id = ?", id, idUser).First(&comment).Error; err != nil {
		return "comment delete failed", err
	}
	record := m.DB.Where("id=?").Delete(&comment)

	if record.RowsAffected == 0 {
		return "comment delete failed", record.Error
	} else {
		return "Comment success", nil
	}
}
