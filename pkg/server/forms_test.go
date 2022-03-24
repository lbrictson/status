package server

import (
	"context"
	"errors"
	"testing"

	"github.com/lbrictson/status/pkg/auth"

	"github.com/lbrictson/status/pkg"
)

type MockStorageBackend struct {
	ReturnError                   bool
	ReturnGetOperatorByEmailValue pkg.Operator
}

func (m MockStorageBackend) SaveOperator(ctx context.Context, operator *pkg.Operator) error {
	//TODO implement me
	panic("implement me")
}

func (m MockStorageBackend) ListOperators(ctx context.Context) ([]pkg.Operator, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockStorageBackend) GetOperatorByEmail(ctx context.Context, email string) (pkg.Operator, error) {
	if m.ReturnError {
		return pkg.Operator{}, errors.New("not found")
	}
	return m.ReturnGetOperatorByEmailValue, nil
}

func Test_validateOperator(t *testing.T) {
	type args struct {
		s        pkg.StorageBackend
		email    string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "validate correct password works for user",
			args: struct {
				s        pkg.StorageBackend
				email    string
				password string
			}{s: MockStorageBackend{
				ReturnError: false,
				ReturnGetOperatorByEmailValue: pkg.Operator{
					Email:          "admin@fake",
					HashedPassword: auth.HashAndSalt("password"),
					Role:           "Admin",
				},
			}, email: "admin@fake", password: "password"},
			want: true,
		},
		{
			name: "make sure invalid password doesn't get logged in",
			args: struct {
				s        pkg.StorageBackend
				email    string
				password string
			}{s: MockStorageBackend{
				ReturnError: false,
				ReturnGetOperatorByEmailValue: pkg.Operator{
					Email:          "admin@fake",
					HashedPassword: auth.HashAndSalt("notPassword"),
					Role:           "Admin",
				},
			}, email: "admin@fake", password: "password"},
			want: false,
		},
		{
			name: "user not found in database",
			args: struct {
				s        pkg.StorageBackend
				email    string
				password string
			}{s: MockStorageBackend{
				ReturnError: true,
				ReturnGetOperatorByEmailValue: pkg.Operator{
					Email:          "fake@fake",
					HashedPassword: auth.HashAndSalt("password"),
					Role:           "Admin",
				},
			}, email: "admin@fake", password: "password"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateOperator(tt.args.s, tt.args.email, tt.args.password); got != tt.want {
				t.Errorf("validateOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}
