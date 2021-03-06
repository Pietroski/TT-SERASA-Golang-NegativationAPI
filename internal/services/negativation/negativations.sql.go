// Code generated by sqlc. DO NOT EDIT.
// source: negativations.sql

package negativations

import (
	"context"
	"time"
)

const deleteNegativated = `-- name: DeleteNegativated :exec
DELETE FROM negativations
WHERE "id" = $1
`

func (q *Queries) DeleteNegativated(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteNegativatedStmt, deleteNegativated, id)
	return err
}

const getNegativatedByID = `-- name: GetNegativatedByID :one
SELECT id, company_document, company_name, customer_document, value, contract, debt_date, inclusion_date FROM "negativations"
WHERE "id" = $1
LIMIT 1
`

func (q *Queries) GetNegativatedByID(ctx context.Context, id int64) (Negativations, error) {
	row := q.queryRow(ctx, q.getNegativatedByIDStmt, getNegativatedByID, id)
	var i Negativations
	err := row.Scan(
		&i.ID,
		&i.CompanyDocument,
		&i.CompanyName,
		&i.CustomerDocument,
		&i.Value,
		&i.Contract,
		&i.DebtDate,
		&i.InclusionDate,
	)
	return i, err
}

const listNegativated = `-- name: ListNegativated :many
SELECT id, company_document, company_name, customer_document, value, contract, debt_date, inclusion_date FROM negativations
ORDER BY "id"
LIMIT $1
OFFSET $2
`

type ListNegativatedParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListNegativated(ctx context.Context, arg ListNegativatedParams) ([]Negativations, error) {
	rows, err := q.query(ctx, q.listNegativatedStmt, listNegativated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Negativations{}
	for rows.Next() {
		var i Negativations
		if err := rows.Scan(
			&i.ID,
			&i.CompanyDocument,
			&i.CompanyName,
			&i.CustomerDocument,
			&i.Value,
			&i.Contract,
			&i.DebtDate,
			&i.InclusionDate,
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

const negativate = `-- name: Negativate :one
INSERT INTO negativations ("company_document",
                           "company_name",
                           "customer_document",
                           "value",
                           "contract",
                           "debt_date",
                           "inclusion_date")
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, company_document, company_name, customer_document, value, contract, debt_date, inclusion_date
`

type NegativateParams struct {
	CompanyDocument  string    `json:"companyDocument"`
	CompanyName      string    `json:"companyName"`
	CustomerDocument string    `json:"customerDocument"`
	Value            float64   `json:"value"`
	Contract         string    `json:"contract"`
	DebtDate         time.Time `json:"debtDate"`
	InclusionDate    time.Time `json:"inclusionDate"`
}

func (q *Queries) Negativate(ctx context.Context, arg NegativateParams) (Negativations, error) {
	row := q.queryRow(ctx, q.negativateStmt, negativate,
		arg.CompanyDocument,
		arg.CompanyName,
		arg.CustomerDocument,
		arg.Value,
		arg.Contract,
		arg.DebtDate,
		arg.InclusionDate,
	)
	var i Negativations
	err := row.Scan(
		&i.ID,
		&i.CompanyDocument,
		&i.CompanyName,
		&i.CustomerDocument,
		&i.Value,
		&i.Contract,
		&i.DebtDate,
		&i.InclusionDate,
	)
	return i, err
}

const updateNegativated = `-- name: UpdateNegativated :one
UPDATE negativations
SET "company_document"  = $2,
    "company_name"      = $3,
    "customer_document" = $4,
    "value"            = $5,
    "contract"         = $6,
    "debt_date"         = $7,
    "inclusion_date"    = $8
WHERE id = $1
RETURNING id, company_document, company_name, customer_document, value, contract, debt_date, inclusion_date
`

type UpdateNegativatedParams struct {
	ID               int64     `json:"id"`
	CompanyDocument  string    `json:"companyDocument"`
	CompanyName      string    `json:"companyName"`
	CustomerDocument string    `json:"customerDocument"`
	Value            float64   `json:"value"`
	Contract         string    `json:"contract"`
	DebtDate         time.Time `json:"debtDate"`
	InclusionDate    time.Time `json:"inclusionDate"`
}

func (q *Queries) UpdateNegativated(ctx context.Context, arg UpdateNegativatedParams) (Negativations, error) {
	row := q.queryRow(ctx, q.updateNegativatedStmt, updateNegativated,
		arg.ID,
		arg.CompanyDocument,
		arg.CompanyName,
		arg.CustomerDocument,
		arg.Value,
		arg.Contract,
		arg.DebtDate,
		arg.InclusionDate,
	)
	var i Negativations
	err := row.Scan(
		&i.ID,
		&i.CompanyDocument,
		&i.CompanyName,
		&i.CustomerDocument,
		&i.Value,
		&i.Contract,
		&i.DebtDate,
		&i.InclusionDate,
	)
	return i, err
}
