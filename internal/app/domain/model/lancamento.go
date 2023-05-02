package model

type LancamentoID int64

type Lancamento struct {
	Id           LancamentoID       `json:"id"`
	Nome         string             `json:"nome"`
	Valor        float64            `json:"valor"`
	Tipo         NaturezaFinanceira `json:"tipo"`
	ExternalInfo string             `json:"external_info,omitempty"`
}
