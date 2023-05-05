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

## Roadmap

### Comentários
- Aplicação multitenant? A principio não pois o serviço poderá ser implantado com facilidade em cada sistema que precisar dele.
- O banco de dados pode ser compartilhado com outras aplicações, a aplicação utiliza um schema proprio para operar.
  - O nome do schema pode ser parametrizado no helm? no futuro, não é prioridade agora.

---

### Checklist

- [x] CRUD de lançamentos
- [ ] SDK http para ser usada por outros serviços
- [ ] suporte a GRPC
- [ ] tratar os de forma assincrona (receber via http, postar na fila e responder 200 se estiver tudo ok, um consumer vai fazer o cadastro dos dados)
- [ ] criar deployment no K8S
- [ ] criar deployment no helm-charts