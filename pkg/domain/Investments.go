package domain

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InvestmentType é uma enumeração que representa o tipo de investimento
type InvestmentLiteral string

// InvestmentId é uma enumeração que representa o ID do investimento
type InvestmentId int

const (
	InvestmentTypeCDIId InvestmentId = iota + 1
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

	InvestmentTypeCDI                InvestmentLiteral = "cdi"                 // Certificado de Depósito Interbancário
	InvestmentTypeCDB                InvestmentLiteral = "cdb"                 // Certificado de Depósito Bancário
	InvestmentTypeLCA                InvestmentLiteral = "lca"                 // Letra de Crédito do Agronegócio
	InvestmentTypeLCI                InvestmentLiteral = "lci"                 // Letra de Crédito Imobiliário
	InvestmentTypePoupanca           InvestmentLiteral = "poupanca"            // Poupança
	InvestmentTypeTesouroDireto      InvestmentLiteral = "tesouro_direto"      // Tesouro Direto
	InvestmentTypeFundosInvestimento InvestmentLiteral = "fundos_investimento" // Fundos de Investimento
	InvestmentTypeAcoes              InvestmentLiteral = "acoes"               // Ações
	InvestmentTypeFIIs               InvestmentLiteral = "fiis"                // Fundos de Investimento Imobiliário
	InvestmentTypeCRI                InvestmentLiteral = "cri"                 // Certificado de Recebíveis Imobiliários
	InvestmentTypeCRA                InvestmentLiteral = "cra"                 // Certificado de Recebíveis do Agronegócio
	InvestmentTypeDebentures         InvestmentLiteral = "debentures"          // Debêntures
	InvestmentTypeBDR                InvestmentLiteral = "bdr"                 // Depositário Brasileiro de Recebíveis
	InvestmentTypeCOE                InvestmentLiteral = "coe"                 // Certificado de Operações Estruturadas
	InvestmentTypeOther              InvestmentLiteral = "other"               // Outros
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

type InvestmentType struct {
	Id         InvestmentId      `json:"id,omitempty" bson:"_id,omitempty" example:"1"`                  // ID do investimento
	Investment InvestmentLiteral `json:"investment,omitempty" bson:"investment,omitempty" example:"cdi"` // Tipo de investimento
}

// Investments são os investimentos que o usuário possui
type Investments struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"678079f6f5080a39a8eedc1e"`             // ID do investimento
	Investment    InvestmentType     `json:"investment_type,omitempty" bson:"investment_type,omitempty"`                       // Tipo de investimento
	Amount        Money              `json:"amount,omitempty" bson:"amount,omitempty"`                                         // Valor do investimento
	Account       Accounts           `json:"account,omitempty" bson:"account,omitempty"`                                       // Conta do usuário
	Recurrence    Recurrence         `json:"recurrence,omitempty" bson:"recurrence,omitempty" example:"monthly"`               // Recurrence do investimento
	RecurrenceDay *RecurrenceDay     `json:"recurrence_day,omitempty" bson:"recurrence_day,omitempty" example:"1"`             // Dia da recorrência
	Description   *string            `json:"description,omitempty" bson:"description,omitempty" example:"Investimento em CDB"` // Descrição do investimento
	CreatedAt     time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" example:"2025-01-01T00:00:00Z"`  // Data de criação do investimento
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty" example:"2025-01-01T00:00:00Z"`  // Data de atualização do investimento
}

// InvestmentRequest é uma estrutura que representa uma requisição para criar um investimento
type InvestmentRequest struct {
	Investment    InvestmentType        `json:"investment_type" bson:"investment_type" validate:"required"`                       // Tipo de investimento
	Amount        Money                 `json:"amount" bson:"amount" validate:"required"`                                         // Valor do investimento
	Account       PartialAccountRequest `json:"account" bson:"account" validate:"required"`                                       // Conta do usuário
	Recurrence    Recurrence            `json:"recurrence" bson:"recurrence" `                                                    // Recurrence do investimento
	RecurrenceDay *RecurrenceDay        `json:"recurrence_day" bson:"recurrence_day" `                                            // Dia da recorrência
	Description   *string               `json:"description,omitempty" bson:"description,omitempty" example:"Investimento em CDB"` // Descrição do investimento
}

// Validate valida os campos do investimento
func (i *InvestmentType) Validate() error {
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

	if i.Account.Id == primitive.NilObjectID {
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
	investmentMap = map[InvestmentId]InvestmentLiteral{
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
	allInvestments = func() []InvestmentType {
		investments := make([]InvestmentType, 0, len(investmentMap))
		for id, investmentType := range investmentMap {
			investments = append(investments, InvestmentType{Id: id, Investment: investmentType})
		}
		return investments
	}()
)

// GetInvestment retorna um investimento
func GetInvestment(id InvestmentId) (InvestmentType, error) {
	investmentType, exists := investmentMap[id]
	if !exists {
		return InvestmentType{}, errors.New("investment not found")
	}

	investment := InvestmentType{Id: id, Investment: investmentType}
	return investment, investment.Validate()
}

// GetInvestments retorna todos os investimentos solicitados
func GetInvestments(ids []InvestmentId) ([]InvestmentType, error) {
	investments := make([]InvestmentType, 0, len(ids))
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
func AllInvestments() []InvestmentType {
	return allInvestments
}
