# B-tree

Implementação básica de árvore B usando inteiros.

## Definindo a entrada

Na pasta b-tree haverá um arquivo chamado 'main.go', lá tem um breve exemplo de como criar sua árvore B. Para criar sua árvore use o método Initbtree(o) que inicializa a árvore B vazia com a ordem 'o' dada.

## Rodando o código

Para que o código rode é preciso ter o Go instalado em sua máquina e que todos os arquivos estejam da maneira como foi baixado no arquivo (especialmente a go.mod). Feito isso, apenas abra seu terminal na pasta b-tree e rode o comando 'go run .'.

## Entendendo o resultado

Ao utilizar o método Print() é retornado a representação em string de como está sua árvore B que é organizado da seguinte forma:

    - Cada linha contém o tipo da página, -1 para raiz e 1 para folhas, e o conteúdo da paǵina em um vetor.

    - A árvore é mostrada nível por nível, ou seja, a primeira linha é sempre a raiz e em seguida são impressas todas suas páginas filhas, dependendo da ordem da raiz. A seguir é impressa as páginas filhas dessas e assim por diante, linha por linha.

## Erros rodando o código

É possível que ocorra que na hora de rodar o código o Go não tenha buscado os pacotes exigidos no códido (neste caso a 'main.go' busca o pacote '\btree'que está no mesmo diretório). Caso isso ocorra, tente rodar em seu terminal os seguintes comandos: 'go get github.com/jamillepp/b-tree/btree' e depois 'go tidy' (ou apenas o último). Caso ocorra qualquer outro nesse nível avise-me.