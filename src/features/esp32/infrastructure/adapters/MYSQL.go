package adapters

import (
	"database/sql"
	"fmt"
	"log"
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/entities"
)

type MYSQL struct {
	conn *sql.DB
}

// Delete implements ports.IEsp32.
func (mysql *MYSQL) Delete(id string) error {
	query := `DELETE FROM esp32 WHERE id = ?`

	result, err := mysql.conn.Prepare(query)

	if err != nil {
		fmt.Printf("error to prepare query")
		return err
	}

	defer result.Close()

	_, errQuery := result.Exec(id)

	if errQuery != nil {
		fmt.Printf("error to exec query")
		return errQuery
	}

	return nil
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

func (mysql *MYSQL) GetByPropietario(id int) ([]entities.Esp32, error) {
	rows, err := mysql.conn.Query("SELECT id, id_propietario FROM esp32 WHERE id_propietario = ?", id)
	if err != nil {
		log.Printf("error querying esp32 devices: %v", err)
		return nil, err
	}
	defer rows.Close()

	var devices []entities.Esp32
	for rows.Next() {
		var esp32 entities.Esp32
		if err := rows.Scan(&esp32.Id, &esp32.IdPropietario); err != nil {
			log.Printf("error scanning esp32 row: %v", err)
			return nil, err
		}
		devices = append(devices, esp32)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error iterating esp32 rows: %v", err)
		return nil, err
	}
	return devices, nil
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
