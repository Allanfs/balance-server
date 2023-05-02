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
	Nome         string
	Valor        float64
	Tipo         model.NaturezaFinanceira `json:"tipo"`
	ExternalInfo string                   `json:"external_info,omitempty"`
}

type NovoLancamentoResponse struct {
	Id model.LancamentoID `json:"id"`
}

// CadastrarLancamento teste
// informaçao aqui
//	@router			/lancamentos [post]
//	@Summary		Cadastra um lançamento
//	@description	Informe o nome, valor (maior que zero) e tipo do lançamento (C, D ou N/A).
//	@description	O campo external_info é opcional. Serve para clientes da API armazenarem informações que lhes sejam uteis.
//	@param			req.body	body		NovoLancamentoRequest	true	"Objeto para cadastrar um novo lançamento"
//	@success		200			{object}	NovoLancamentoResponse
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
		Nome:         l.Nome,
		Valor:        l.Valor,
		Tipo:         l.Tipo,
		ExternalInfo: l.ExternalInfo,
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

	lanc, err := lancamentos.ListarLancamentos(c.Context())
	if err != nil {
		return c.
			Status(http.StatusInternalServerError).
			JSON(HttpError{err, "Falha ao listar lançamentos"})
	}

	return c.JSON(lanc)
}

// Lancaemtno model info
//	@Description	User account information
//	@Description	with user id and username
type Lancamento struct {
	Id           model.LancamentoID       `json:"id"`
	Nome         string                   `json:"nome"`
	Valor        float64                  `json:"valor"`
	Tipo         model.NaturezaFinanceira `json:"tipo"`
	ExternalInfo string                   `json:"external_info,omitempty"`
}

func ObterLancamento(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(HttpError{err, ":id deve ser inteiro"})
	}

	lancamento, err := lancamentos.ObterLancamento(c.Context(), model.LancamentoID(id))

	c.Context().Logger().Printf("TODO: validar erro de não encontrado")

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(HttpError{err, "Falha ao obter lançamento"})
	}
	return c.Status(http.StatusOK).JSON(lancamento)
}
