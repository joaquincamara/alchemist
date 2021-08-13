package aboutMe

type Repository interface {
	Add(aboutMe *AboutMe) error
	FindAll() ([]*AboutMe, error)
	Update(aboutMe *AboutMe) error
	Delete(id int) error
}
	