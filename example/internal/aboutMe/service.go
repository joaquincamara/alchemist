package aboutMe

type Service interface {
	Add(aboutMe *AboutMe) error
	FindAll() ([]*AboutMe, error)
	Update(aboutMe *AboutMe) error
	Delete(id int) error
}

type service struct {
	aboutMeRepo Repository
}

func NewAboutMeService(aboutMeRepo Repository) Service {
	return &service{aboutMeRepo: aboutMeRepo}
}
	
func (s *service) Add(aboutMe *AboutMe) error {
	return s.aboutMeRepo.Add(aboutMe)
}
	
func (s *service) FindAll() ([]*AboutMe, error) {
	return s.aboutMeRepo.FindAll()
}
	
func (s *service) Delete(id int) error {
	return s.aboutMeRepo.Delete(id)
}
	
func (s *service) Update(aboutMe *AboutMe) error {
	return s.aboutMeRepo.Update(aboutMe)
}
