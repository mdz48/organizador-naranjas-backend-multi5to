package infrastructure

import (
	"database/sql"
	"errors"
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (m *MySQL) Create(caja domain.Caja) (domain.Caja, error) {
	result, err := m.db.Prepare("INSERT INTO cajas (descripcion, peso_total, precio, hora_inicio, hora_fin, lote_fk, encargado_fk, cantidad) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return domain.Caja{}, err
	}
	defer result.Close()

	res, err := result.Exec(caja.Descripcion, caja.PesoTotal, caja.Precio, caja.HoraInicio, caja.HoraFin, caja.LoteFK, caja.EncargadoFK, caja.Cantidad)
	if err != nil {
		return domain.Caja{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Caja{}, err
	}

	caja.ID = int(id)
	return caja, nil
}

func (m *MySQL) GetAll() ([]domain.Caja, error) {
	rows, err := m.db.Query("SELECT id, descripcion, peso_total, precio, hora_inicio, hora_fin, lote_fk, encargado_fk, cantidad FROM cajas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cajas []domain.Caja
	for rows.Next() {
		var caja domain.Caja
		if err := rows.Scan(&caja.ID, &caja.Descripcion, &caja.PesoTotal, &caja.Precio, &caja.HoraInicio, &caja.HoraFin, &caja.LoteFK, &caja.EncargadoFK, &caja.Cantidad); err != nil {
			return nil, err
		}
		cajas = append(cajas, caja)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cajas, nil
}

func (m *MySQL) GetById(id int) (domain.Caja, error) {
	var caja domain.Caja

	err := m.db.QueryRow("SELECT id, descripcion, peso_total, precio, hora_inicio, hora_fin, lote_fk, encargado_fk, cantidad FROM cajas WHERE id = ?", id).
		Scan(&caja.ID, &caja.Descripcion, &caja.PesoTotal, &caja.Precio, &caja.HoraInicio, &caja.HoraFin, &caja.LoteFK, &caja.EncargadoFK, &caja.Cantidad)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Caja{}, errors.New("caja not found")
		}
		return domain.Caja{}, err
	}

	return caja, nil
}

func (m *MySQL) GetByDescripcion(descripcion string) (domain.Caja, error) {
	query := "SELECT id, descripcion, peso_total, precio, lote_fk, encargado_fk, cantidad FROM cajas WHERE descripcion = ?"
	row := m.db.QueryRow(query, descripcion)

	var caja domain.Caja
	if err := row.Scan(&caja.ID, &caja.Descripcion, &caja.PesoTotal, &caja.Precio, &caja.LoteFK, &caja.EncargadoFK); err != nil {
		if err == sql.ErrNoRows {
			return domain.Caja{}, nil
		}
		return domain.Caja{}, err
	}

	return caja, nil
}

func (m *MySQL) Update(caja domain.Caja) (domain.Caja, error) {
	stmt, err := m.db.Prepare("UPDATE cajas SET descripcion = ?, peso_total = ?, precio = ?, hora_inicio = ?, hora_fin = ?, lote_fk = ?, encargado_fk = ?, cantidad = ? WHERE id = ?")
	if err != nil {
		return domain.Caja{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(caja.Descripcion, caja.PesoTotal, caja.Precio, caja.HoraInicio, caja.HoraFin, caja.LoteFK, caja.EncargadoFK, caja.Cantidad, caja.ID)
	if err != nil {
		return domain.Caja{}, err
	}

	return caja, nil
}

func (m *MySQL) Delete(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM cajas WHERE id = ?")
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

func (m *MySQL) GetByLote(loteID int) ([]domain.Caja, error) {
	rows, err := m.db.Query("SELECT id, descripcion, peso_total, precio, hora_inicio, hora_fin, lote_fk, encargado_fk, cantidad FROM cajas WHERE lote_fk = ?", loteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cajas []domain.Caja
	for rows.Next() {
		var caja domain.Caja
		if err := rows.Scan(&caja.ID, &caja.Descripcion, &caja.PesoTotal, &caja.Precio, &caja.HoraInicio, &caja.HoraFin, &caja.LoteFK, &caja.EncargadoFK, &caja.Cantidad); err != nil {
			return nil, err
		}
		cajas = append(cajas, caja)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cajas, nil
}

func (m *MySQL) GetByEncargado(encargadoId int) ([]domain.Caja, error) {
	rows, err := m.db.Query("SELECT id, descripcion, peso_total, precio, hora_inicio, hora_fin, lote_fk, encargado_fk, cantidad FROM cajas WHERE encargado_fk = ?", encargadoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cajas []domain.Caja
	for rows.Next() {
		var caja domain.Caja
		if err := rows.Scan(&caja.ID, &caja.Descripcion, &caja.PesoTotal, &caja.Precio, &caja.HoraInicio, &caja.HoraFin, &caja.LoteFK, &caja.EncargadoFK, &caja.Cantidad); err != nil {
			return nil, err
		}
		cajas = append(cajas, caja)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cajas, nil
}
