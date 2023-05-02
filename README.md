# Balance Server | Serviço de Saldo
O serviço realiza operações de CRUD para informações de saldo de forma genérica, se resumindo a forma mais básica.

Exemplos de uso:
- lista de receitas e despesas
- lista de itens de um pedido em um comerico

## ORM - Não usa

Utiliza diretamente a lib do Go através do [sqlc](https://sqlc.dev/).

Dentro de [scripts](./scripts) estão todas as [consultas mapeadas](./scripts/query.sql) e que são geradas por ele.

## Taskfile

Veja as operações possíveis
```bash
task --list-all
``` 

## Swagger

[Swaggo](https://github.com/swaggo/swag) 
```bash
swag fmt && swag init
```