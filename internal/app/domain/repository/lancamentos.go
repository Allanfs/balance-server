package repository

import (
	"time"

	"github.com/allanfs/balance-server/internal/app/domain/model"
)

var Lancamentos LancamentosRepository

type LancamentosRepository interface {
	CadastrarLancamento(model.Lancamento) (model.LancamentoID, error)

	AlterarLancamento(model.LancamentoID, model.Lancamento) error

	BuscarLancamentos() ([]model.Lancamento, error)
	BuscarLancamentosDoDia(time.Time) ([]model.Lancamento, error)
	BuscarLancamentoPorId(model.LancamentoID) (model.Lancamento, error)

	RemoverLancamento(model.LancamentoID) error
}

type lancamentosRepositoryDatabase struct{}
