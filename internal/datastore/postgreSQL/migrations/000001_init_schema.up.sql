CREATE TABLE IF NOT EXISTS "negativations"
(
    "id"               BIGSERIAL PRIMARY KEY,
    "company_document"  VARCHAR(255) NOT NULL,
    "company_name"      VARCHAR(255) NOT NULL,
    "customer_document" VARCHAR(255) NOT NULL,
    "value"            float8       NOT NULL,
    "contract"         VARCHAR(255) NOT NULL,
    "debt_date"         DATE         NOT NULL,
    "inclusion_date"    DATE         NOT NULL
);
