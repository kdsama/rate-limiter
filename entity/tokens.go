package entity

type UserToken struct {
	User_ID   string `json:"user_id" bson:"user_id"`
	Token     string `json:"token" bson:"token"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

func NewUserToken(user_id string, token string, timestamp int64) *UserToken {
	return &UserToken{User_ID: user_id, Token: token, CreatedAt: timestamp, UpdatedAt: timestamp}
}
