package coolFeatures

type Repository interface {
	Add(coolFeatures *CoolFeatures) error
	FindAll() ([]*CoolFeatures, error)
	Update(coolFeatures *CoolFeatures) error
	Delete(id int) error
}
	