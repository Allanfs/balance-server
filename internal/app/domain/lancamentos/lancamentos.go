package lancamentos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/allanfs/balance-server/internal/app/domain/model"
	"github.com/allanfs/balance-server/internal/app/domain/repository"
)

func CadastrarLancamento(ctx context.Context, l *model.Lancamento) error {

	id, err := repository.Lancamentos.CadastrarLancamento(ctx, *l)

	l.Id = id

	return err
}

func ObterLancamento(ctx context.Context, id model.LancamentoID) (*model.Lancamento, error) {

	l, err := repository.Lancamentos.BuscarLancamentoPorId(ctx, id)

	return l, err
}

func ListarLancamentos(ctx context.Context) ([]*model.Lancamento, error) {

	l, err := repository.Lancamentos.BuscarLancamentos(ctx)

	if err != nil {
		return nil, err
	}

	return l, err
}

func Deletar(ctx context.Context, id model.LancamentoID) error {

	err := repository.Lancamentos.RemoverLancamento(ctx, id)

	return err
}

func Atualizar(ctx context.Context, id model.LancamentoID, l model.Lancamento) error {

	existing, err := repository.Lancamentos.BuscarLancamentoPorId(ctx, id)

	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("lancamento não encontrado: %w", err)
	} else if err != nil {
		return err
	}

	if existing == nil {
		return nil
	}

	l.Id = id
	err = repository.Lancamentos.AlterarLancamento(ctx, id, l)

	return err
}
