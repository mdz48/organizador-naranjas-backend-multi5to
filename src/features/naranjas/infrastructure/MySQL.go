package infrastructure

import (
	"database/sql"
	"errors"
	"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (m *MySQL) Create(naranja domain.Naranja) (domain.Naranja, error) {
	result, err := m.db.Prepare("INSERT INTO naranjas (peso, tamano, color, hora, caja_fk, esp32_fk) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return domain.Naranja{}, err
	}
	defer result.Close()

	res, err := result.Exec(naranja.Peso, naranja.Tamaño, naranja.Color, naranja.Hora, naranja.CajaFK, naranja.Esp32FK)
	if err != nil {
		return domain.Naranja{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Naranja{}, err
	}

	naranja.ID = int(id)
	return naranja, nil
}

func (m *MySQL) GetById(id int) (domain.Naranja, error) {
	var naranja domain.Naranja

	err := m.db.QueryRow("SELECT id, peso, tamaño, color, hora, caja_fk, esp32_fk FROM naranjas WHERE id = ?", id).
		Scan(&naranja.ID, &naranja.Peso, &naranja.Tamaño, &naranja.Color, &naranja.Hora, &naranja.CajaFK, &naranja.Esp32FK)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Naranja{}, errors.New("naranja not found")
		}
		return domain.Naranja{}, err
	}

	return naranja, nil
}

func (m *MySQL) GetByCaja(cajaId int) ([]domain.Naranja, error) {
	rows, err := m.db.Query("SELECT id, peso, tamaño, color, hora, caja_fk, esp32_fk FROM naranjas WHERE caja_fk = ?", cajaId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var naranjas []domain.Naranja
	for rows.Next() {
		var naranja domain.Naranja
		if err := rows.Scan(&naranja.ID, &naranja.Peso, &naranja.Tamaño, &naranja.Color, &naranja.Hora, &naranja.CajaFK, &naranja.Esp32FK); err != nil {
			return nil, err
		}
		naranjas = append(naranjas, naranja)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return naranjas, nil
}

func (m *MySQL) GetAll() ([]domain.Naranja, error) {
	rows, err := m.db.Query("SELECT id, peso, tamaño, color, hora, caja_fk, esp32_fk FROM naranjas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var naranjas []domain.Naranja
	for rows.Next() {
		var naranja domain.Naranja
		if err := rows.Scan(&naranja.ID, &naranja.Peso, &naranja.Tamaño, &naranja.Color, &naranja.Hora, &naranja.CajaFK, &naranja.Esp32FK); err != nil {
			return nil, err
		}
		naranjas = append(naranjas, naranja)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return naranjas, nil
}

func (m *MySQL) Update(naranja domain.Naranja) (domain.Naranja, error) {
	stmt, err := m.db.Prepare("UPDATE naranjas SET peso = ?, tamaño = ?, color = ?, hora = ?, caja_fk = ?, esp32_fk = ? WHERE id = ?")
	if err != nil {
		return domain.Naranja{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(naranja.Peso, naranja.Tamaño, naranja.Color, naranja.Hora, naranja.CajaFK, naranja.Esp32FK, naranja.ID)
	if err != nil {
		return domain.Naranja{}, err
	}

	return naranja, nil
}

func (m *MySQL) Delete(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM naranjas WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no record found to delete")
	}

	return nil
}

func (m *MySQL) GetByEsp32(esp32Id string) ([]domain.Naranja, error) {
	rows, err := m.db.Query("SELECT id, peso, tamaño, color, hora, caja_fk, esp32_fk FROM naranjas WHERE esp32_fk = ?", esp32Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var naranjas []domain.Naranja
	for rows.Next() {
		var naranja domain.Naranja
		if err := rows.Scan(&naranja.ID, &naranja.Peso, &naranja.Tamaño, &naranja.Color, &naranja.Hora, &naranja.CajaFK, &naranja.Esp32FK); err != nil {
			return nil, err
		}
		naranjas = append(naranjas, naranja)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return naranjas, nil
}
