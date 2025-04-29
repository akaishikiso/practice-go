package menu

import "practice-go/backend/domain/menu" 

type Service struct {
	Repo menu.MenuRepository
}

func (s *Service) Create(name, date, note string) (*menu.Menu, error) {
	m := &menu.Menu{
		Name: name,
		Date: date,
		Note: note,
	}
	return m, s.Repo.Save(m)
}

func (s *Service) GetAll() ([]*menu.Menu, error) {
	return s.Repo.FindAll()
}

func (s *Service) Delete(id int) error {
	return s.Repo.Delete(id)
}