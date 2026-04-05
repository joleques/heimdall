# Status das PRs — Heimdall

## Objetivo
Documento de continuidade para retomar o projeto em qualquer chat sem perder contexto.

## Contexto Atual
- Linguagem ubíqua consolidada: `Assistant`.
- Raiz da codebase: `./src`.
- MVP atual: `init`, `start`, `list-lib`, `install <assistant>`.

## PRs Concluídas

### PR-01 — Bootstrap Base
**Status:** concluída

Entregas:
- Estrutura inicial do projeto Go em camadas.
- `go.mod`.
- Domínio base (`AssistantSpec`, policy, platform).
- Testes unitários iniciais.

### PR-02 — Parser CLI (fase anterior)
**Status:** concluída

Entregas:
- Parser do fluxo `install`.
- Wiring inicial em `application`.
- Testes de parser e dispatch.

### PR-03 — Catálogo + Instalação Base
**Status:** concluída

Entregas:
- Gateway de catálogo com parse YAML canônico.
- Gateway de instalação em filesystem com idempotência e política de `AGENTS.md`.
- Use cases reais de instalação.

### PR-RA — Refatoração de Realinhamento
**Status:** concluída

Motivo:
- Pivot estratégico para domínio centrado em `Assistant`.

Entregas:
- Remoção da superfície pública antiga (`install skills` e `install all`).
- Fluxo principal consolidado em `install assistant`.
- `AssistantSpec` evoluído com `skills` associadas.
- Use case de instalação focado em Assistant + skills associadas automáticas.

### PR-04 — `heimdall init <target>`
**Status:** concluída

Entregas:
- Novo comando `init` com parser próprio.
- Caso de uso `InitTarget`.
- Implementação de criação de layout por target no filesystem gateway.
- Integração com política de `AGENTS.md`.
- Testes de parser, app dispatch, use_case e infra.

## PRs Planejadas (Próximas)

### PR-05 — `/heimdall start`
**Status:** concluída

Entregas:
- Fluxo de inicialização de contexto do projeto.
- Coleta de dados em modo não interativo (`--title`, `--description`, `--doc`) e interativo via `stdin`.
- Persistência do contexto em estrutura canônica `.heimdall/context/`.
- Manifesto YAML (`project-context.yaml`) e resumo Markdown (`README.md`).
- Material base persistido em `.heimdall/context/docs/` com suporte a arquivo real ou anotação textual.
- Testes de parser, dispatch, use case, domínio e infra.

### PR-06 — `/heimdall list-lib`
**Status:** concluída

Entregas:
- Listagem da biblioteca de Assistants.
- Exposição do comando `heimdall list-lib` na CLI.
- Exibição de metadados (`id`, `name`, `description`, `skills`).
- Enriquecimento do catálogo com metadados dos Assistants parseados do YAML.
- Testes de parser, app dispatch, use case e infra.

### PR-07 — `/heimdall install <assistant>` hardening
**Status:** concluída

Entregas:
- Hardening do fluxo de instalação para todos targets suportados.
- Mensagens de erro mais úteis para catálogo vazio e Assistant inexistente.
- Saída de CLI com detalhes de `installed`, `skipped`, `failed` e `warning`.
- Propagação explícita de warnings para skills associadas ausentes.
- Testes adicionais de parser, app dispatch e use case para cenários críticos.

## Estado Técnico Atual (Resumo)
- `go test ./...`: passando
- `go build ./...`: passando

## Próximo Passo Recomendado
Entrega 1 do MVP está funcional; próximo ciclo recomendado é consolidar documentação de uso e definir a evolução pós-MVP (criação/edição/remoção de Assistants).
