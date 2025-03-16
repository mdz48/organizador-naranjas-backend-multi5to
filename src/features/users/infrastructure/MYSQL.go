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
	result, err := mysql.db.Exec("INSERT INTO users (name, password, rol, email, username) VALUES (?,?,?,?,?)", user.Name, user.Password, user.Rol, user.Email, user.Username)
	if err != nil {
		return &entities.User{}, err
	}
	id, errId := result.LastInsertId()
	if errId != nil {
		return &entities.User{}, errId
	}
	user.ID = int32(id)
	return user, err
}

func (mysql *MYSQL) LogIn(userLog *entities.UserLogIn) (*entities.User, error) {
	var user entities.User
	err := mysql.db.QueryRow("SELECT id, name, password, rol, email, username FROM users WHERE username = ?", userLog.Username).Scan(&user.ID, &user.Name, &user.Password, &user.Rol, &user.Email, &user.Username)
	if err != nil {
		return &entities.User{}, err
	}
	return &user, nil
}

func (mysql *MYSQL) Update(user *entities.User) (*entities.User, error) {
	result, err := mysql.db.Exec("UPDATE users SET name = ?, password = ?, rol = ?, email = ?, username = ? WHERE id = ?",
		user.Name, user.Password, user.Rol, user.Email, user.Username, user.ID)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	return user, nil
}

func (mysql *MYSQL) Delete(user *entities.User) (*entities.User, error) {
	result, err := mysql.db.Exec("DELETE FROM users WHERE id = ?", user.ID)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	return user, nil
}
