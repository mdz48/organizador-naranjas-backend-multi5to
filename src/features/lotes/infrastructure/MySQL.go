package infrastructure

import (
	"database/sql"
	"log"
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) Create(lote domain.Lote) (domain.Lote, error) {
	result, err := mysql.db.Exec("INSERT INTO lote (fecha, observaciones) VALUES (?,?)", lote.Fecha, lote.Observaciones)

	if err != nil {
		return domain.Lote{}, err
	}

	id, errId := result.LastInsertId()

	if errId != nil {
		return domain.Lote{}, errId
	}

	lote.ID = int(id)

	return lote, nil
}

func (mysql *MySQL) GetAll() ([]domain.Lote, error) {
	var lotes []domain.Lote
	result, err := mysql.db.Query("SELECT * FROM lote")

	if err != nil {
		return nil, err
	}

	for result.Next() {
		var lote domain.Lote
		errScan := result.Scan(&lote.ID, &lote.Fecha, &lote.Observaciones)
		if errScan != nil {
			log.Printf("error to scan lote!")
		}

		lotes = append(lotes, lote)
	}

	return lotes, nil
}

func (mysql *MySQL) GetById(id int) (domain.Lote, error) {
	var lote domain.Lote
	result := mysql.db.QueryRow("SELECT * FROM lote WHERE id = ?", id)

	if err := result.Err(); err != nil {
		return domain.Lote{}, err
	}

	errScan := result.Scan(&lote.ID, &lote.Fecha, &lote.Observaciones)

	if errScan != nil {
		return domain.Lote{}, errScan
	}

	return lote, nil
}

func (mysql *MySQL) Delete(id int) error {
	_, err := mysql.db.Exec("DELETE FROM lote WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func (mysql *MySQL) Update(id int, lote domain.Lote) (domain.Lote, error) {
	log.Printf("message %s", lote);
	_, err := mysql.db.Query("UPDATE lote SET fecha = ?, observaciones = ? WHERE id = ?", lote.Fecha, lote.Observaciones, id)

	if err != nil {
		return domain.Lote{}, err
	}

	return lote, err
}

func (mysql *MySQL) GetByDate(date string) ([]domain.Lote, error) {
	var lotes []domain.Lote
	result, err := mysql.db.Query("SELECT * FROM lote WHERE fecha = ?", date)

	if err != nil {
		return nil, err
	}

	for result.Next() {
		var lote domain.Lote

		errScan := result.Scan(&lote.ID, &lote.Fecha, &lote.Observaciones)
		if errScan != nil {
			return nil, errScan
		}

		lotes = append(lotes, lote)
	}

	return lotes, nil
}
