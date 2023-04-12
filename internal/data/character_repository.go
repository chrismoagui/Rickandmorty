package data

import (
	"context"
	"fmt"
	"gomysql/pkg/character"
)

type CharacterRepository struct {
	Data *Data
}

func (ch *CharacterRepository) GetAll(ctx context.Context) ([]character.Character, error) {
	fmt.Println("GetAll")
	q := `SELECT *  rickandmorty.character;`
	fmt.Println(q)

	rows, err := ch.Data.DB.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var characters []character.Character
	for rows.Next() {
		var c character.Character
		rows.Scan(&c.ID, &c.Nombre, &c.Estado, &c.Especie)
		characters = append(characters, c)
	}

	return characters, nil
}

//obtener un solo usuario

func (ch *CharacterRepository) GetOne(ctx context.Context, idcharacter uint) (character.Character, error) {
	q := `
    SELECT idcharacter, nombre, estado, especie
        FROM rickandmorty.character
		 WHERE idcharacter = $1;
    `

	row := ch.Data.DB.QueryRowContext(ctx, q, idcharacter)

	var c character.Character
	err := row.Scan(&c.ID, &c.Nombre, &c.Estado, &c.Especie)

	if err != nil {
		return character.Character{}, err
	}

	return c, nil
}

func (ch *CharacterRepository) GetByUsername(ctx context.Context, nombre string) (character.Character, error) {
	q := `
    SELECT idcharacter, nombre, estado, especie
        FROM character WHERE nombre = $1;
    `

	row := ch.Data.DB.QueryRowContext(ctx, q, nombre)

	var c character.Character
	err := row.Scan(&c.ID, &c.Nombre, &c.Estado, &c.Especie)
	if err != nil {
		return character.Character{}, err
	}

	return c, nil
}

func (ch *CharacterRepository) Create(ctx context.Context, c *character.Character) error {
	q := `
    INSERT INTO character (nombre,estado,especie)
        VALUES ($1, $2, $3)
        RETURNING idcharacter;
    `

	stmt, err := ch.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, c.Nombre, c.Estado, c.Especie)

	err = row.Scan(&c.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ch *CharacterRepository) Update(ctx context.Context, idcharacter uint, c character.Character) (character.Character, error) {
	q := `
    UPDATE character set nombre=$1, estado=$2, especie=$3
        WHERE idcharacter=$4;
    `

	stmt, err := ch.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return c, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, c.Nombre, c.Estado, c.Especie,
		idcharacter,
	)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (ch *CharacterRepository) Delete(ctx context.Context, idcharacter uint) error {
	q := `DELETE FROM chracter WHERE id=$1;`

	stmt, err := ch.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, idcharacter)
	if err != nil {
		return err
	}

	return nil
}
