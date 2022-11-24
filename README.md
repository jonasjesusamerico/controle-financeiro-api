<div align="right">
<img height="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" />
</div>

# CONTROLE FINANCEIRO API

* Aplição tem como objetivo aplicar os conhecimento adquirido ao longo do tempo com a linguagem GOlang.

## TECNOLOGIAS
* Desenvolvimento
    * [GOlang](https://go.dev/) - Liguagem para desenvolvimento
    * [Gin-Gonic](https://gin-gonic.com/docs/) - Lib responsável pelo controle de rota da aplicação
    * [Package validator](https://github.com/go-playground/validator) - Responsável para validação na construção das struct
* Baco de dados
    * [PostgreSQL](https://www.postgresql.org/) - Banco de dados

## ETAPAS

- [x] Início do projeto
- [x] Criar Dockerfile junto com docker-compose.yml
- [x] Configura o banco de dados Postgres
- [ ] ~~Verificar opção de migration para versionamento do sql~~ (Criar novo projeto apenas para executar as migrations)
- [x] Adicionar router e criar a primeira rota
    - [x] Padronizar rotas
- [x] Adicionar arquivo .env 
- [x] Implementar autenticação com token jwt
- [x] Criar rota login para autenticação
- [ ] Adicionar documentação dos endpoint com swagger