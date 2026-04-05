# Pontos Estratégicos — Decisões Consolidadas

## Objetivo
Registrar as decisões estratégicas discutidas e aprovadas para guiar o próximo ciclo.

## Decisão 1 — Direcionamento do Produto
**Aprovado:** Heimdall é um produto orientado a **Assistants**.

- Skills continuam relevantes, mas como capacidade interna/associada.
- A experiência principal para usuário final é centrada em Assistant.

## Decisão 2 — Escopo Principal de Criação
**Aprovado:** foco em criação e gestão de **Assistants**.

- Operações de usuário: instalar, listar, usar e futuramente criar/editar/excluir Assistants.

## Decisão 3 — Instalação de Dependências
**Aprovado:** ao instalar um Assistant, instalar automaticamente as skills associadas (padrão).

- Objetivo: evitar setup incompleto e reduzir fricção.

## Decisão 4 — Linguagem Ubíqua (DDD)
**Aprovado:** termo único de domínio e produto = `Assistant`.

- Evitar termos paralelos para a mesma entidade principal.
- Manter consistência em docs, código e experiência de uso.

## Impacto Imediato no Planejamento
1. Backlog passa a seguir jornadas do MVP de Assistants (`init`, `start`, `list-lib`, `install`).
2. Spec canônica reforça Assistant como entidade principal.
3. Comunicação do produto evita alternância de termos concorrentes.
