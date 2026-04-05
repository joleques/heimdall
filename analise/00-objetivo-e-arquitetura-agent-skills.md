# Análise de Continuidade — Projeto `Heimdall`

## Objetivo do Trabalho
Construir uma aplicação CLI chamada `heimdall` para gestão de **Assistants** (instalação, inicialização e operação), em diferentes plataformas de consumo.

Escopo atual (plataformas suportadas neste momento):
- `codex`
- `antigravity`
- `claude`
- `cursor`

## Linguagem Ubíqua (DDD)
Decisão explícita: o termo canônico de produto e domínio é **Assistant**.

- Não usamos termos concorrentes no core (ex.: "squad" como entidade principal).
- Skills continuam existindo como capacidades internas de um Assistant.
- Workflow, no contexto do produto, é a execução/orquestração de um Assistant.

## Problema que estamos resolvendo
Hoje o repositório foi pensado com foco em assets de `.agent/`.
Queremos evoluir para um produto instalável e gerenciável, com foco no usuário final:
1. iniciar ambiente por target,
2. configurar contexto do projeto,
3. explorar biblioteca de Assistants,
4. instalar e usar Assistants prontos,
5. futuramente criar/editar/excluir Assistants próprios.

## Decisões já tomadas
1. Linguagem ubíqua única: **Assistant**.
2. Ao instalar um Assistant, suas skills associadas devem ser instaladas automaticamente (padrão).
3. `AGENTS.md` é opcional no destino com política:
   - `skip`
   - `if-missing` (default)
   - `overwrite`
4. Linguagem de implementação: **Golang**.
5. Raiz da codebase: `./src`.

## Jornada MVP (Entrega 1)
1. Inicialização de estrutura por target:
   - `heimdall init [codex|antigravity|cursor|claude]`
2. Inicialização de contexto do projeto (chat command por target):
   - `/heimdall start`
   - coleta: título do projeto, descrição/contexto, documentação base.
3. Listagem da biblioteca de Assistants:
   - `/heimdall list-lib`
4. Instalação de Assistant:
   - `/heimdall install [assistant]`

## Arquitetura proposta

### Estrutura de alto nível
```text
heimdall/
├── src/
│   ├── domain/
│   ├── use_case/
│   ├── application/
│   ├── infra/
│   ├── shared/
│   ├── tests/
│   ├── templates/
│   │   └── default/
│   │       ├── AGENTS.md
│   │       └── .agent/
│   │           ├── skills/
│   │           └── workflows/
│   └── cmd/
│       └── heimdall/
│           └── main.go
├── analise/
├── go.mod
└── README.md
```

### Papéis por camada
- `src/domain/`: entidades e regras centrais (`Assistant`, `AssistantMember`, `TargetPlatform`, políticas de AGENTS).
- `src/use_case/`: casos de uso (`InitTarget`, `StartProjectContext`, `ListAssistantLibrary`, `InstallAssistant`).
- `src/application/`: entrada CLI/chat adapters, parsing e mapeamento para DTOs de casos de uso.
- `src/infra/`: filesystem, catálogo, transformadores por plataforma e instaladores por target.
- `src/shared/`: erros comuns, logging e utilitários transversais.
- `src/tests/`: testes unitários e de integração.

## Estratégia de templates
- Conteúdo-fonte de Assistants em `src/templates/<pacote>/.agent/workflows/...`.
- Skills de suporte em `src/templates/<pacote>/.agent/skills/...`.
- `AGENTS.md` em `src/templates/<pacote>/AGENTS.md`.
- Pacote inicial: `src/templates/default`.

## Estratégia de target (resumo)
- `target` define formato de saída e destino padrão.
- Objetivo de UX: ao final do fluxo, diretório pronto para uso no target escolhido.

## Backlog imediato (próximos passos)
1. Replanejar backlog técnico para jornadas do MVP de Assistants.
2. Consolidar contratos de Assistant como entidade principal.
3. Implementar `init` por target.
4. Implementar `start` com coleta de contexto do projeto.
5. Implementar `list-lib`.
6. Implementar `install assistant` com instalação automática de skills associadas.

## Definição de pronto da Entrega 1
- Jornadas MVP funcionais (`init`, `start`, `list-lib`, `install assistant`).
- Linguagem ubíqua consistente com `Assistant` em toda a stack.
- Instalação deixa diretório pronto para uso no target.
- Documentação de uso no README.
