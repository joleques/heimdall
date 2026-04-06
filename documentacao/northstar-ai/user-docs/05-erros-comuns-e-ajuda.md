# Erros comuns e ajuda

## `invalid target platform`

### Causa provavel

Voce informou um target fora da lista suportada.

### Como corrigir

Use um destes valores:

- `codex`
- `antigravity`
- `claude`
- `cursor`

Exemplo:

```bash
./northstar init codex
```

## `missing value for --output`

### Causa provavel

Voce informou a flag `--output` sem passar o diretório.

### Como corrigir

Passe o valor logo depois da flag:

```bash
./northstar init codex --output /tmp/meu-projeto
```

## Apos o `init`, eu ainda preciso usar o terminal?

### Causa provavel da duvida

A documentacao anterior dava a entender que o fluxo completo devia continuar na CLI.

### Como corrigir

No fluxo recomendado do produto, nao.

Depois do `init`, siga pelo agente de IA usando:

- `northstar-start`
- `northstar-list-lib`
- `northstar-install`

Use os comandos de terminal `start`, `list-lib` e `install` apenas como referencia tecnica ou apoio interno, nao como trilha principal do usuario final.

## `missing value for --category`

### Causa provavel

Voce chamou `list-lib` ou `install` com `--category`, mas nao informou a categoria.

### Como corrigir

Exemplo valido:

```text
$northstar-list-lib
Liste os assistants e skills da categoria documentation.
```

Ou:

```text
$northstar-install
Instale os assets da categoria documentation.
```

## `unknown option`

### Causa provavel

A CLI recebeu uma flag que nao existe para aquele comando.

### Como corrigir

Revise a sintaxe do comando. Para usuario final, o principal e:

```bash
./northstar init codex --agents-policy overwrite --force
```

Depois disso, prefira usar o agente de IA com as tools de plataforma.

## `project context not found`

### Causa provavel

Voce ainda nao executou `init` no projeto, entao as tools de plataforma e o contexto base nao foram preparados.

### Como corrigir

Rode primeiro:

```bash
./northstar init codex
```

Depois, no agente:

```text
$northstar-start
Registre o contexto inicial do projeto.
```

## `commands 'install skills' and 'install all' were removed`

### Causa provavel

Voce usou uma sintaxe antiga da CLI. Isso pode aparecer em testes, automacoes ou material tecnico interno, mas nao deve ser o fluxo principal do usuario final.

### Como corrigir

Troque:

```bash
./northstar install skills codex skill-a
```

Por:

```text
$northstar-install
Instale o asset skill-a.
```

Ou, se quiser instalar por categoria:

```text
$northstar-install
Instale os assets da categoria documentation.
```

## O binario existe, mas o nome nao bate com o comando do guia

### Causa provavel

Os binarios distribuidos em `dist/` usam nomes de artefato como `northstar-ai-latest-*` e `northstar-ai-beta-*`, enquanto o comando de uso no dia a dia continua sendo `northstar`.

### Como corrigir

Use o caminho real do arquivo que voce tem em maos:

```bash
./dist/latest/northstar-ai-latest-linux-amd64 init codex
```

Ou gere seu proprio binario com o nome `northstar`:

```bash
go build -o northstar ./src/cmd/northstar
```

## Quando pedir ajuda ao time

Procure ajuda quando:

- nenhum dos targets suportados atende ao seu fluxo
- a biblioteca esperada nao aparece quando voce usa `northstar-list-lib`
- voce precisa de um instalador oficial para distribuicao interna
- a diferenca entre o nome do comando `northstar` e o nome dos artefatos em `dist/` impacta a adocao do produto

## Informacoes uteis para enviar no pedido de ajuda

Inclua:

- sistema operacional
- comando executado
- mensagem de erro completa
- target usado, como `codex` ou `claude`
- se voce estava usando binario de `dist/` ou build local
