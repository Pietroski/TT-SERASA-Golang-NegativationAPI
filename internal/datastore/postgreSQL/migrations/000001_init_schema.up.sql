CREATE TABLE IF NOT EXISTS "negativations"
(
    "id"               BIGSERIAL PRIMARY KEY,
    "companyDocument"  VARCHAR(255) NOT NULL,
    "companyName"      VARCHAR(255) NOT NULL,
    "customerDocument" VARCHAR(255) NOT NULL,
    "value"            float8       NOT NULL,
    "contract"         VARCHAR(255) NOT NULL,
    "debtDate"         DATE         NOT NULL,
    "inclusionDate"    DATE         NOT NULL
);
