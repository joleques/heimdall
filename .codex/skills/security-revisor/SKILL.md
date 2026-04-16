---
name: security-revisor
description: Revisa seguranca da entrega implementada, com foco em LGPD, validacao de entrada, riscos de injecao e segredo hardcoded.
---

# Security Revisor

Use esta skill apos `bug` ou `implementacao`, antes da entrega final.
Nao use para `analise`.

## Objetivo

Validar riscos basicos de seguranca na mudanca e bloquear entrega quando houver falha critica.

## Coleta minima de contexto

1. Tipo da demanda (`bug` ou `implementacao`)
2. Arquivos alterados
3. Testes executados e resultado
4. Variaveis/configuracoes novas de ambiente

## Checklist obrigatorio

### LGPD e dados sensiveis

- Dados pessoais e sensiveis possuem necessidade clara de coleta/processamento.
- Nao ha exposicao indevida de PII em logs, erros ou payloads.
- Quando aplicavel, ha mascaramento/anonimizacao.

### Validacao e injecao

- Entradas externas sao validadas e saneadas.
- Nao ha construcao insegura de query/comando (SQL/NoSQL/command injection).
- Nao ha concatenacao insegura para interpretadores.

### Segredos e configuracao

- Nao ha segredo hardcoded (token, senha, chave privada, credencial).
- Segredos ficam em ambiente/cofre, nao em codigo ou documento de entrega.

## Criticidade

Classifique achados com:

- `critico`: bloqueia entrega
- `alto`: deve corrigir antes de concluir
- `medio`: corrigir no ciclo atual quando possivel
- `baixo`: recomendacao

## Saida estruturada obrigatoria

Responda sempre com:

- `status`: `PASS` ou `FAIL`
- `achados`: lista com severidade, arquivo e descricao
- `evidencias`: trechos/comandos que sustentam a conclusao
- `acoes`: correcoes obrigatorias antes de concluir

## Regra de bloqueio

Se houver achado `critico` ou `alto` nao mitigado, o resultado da skill deve ser `FAIL`.
