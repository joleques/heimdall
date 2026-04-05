# Backlog Técnico — Entrega 1 (MVP Assistants)

## Premissa Estrutural
A raiz da codebase é `./src`.

## Premissa de Domínio
Linguagem ubíqua única: `Assistant`.

## Objetivo
Entregar jornadas base do produto para uso real:
1. `heimdall init <target>`
2. `/heimdall start`
3. `/heimdall list-lib`
4. `/heimdall install <assistant>`

Targets:
- `codex`
- `antigravity`
- `claude`
- `cursor`

## Épico 1 — Modelagem de Domínio Assistant
### Tarefas
1. Consolidar `AssistantSpec` como entidade principal.
2. Adicionar campo `skills` associadas ao Assistant.
3. Validar unicidade e consistência de `skills`.

### Critérios de aceite
1. `Assistant` é o termo dominante em domínio e aplicação.
2. Contrato canônico cobre instalação conjunta Assistant + skills.

## Épico 2 — Init por Target
### Tarefas
1. Implementar comando `heimdall init <target>`.
2. Criar layout padrão por target.
3. Aplicar política de `AGENTS.md`.

### Critérios de aceite
1. Diretório fica pronto para uso no target escolhido.
2. Fluxo é idempotente por padrão.

## Épico 3 — Start de Contexto do Projeto
### Tarefas
1. Implementar fluxo `/heimdall start`.
2. Coletar: título, contexto/descrição, documentação inicial.
3. Persistir contexto em estrutura canônica do projeto.

### Critérios de aceite
1. Contexto fica disponível para os Assistants do projeto.
2. Fluxo funciona em modo não interativo (flags) e interativo (quando aplicável).

## Épico 4 — Biblioteca de Assistants
### Tarefas
1. Implementar `/heimdall list-lib`.
2. Exibir `id`, `name`, `description`.
3. Exibir skills associadas por Assistant.

### Critérios de aceite
1. Lista carregada de `src/templates/default`.
2. Erros de catálogo retornam mensagem clara.

## Épico 5 — Instalação de Assistant
### Tarefas
1. Implementar `/heimdall install <assistant>`.
2. Resolver Assistant no catálogo.
3. Instalar Assistant + skills associadas automaticamente.
4. Exibir resumo (`installed/skipped/failed/warnings`).

### Critérios de aceite
1. Instalação automática de skills associadas por padrão.
2. Idempotência em execução repetida sem `--force`.

## Épico 6 — Testes (TDD + Integração)
### Unitários
1. Parser de comandos do MVP.
2. Validação de `AssistantSpec` com `skills`.
3. Seleção de catálogo por Assistant.

### Integração
1. `init` em cada target.
2. `list-lib` com catálogo válido.
3. `install` com Assistant válido (inclui skills associadas).
4. Política `AGENTS.md` (`skip|if-missing|overwrite`).

### Critérios de aceite
1. `go test ./...` e `go build ./...` passando.
2. Cenários críticos de jornada cobertos.

## Sequência sugerida de PRs
1. PR-04: `init <target>` + política `AGENTS.md`.
2. PR-05: `/heimdall start` + persistência de contexto.
3. PR-06: `/heimdall list-lib` + metadados de Assistants.
4. PR-07: `/heimdall install <assistant>` + skills associadas + hardening.

## Definição de pronto da Entrega 1
1. Jornadas `init`, `start`, `list-lib`, `install` funcionais.
2. Linguagem ubíqua consistente com `Assistant`.
3. Instalação de Assistant instala skills associadas automaticamente.
4. Documentação atualizada para uso do MVP.
