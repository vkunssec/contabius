package domain

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InvestmentType é uma enumeração que representa o tipo de investimento
type InvestmentType string

// InvestmentId é uma enumeração que representa o ID do investimento
type InvestmentId int

const (
	InvestmentTypeCDIId InvestmentId = iota
	InvestmentTypeCDBId
	InvestmentTypeLCAId
	InvestmentTypeLCIId
	InvestmentTypePoupancaId
	InvestmentTypeTesouroDiretoId
	InvestmentTypeFundosInvestimentoId
	InvestmentTypeAcoesId
	InvestmentTypeFIIsId
	InvestmentTypeCRIId
	InvestmentTypeCRAId
	InvestmentTypeDebenturesId
	InvestmentTypeBDRId
	InvestmentTypeCOEId
	InvestmentTypeOtherId

	InvestmentTypeCDI                InvestmentType = "cdi"                 // Certificado de Depósito Interbancário
	InvestmentTypeCDB                InvestmentType = "cdb"                 // Certificado de Depósito Bancário
	InvestmentTypeLCA                InvestmentType = "lca"                 // Letra de Crédito do Agronegócio
	InvestmentTypeLCI                InvestmentType = "lci"                 // Letra de Crédito Imobiliário
	InvestmentTypePoupanca           InvestmentType = "poupanca"            // Poupança
	InvestmentTypeTesouroDireto      InvestmentType = "tesouro_direto"      // Tesouro Direto
	InvestmentTypeFundosInvestimento InvestmentType = "fundos_investimento" // Fundos de Investimento
	InvestmentTypeAcoes              InvestmentType = "acoes"               // Ações
	InvestmentTypeFIIs               InvestmentType = "fiis"                // Fundos de Investimento Imobiliário
	InvestmentTypeCRI                InvestmentType = "cri"                 // Certificado de Recebíveis Imobiliários
	InvestmentTypeCRA                InvestmentType = "cra"                 // Certificado de Recebíveis do Agronegócio
	InvestmentTypeDebentures         InvestmentType = "debentures"          // Debêntures
	InvestmentTypeBDR                InvestmentType = "bdr"                 // Depositário Brasileiro de Recebíveis
	InvestmentTypeCOE                InvestmentType = "coe"                 // Certificado de Operações Estruturadas
	InvestmentTypeOther              InvestmentType = "other"               // Outros
)

// Recurrence é uma enumeração que representa a recorrência do investimento
type Recurrence string

const (
	RecurrenceMonthly  Recurrence = "monthly"  // Mensal
	RecurrenceYearly   Recurrence = "yearly"   // Anual
	RecurrenceSporadic Recurrence = "sporadic" // Esporádico
)

// RecurrenceDay é uma enumeração que representa o dia da recorrência
type RecurrenceDay int

type Investment struct {
	Id         InvestmentId   `json:"id" bson:"id" example:"1"`                   // ID do investimento
	Investment InvestmentType `json:"investment" bson:"investment" example:"cdi"` // Tipo de investimento
}

// Investments são os investimentos que o usuário possui
type Investments struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty" example:"678079f6f5080a39a8eedc1e"`            // ID do investimento
	Investment    Investment         `json:"investment" bson:"investment"`                                          // Tipo de investimento
	Amount        Money              `json:"amount" bson:"amount"`                                                  // Valor do investimento
	Account       Accounts           `json:"account" bson:"account"`                                                // Conta do usuário
	Recurrence    Recurrence         `json:"recurrence" bson:"recurrence" example:"monthly"`                        // Recurrence do investimento
	RecurrenceDay *RecurrenceDay     `json:"recurrence_day" bson:"recurrence_day" example:"1"`                      // Dia da recorrência
	Description   *string            `json:"description" bson:"description" example:"Investimento em CDB"`          // Descrição do investimento
	CreatedAt     time.Time          `json:"created_at" bson:"created_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de criação do investimento
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de atualização do investimento
}

// Validate valida os campos do investimento
func (i *Investment) Validate() error {
	if i.Investment == "" {
		return errors.New("investment is required")
	}
	return nil
}

// Validate valida os campos do investimento
func (i *Investments) Validate() error {
	switch i.Recurrence {
	case RecurrenceMonthly:
		if i.RecurrenceDay == nil {
			return errors.New("recurrence_day is required for monthly recurrence")
		}
	case RecurrenceYearly:
		if i.RecurrenceDay != nil {
			return errors.New("recurrence_day is not allowed for yearly recurrence")
		}
	case RecurrenceSporadic:
		if i.RecurrenceDay != nil {
			return errors.New("recurrence_day is not allowed for sporadic recurrence")
		}
	}

	if i.RecurrenceDay != nil && (*i.RecurrenceDay < 1 || *i.RecurrenceDay > 31) {
		return errors.New("recurrence_day must be between 1 and 31")
	}

	if i.Amount.IsZero() {
		return errors.New("amount must be greater than 0")
	}

	if i.Account.Id == "" {
		return errors.New("account is required")
	}

	return nil
}

// GetRecurrenceDay retorna o dia da recorrência
func (i *Investments) GetRecurrenceDay() int {
	if i.RecurrenceDay == nil {
		return 0
	}
	return int(*i.RecurrenceDay)
}

var (
	// investmentMap armazena o mapeamento de IDs para Investments
	investmentMap = map[InvestmentId]InvestmentType{
		InvestmentTypeCDIId:                InvestmentTypeCDI,
		InvestmentTypeCDBId:                InvestmentTypeCDB,
		InvestmentTypeLCAId:                InvestmentTypeLCA,
		InvestmentTypeLCIId:                InvestmentTypeLCI,
		InvestmentTypePoupancaId:           InvestmentTypePoupanca,
		InvestmentTypeTesouroDiretoId:      InvestmentTypeTesouroDireto,
		InvestmentTypeFundosInvestimentoId: InvestmentTypeFundosInvestimento,
		InvestmentTypeAcoesId:              InvestmentTypeAcoes,
		InvestmentTypeFIIsId:               InvestmentTypeFIIs,
		InvestmentTypeCRIId:                InvestmentTypeCRI,
		InvestmentTypeCRAId:                InvestmentTypeCRA,
		InvestmentTypeDebenturesId:         InvestmentTypeDebentures,
		InvestmentTypeBDRId:                InvestmentTypeBDR,
		InvestmentTypeCOEId:                InvestmentTypeCOE,
		InvestmentTypeOtherId:              InvestmentTypeOther,
	}

	// allInvestments armazena a lista completa de investimentos
	allInvestments = func() []Investment {
		investments := make([]Investment, 0, len(investmentMap))
		for id, investmentType := range investmentMap {
			investments = append(investments, Investment{Id: id, Investment: investmentType})
		}
		return investments
	}()
)

// GetInvestment retorna um investimento
func GetInvestment(id InvestmentId) (Investment, error) {
	investmentType, exists := investmentMap[id]
	if !exists {
		return Investment{}, errors.New("investment not found")
	}

	investment := Investment{Id: id, Investment: investmentType}
	return investment, investment.Validate()
}

// GetInvestments retorna todos os investimentos solicitados
func GetInvestments(ids []InvestmentId) ([]Investment, error) {
	investments := make([]Investment, 0, len(ids))
	for _, id := range ids {
		investment, err := GetInvestment(id)
		if err != nil {
			return nil, err // retorna erro imediatamente se encontrar um investimento inválido
		}
		investments = append(investments, investment)
	}
	return investments, nil
}

// AllInvestments retorna todos os investimentos
func AllInvestments() []Investment {
	return allInvestments
}
