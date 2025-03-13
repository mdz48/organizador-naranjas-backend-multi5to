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
	query := "INSERT INTO cajas (descripcion, peso_total, precio, lote_fk, encargado_fk) VALUES (?, ?, ?, ?, ?)"
	result, err := m.db.Exec(query, caja.Descripcion, caja.PesoTotal, caja.Precio, caja.LoteFK, caja.EncargadoFK)
	if err != nil {
		return domain.Caja{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Caja{}, err
	}

	caja.ID = int(id)
	return caja, nil
}

func (m *MySQL) GetAll() ([]domain.Caja, error) {
	query := "SELECT id, descripcion, peso_total, precio, lote_fk, encargado_fk FROM cajas"
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cajas []domain.Caja
	for rows.Next() {
		var caja domain.Caja
		if err := rows.Scan(&caja.ID, &caja.Descripcion, &caja.PesoTotal, &caja.Precio, &caja.LoteFK, &caja.EncargadoFK); err != nil {
			return nil, err
		}
		cajas = append(cajas, caja)
	}

	return cajas, nil
}

func (m *MySQL) GetById(id int) (domain.Caja, error) {
	query := "SELECT id, descripcion, peso_total, precio, lote_fk, encargado_fk FROM cajas WHERE id = ?"
	row := m.db.QueryRow(query, id)

	var caja domain.Caja
	if err := row.Scan(&caja.ID, &caja.Descripcion, &caja.PesoTotal, &caja.Precio, &caja.LoteFK, &caja.EncargadoFK); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Caja{}, nil
		}
		return domain.Caja{}, err
	}

	return caja, nil
}

func (m *MySQL) GetByDescripcion(descripcion string) (domain.Caja, error) {
	query := "SELECT id, descripcion, peso_total, precio, lote_fk, encargado_fk FROM cajas WHERE descripcion = ?"
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
	query := "UPDATE cajas SET descripcion = ?, peso_total = ?, precio = ?, lote_fk = ?, encargado_fk = ? WHERE id = ?"
	_, err := m.db.Exec(query, caja.Descripcion, caja.PesoTotal, caja.Precio, caja.LoteFK, caja.EncargadoFK, caja.ID)
	if err != nil {
		return domain.Caja{}, err
	}

	return caja, nil
}

func (m *MySQL) Delete(id int) error {
	_, err := m.db.Exec("DELETE FROM cajas WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MySQL) GetByLote(loteID int) ([]domain.Caja, error) {
	query := "SELECT id, descripcion, peso_total, precio, lote_fk, encargado_fk FROM cajas WHERE lote_fk = ?"
	rows, err := m.db.Query(query, loteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cajas []domain.Caja
	for rows.Next() {
		var caja domain.Caja
		if err := rows.Scan(&caja.ID, &caja.Descripcion, &caja.PesoTotal, &caja.Precio, &caja.LoteFK, &caja.EncargadoFK); err != nil {
			return nil, err
		}
		cajas = append(cajas, caja)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cajas, nil
}
