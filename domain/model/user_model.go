package model

type UserModel struct {
	Name     string `bson:"name" json:"name" validate:"min=4"`
	Username string `bson:"username" json:"username" validate:"min=6"`
	Password string `bson:"password" json:"password" validate:"min=8"`
	Modify   Modify `bson:"modify" json:"modify"`
}
