# golang-email

## Descrição do que foi feito até agora

- **Módulo Go configurado**  
  - Arquivo `go.mod` criado com o módulo `emailn` e dependências para testes (`testify`) e geração de IDs (`xid`).

- **Domínio de campanha de e-mail (`internal/domain/compaign`)**  
  - **Entidade `Compaign`** com os campos:
    - `ID` (string, gerado automaticamente com `xid`),
    - `Name` (nome da campanha),
    - `Content` (conteúdo do e-mail),
    - `CreatedAt` e `UpdatedAt` (datas de criação e atualização),
    - `Contacts` (lista de contatos, cada um com `Email`).
  - **Função `NewCompaign`** responsável por criar uma nova campanha:
    - valida se `name`, `content` e a lista de `contacts` foram informados,
    - converte a lista de `string` para `[]Contact`,
    - gera o ID e preenche as datas de criação/atualização.

- **Contrato de entrada (`internal/contract`)**  
  - DTO `NewCampaignDto` contendo `Name`, `Content` e `Contacts []string` para representar os dados necessários para criar uma nova campanha.

- **Serviço de campanhas (`CompaignService`)**  
  - Interface `CompaignRepository` com o método `Save(*Compaign) error` para abstrair o repositório.
  - Struct `CompaignService` que recebe um `CompaignRepository`.
  - Método `Create(newCampaignDto contract.NewCampaignDto)`:
    - monta a entidade de domínio usando `NewCompaign`,
    - chama `repository.Save` para persistir a campanha,
    - retorna o `ID` da campanha criada ou um erro.

- **Testes de unidade**  
  - Testes para `NewCompaign` verificando:
    - preenchimento correto de `Name`, `Content` e `Contacts`,
    - geração de `CreatedAt`.
  - Testes para `CompaignService.Create` usando `testify` e `mock`:
    - criação de campanha válida retornando um ID não nulo e sem erro,
    - validação de erro de domínio (por exemplo, `name` vazio),
    - verificação se o repositório `Save` é chamado com os dados corretos,
    - simulação de erro de banco de dados retornado pelo repositório.

- **`main.go`**  
  - Arquivo de entrada criado, ainda sem implementação da API/CLI (função `main` vazia), preparado para futuras integrações (ex.: HTTP API para criar campanhas).