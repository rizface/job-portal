package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"job-portal/app/model/request"
	"job-portal/app/setup"
	"job-portal/core/repository/auth_repo"
	"job-portal/core/service/auth_service/applicant_auth_service"
	"job-portal/core/service/auth_service/company_auth_service"
	"job-portal/helper"
	"job-portal/route"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func init() {
	helper.LoadConfig("../.env")
}

func TestRegister(t *testing.T) {
	//helper.LoadConfig(".env_test")
	db := helper.Connection()
	t.Run("success", func(t *testing.T) {
		t.Run("company", func(t *testing.T) {
			repo := auth_repo.NewAuth()
			result,err := repo.Register(db,context.Background(),request.Auth{
				Email:    "company@gmail.com",
				Password: "rahasia",
				Level:    "company",
			},"companies")
			assert.True(t, result)
			assert.Nil(t, err)
		})
		t.Run("applicant", func(t *testing.T) {
			repo := auth_repo.NewAuth()
			result,err := repo.Register(db,context.TODO(),request.Auth{
				Email:    "applicant@gmail.com",
				Password: "rahasia",
				Level:    "applicant",
			},"applicants")
			assert.True(t, result)
			assert.Nil(t, err)
		})
	})

	t.Run("success with controller", func(t *testing.T) {
		router := setup.Auth()
		rand.Seed(time.Now().Unix())
		email := strconv.Itoa(rand.Intn(1000)) + "@gmail.com"
		t.Run("applicant", func(t *testing.T) {
			request := request.Auth{
				Email:    email,
				Password: "rahasia",
			}
			payload,_ := json.Marshal(request)
			requestC := httptest.NewRequest(http.MethodPost,route.REGISTER_APPLICANT,bytes.NewReader(payload))
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder,requestC)
			assert.Equal(t, http.StatusOK,recorder.Code)
		})
		t.Run("company", func(t *testing.T) {
			request := request.Auth{
				Email:    email,
				Password: "rahasia",
			}
			payload,_ := json.Marshal(request)
			requestC := httptest.NewRequest(http.MethodPost,route.REGISTER_COMPANY,bytes.NewReader(payload))
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder,requestC)
			assert.Equal(t, http.StatusOK,recorder.Code)
		})
	})

	t.Run("failed empty request", func(t *testing.T) {
		router := setup.Auth()
		rand.Seed(time.Now().Unix())
		t.Run("applicant", func(t *testing.T) {
			request := request.Auth{
				Email:    "",
				Password: "rahasia",
			}
			payload,_ := json.Marshal(request)
			requestC := httptest.NewRequest(http.MethodPost,route.REGISTER_APPLICANT,bytes.NewReader(payload))
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder,requestC)
			assert.Equal(t, http.StatusBadRequest,recorder.Code)
		})
		t.Run("company", func(t *testing.T) {
			request := request.Auth{
				Email:    "",
				Password: "rahasia",
			}
			payload,_ := json.Marshal(request)
			requestC := httptest.NewRequest(http.MethodPost,route.REGISTER_COMPANY,bytes.NewReader(payload))
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder,requestC)
			assert.Equal(t, http.StatusBadRequest,recorder.Code)
		})
	})
}

func TestLogin(t *testing.T) {
	//helper.LoadConfig(".env_test")
	db := helper.Connection()
	valid := validator.New()
	t.Run("success", func(t *testing.T) {
		t.Run("applicant", func(t *testing.T) {
			auth := applicant_auth_service.NewAuth(db,valid,auth_repo.NewAuth())
			auth.Login(request.Auth{
				Email:    "applicant@gmail.com",
				Password: "rahasia",
			},"applicants")
		})
		t.Run("company", func(t *testing.T) {
			auth := company_auth_service.NewAuth(db,valid,auth_repo.NewAuth())
			auth.Login(request.Auth{
				Email:    "company@gmail.com",
				Password: "rahasia",
			},"companies")
		})
	})

	t.Run("success with controller", func(t *testing.T) {
		router := setup.Auth()
		t.Run("applicant", func(t *testing.T) {
			request := request.Auth{
				Email:    "applicant@gmail.com",
				Password: "rahasia",
			}
			payload,_ := json.Marshal(request)
			recorder := httptest.NewRecorder()
			requestC := httptest.NewRequest(http.MethodPost,route.LOGIN_APPLICANT,bytes.NewReader(payload))
			router.ServeHTTP(recorder,requestC)
			assert.Equal(t, http.StatusOK,recorder.Code)

		})
		t.Run("company", func(t *testing.T) {
			request := request.Auth{
				Email:    "company@gmail.com",
				Password: "rahasia",
			}
			payload,_ := json.Marshal(request)
			recorder := httptest.NewRecorder()
			requestC := httptest.NewRequest(http.MethodPost,route.LOGIN_COMPANY,bytes.NewReader(payload))
			router.ServeHTTP(recorder,requestC)
			assert.Equal(t, http.StatusOK,recorder.Code)
		})
	})

	t.Run("failed account not found", func(t *testing.T) {
		router := setup.Auth()
		t.Run("applicant", func(t *testing.T) {
			request := request.Auth{
				Email:    "app@gmail.com",
				Password: "rahasia",
			}
			payload,_ := json.Marshal(request)
			recorder := httptest.NewRecorder()
			requestC := httptest.NewRequest(http.MethodPost,route.LOGIN_APPLICANT,bytes.NewReader(payload))
			router.ServeHTTP(recorder,requestC)
			assert.Equal(t, http.StatusNotFound,recorder.Code)

		})
		t.Run("company", func(t *testing.T) {
			request := request.Auth{
				Email:    "app@gmail.com",
				Password: "rahasia",
			}
			payload,_ := json.Marshal(request)
			recorder := httptest.NewRecorder()
			requestC := httptest.NewRequest(http.MethodPost,route.LOGIN_COMPANY,bytes.NewReader(payload))
			router.ServeHTTP(recorder,requestC)
			assert.Equal(t, http.StatusNotFound,recorder.Code)
		})
	})
}
