-- name: Negativate :one
INSERT INTO negativations ("companyDocument",
                           "companyName",
                           "customerDocument",
                           "value",
                           "contract",
                           "debtDate",
                           "inclusionDate")
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateNegativated :one
UPDATE negativations
SET "companyDocument"  = $2,
    "companyName"      = $3,
    "customerDocument" = $4,
    "value"            = $5,
    "contract"         = $6,
    "debtDate"         = $7,
    "inclusionDate"    = $8
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
