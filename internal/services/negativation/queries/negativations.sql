-- name: Negativate :one
INSERT INTO negativations ("company_document",
                           "company_name",
                           "customer_document",
                           "value",
                           "contract",
                           "debt_date",
                           "inclusion_date")
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateNegativated :one
UPDATE negativations
SET "company_document"  = $2,
    "company_name"      = $3,
    "customer_document" = $4,
    "value"            = $5,
    "contract"         = $6,
    "debt_date"         = $7,
    "inclusion_date"    = $8
WHERE id = $1
RETURNING *;

-- name: GetNegativatedByID :one
SELECT * FROM "negativations"
WHERE "id" = $1
LIMIT 1;

-- name: ListNegativated :many
SELECT * FROM negativations
ORDER BY "id"
LIMIT $1
OFFSET $2;

-- name: DeleteNegativated :exec
DELETE FROM negativations
WHERE "id" = $1;
