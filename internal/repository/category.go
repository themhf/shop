package repository

import (
	"database/sql"
	"shop/internal/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (r *CategoryRepository) Insert(category *model.Category) error {
	query := `
	INSERT INTO categories
	(parent_id, title, slug, priority, level, is_active, description)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := r.db.Exec(
		query,
		category.ParentID,
		category.Title,
		category.Slug,
		category.Priority,
		category.Level,
		category.IsActive,
		category.Description,
	)

	return err

}

func (r *CategoryRepository) GetByID(id int64) (*model.Category, error) {
	query := `
		SELECT id, parent_id, title, slug, priority, level,
		       is_active, description
		FROM categories
		WHERE id = $1
	`

	var category model.Category

	err := r.db.QueryRow(query, id).Scan(
		&category.ID,
		&category.ParentID,
		&category.Title,
		&category.Slug,
		&category.Priority,
		&category.Level,
		&category.IsActive,
		&category.Description,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) GetByParent(parentID *int64) ([]model.Category, error) {

	query := `
		SELECT id, parent_id, title, slug, priority, level,
		       is_active, description
		FROM categories
		WHERE parent_id = $1
		ORDER BY priority
	`

	rows, err := r.db.Query(query, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category

	for rows.Next() {
		var category model.Category

		err := rows.Scan(
			&category.ID,
			&category.ParentID,
			&category.Title,
			&category.Slug,
			&category.Priority,
			&category.Level,
			&category.IsActive,
			&category.Description,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) Update(category *model.Category) error {

	query := `
		UPDATE categories
		SET
			parent_id = $1,
			title = $2,
			slug = $3,
			priority = $4,
			level = $5,
			is_active = $6,
			description = $7
		WHERE id = $8
	`

	_, err := r.db.Exec(
		query,
		category.ParentID,
		category.Title,
		category.Slug,
		category.Priority,
		category.Level,
		category.IsActive,
		category.Description,
		category.ID,
	)

	return err
}

func (r *CategoryRepository) Delete(id int64) error {

	query := `
		DELETE FROM categories
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)

	return err
}

func (r *CategoryRepository) UpdateStatus(id int64, active bool) error {

	query := `
		UPDATE categories
		SET is_active = $1
		WHERE id = $2
	`

	_, err := r.db.Exec(query, active, id)

	return err
}

func (r *CategoryRepository) ExistsByTitle(parentID *int64, title string) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM categories
			WHERE parent_id IS NOT DISTINCT FROM $1
			  AND title = $2
		)
	`

	var exists bool

	err := r.db.QueryRow(query, parentID, title).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
