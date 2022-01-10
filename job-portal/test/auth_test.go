package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	request2 "job-portal/app/model/request"
	"job-portal/app/setup"
	"job-portal/route"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	t.Skip("skip karena memakan resource email service")
	router := setup.Auth()
	t.Run("register applicant", func(t *testing.T) {
		payload := generateAuthRequest()
		t.Run("success", func(t *testing.T) {
			payloadJson,_ := json.Marshal(payload)
			request := httptest.NewRequest(http.MethodPost,route.REGISTER_APPLICANT,bytes.NewReader(payloadJson))
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder,request)
			assert.Equal(t, http.StatusOK,recorder.Code)
		})
		t.Run("failed", func(t *testing.T) {
			t.Run("duplicate", func(t *testing.T) {
				//payload := request2.Auth{
				//	Email:    "malfarizzai33@gmail.com",
				//	Password: "rahasia",
				//}
				payloadJson,_ := json.Marshal(payload)
				request := httptest.NewRequest(http.MethodPost,route.REGISTER_APPLICANT,bytes.NewReader(payloadJson))
				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusConflict,recorder.Code)
			})
			t.Run("bad request", func(t *testing.T) {
				payload := request2.Auth{}
				payloadJson,_ := json.Marshal(payload)
				request := httptest.NewRequest(http.MethodPost,route.REGISTER_APPLICANT,bytes.NewReader(payloadJson))
				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusBadRequest,recorder.Code)
			})
		})
	})

	t.Run("register company", func(t *testing.T) {
		payload := generateAuthRequest()
		t.Run("success", func(t *testing.T) {
			payloadJson,_ := json.Marshal(payload)
			request := httptest.NewRequest(http.MethodPost,route.REGISTER_COMPANY,bytes.NewReader(payloadJson))
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder,request)
			assert.Equal(t, http.StatusOK,recorder.Code)
		})
		t.Run("failed", func(t *testing.T) {
			t.Run("duplicate", func(t *testing.T) {
				payloadJson,_ := json.Marshal(payload)
				request := httptest.NewRequest(http.MethodPost,route.REGISTER_COMPANY,bytes.NewReader(payloadJson))
				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusConflict,recorder.Code)
			})
			t.Run("bad request", func(t *testing.T) {
				payload := request2.Auth{}
				payloadJson,_ := json.Marshal(payload)
				request := httptest.NewRequest(http.MethodPost,route.REGISTER_COMPANY,bytes.NewReader(payloadJson))
				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusBadRequest,recorder.Code)
			})
		})
	})

}


