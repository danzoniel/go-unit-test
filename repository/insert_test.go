package repository

import (
	"errors"
	"testing"
)

func TestInsertUser(t *testing.T) {
	// Casos de teste
	testCases := []struct {
		name        string
		user        User
		expectedErr error
	}{
		{
			name:        "Usuário inserido com sucesso",
			user:        User{Name: "John", Email: "john@example.com"},
			expectedErr: nil,
		},
		{
			name:        "Usuário já existe",
			user:        User{Name: "Existing", Email: "existing@example.com"},
			expectedErr: errors.New("usuário já existe"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userService := &UserServiceImpl{}

			err := userService.InsertUser(tc.user)

			if (err == nil && tc.expectedErr != nil) || (err != nil && tc.expectedErr == nil) {
				t.Fatalf("esperava-se o erro '%v', mas obteve-se '%v'", tc.expectedErr, err)
			}
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Fatalf("esperava-se o erro '%v', mas obteve-se '%v'", tc.expectedErr, err)
			}
		})
	}
}
