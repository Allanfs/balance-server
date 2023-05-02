package lancamentos

import (
	"context"

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
