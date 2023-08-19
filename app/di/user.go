package di

import (
	"githut.com/warley-juneo/bexs-bank-challenger/adapter/http/userservice"
	"githut.com/warley-juneo/bexs-bank-challenger/adapter/postgres"
	"githut.com/warley-juneo/bexs-bank-challenger/adapter/http/userrepository"
	"githut.com/warley-juneo/bexs-bank-challenger/core/domain"
	"githut.com/warley-juneo/bexs-bank-challenger/core/userusecase"
)


// ConfigUserDI return a UserService abstraction with dependency injection configuration
func ConfigUserDI(conn postgres.PoolInterface) domain.UserService {
	userRepository := userrepository.New(conn)
	userUseCase := userusecase.New(userRepository)
	userService := userservice.New(userUseCase)

	return userService
}