package model

type UserModel struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Modify   Modify `bson:"modify" json:"modify"`
}
