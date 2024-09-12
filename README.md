# Hire-Go
App for connect clients and professionals.

Estrutura do projeto
Nesse post já vamos deixar nossa estrutura inicial pronta, vamos separa da seguinte maneira:

cmd: Aqui vamos deixar nossos arquivos main.go, responsáveis por iniciar nossa aplicação.
config: Vamos salvar algumas configs aqui, como envs, logs.
internal: Aqui é onde vai ficar nossa regra de negócio
internal/dto: Onde vamos determinar os tipos de dados que vamos permitir entrar na aplicação
internal/handler: Essa pasta vai ficar nossos arquivos de roteamento (pode chamar de controller se preferir)
internal/database: Essa pasta é onde vamos salvar tudo que for relacionado ao banco de dados
internal/database/migrations: Vamos salvar nossas migrations aqui
internal/database/queries: Onde vai ficar nossas queries sql de consulta ao banco
internal/repository: Aqui onde vai ficar nossa camada de repositórios.
internal/service: Por último, nossa camada de service, onde a regra de negócio vai ficar (pode chamar de usecase se preferir).

Techs: 
API Client: Gin
ORM: Gorm
DB: PostgreSQL
Containization: Docker
Container-orchestration: K8s
Logger: zerolog
Docs: Swaggo
Migrations-db: golang-migrate
Validator: go-playground/validator