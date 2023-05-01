package model

type LancamentoID int64

type Lancamento struct {
	Id    LancamentoID
	Nome  string
	Valor float64
	Tipo  NaturezaFinanceira
}
