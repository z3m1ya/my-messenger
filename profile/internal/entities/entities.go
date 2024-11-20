package entities

type Profile struct {
	Id      string
	Name    string
	Email   string
	Friends []string
}

type FriendRequest struct {
	UserId   string
	TargetId string
}
