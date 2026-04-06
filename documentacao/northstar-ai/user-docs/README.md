# Northstar AI

> Produto para criar squads de trabalho com IA a partir da skill principal `northstar-squad-builder`.

## Comece por aqui

1. Leia [O que e](./01-o-que-e.md)
2. Siga [Instalacao](./02-instalacao.md)
3. Execute [Primeiros passos](./03-primeiros-passos.md)
4. Consulte [Exemplos de uso](./04-exemplos-de-uso.md)
5. Use [Erros comuns e ajuda](./05-erros-comuns-e-ajuda.md) se algo sair do roteiro

## Este material ajuda voce a

- entender se o Northstar AI faz sentido para o seu projeto
- instalar no Linux, macOS e Windows
- executar o primeiro fluxo util com seguranca
- usar a skill principal do produto para montar squads
- resolver erros comuns sem depender de contexto interno do codigo

## Para quem este guia foi escrito

Este guia e para pessoas e times que querem criar squads de trabalho com IA para resolver demandas reais sem sair montando assistents e skills no improviso.

Voce nao precisa conhecer a arquitetura interna do produto. Basta entender o objetivo da demanda e conseguir conversar com o agente.

## Rota recomendada para iniciantes

1. Instale ou gere o executavel
2. Rode `init` para preparar o target do seu agente
3. Abra o agente de IA do target inicializado
4. Use `northstar-squad-builder` para criar sua squad de trabalho

## Como o produto deve ser usado

No fluxo de produto recomendado, o terminal entra apenas no primeiro passo com `init`.

Depois disso, o proprio `init` instala os artefatos de plataforma que tornam `northstar-squad-builder` disponivel no ambiente do agente.

E esse e o ponto mais importante da doc: depois de instalado, `northstar-squad-builder` passa a ser a skill principal do produto.

Ela e a porta de entrada para o trabalho real. E a partir dela que as squads sao criadas, refinadas, reaproveitadas e instaladas.

Na pratica:

- no terminal: `northstar init <target>`
- via agente: comece por `northstar-squad-builder`
- apoios da plataforma: `northstar-start`, `northstar-list-lib` e `northstar-install`

## A skill principal do produto

Se voce ler apenas uma coisa sobre o Northstar AI, leia isto:

`northstar-squad-builder` e a skill principal do produto.

E ela existe para:

- entrevistar o usuario antes de propor qualquer solucao
- entender a intencao real da demanda
- procurar reuso de skills e assistents existentes
- definir a complexidade da necessidade
- propor a melhor composicao de squad
- criar ou atualizar os itens necessarios
- instalar os itens aprovados no target ativo

Em outras palavras, Northstar AI nao e um produto para “instalar coisinhas de IA no projeto”.

Ele existe para transformar uma necessidade de trabalho em uma squad operacional, com:

- uma lideranca
- especialistas com responsabilidade clara
- reuso da biblioteca existente
- instalacao no target correto

## Modelo mental do produto

O Northstar AI trabalha com este modelo:

- `skill` = especialista da squad
- `assistent` = lideranca da squad
- `northstar-squad-builder` = skill principal que desenha essa composicao

Isso significa que o produto nao deveria ser lido como uma lista de comandos.

Ele deve ser entendido como um sistema para formar squads de trabalho a partir de uma necessidade real.

## O sentido do produto

O sentido de existir do Northstar AI e simples:

ajudar voce a sair de uma demanda solta para uma squad de trabalho instalada e pronta para operar.

O `init` so prepara o terreno.

O valor central aparece quando a skill `northstar-squad-builder` entra em acao.

## Observacao importante sobre os nomes dos binarios

No codigo e no README principal do projeto, o produto e tratado como `Northstar AI` e o comando aparece como `northstar`.

Nos binarios precompilados atuais da pasta `dist/`, os arquivos seguem o padrao `northstar-ai-latest-*` e `northstar-ai-beta-*`.
