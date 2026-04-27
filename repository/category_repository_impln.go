package repository

import (
	"context"
	"database/sql"
	"errors"
	"perpus_golang/helper"
	"perpus_golang/model/domain"
)

type CategoryRepositoryImpln struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpln{}
}

func (repository *CategoryRepositoryImpln) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	//Menyiapkan code sql
	SQL := `INSERT INTO category(name) VALUES (?)`
	//Mengeksekusi SQL,context, dan transaksion
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.ErrorT(err)

	//Mengambil balik id dari auto increment SQL
	id, err := result.LastInsertId()
	helper.ErrorT(err)

	//Menyimpan id ke category
	category.Id = int(id)

	return category
}
func (repository *CategoryRepositoryImpln) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	//Menyiapkan code sql

	SQL := "UPDATE category SET name = ? WHERE id = ?"
	//Mengeksekusi SQL,context, dan  transaksi
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.ErrorT(err)

	return category
}
func (repository *CategoryRepositoryImpln) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.ErrorT(err)

}
func (repository *CategoryRepositoryImpln) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {

	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.ErrorT(err)

	defer rows.Close()
	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.ErrorT(err)
		return category, nil
	} else {

		return category, errors.New("Data Not Found")

	}
}
func (repository *CategoryRepositoryImpln) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM category"

	rows, err := tx.QueryContext(ctx, SQL, nil)
	helper.ErrorT(err)
	defer rows.Close()
	var categories = []domain.Category{}
	for rows.Next() {
		category := domain.Category{}

		err := rows.Scan(&category.Id, &category.Name)
		helper.ErrorT(err)

		categories = append(categories, category)

	}
	return categories
}
