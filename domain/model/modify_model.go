package model

type Modify struct {
	CreatedAt int `bson:"created_at" json:"created_at"`
	UpdatedAt int `bson:"updated_at" json:"updated_at"`
	DeletedAt int `bson:"deleted_at" json:"deleted_at"`
}
