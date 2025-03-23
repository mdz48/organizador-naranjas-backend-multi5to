package adapters

import (
	"database/sql"
	"log"
	"organizador-naranjas-backend-multi5to/src/features/sp32/domain/entities"
)

type MYSQL struct {
	conn *sql.DB
}

func NewMysql(conn *sql.DB) *MYSQL {
	return &MYSQL{
		conn: conn,
	}
}

func (mysql *MYSQL) Save(sp32 *entities.Sp32) (*entities.Sp32, error) {
	result, err := mysql.conn.Prepare("INSERT INTO sp32 (id_propietario) VALUES (?)")

	if err != nil {
		log.Printf("error to prepare query")
		return &entities.Sp32{}, err
	}

	resultInsert, errInsert := result.Exec(sp32.Id_propetario)

	if errInsert != nil {
		log.Printf("error to insert into sp32")
		return &entities.Sp32{}, errInsert
	}

	id, errId := resultInsert.LastInsertId()

	if errId != nil {
		log.Printf("error to get id")
	}

	sp32.Id = int(id)

	return sp32, nil
}