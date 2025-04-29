package menu

type MenuRepository interface {
	Save(menu *Menu) error
	FindAll() ([]*Menu, error)
	Delete(id int) error
}