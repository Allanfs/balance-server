package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/allanfs/balance-server/internal/app/domain/model"
	"github.com/allanfs/balance-server/internal/app/outbound/db"
)

var Lancamentos LancamentosRepository

func NewLancamentosRepository(sqldb *sql.DB) LancamentosRepository {
	return &lancamentosRepositorySqlc{
		sqlc: db.New(sqldb),
	}
}

type LancamentosRepository interface {
	CadastrarLancamento(context.Context, model.Lancamento) (model.LancamentoID, error)

	AlterarLancamento(context.Context, model.LancamentoID, model.Lancamento) error

	BuscarLancamentos(context.Context) ([]*model.Lancamento, error)
	BuscarLancamentosDoDia(context.Context, time.Time) ([]*model.Lancamento, error)
	BuscarLancamentoPorId(context.Context, model.LancamentoID) (*model.Lancamento, error)

	RemoverLancamento(context.Context, model.LancamentoID) error
}

type lancamentosRepositorySqlc struct {
	sqlc *db.Queries
}

func (l *lancamentosRepositorySqlc) CadastrarLancamento(ctx context.Context, lancamento model.Lancamento) (model.LancamentoID, error) {
	id, err := l.sqlc.CreateEntry(ctx, db.CreateEntryParams{
		Name:         lancamento.Nome,
		Amount:       fmt.Sprintf("%.3f", lancamento.Valor),
		EntryType:    string(lancamento.Tipo),
		ExternalInfo: sql.NullString{String: lancamento.ExternalInfo, Valid: true},
	})

	if err != nil {
		return 0, err
	}

	return model.LancamentoID(id), nil
}

func (l *lancamentosRepositorySqlc) AlterarLancamento(_ context.Context, _ model.LancamentoID, _ model.Lancamento) error {
	panic("not implemented") // TODO: Implement
}

func (l *lancamentosRepositorySqlc) BuscarLancamentos(ctx context.Context) ([]*model.Lancamento, error) {
	db, err := l.sqlc.GetEntries(ctx)

	if err != nil {
		return nil, err
	}

	var lancamentos []*model.Lancamento
	for _, r := range db {

		lanc, err := l.toModel(r)
		if err != nil {
			return nil, err
		}
		lancamentos = append(lancamentos, lanc)
	}

	return lancamentos, nil
}

func (l *lancamentosRepositorySqlc) BuscarLancamentosDoDia(_ context.Context, _ time.Time) ([]*model.Lancamento, error) {
	panic("not implemented") // TODO: Implement
}

func (l *lancamentosRepositorySqlc) BuscarLancamentoPorId(ctx context.Context, id model.LancamentoID) (*model.Lancamento, error) {
	r, err := l.sqlc.GetEntry(ctx, int64(id))

	if err != nil {
		return nil, err
	}
	return l.toModel(r)

}

func (l *lancamentosRepositorySqlc) RemoverLancamento(_ context.Context, _ model.LancamentoID) error {
	panic("not implemented") // TODO: Implement
}

func (l *lancamentosRepositorySqlc) toFloat(s string) (float64, error) {
	valor, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, fmt.Errorf("value is not a valid float: %w", err)
	}

	return valor, nil
}

func (l *lancamentosRepositorySqlc) toModel(be db.BalanceEntry) (*model.Lancamento, error) {
	valor, err := l.toFloat(be.Amount)
	if err != nil {
		return nil, err
	}
	return &model.Lancamento{
		Id:           model.LancamentoID(be.ID),
		Nome:         be.Name,
		Valor:        valor,
		Tipo:         model.NaturezaFinanceira(be.EntryType),
		ExternalInfo: be.ExternalInfo.String,
	}, nil
}
