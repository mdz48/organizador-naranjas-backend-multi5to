package adapters

import (
	"database/sql"
	"log"
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/entities"
)

type MYSQL struct {
	conn *sql.DB
}

func NewMysql(conn *sql.DB) *MYSQL {
	return &MYSQL{
		conn: conn,
	}
}

func (mysql *MYSQL) Save(esp32 *entities.Esp32) (*entities.Esp32, error) {
	result, err := mysql.conn.Prepare("INSERT INTO esp32 (id, id_propietario) VALUES (?, ?)")
	if err != nil {
		log.Printf("error to prepare query: %v", err)
		return &entities.Esp32{}, err
	}
	defer result.Close()

	// Corregido el nombre del campo
	_, errInsert := result.Exec(esp32.Id, esp32.IdPropietario)
	if errInsert != nil {
		log.Printf("error to insert into esp32: %v", errInsert)
		return &entities.Esp32{}, errInsert
	}

	return esp32, nil
}

func (mysql *MYSQL) GetByPropietario(id int) (*entities.Esp32, error) {
	row := mysql.conn.QueryRow("SELECT id, id_propietario FROM esp32 WHERE id_propietario = ?", id)

	var esp32 entities.Esp32
	err := row.Scan(&esp32.Id, &esp32.IdPropietario)
	if err != nil {
		log.Printf("error to get esp32 by username: %v", err)
		return &entities.Esp32{}, err
	}

	return &esp32, nil
}

func (mysql *MYSQL) Delete(id string) error {
	result, err := mysql.conn.Prepare("DELETE FROM esp32 WHERE id = ?")
	if err != nil {
		log.Printf("error to prepare query: %v", err)
		return err
	}
	defer result.Close()

	_, errDelete := result.Exec(id)
	if errDelete != nil {
		log.Printf("error to delete esp32: %v", errDelete)
		return errDelete
	}

	return nil
}

// Añadido método para obtener por ID
func (mysql *MYSQL) GetById(id string) (*entities.Esp32, error) {
	row := mysql.conn.QueryRow("SELECT id, id_propietario FROM esp32 WHERE id = ?", id)

	var esp32 entities.Esp32
	err := row.Scan(&esp32.Id, &esp32.IdPropietario)
	if err != nil {
		log.Printf("error to get esp32 by id: %v", err)
		return &entities.Esp32{}, err
	}

	return &esp32, nil
}
