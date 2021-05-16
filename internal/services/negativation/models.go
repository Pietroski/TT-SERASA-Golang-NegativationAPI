// Code generated by sqlc. DO NOT EDIT.

package negativations

import (
	"time"
)

type Negativations struct {
	ID               int64     `json:"id"`
	CompanyDocument  string    `json:"companyDocument"`
	CompanyName      string    `json:"companyName"`
	CustomerDocument string    `json:"customerDocument"`
	Value            float64   `json:"value"`
	Contract         string    `json:"contract"`
	DebtDate         time.Time `json:"debtDate"`
	InclusionDate    time.Time `json:"inclusionDate"`
}
