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

func NewProfileRepository(db *sqlx.DB) ProfileRepository {
	return &profileRepository{db: db}
}
