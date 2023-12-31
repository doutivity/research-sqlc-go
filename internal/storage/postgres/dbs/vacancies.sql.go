// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: vacancies.sql

package dbs

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const vacancyByCompanies = `-- name: VacancyByCompanies :many
SELECT v.id    AS vacancy_id,
       v.title AS vacancy_title,
       c.id    AS company_id,
       c.alias AS company_alias,
       c.name  AS company_name
FROM vacancies v
         INNER JOIN companies c ON (v.company_id = c.id)
WHERE c.id = ANY ($1::BIGINT[])
`

type VacancyByCompaniesRow struct {
	VacancyID    int64
	VacancyTitle string
	CompanyID    int64
	CompanyAlias string
	CompanyName  string
}

func (q *Queries) VacancyByCompanies(ctx context.Context, companies []int64) ([]VacancyByCompaniesRow, error) {
	rows, err := q.query(ctx, q.vacancyByCompaniesStmt, vacancyByCompanies, pq.Array(companies))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []VacancyByCompaniesRow
	for rows.Next() {
		var i VacancyByCompaniesRow
		if err := rows.Scan(
			&i.VacancyID,
			&i.VacancyTitle,
			&i.CompanyID,
			&i.CompanyAlias,
			&i.CompanyName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const vacancyNew = `-- name: VacancyNew :exec
INSERT INTO vacancies (company_id, title, created_by, created_at)
VALUES ($1, $2, $3, $4)
`

type VacancyNewParams struct {
	CompanyID int64
	Title     string
	CreatedBy int64
	CreatedAt time.Time
}

func (q *Queries) VacancyNew(ctx context.Context, arg VacancyNewParams) error {
	_, err := q.exec(ctx, q.vacancyNewStmt, vacancyNew,
		arg.CompanyID,
		arg.Title,
		arg.CreatedBy,
		arg.CreatedAt,
	)
	return err
}
