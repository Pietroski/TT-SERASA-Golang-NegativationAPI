package negativations

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/test/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomNegativation() NegativateParams {
	np := NegativateParams{
		CompanyDocument:  mock.RandomData.GenerateRandomDocument(14),
		CompanyName:      mock.RandomData.GenerateCompanyName(),
		CustomerDocument: mock.RandomData.GenerateRandomDocument(11),
		Value:            mock.RandomData.GenerateRandomDebt(0, 5_000),
		Contract:         mock.RandomData.GenerateRandomContract(),
		DebtDate:         mock.RandomData.GenerateRandomDate(), // time.Time{},
		InclusionDate:    mock.RandomData.GenerateRandomDate(), // time.Time{},
	}

	return np
}

func CreateNegativationOnDB(t *testing.T) Negativations {
	negativationArgs := CreateRandomNegativation()

	ctx := context.Background()
	negativation, err := testQueries.Negativate(ctx, negativationArgs)

	require.NoError(t, err)
	require.NotEmpty(t, negativation)

	require.NotZero(t, negativation.ID)
	require.NotZero(t, negativation.Value)

	return negativation
}

func TestNegativate(t *testing.T) {
	CreateNegativationOnDB(t)
}

func TestGetNegativatedByID(t *testing.T) {
	fmt.Println("TestGetNegativatedByID")

	neg := CreateNegativationOnDB(t)

	ctx := context.Background()
	negativation, err := testQueries.GetNegativatedByID(ctx, neg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, negativation)

	require.Equal(t, neg.ID, negativation.ID)
	require.Equal(t, neg.CompanyDocument, negativation.CompanyDocument)
	require.Equal(t, neg.CompanyName, negativation.CompanyName)
	require.Equal(t, neg.CustomerDocument, negativation.CustomerDocument)
	require.Equal(t, neg.Value, negativation.Value)
	require.Equal(t, neg.DebtDate, negativation.DebtDate)
	require.Equal(t, neg.InclusionDate, negativation.InclusionDate)
}

func TestUpdateNegativated(t *testing.T) {
	fmt.Println("TestUpdateNegativated")

	neg := CreateNegativationOnDB(t)

	var newCompName = "ABC L.T.D.A"

	args := UpdateNegativatedParams(neg)
	args.CompanyName = newCompName

	ctx := context.Background()
	updatedNeg, err := 	testQueries.UpdateNegativated(ctx, args)
	require.NoError(t, err)
	require.NotEmpty(t, updatedNeg)

	require.Equal(t, newCompName, updatedNeg.CompanyName)
	require.NotEqual(t, neg.CompanyName, updatedNeg.CompanyName)

	require.Equal(t, neg.ID, updatedNeg.ID)
	require.Equal(t, neg.CompanyDocument, updatedNeg.CompanyDocument)

	require.Equal(t, neg.CustomerDocument, updatedNeg.CustomerDocument)
	require.Equal(t, neg.Contract, updatedNeg.Contract)
	require.Equal(t, neg.Value, updatedNeg.Value)
	require.Equal(t, neg.DebtDate, updatedNeg.DebtDate)
	require.Equal(t, neg.InclusionDate, updatedNeg.InclusionDate)

}

func TestListNegativated(t *testing.T) {
	fmt.Println("TestListNegativated")

	for i := 0; i < 10; i++ {
		CreateNegativationOnDB(t)
	}

	arg := ListNegativatedParams{
		Limit:  5,
		Offset: 5,
	}

	ctx := context.Background()
	negativations, err := testQueries.ListNegativated(ctx, arg)
	require.NoError(t, err)
	require.Len(t, negativations, 5)

	for _, neg := range negativations {
		require.NotEmpty(t, neg)
	}
}

func TestDeleteNegativated(t *testing.T) {
	fmt.Println("TestDeleteNegativated")

	neg := CreateNegativationOnDB(t)
	ctx := context.Background()
	err := testQueries.DeleteNegativated(ctx, neg.ID)
	require.NoError(t, err)

	suppNeg, err := testQueries.GetNegativatedByID(ctx, neg.ID)
	require.Error(t, err)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, suppNeg)
}
