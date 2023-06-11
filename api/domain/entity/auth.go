package entity

type AuthenticateOwner struct {
	OwnerID string `json:"owner_id"`
	Passwd  string `json:"passwd"`
}
