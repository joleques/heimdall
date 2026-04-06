# Instalacao

## Antes de instalar

Hoje o repositório mostra dois caminhos reais de uso:

1. gerar o executavel a partir do codigo-fonte com Go
2. usar um binario precompilado da pasta `dist/`

Nao ha, neste snapshot do projeto, instalador nativo, pacote Homebrew, `.msi`, `.deb` ou equivalente. Milagre de onboarding ainda nao foi empacotado.

## O que voce faz no terminal e o que fica para o agente

No fluxo recomendado do produto:

- o terminal e usado para executar `init`
- o restante acontece via agente de IA

O motivo e simples: o `init` prepara o target e instala as tools de plataforma. Depois disso, o usuario consegue continuar por conversa com o agente, sem precisar seguir executando `start`, `list-lib` e `install` manualmente no terminal.

Entre essas tools, a mais importante para o valor do produto e `northstar-squad-builder`, que ajuda a montar a squad de trabalho adequada para a necessidade do usuario.

## Linux

### Opcao 1: gerar a partir do codigo-fonte

#### Pre-requisitos

- Go `1.25.0` ou compativel com o projeto
- terminal com acesso ao repositorio

#### Como instalar

Na raiz do projeto:

```bash
go build -o northstar ./src/cmd/northstar
```

#### Como validar se funcionou

Execute:

```bash
./northstar init codex
```

Resultado esperado:

- a CLI executa sem erro de comando nao encontrado
- a pasta `.codex/skills` e criada no diretório escolhido
- o arquivo `.northstar/context/project-context.yaml` e criado

### Opcao 2: usar o binario precompilado do repositorio

Arquivos disponiveis:

- `dist/latest/northstar-ai-latest-linux-amd64`
- `dist/latest/northstar-ai-latest-linux-arm64`
- `dist/beta/northstar-ai-beta-linux-amd64`
- `dist/beta/northstar-ai-beta-linux-arm64`

#### Como instalar

Escolha o arquivo correto para a arquitetura da sua maquina e torne-o executavel:

```bash
chmod +x ./dist/latest/northstar-ai-latest-linux-amd64
```

#### Como validar se funcionou

Execute o binario diretamente:

```bash
./dist/latest/northstar-ai-latest-linux-amd64 init codex
```

Resultado esperado:

- a CLI executa normalmente
- o projeto recebe a estrutura inicial do target

#### Problemas comuns na instalacao

- `Permission denied`
  Causa provavel: o binario nao recebeu permissao de execucao.
  Como corrigir:

```bash
chmod +x ./dist/latest/northstar-ai-latest-linux-amd64
```

- `command not found`
  Causa provavel: voce tentou chamar `northstar` sem ter gerado um binario com esse nome nem colocado o executavel no `PATH`.
  Como corrigir: execute pelo caminho do arquivo, por exemplo `./northstar` ou `./dist/latest/northstar-ai-latest-linux-amd64`.

## macOS

### Opcao 1: gerar a partir do codigo-fonte

#### Pre-requisitos

- Go `1.25.0` ou compativel com o projeto
- terminal com acesso ao repositorio

#### Como instalar

Na raiz do projeto:

```bash
go build -o northstar ./src/cmd/northstar
```

#### Como validar se funcionou

Execute:

```bash
./northstar init codex
```

Resultado esperado:

- a CLI abre e processa o comando
- os arquivos iniciais do target sao criados no projeto

### Opcao 2: usar binario precompilado

Arquivos disponiveis:

- `dist/latest/northstar-ai-latest-darwin-amd64`
- `dist/latest/northstar-ai-latest-darwin-arm64`
- `dist/beta/northstar-ai-beta-darwin-amd64`
- `dist/beta/northstar-ai-beta-darwin-arm64`

#### Como instalar

Torne o arquivo executavel:

```bash
chmod +x ./dist/latest/northstar-ai-latest-darwin-arm64
```

#### Como validar se funcionou

```bash
./dist/latest/northstar-ai-latest-darwin-arm64 init codex
```

Resultado esperado:

- o comando roda
- a estrutura inicial do target e criada

#### Problemas comuns na instalacao

- erro de permissao ao executar
  Causa provavel: falta de permissao de execucao.
  Como corrigir:

```bash
chmod +x ./dist/latest/northstar-ai-latest-darwin-arm64
```

- voce escolheu o binario errado
  Causa provavel: confusao entre `amd64` e `arm64`.
  Como corrigir: baixe ou use o arquivo correspondente ao seu Mac.

## Windows

### Opcao 1: gerar a partir do codigo-fonte

#### Pre-requisitos

- Go `1.25.0` ou compativel com o projeto
- Prompt de Comando ou PowerShell

#### Como instalar

Na raiz do projeto:

```powershell
go build -o northstar.exe ./src/cmd/northstar
```

#### Como validar se funcionou

```powershell
.\northstar.exe init codex
```

Resultado esperado:

- o executavel abre sem erro
- a estrutura inicial do target e criada no projeto

### Opcao 2: usar binario precompilado

Arquivos disponiveis:

- `dist/latest/northstar-ai-latest-windows-amd64.exe`
- `dist/latest/northstar-ai-latest-windows-arm64.exe`
- `dist/beta/northstar-ai-beta-windows-amd64.exe`
- `dist/beta/northstar-ai-beta-windows-arm64.exe`

#### Como instalar

Nao ha instalador dedicado neste snapshot. Use o arquivo `.exe` diretamente.

#### Como validar se funcionou

```powershell
.\dist\latest\northstar-ai-latest-windows-amd64.exe init codex
```

Resultado esperado:

- a CLI executa
- o target e inicializado no projeto

#### Problemas comuns na instalacao

- Windows bloqueia o binario baixado
  Causa provavel: o sistema marcou o arquivo como vindo da internet.
  Como corrigir: desbloqueie o arquivo nas propriedades do Windows ou mova o binario para um diretório confiavel antes de executar.

- comando nao encontrado
  Causa provavel: voce tentou rodar `northstar` sem incluir o executavel no `PATH`.
  Como corrigir: execute pelo caminho completo, por exemplo `.\northstar.exe`.

## Como validar o projeto antes de distribuir a CLI

Se voce estiver preparando o Northstar AI localmente, o proprio repositório recomenda:

```bash
go test ./...
go build ./...
```

## Lacunas conhecidas de instalacao

- os binarios em `dist/` usam nomes de artefato como `northstar-ai-latest-*` e `northstar-ai-beta-*`, enquanto o comando de uso continua sendo `northstar`
- nao ha processo oficial documentado para instalacao global no sistema
- nao ha pagina de releases ou URL de download documentada nas fontes analisadas
