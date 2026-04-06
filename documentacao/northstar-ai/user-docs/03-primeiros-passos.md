# Primeiros passos

## Objetivo deste guia

Aqui voce vai do zero ao primeiro resultado util:

1. preparar o target do agente
2. abrir o agente com as tools de plataforma disponiveis
3. registrar o contexto do projeto via agente
4. montar a primeira squad via agente
5. instalar os assets aprovados via agente

Os exemplos abaixo usam `northstar` como nome do comando.

Se voce estiver usando os binarios de `dist/`, troque `northstar` pelo arquivo correspondente, como `./dist/latest/northstar-ai-latest-linux-amd64` ou `.\dist\latest\northstar-ai-latest-windows-amd64.exe`.

## Passo 1: instalar ou gerar o executavel

### Objetivo

Ter a CLI pronta para ser usada.

### Comando

```bash
go build -o northstar ./src/cmd/northstar
```

### Exemplo real

```bash
./northstar init codex
```

### Resultado esperado

O comando roda sem erro de executavel ausente.

## Passo 2: inicializar o target

### Objetivo

Preparar a estrutura base para o target do agente no seu projeto.

### Comando

```bash
northstar init codex
```

### Exemplo real

```bash
./northstar init codex
```

### Resultado esperado

No target `codex`, o projeto passa a ter pelo menos:

```text
.codex/
└── skills/

.northstar/
├── context/
│   └── project-context.yaml
└── template/
```

Tambem e esperado que o `.gitignore` passe a incluir `.northstar`.

## Passo 3: abrir o agente de IA

### Objetivo

Usar o ambiente do target que acabou de ser preparado.

### Acao

Abra o agente de IA correspondente ao target inicializado.

Exemplos:

- se voce rodou `init codex`, siga no Codex
- se voce rodou `init claude`, siga no Claude
- se voce rodou `init cursor`, siga no Cursor

### Resultado esperado

As tools de plataforma instaladas pelo `init` ficam disponiveis para uso na conversa com o agente.

## Passo 4: registrar o contexto do projeto via agente

### Objetivo

Salvar as informacoes basicas que descrevem o projeto.

### Acao

Peca para o agente executar a tool `northstar-start`.

### Exemplo real

```text
$northstar-start
Registre o contexto do projeto com o titulo "Northstar AI", descricao "CLI para preparar projetos para uso com Assistants de IA" e documentacao README.md.
```

### Resultado esperado

O arquivo `.northstar/context/project-context.yaml` e atualizado com:

- target do projeto
- titulo
- descricao
- documentacao informada com `--doc`

O target ja vem do `init`, entao o agente consegue preencher esse fluxo usando o contexto do projeto.

## Passo 5: listar a biblioteca disponivel via agente

### Objetivo

Descobrir quais assistants e skills a biblioteca padrao oferece para reaproveitamento.

### Acao

Peca para o agente executar a tool `northstar-list-lib`.

### Exemplo real

```text
$northstar-list-lib
Liste os assistants e skills da categoria documentation.
```

### Resultado esperado

Voce recebe a lista de assistants disponiveis e, quando pedir skills, tambem a lista de skills.

O filtro por categoria ajuda a mostrar apenas um grupo, como `documentation`.

## Passo 6: montar a primeira squad via agente

### Objetivo

Transformar a necessidade do usuario em uma composicao de lideranca e especialistas.

### Acao

Peca para o agente executar a tool `northstar-squad-builder`.

### Exemplo real

```text
$northstar-squad-builder
Quero montar uma squad para documentar um produto para usuario final, aproveitando ao maximo as skills ja existentes e criando algo novo so se realmente precisar.
```

### Resultado esperado

O agente:

- entrevista voce para entender a demanda
- busca reuso na biblioteca local
- classifica a complexidade
- propoe nome da squad, lideranca e membros
- pede sua aprovacao antes de consolidar

## Passo 7: instalar assistants ou skills via agente

### Objetivo

Materializar no projeto a squad ou os assets aprovados.

### Acao

Peca para o agente executar a tool `northstar-install`.

### Exemplo real

```text
$northstar-install
Instale os assets da categoria documentation neste projeto.
```

### Resultado esperado

O Northstar AI instala os assets correspondentes ao target atual e aplica as skills associadas quando necessario.

## Passo 8: confirmar que o projeto esta pronto

### Objetivo

Verificar se o setup inicial terminou.

### Acao

Confirme estes itens no seu projeto:

- existe uma pasta do target, como `.codex/`
- existe `.northstar/context/project-context.yaml`
- a biblioteca pode ser consultada via agente com `northstar-list-lib`
- uma proposta de squad pode ser feita via `northstar-squad-builder`
- pelo menos um assistant ou skill foi instalado com sucesso via `northstar-install`

### Resultado esperado

Seu projeto esta preparado para continuar o uso pelo agente de IA. No fluxo recomendado do produto, o terminal ja cumpriu seu papel no `init`.
