package data

import (
	"context"
	"fmt"
	"gomysql/pkg/location"
)

type LocationRepository struct {
	Data *Data
}

func (lo *LocationRepository) GetOne(ctx context.Context, idlocation uint) (location.Location, error) {
	fmt.Println("idlocation", idlocation)
	q := ` SELECT idlocation, nombre, tipo, dimension
	FROM rickandmorty.location WHERE idlocation = ?;`

	row := lo.Data.DB.QueryRowContext(ctx, q, idlocation)
	var l location.Location

	err := row.Scan(&l.ID, &l.Nombre, &l.Tipo, &l.Dimension)

	if err != nil {
		return location.Location{}, err
	}

	return l, nil
}

func (lo *LocationRepository) GetAll(ctx context.Context) ([]location.Location, error) {
	q := `SELECT idlocation, nombre, tipo, dimension FROM location;`

	rows, err := lo.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var locations []location.Location
	for rows.Next() {
		var l location.Location
		rows.Scan(&l.ID, &l.Nombre, &l.Tipo, &l.Dimension)
		locations = append(locations, l)
	}

	return locations, nil
}

func (lo *LocationRepository) GetByUser(ctx context.Context, idlocation uint) ([]location.Location, error) {
	q := `
    SELECT idlocation, nombre, tipo, dimension FROM location
        WHERE idlocation = ?;
    `

	rows, err := lo.Data.DB.QueryContext(ctx, q, idlocation)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var locations []location.Location
	for rows.Next() {
		var l location.Location
		rows.Scan(&l.ID, &l.Nombre, &l.Tipo, &l.Dimension)
		locations = append(locations, l)
	}

	return locations, nil
}

/*
	INSERTAR
*/

func (lo *LocationRepository) Create(ctx context.Context, l *location.Location) error {
	q := `
    INSERT INTO location ( nombre, tipo, dimension)
        VALUES ($1, $2, $3)
        RETURNING idlocation;
    `

	stmt, err := lo.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, &l.ID, &l.Nombre, &l.Tipo, &l.Dimension)

	err = row.Scan(&l.ID)
	if err != nil {
		return err
	}

	return nil
}

/*
	UPDATE
*/

func (lo *LocationRepository) Update(ctx context.Context, idlocation uint, l location.Location) (location.Location, error) {
	q := `UPDATE location set name=$1, price=$2, description=$3, updated_at=$4 WHERE id=?; `

	stmt, err := lo.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return l, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, &l.ID, &l.Nombre, &l.Tipo, &l.Dimension, idlocation,
	)
	if err != nil {
		return l, err
	}

	return l, nil
}

/*
	DELETE
*/

func (lo *LocationRepository) Delete(ctx context.Context, idlocation uint) error {
	q := `DELETE FROM location WHERE idlocation=$1;`

	stmt, err := lo.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, idlocation)
	if err != nil {
		return err
	}

	return nil
}
