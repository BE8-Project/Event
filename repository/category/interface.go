package category

import "event/entity"

type CategoryModel interface {
	Insert(categ entity.Category, id uint) (string, error)
	Get() ([]entity.Category, error)
	Delete(id_user, id_categ uint) (string, error)
}
