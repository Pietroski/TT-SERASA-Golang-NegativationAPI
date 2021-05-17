package test

import (
	"fmt"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/factories"
	negativations "github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/services/negativation"
	mock_negativation "github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/services/negativation/mock"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func RandomNegativationWithoutID() negativations.NegativateParams {
	np := negativations.NegativateParams{
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

func CreateRandomNegativation() negativations.Negativations {
	np := negativations.Negativations{
		ID:               mock.RandomData.GenerateRandomID(int64(6543321)),
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

func TestNegativate(t *testing.T) {
	//rn := RandomNegativationWithoutID()
	//nrn := negativations.Negativations{
	//	ID:               1,
	//	CompanyDocument:  rn.CompanyDocument,
	//	CompanyName:      rn.CompanyName,
	//	CustomerDocument: rn.CustomerDocument,
	//	Value:            rn.Value,
	//	Contract:         rn.Contract,
	//	DebtDate:         rn.DebtDate,
	//	InclusionDate:    rn.InclusionDate,
	//}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_negativation.NewMockStore(ctrl)

	// build stubs
	//store.EXPECT().Negativate(gomock.Any(), rn).Times(1).Return(nrn, nil)

	// start test server and send request
	server := factories.Negativations.NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/v2/negativate")
	req, err := http.NewRequest(http.MethodPost, url, nil)
	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, req)

	// check response
	//require.Equal(t, http.StatusCreated, recorder.Code)
}

func TestGetNegativatedByID(t *testing.T) {
	rn := CreateRandomNegativation()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_negativation.NewMockStore(ctrl)

	// build stubs
	//store.EXPECT().GetNegativatedByID(gomock.Any(), gomock.Eq(rn.ID)).Times(1).Return(rn, nil)

	// start test server and send request
	server := factories.Negativations.NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/negativated/%d", rn.ID)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, req)

	// check response
	//require.Equal(t, http.StatusCreated, recorder.Code)
}