package repository

import "github.com/vkunssec/contabius/pkg/domain"

// GetInvestmentTypes retorna um investimento
func GetInvestmentTypes(ids []domain.InvestmentId) ([]domain.Investment, error) {
	if len(ids) == 0 {
		return domain.AllInvestments(), nil
	}

	return domain.GetInvestments(ids)
}
