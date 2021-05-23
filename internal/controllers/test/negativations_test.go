package test

import (
	"bytes"
	"encoding/json"
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
	rn := RandomNegativationWithoutID()
	nrn := negativations.Negativations{
		ID:               1,
		CompanyDocument:  rn.CompanyDocument,
		CompanyName:      rn.CompanyName,
		CustomerDocument: rn.CustomerDocument,
		Value:            rn.Value,
		Contract:         rn.Contract,
		DebtDate:         rn.DebtDate,
		InclusionDate:    rn.InclusionDate,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_negativation.NewMockStore(ctrl)

	// build stubs
	store.EXPECT().Negativate(gomock.Any(), gomock.Eq(rn)).Times(1).Return(nrn, nil)

	// start test server and send request
	server := factories.Negativations.NewServer(store)
	recorder := httptest.NewRecorder()

	postReq, err := json.Marshal(rn)
	require.NoError(t, err)
	url := fmt.Sprintf("/v2/negativated")
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(postReq))
	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, req)

	// check response
	require.Equal(t, http.StatusCreated, recorder.Code)
}

func TestGetNegativatedByID(t *testing.T) {
	rn := CreateRandomNegativation()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_negativation.NewMockStore(ctrl)

	// build stubs
	store.
		EXPECT().
		GetNegativatedByID(gomock.Any(), gomock.Eq(rn.ID)).
		Times(1).
		Return(rn, nil)

	// start test server and send request
	server := factories.Negativations.NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/v2/negativated/%d", rn.ID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, req)

	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
}

func TestDeleteNegativated(t *testing.T) {
	rn := CreateRandomNegativation()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_negativation.NewMockStore(ctrl)

	// build stubs
	store.EXPECT().DeleteNegativated(gomock.Any(), gomock.Eq(rn.ID)).Times(1).Return(nil)

	// start test server and send request
	server := factories.Negativations.NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/v2/negativated/%d", rn.ID)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, req)

	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
}

func TestUpdateNegativated(t *testing.T) {
	rn := CreateRandomNegativation()
	rnp := negativations.UpdateNegativatedParams{
		ID:               rn.ID,
		CompanyDocument:  rn.CompanyDocument,
		CompanyName:      rn.CompanyName,
		CustomerDocument: rn.CustomerDocument,
		Value:            mock.RandomData.GenerateRandomDebt(0, 5_000),
		Contract:         rn.Contract,
		DebtDate:         rn.DebtDate,
		InclusionDate:    rn.InclusionDate,
	}
	rnr := negativations.Negativations{
		ID:               rn.ID,
		CompanyDocument:  rn.CompanyDocument,
		CompanyName:      rn.CompanyName,
		CustomerDocument: rn.CustomerDocument,
		Value:            rnp.Value,
		Contract:         rn.Contract,
		DebtDate:         rn.DebtDate,
		InclusionDate:    rn.InclusionDate,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_negativation.NewMockStore(ctrl)

	// build stubs
	store.EXPECT().UpdateNegativated(gomock.Any(), gomock.Eq(rnp)).Times(1).Return(rnr, nil)

	// start test server and send request
	server := factories.Negativations.NewServer(store)
	recorder := httptest.NewRecorder()

	putReq, err := json.Marshal(rnp)
	require.NoError(t, err)
	url := fmt.Sprintf("/v2/negativated")
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(putReq))
	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, req)

	// check response
	require.Equal(t, http.StatusCreated, recorder.Code)
}

func TestListNegativated(t *testing.T) {
	pageID := int32(1)
	pageSize := int32(5)

	rn := negativations.ListNegativatedParams{
		Limit:  pageSize,
		Offset: (pageID - 1) * pageSize,
	}

	rnList := make([]negativations.Negativations, pageSize)

	for i := 0; i < int(pageSize); i++ {
		rnList[i] = CreateRandomNegativation()
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_negativation.NewMockStore(ctrl)

	// build stubs
	store.EXPECT().ListNegativated(gomock.Any(), gomock.Eq(rn)).Times(1).Return(rnList, nil)

	// start test server and send request
	server := factories.Negativations.NewServer(store)
	recorder := httptest.NewRecorder()

	pagination, err := json.Marshal(rn)
	require.NoError(t, err)
	url := fmt.Sprintf("/v2/negativated?page_number=%d&page_size=%d", pageID, pageSize)
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(pagination))
	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, req)

	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
}
