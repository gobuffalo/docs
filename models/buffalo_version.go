package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
)

type BuffaloVersion struct {
	ID         int       `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Version    string    `json:"version" db:"version"`
	PreRelease bool      `json:"pre_release" db:"pre_release"`
	Token      string    `json:"token" db:"-"`
}

// String is not required by pop and may be deleted
func (b BuffaloVersion) String() string {
	bm, _ := json.Marshal(b)
	return string(bm)
}

// BuffaloVersions is not required by pop and may be deleted
type BuffaloVersions []BuffaloVersion

// String is not required by pop and may be deleted
func (b BuffaloVersions) String() string {
	bm, _ := json.Marshal(b)
	return string(bm)
}

func CurrentBuffaloVersion(tx *pop.Connection) *BuffaloVersion {
	v := &BuffaloVersion{}
	err := tx.Order("version desc").First(v)
	if err != nil {
		v.Version = "0.0.0"
	}
	return v
}
