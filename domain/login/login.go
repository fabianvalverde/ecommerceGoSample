package login

import "github.com/fabianvalverde/ecommerceGoSample/model"

// Puerto de entrada a login
type UseCase interface {
	Login(email, password, jwtSecretKey string) (model.User, string, error)
}

// Puerto de entrada a user
type UseCaseUser interface {
	Login(email, password string) (model.User, error)
}
