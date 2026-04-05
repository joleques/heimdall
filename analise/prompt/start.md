ok, vamos dar 2 passos atras e repensar o projeto, vou trazer minha ideia e rediscutimos e ja olhando os pntos estrategicos de antes.

a ideia inicial foi instalar skills e workflows (assitentes) em diversos targets (codex, antigravity, cursor e claude), so que no meio do caminho me veio q o legal sera somente assistentes e depois que na real o legal seria se o usuario pudesse criar suas squas cognitivas, onde cada membro da squad tivesse uma competencia (skill) para desempenhar um papel e aquela squad realizasse um trabalho ou processo sendo assim um workflow (assistente). Exemplo: preciso criar uma squad capaz de montar posts para o instagram (isso é o assitente), para isso é necessario alguem que faça pesquisas sobre o tema proposto (skill pesquisador), alguem que monte o design do post (skill desginer) alguem que revise para ver se esta de acordo com o proposto (skill validador) e por ultimo o gerado de png para o post (skill gerador de imagens).
A ideia é ter alguns assistentes prontos para o usuario trabalhar, que são os workflows que ja temos, mas o usuario podera criar o deles, editar excluir ou seja fazer toda a gestão dele.

# Jornadas do usuario (exemplos usados pensando no antigravity, mas cada target tem sua propria sintax de comunicação no chat)

- Como seria a jornada de inicialização, o usuario entra no dir que ele deseja trabalhar com o assistentes e diria no terminal algo assim heimdall init [codex|Antigravity|cursor|claude] ai toda a estrutura para o usuario.
- Jornada de inicialização do Heimdall (exemplo antigravity mas poderi ser outro da lista), o usuario digitaria /heimdall start, perguntaria titulo do projeto, descrição/contexto do projeto, documentação do projeto. com isso ele poderia montar todo o contexto do projeto  para os agentes de IA saberem o contexto de trabalho deles.
Jornada listar agentes existentes 
- Jornada listar assistentes existentes, o usuario digitaria no chat /heimdall list-lib, e teria a lista da biblioteca de assistentes para ele poder instalar com as descrições dos assitentes.
- Jornada instalação de assitentes, o usuario o usuario digitaria no chat /heimdall install [assitente], então aquele assistente seria instalado e o usuario poderia usa-lo, cada membro da squad teria um nome e o assitente também, usar nomes comuns pessoas e podendo ser feminino e masculino escolher de forma aleatoria

Na entrega 1 (MVP), podemos trabalhar com essas jornadas e depois vamos evoluindo.