package repository

type IAuthRepository interface {
	AuthenticateOwner(ownerID, password string) (bool, error)
}
