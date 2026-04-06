# O que e o Northstar AI

## Resumo rapido

Northstar AI e um produto para criar squads de trabalho com IA dentro do seu projeto.

Ele usa um modelo simples:

- `skill` = especialista que executa uma funcao especifica
- `assistent` = lideranca que coordena a squad

Em vez de deixar cada demanda virar uma gambiarra artesanal de prompts, arquivos e ferramentas soltas, o Northstar AI ajuda a transformar a necessidade do usuario em uma squad coerente, reutilizavel e instalada no target correto.

## Qual problema o produto resolve

O problema central que o Northstar AI resolve e este:

como montar a combinacao certa de lideranca e especialistas para uma demanda, sem inventar papeis do nada, sem duplicar o que ja existe e sem espalhar responsabilidade sem dono.

Para isso, o produto ajuda voce a:

- inicializar o ambiente do target
- registrar o contexto do projeto
- descobrir o que ja existe para reuso
- montar squads com base em complexidade e responsabilidade
- instalar a squad ou os assets necessarios no projeto

## Quem deve usar

Faz sentido para:

- times que querem criar squads de IA com responsabilidade bem definida
- pessoas que usam `codex`, `claude`, `cursor` ou `antigravity` como target
- quem quer evitar criar assistents e skills no improviso
- quem quer reaproveitar a biblioteca existente antes de sair clonando ideia com nome novo

## Quando vale a pena usar

Use o Northstar AI quando voce quiser:

- transformar uma necessidade em uma squad de trabalho
- decidir se vale reaproveitar uma skill existente ou criar algo novo
- registrar contexto do projeto em um formato padrao
- instalar a composicao certa de lideranca e especialistas no target

## Targets suportados no MVP

O projeto aceita estes targets:

- `codex`
- `antigravity`
- `claude`
- `cursor`

## Beneficios principais

- squads montadas com mais criterio e menos achismo
- reaproveitamento da biblioteca local antes de criar coisa nova
- registro padrao do contexto do projeto
- instalacao no target certo com menos trabalho manual

## Como a experiencia funciona no produto

O uso recomendado do Northstar AI e dividido em duas etapas:

- terminal, uma unica vez: `init`
- agente de IA, para o restante do fluxo

Depois do `init`, o projeto ja fica com as tools de plataforma instaladas. Isso permite que tarefas como registrar contexto, buscar reuso, montar squads e instalar assets sejam feitas pela conversa com o agente.

## A tool mais importante do produto

A tool que melhor representa o sentido do Northstar AI e `northstar-squad-builder`.

Ela existe para:

- entrevistar o usuario antes de propor qualquer solucao
- identificar a intencao e o nivel de complexidade da demanda
- procurar reuso na biblioteca local
- propor a composicao da squad
- materializar e instalar os itens aprovados

Se fosse para resumir o produto em uma frase honesta, seria esta:

Northstar AI existe para montar squads de IA orientadas a trabalho real, e nao apenas para instalar arquivos em pastas.

## O que ele nao tenta ser

Este guia nao trata o Northstar AI como uma plataforma visual ou um painel web, porque ele nao e isso.
