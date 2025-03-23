package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/core"

	cajasUseCases "organizador-naranjas-backend-multi5to/src/features/cajas/application"
	cajasInfrastructure "organizador-naranjas-backend-multi5to/src/features/cajas/infrastructure"
	cajasControllers "organizador-naranjas-backend-multi5to/src/features/cajas/infrastructure/controllers"

	lotesUseCases "organizador-naranjas-backend-multi5to/src/features/lotes/application"
	lotesInfrastructure "organizador-naranjas-backend-multi5to/src/features/lotes/infrastructure"
	lotesControllers "organizador-naranjas-backend-multi5to/src/features/lotes/infrastructure/controllers"

	naranjasUseCases "organizador-naranjas-backend-multi5to/src/features/naranjas/application"
	naranjasInfrastructure "organizador-naranjas-backend-multi5to/src/features/naranjas/infrastructure"
	naranjasControllers "organizador-naranjas-backend-multi5to/src/features/naranjas/infrastructure/controllers"

	usersUseCases "organizador-naranjas-backend-multi5to/src/features/users/application"
	usersInfrastructure "organizador-naranjas-backend-multi5to/src/features/users/infrastructure"
	usersControllers "organizador-naranjas-backend-multi5to/src/features/users/infrastructure/controllers"

	sp32UseCases "organizador-naranjas-backend-multi5to/src/features/sp32/application"
	sp32Infrastructure "organizador-naranjas-backend-multi5to/src/features/sp32/infrastructure"
	sp32Adapters "organizador-naranjas-backend-multi5to/src/features/sp32/infrastructure/adapters"
	sp32Controllers "organizador-naranjas-backend-multi5to/src/features/sp32/infrastructure/controllers"
)

type Dependencies struct {
	engine *gin.Engine
}

func NewDependencies() *Dependencies {
	return &Dependencies{
		engine: gin.Default(),
	}
}

func (d *Dependencies) Run() error {
	database := core.NewDatabase()

	cajasDatabase := cajasInfrastructure.NewMySQL(database.Conn)
	getAllUseCase := cajasUseCases.NewGetAllCajaUseCase(cajasDatabase)
	getAllController := cajasControllers.NewGetAllCajaController(getAllUseCase)
	createCajaUseCase := cajasUseCases.NewCreateCajaUseCase(cajasDatabase)
	createCajaController := cajasControllers.NewCreateCajaController(createCajaUseCase)
	updateCajasUseCase := cajasUseCases.NewUpdateCajaUseCase(cajasDatabase)
	updateCajasController := cajasControllers.NewUpdateCajaController(updateCajasUseCase)
	deleteCajasUseCase := cajasUseCases.NewDeleteCajaUseCase(cajasDatabase)
	deleteCajasController := cajasControllers.NewDeleteCajaController(deleteCajasUseCase)
	cajasRoutes := cajasInfrastructure.NewCajasRoutes(d.engine, getAllController, createCajaController, updateCajasController, deleteCajasController)

	naranjaDatabase := naranjasInfrastructure.NewMySQL(database.Conn)
	createNaranjaUseCase := naranjasUseCases.NewCreateCajaUseCase(naranjaDatabase)
	createNaranjaController := naranjasControllers.NewCreateNaranjaController(createNaranjaUseCase)
	getAllNaranjaUseCase := naranjasUseCases.NewGetAllUseCase(naranjaDatabase)
	getAllNaranjasController := naranjasControllers.NewGetAllController(getAllNaranjaUseCase)
	updateNaranjaUseCase := naranjasUseCases.NewUpdateNaranjaUseCase(naranjaDatabase)
	updateContollers := naranjasControllers.NewUpdateNaranjaController(updateNaranjaUseCase)
	naranjasRoutes := naranjasInfrastructure.NewNaranjasRoutes(d.engine, createNaranjaController, getAllNaranjasController, updateContollers)

	userDataBase := usersInfrastructure.NewMysql(database.Conn)
	createUser := usersUseCases.NewSaveUser(userDataBase)
	logInUser := usersUseCases.NewLogInUseCase(userDataBase)
	createUserController := usersControllers.NewCreateUserController(createUser)
	logInController := usersControllers.NewLoginController(logInUser)
	userUpdate := usersUseCases.NewUpdateUserUseCase(userDataBase)
	updateUserController := usersControllers.NewUpdateUserController(userUpdate)
	deleteUserUseCase := usersUseCases.NewDeleteUserUseCase(userDataBase)
	deleteUserController := usersControllers.NewDeleteUserController(deleteUserUseCase)
	getUsersUseCase := usersUseCases.NewGetUsersUseCase(userDataBase)
	getUsersController := usersControllers.NewGetUsersController(getUsersUseCase)
	getUserByIdUseCase := usersUseCases.NewGetUserByIDUseCase(userDataBase)
	getUserByIdController := usersControllers.NewGetUserByIdController(getUserByIdUseCase)
	getUserByUsernameUseCase := usersUseCases.NewGetUserByUsernameUseCase(userDataBase)
	getUserByUsernameController := usersControllers.NewGetByUsernameController(getUserByUsernameUseCase)
	userRoutes := usersInfrastructure.NewUserRoutes(d.engine, createUserController, logInController, updateUserController, deleteUserController, getUserByIdController, getUserByUsernameController, getUsersController)

	lotesDatabase := lotesInfrastructure.NewMySQL(database.Conn)
	createLoteUseCase := lotesUseCases.NewCreateLoteUseCase(lotesDatabase)
	listAllLotesUseCase := lotesUseCases.NewListLotesUseCase(lotesDatabase)
	listLoteIdUseCase := lotesUseCases.NewListLoteIdUseCase(lotesDatabase)
	listLoteDateUseCase := lotesUseCases.NewListLoteDateUseCase(lotesDatabase)
	deleteLoteUseCase := lotesUseCases.NewRemoveLoteUseCase(lotesDatabase)
	updateLoteUseCase := lotesUseCases.NewUpdateLoteUseCase(lotesDatabase)
	createLoteController := lotesControllers.NewCreateLoteController(createLoteUseCase)
	listAllLotesController := lotesControllers.NewGetAllLotesController(listAllLotesUseCase)
	listLoteIdController := lotesControllers.NewListLoteIdController(listLoteIdUseCase)
	listLoteDateController := lotesControllers.NewListLoteDateController(listLoteDateUseCase)
	deleteLoteController := lotesControllers.NewDeleteLoteController(deleteLoteUseCase)
	updateLoteControlerr := lotesControllers.NewUpdateLoteController(updateLoteUseCase)
	lotesRoutes := lotesInfrastructure.NewLotesRoutes(d.engine, createLoteController, listAllLotesController, listLoteIdController, listLoteDateController, deleteLoteController, updateLoteControlerr)

	sp32Mysql := sp32Adapters.NewMysql(database.Conn)
	saveSp32UseCase := sp32UseCases.NewSaveSp32(sp32Mysql)
	createSp32Controller := sp32Controllers.NewCreateSp32Controller(saveSp32UseCase)
	sp32Routes := sp32Infrastructure.NewSp32Routes(d.engine.RouterGroup,createSp32Controller)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	d.engine.Use(cors.New(config))

	cajasRoutes.SetupRoutes()
	sp32Routes.Run()
	naranjasRoutes.SetupRoutes()
	userRoutes.SetupRoutes()
	lotesRoutes.SetupRoutes()

	return d.engine.Run(":8082")
}
