# Workflows no Codex

No Codex, workflows não são executados como entidade nativa. Neste repositório, os arquivos em `.agent/workflows/` ficam como playbook de referência e a execução real foi convertida para skills orquestradoras em `.agent/skills/`.

## Mapeamento

- `analise-migracao-aws.md` -> `workflow-analise-migracao-aws`
- `doc-api.md` -> `workflow-doc-api`
- `doc-produto.md` -> `workflow-doc-produto`
- `fine-tuning-gemini.md` -> `workflow-fine-tuning-gemini`
- `init-bounded-context.md` -> `workflow-init-bounded-context`
- `init-project.md` -> `workflow-init-project`
- `write-tech-article.md` -> `workflow-write-tech-article`

## Observações de adaptação

- `notify_user` e `PathsToReview`: substituídos por resposta direta no chat.
- `read_url_content`/`search_web`: substituir por `web.search_query` + `web.open`.
- Slash commands (ex: `/doc-api`): tratados como gatilho textual na skill.
