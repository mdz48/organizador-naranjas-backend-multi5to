package infrastructure

import (
	"database/sql"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
)

type MYSQL struct {
	db *sql.DB
}

func NewMysql(db *sql.DB) *MYSQL {
	return &MYSQL{
		db: db,
	}
}

func (mysql *MYSQL) Save(user *entities.User) (*entities.User, error) {
	result, err := mysql.db.Exec("INSERT INTO users (name, password, rol) VALUES (?,?,?)", user.Name, user.Password, user.Rol)

	if err != nil {
		return &entities.User{}, err; 
	}

	id, errId := result.LastInsertId(); 

	if errId != nil {
		return &entities.User{}, errId; 
	}

	user.ID = int32(id); 

	return  user, err; 
}


func (mysql *MYSQL) LogIn(userLog *entities.UserLogIn) (*entities.User, error) {
	var user entities.User

	result := mysql.db.QueryRow("SELECT * FROM users WHERE name = ?", userLog.UserName);

	if errSearch := result.Err(); errSearch != nil {
		return &entities.User{}, errSearch; 
	}

	errScan := result.Scan(&user.ID, &user.Name, &user.Password, &user.Rol);

	if errScan != nil {
		return &entities.User{}, errScan; 
	}

	return &user, nil;
}