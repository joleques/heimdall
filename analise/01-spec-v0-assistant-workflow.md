# Spec v0 — Contrato Canônico de Assistant

## Objetivo
Definir o contrato mínimo de **Assistant** como entidade principal do produto, com transformação previsível para targets:
- `antigravity`
- `codex`
- `claude`
- `cursor`

## Linguagem Ubíqua
- Termo canônico único: `Assistant`.
- `Workflow` representa a execução/orquestração de um Assistant.
- `Skill` representa capacidades internas/associadas de um Assistant.

## Escopo da v0
- Modelo canônico de Assistant.
- Validação obrigatória mínima.
- Relação Assistant -> Skills associadas.
- Mapeamento de saída por target.

## Entidade Canônica (`AssistantSpec`)

### Campos obrigatórios
1. `id` (`string`)
2. `name` (`string`)
3. `description` (`string`)
4. `instructions` (`string`)

### Campos opcionais
1. `version` (`string`, default: `"0.1.0"`)
2. `inputs` (`array<InputSpec>`, default: `[]`)
3. `tools` (`array<string>`, default: `[]`)
4. `tags` (`array<string>`, default: `[]`)
5. `metadata` (`map<string,string>`, default: `{}`)
6. `skills` (`array<string>`, default: `[]`) — lista de skills associadas ao Assistant.

### Subtipo `InputSpec`
1. `name` (`string`, obrigatório)
2. `description` (`string`, obrigatório)
3. `required` (`bool`, default: `false`)
4. `default` (`string`, opcional)

## Regras de validação (v0)
1. `id` em `kebab-case`, regex: `^[a-z0-9]+(-[a-z0-9]+)*$`.
2. `name`, `description`, `instructions` não podem ser vazios.
3. `instructions` deve ter pelo menos 30 caracteres.
4. `inputs[].name` deve ser único por Assistant.
5. `tools[]` não pode ter duplicatas.
6. `tags[]` não pode ter duplicatas.
7. `skills[]` não pode ter duplicatas.

## Formato-fonte no template
Origem recomendada:
- `src/templates/<pacote>/.agent/workflows/<assistant-id>.yaml`

YAML esperado:
```yaml
id: instagram-post-studio
name: Instagram Post Studio
description: Assistant para criar posts completos de Instagram.
instructions: |
  Você coordena pesquisa, design, validação e geração de imagem para posts.
version: 0.1.0
skills:
  - researcher
  - designer
  - validator
  - image-generator
inputs:
  - name: tema
    description: Tema principal do post.
    required: true
tools:
  - shell
tags:
  - social-media
metadata:
  owner: growth-team
```

## Regra crítica de instalação
Ao instalar um Assistant, as `skills` associadas devem ser instaladas automaticamente por padrão.

## Pipeline de transformação
1. `Load`: ler YAML canônico do Assistant.
2. `Validate`: validar contrato.
3. `Normalize`: aplicar defaults e deduplicações.
4. `ResolveSkills`: resolver skills associadas no catálogo.
5. `Transform(target)`: converter para formato da plataforma.
6. `Install(targetLayout)`: gravar Assistant e skills no destino.
7. `Report`: retornar `installed/skipped/failed/warnings`.

## Regras de destino e UX
1. `target` define formato e destino padrão.
2. Usuário não precisa montar estrutura manualmente.
3. Fluxo deve deixar diretório pronto para uso no target.

## Política de `AGENTS.md`
1. `skip`
2. `if-missing` (default)
3. `overwrite`

## Casos de teste mínimos (v0)
1. Instalar 1 Assistant com skills associadas em cada target.
2. Instalar todos Assistants (`install assistant <target>` sem lista).
3. Erro por `id` inválido.
4. Erro por `instructions` vazio.
5. Warning por tool não suportada.
6. Idempotência em execução repetida.
7. Política default `AGENTS.md = if-missing`.
