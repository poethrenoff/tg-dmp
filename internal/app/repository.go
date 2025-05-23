package app

import "github.com/jmoiron/sqlx"

type ProfileRepository interface {
	GetProfileById(id string) (Profile, error)
}

type profileRepository struct {
	db *sqlx.DB
}

func (r *profileRepository) GetProfileById(id string) (Profile, error) {
	var profile Profile
	err := r.db.Get(&profile, "SELECT * FROM dmp_profile WHERE id=$1 LIMIT 1", id)
	return profile, err

}

type Repository struct {
	ProfileRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ProfileRepository: &profileRepository{db: db},
	}
}
