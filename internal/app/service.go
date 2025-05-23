package app

type ProfileService interface {
	GetProfileById(id string) (Profile, error)
}

type profileService struct {
	repo ProfileRepository
}

func (s *profileService) GetProfileById(id string) (Profile, error) {
	profile, err := s.repo.GetProfileById(id)
	return profile, err
}

func NewProfileService(r ProfileRepository) ProfileService {
	return &profileService{repo: r}
}
