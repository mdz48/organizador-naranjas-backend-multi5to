package main

import (
	"log"
	"organizador-naranjas-backend-multi5to/src/core"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

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

	esp32UseCases "organizador-naranjas-backend-multi5to/src/features/esp32/application"
	esp32Infrastructure "organizador-naranjas-backend-multi5to/src/features/esp32/infrastructure"
	esp32Adapter "organizador-naranjas-backend-multi5to/src/features/esp32/infrastructure/adapters"
	esp32Controllers "organizador-naranjas-backend-multi5to/src/features/esp32/infrastructure/controllers"
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

	// Inicializar RabbitMQ (puede devolver nil si hay error)
	var rabbitMQ *core.RabbitMQ
	rabbitMQTmp, err := core.NewRabbitMQ()
	if err != nil {
		log.Printf("Advertencia: No se pudo conectar a RabbitMQ: %v", err)
		// Continuamos con rabbitMQ = nil
	} else {
		rabbitMQ = rabbitMQTmp
		// No uses defer rabbitMQ.Close() aquí, cerrará la conexión
		// antes de que se use
	}

	cajasDatabase := cajasInfrastructure.NewMySQL(database.Conn)
	getAllUseCase := cajasUseCases.NewGetAllCajaUseCase(cajasDatabase)
	getAllController := cajasControllers.NewGetAllCajaController(getAllUseCase)
	createCajaUseCase := cajasUseCases.NewCreateCajaUseCase(cajasDatabase)
	createCajaController := cajasControllers.NewCreateCajaController(createCajaUseCase)
	updateCajasUseCase := cajasUseCases.NewUpdateCajaUseCase(cajasDatabase)
	updateCajasController := cajasControllers.NewUpdateCajaController(updateCajasUseCase)
	deleteCajasUseCase := cajasUseCases.NewDeleteCajaUseCase(cajasDatabase)
	deleteCajasController := cajasControllers.NewDeleteCajaController(deleteCajasUseCase)
	getTop3CajasByLoteUseCase := cajasUseCases.NewGetTop3CajasByLoteUseCase(cajasDatabase)
	getTop3CajasByLoteController := cajasControllers.NewGetTop3CajasByLoteController(getTop3CajasByLoteUseCase)

	cajasRoutes := cajasInfrastructure.NewCajasRoutes(d.engine, getAllController, createCajaController, updateCajasController, deleteCajasController, getTop3CajasByLoteController)

	findActiveCajaByEsp32UseCase := cajasUseCases.NewFindActiveCajaByEsp32UseCase(cajasDatabase)

	// Crear el productor (funcionará incluso con rabbitMQ = nil)
	naranjaProducer := naranjasInfrastructure.NewProducer(rabbitMQ)

	naranjaDatabase := naranjasInfrastructure.NewMySQL(database.Conn)
	createNaranjaUseCase := naranjasUseCases.NewCreateNaranjaUseCase(
		naranjaDatabase,
		naranjaProducer,
		findActiveCajaByEsp32UseCase,
	)
	createNaranjaController := naranjasControllers.NewCreateNaranjaController(createNaranjaUseCase)
	getAllNaranjaUseCase := naranjasUseCases.NewGetAllUseCase(naranjaDatabase)
	getAllNaranjasController := naranjasControllers.NewGetAllController(getAllNaranjaUseCase)
	updateNaranjaUseCase := naranjasUseCases.NewUpdateNaranjaUseCase(naranjaDatabase)
	updateContollers := naranjasControllers.NewUpdateNaranjaController(updateNaranjaUseCase)
	getByCajaNaranjaUseCase := naranjasUseCases.NewGetByCajaNaranjaUseCase(naranjaDatabase)
	getByCajaNaranjaController := naranjasControllers.NewGetByCajaNaranjaController(getByCajaNaranjaUseCase)
	getByEsp32NaranjaUseCase := naranjasUseCases.NewGetByEsp32NaranjaUseCase(naranjaDatabase)
	getByEsp32NaranjaController := naranjasControllers.NewGetByEsp32NaranjaController(getByEsp32NaranjaUseCase)
	naranjasRoutes := naranjasInfrastructure.NewNaranjasRoutes(d.engine, createNaranjaController, getAllNaranjasController, updateContollers, getByEsp32NaranjaController, getByCajaNaranjaController)

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
	getUsersByJefeUseCase := usersUseCases.NewGetAllByJefeUseCase(userDataBase)
	getUsersByJefeController := usersControllers.NewGetAllByJefeController(getUsersByJefeUseCase)
	userRoutes := usersInfrastructure.NewUserRoutes(d.engine, createUserController, logInController, updateUserController, deleteUserController, getUserByIdController, getUserByUsernameController, getUsersController, getUsersByJefeController)

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
	updateLoteStatusUseCase := lotesUseCases.NewUpdateLoteStatusUseCase(lotesDatabase, cajasDatabase)
	updateLoteStatusController := lotesControllers.NewUpdateLoteStatusController(updateLoteStatusUseCase)
	getLoteByUserUseCase := lotesUseCases.NewGetLotesByUserUseCase(lotesDatabase)
	getLoteByUserController := lotesControllers.NewGetLotesByUserController(getLoteByUserUseCase)

	// Necesito esto JUSTO ACÁ, NO MOVER
	esp32Database := esp32Adapter.NewMysql(database.Conn)

	// Inicializar el caso de uso que crea lotes con cajas
	createLoteWithCajasUseCase := lotesUseCases.NewCreateLoteWithCajasUseCase(
		lotesDatabase,
		cajasDatabase,
		esp32Database,
	)
	createLoteWithCajasController := lotesControllers.NewCreateLoteWithCajasController(createLoteWithCajasUseCase)

	getLoteWithCajasUseCase := lotesUseCases.NewGetLoteWithCajasUseCase(lotesDatabase, cajasDatabase)
	getLoteWithCajasController := lotesControllers.NewGetLoteWithCajasController(getLoteWithCajasUseCase)
	getAllLotesWithCajasUseCase := lotesUseCases.NewGetAllLotesWithCajasUseCase(lotesDatabase, cajasDatabase)
	getAllLotesWithCajasController := lotesControllers.NewGetAllLotesWithCajasController(getAllLotesWithCajasUseCase)
	getLotesWithCajasByUserUseCase := lotesUseCases.NewGetLotesWithCajasByUserUseCase(lotesDatabase, cajasDatabase)
	getLotesWithCajasByUserController := lotesControllers.NewGetLotesWithCajasByUserController(getLotesWithCajasByUserUseCase)
	getLotesWithCajasByUserUseCaseWithRanges := lotesUseCases.NewGetLotesWithCajasByUserDateRangeUseCase(lotesDatabase, cajasDatabase)
	getLotesWithCajasByUserControllerWithRanges := lotesControllers.NewGetLotesWithCajasByUserDateRangeController(getLotesWithCajasByUserUseCaseWithRanges)

	lotesRoutes := lotesInfrastructure.NewLotesRoutes(
		d.engine,
		createLoteController,
		listAllLotesController,
		listLoteIdController,
		listLoteDateController,
		deleteLoteController,
		updateLoteControlerr,
		updateLoteStatusController,
		getLoteByUserController,
		createLoteWithCajasController,
		getLoteWithCajasController,
		getAllLotesWithCajasController,
		getLotesWithCajasByUserController,
		getLotesWithCajasByUserControllerWithRanges,
	)

	// Crear el productor de ESP32 (funcionará incluso con rabbitMQ = nil)
	esp32Producer := esp32Adapter.NewRabbitMQProducer(rabbitMQ)

	createEsp32UseCase := esp32UseCases.NewSaveEsp32(esp32Database)
	createEsp32Controller := esp32Controllers.NewCreateEsp32Controller(createEsp32UseCase)
	getEsp32ByUsernameUseCase := esp32UseCases.NewGetEsp32ByPropietarioUseCase(esp32Database)
	getEsp32ByUsernameController := esp32Controllers.NewGetEsp32ByPropietarioController(getEsp32ByUsernameUseCase)
	deleteEsp32UseCase := esp32UseCases.NewDeleteEsp32UseCase(esp32Database)
	deleteEsp32Controller := esp32Controllers.NewDeleteEsp32Controller(deleteEsp32UseCase)
	updateEsp32StatusUseCase := esp32UseCases.NewUpdateEsp32StatusUseCase(
		esp32Database,
		cajasDatabase,
		updateLoteStatusUseCase,
		esp32Producer,
	)
	updateEsp32StatusController := esp32Controllers.NewUpdateEsp32StatusController(updateEsp32StatusUseCase)
	getEsp32ByPropietarioAndStatusUseCase := esp32UseCases.NewGetEsp32ByPropietarioAndStatusUseCase(esp32Database)
	getEsp32ByPropietarioAndStatusController := esp32Controllers.NewGetEsp32ByPropietarioAndStatusController(getEsp32ByPropietarioAndStatusUseCase)
	sp32Routes := esp32Infrastructure.NewEsp32Routes(d.engine, createEsp32Controller, getEsp32ByUsernameController, deleteEsp32Controller, updateEsp32StatusController, getEsp32ByPropietarioAndStatusController)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	d.engine.Use(cors.New(config))

	cajasRoutes.SetupRoutes()
	sp32Routes.Run()
	naranjasRoutes.SetupRoutes()
	userRoutes.SetupRoutes()
	lotesRoutes.SetupRoutes()

	return d.engine.Run()
}
