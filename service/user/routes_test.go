package user

import (
	"bytes"
	"encoding/json"
	"go-backend-api/types"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestUserServicesHandlers(t *testing.T){
	userStore := mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Rick",
			LastName: "Roll",
			Email: "",
			Password: "123",
		}
		marshalled, _ := json.Marshal(payload) 
		
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest{
			t.Error("expired status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct {}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error){
	return  nil, nil 
}

func (m *mockUserStore) CreateUser(id int) (types.User, error){
	return  nil, nil 
}