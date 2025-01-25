package repository

import (
	"github.com/vkunssec/contabius/pkg/domain"
)

func GetMethod(ids []domain.MethodId) ([]domain.Methods, error) {
	if len(ids) == 0 {
		return domain.AllMethods(), nil
	}
	return domain.GetMethods(ids)
}
