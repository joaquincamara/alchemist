package coolFeatures

type Service interface {
	Add(coolFeatures *CoolFeatures) error
	FindAll() ([]*CoolFeatures, error)
	Update(coolFeatures *CoolFeatures) error
	Delete(id int) error
}

type service struct {
	coolFeaturesRepo Repository
}

func NewCoolFeaturesService(coolFeaturesRepo Repository) Service {
	return &service{coolFeaturesRepo: coolFeaturesRepo}
}
	
func (s *service) Add(coolFeatures *CoolFeatures) error {
	return s.coolFeaturesRepo.Add(coolFeatures)
}
	
func (s *service) FindAll() ([]*CoolFeatures, error) {
	return s.coolFeaturesRepo.FindAll()
}
	
func (s *service) Delete(id int) error {
	return s.coolFeaturesRepo.Delete(id)
}
	
func (s *service) Update(coolFeatures *CoolFeatures) error {
	return s.coolFeaturesRepo.Update(coolFeatures)
}
