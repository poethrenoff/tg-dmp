package app

type ProfileService interface {
	GetProfileById(id string) (Profile, error)
}

type profileService struct {
	repository ProfileRepository
}

func (s *profileService) GetProfileById(id string) (Profile, error) {
	profile, err := s.repository.GetProfileById(id)
	return profile, err
}

type Service struct {
	ProfileService
}

func NewService(r *Repository) *Service {
	return &Service{
		ProfileService: &profileService{repository: r.ProfileRepository},
	}
}
