package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.GetAll()
}

func (s *Service) GetUserByID(id int) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) CreateUser(user *User) error {
	return s.repo.Create(user)
}

func (s *Service) UpdateUser(user *User) error {
	return s.repo.Update(user)
}

func (s *Service) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
