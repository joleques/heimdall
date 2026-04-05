---
description: Orquestra a análise e revisão de impacto para migração AWS através das skills migration-analyzer e migration-analyzer-revisor
---

# 🚀 Workflow: Análise de Migração AWS (`/analise-migracao-aws`)

Este workflow orquestra a criação do ambiente de contexto isolado (via Symlinks) e a execução em cascata dupla (**Maker/Checker**) para gerar os artefatos vitais de migração de uma arquitetura para a nuvem.

## Parâmetros de Entrada
- **Contexto (Obrigatório):** O nome do cluster/Bounded Context para nomear a pasta isolada (ex: `financeiro`, `pedidos`).
- **Lista de Repositórios (Obrigatório):** Caminhos absolutos de onde estão os códigos originais. **Mínimo:** 1 repositório. **Máximo:** Ilimitado.
- **Padrões da Arquitetura (Obrigatório):** Formato das *environment variables* injetadas e URLs predefinidas da corporação (ex: domínios `.internal`).

---

## 🛠️ Passos de Execução

1. **Validação Rígida de Setup**
   Certifique-se que o usuário passou o **mínimo de 1 repositório válido**. Se faltou input, aborte e cobre o repasse correto.

2. **Setup Rápido via SO (KISS/YAGNI)**
   Crie as subpastas `migracao-{contexto}/repos` e `migracao-{contexto}/analise/` usando bash.
// turbo
3. **Criação Nativa de Symlinks**
   Para CADA repositório fornecido, lance comandos de Linux via terminal (`ln -s /caminho/origem ./repos/origem`) para plugar as referências com segurança e não clonar arquivos brutos.

4. **Geração Inicial (Maker Phase)**
   Acione a skill `migration-analyzer` designando o escopo estrito à nova pasta `/repos` e aos seus diretórios `/infra`. Solicite a geração dos 3 relatórios primários na pasta `/analise/`.

5. **Loop de Refinamento (Checker Phase — Máx 5x)**
   - Invoque imediatamente a skill revisora `migration-analyzer-revisor` apontando para o resultado recém-produzido.
   - Pule os laudos se constarem erros: Caso o *Checker* aponte reprovação `[REPROVADO ❌]`, retroalimente a skill `migration-analyzer` com a bronca técnica para refatorar o `migration-waves.md` ou os alertas omitidos.
   - Repita o loop até validação `[APROVADO ✅]` pelo *Checker* ou exaustão das 5 tentativas de orquestração.

6. **Handover Final**
   Notifique o usuário com um sumário conciso do Laudo Aprovado e entregue os paths físicos para a conferência humana.
