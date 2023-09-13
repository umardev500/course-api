package model

type Modify struct {
	CreatedAt int64 `bson:"created_at" json:"created_at"`
	UpdatedAt int64 `bson:"updated_at" json:"updated_at,omitempty"`
	DeletedAt int64 `bson:"deleted_at" json:"deleted_at,omitempty"`
}
