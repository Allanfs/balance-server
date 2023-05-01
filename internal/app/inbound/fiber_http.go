package inbound

import (
	"net/http"

	"github.com/allanfs/balance-server/internal/app/domain/lancamentos"
	"github.com/allanfs/balance-server/internal/app/domain/model"
	"github.com/gofiber/fiber/v2"
)

func SetLancamentosRoutes(app fiber.Router) {
	app.Post("/", CadastrarLancamento)

	app.Get("/", ListarLancamento)

	app.Get("/:id", ObterLancamento)
}

type NovoLancamentoRequest struct {
	Nome  string
	Valor float64
	Tipo  model.NaturezaFinanceira `json:"tipo"`
}

type NovoLancamentoResponse struct {
	Id model.LancamentoID `json:"id"`
}

func CadastrarLancamento(c *fiber.Ctx) error {

	l := new(NovoLancamentoRequest)

	// TODO: validações
	// nome não pode ser vazio
	// valor deve ser maior que zero
	// tipo tem que ser C, D ou N/A

	if err := c.BodyParser(l); err != nil {
		return c.JSON(HttpError{err, "Request body inválido"})
	}

	lancamento := model.Lancamento{
		Nome:  l.Nome,
		Valor: l.Valor,
		Tipo:  l.Tipo,
	}

	err := lancamentos.CadastrarLancamento(c.Context(), &lancamento)
	if err != nil {
		return c.
			Status(http.StatusInternalServerError).
			JSON(HttpError{err, "Falha ao cadastrar lançamento"})
	}

	return c.
		Status(http.StatusCreated).
		JSON(NovoLancamentoResponse{lancamento.Id})

}

func ListarLancamento(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotImplemented)
}

func ObterLancamento(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotImplemented)
}
