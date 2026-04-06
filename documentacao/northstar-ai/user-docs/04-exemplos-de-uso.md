# Exemplos de uso

## Exemplo minimo

Prepare um projeto para `codex`:

```bash
./northstar init codex
```

Depois, no agente:

```text
$northstar-start
Registre o contexto inicial do projeto com titulo "Projeto Exemplo", descricao "Contexto inicial do projeto" e documentacao README.md.

$northstar-squad-builder
Quero montar uma squad para me ajudar a estruturar a documentacao do produto.
```

Resultado esperado:

- target inicializado
- contexto do projeto salvo via agente
- primeira proposta de squad apresentada via agente

## Exemplo recomendado

Preparar no terminal e continuar tudo pelo agente:

```bash
./northstar init codex
```

No agente:

```text
$northstar-start
Registre o contexto do projeto com titulo "Northstar AI", descricao "CLI para preparar projetos para uso com Assistants de IA" e documentacao README.md.

$northstar-list-lib
Liste os assistants e skills da categoria documentation.

$northstar-squad-builder
Quero montar uma squad para documentacao de produto para usuario final com maximo reuso da biblioteca atual.

$northstar-install
Instale os assets da categoria documentation.
```

Resultado esperado:

- contexto salvo em `.northstar/context/project-context.yaml`
- lista filtrada de assistants e skills de documentacao via agente
- proposta de squad com lideranca e especialistas
- instalacao dos assets relevantes para o target atual via agente

## Fluxo tipico de um usuario iniciante

### 1. Iniciar o target

```bash
./northstar init claude
```

### 2. Registrar o contexto no agente

```text
$northstar-start
Registre o contexto com titulo "Portal do Cliente", descricao "Projeto para atendimento com agentes" e documentacao README.md.
```

### 3. Ver a biblioteca disponivel no agente

```text
$northstar-list-lib
Liste a biblioteca disponivel para este projeto.
```

### 4. Montar a squad no agente

```text
$northstar-squad-builder
Quero uma squad para atender demandas de documentacao de produto com lideranca clara e especialistas reaproveitados quando possivel.
```

### 5. Instalar os itens aprovados

```text
$northstar-install
Instale os assets aprovados para esta squad neste projeto.
```

## Fluxo com variacoes comuns

### Pedindo instalacao por categoria no agente

```text
$northstar-install
Instale todos os assets da categoria media.
```

### Pedindo descoberta de catalogo no agente

```text
$northstar-list-lib
Liste apenas a categoria documentation com skills incluidas.
```

### Pedindo montagem de squad no agente

```text
$northstar-squad-builder
Quero montar uma squad para produzir conteudo tecnico com revisao e reaproveitamento da biblioteca existente.
```

### Pedindo registro de contexto no agente

```text
$northstar-start
Atualize o contexto do projeto com a descricao "Projeto para atendimento com agentes".
```

## Exemplo com binario precompilado

No Linux:

```bash
chmod +x ./dist/latest/northstar-ai-latest-linux-amd64
./dist/latest/northstar-ai-latest-linux-amd64 init codex
```

No Windows:

```powershell
.\dist\latest\northstar-ai-latest-windows-amd64.exe init codex
```
