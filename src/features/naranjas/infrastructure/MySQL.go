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
	query := "INSERT INTO cajas (descripcion, peso_total, precio, lote_fk, encargado_fk) VALUES (?, ?, ?, ?, ?)"
	queryCaja := "SELECT id FROM cajas WHERE hora_creacion < ? AND hora_fin > ?"

    var idCaja int
    err := m.db.QueryRow(queryCaja, naranja.Hora, naranja.Hora).Scan(&idCaja)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return domain.Naranja{}, errors.New("no se encontró una caja disponible para la hora especificada")
        }
        return domain.Naranja{}, err
    }

	result, err := m.db.Exec(query, naranja.Peso, naranja.Tamaño, naranja.Color, naranja.Hora, idCaja)
	if err != nil {
		return domain.Naranja{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Naranja{}, err
	}

	naranja.ID = int(id)
	return naranja, nil
}
/*
func (m *MySQL) GetAll() ([]cajas.Caja, error) {
	query := "SELECT id, descripcion, peso_total, precio, lote_fk, encargado_fk FROM cajas"
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cajas []cajas.Caja
	for rows.Next() {
		var caja CajasRoutes
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
*/