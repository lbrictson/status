package storage

import (
	"context"
	"errors"

	"github.com/lbrictson/status/ent"

	"github.com/lbrictson/status/ent/operator"

	"github.com/lbrictson/status/pkg"
)

func convertEntOperatorToPkgOperator(o *ent.Operator) pkg.Operator {
	return pkg.Operator{
		Email:          o.Email,
		HashedPassword: o.HashedPassword,
		Role:           o.Role,
	}
}

func (s *Store) SaveOperator(ctx context.Context, operator *pkg.Operator) error {
	o, err := s.client.Operator.Create().
		SetEmail(operator.Email).
		SetRole(operator.Role).
		SetHashedPassword(operator.HashedPassword).
		Save(ctx)
	if err != nil {
		return err
	}
	operator.Email = o.Email
	operator.Role = o.Role
	operator.HashedPassword = o.HashedPassword
	return nil
}

func (s *Store) ListOperators(ctx context.Context) ([]pkg.Operator, error) {
	operatorList, err := s.client.Operator.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	cleanOperatorList := []pkg.Operator{}
	for _, x := range operatorList {
		cleanOperatorList = append(cleanOperatorList, convertEntOperatorToPkgOperator(x))
	}
	return cleanOperatorList, nil
}

func (s *Store) GetOperatorByEmail(ctx context.Context, email string) (pkg.Operator, error) {
	o, err := s.client.Operator.Query().Where(operator.EmailEQ(email)).Only(ctx)
	if err != nil {
		return pkg.Operator{}, err
	}
	if o == nil {
		return pkg.Operator{}, errors.New("not found")
	}
	if o.Email != email {
		return pkg.Operator{}, errors.New("not found")
	}
	return convertEntOperatorToPkgOperator(o), nil
}
