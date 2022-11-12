# Golang doc (Itens que estudei e entendi)
## Aqui terá alguns exemplos do código go para ser usado em projetos futuros (de como construir um código em go)

> "A persistencia supera o talento"

[![Linkedin Badge](https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https:/https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)](https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)
[![Instagram Badge](https://img.shields.io/badge/-Instagram-a43b9d?style=flat-square&logo=Instagram&logoColor=white&link=https://https://www.instagram.com/vihhbz/?hl=pt-br/)](https://www.instagram.com/vihhstx/?hl=pt-br/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:vitorbrussolo@gmail.com)](mailto:vitorbrussolo@gmail.com)

# Estruturas de pastas e GOPATH

### /PKG
```sh
Lugar onde o go guarda todas as dependencias ja instaladas em seus projetos.
(o go nao cria uma pasta build.gradlew onde fica todas dependencias ja baixadas)
o go quando voce instala alguma coisa ele literalmente da um 'git clone' da dependencia por isso que o modulo tem que ter a url identica ao repositório
ou seja ele vai até o repositório, da um git clone e joga para dentro da /pkg
no que isso pode ajudar ? -> vamos supor que dois projetos utiliza a mesma dependencia, o go como aponta pra pkg pra pegar essa dependencia, voce nao vai precisar baixar"clonar"
duas vezes, ele ja vai ta instalado
```

### /src
```sh
lugar onde todo o seu codigo go deve ser guardado
com isso ele tem a estrutura de pastas certinha pra compartilha dependencia
as vezes voce pode utilizar o pacote de um projeto em outro projeto só organizando certinho as pastas (se os dois projetos tiverem no GOPATH, ou seja localmente)
```

### /bin 
```sh
Lugar onde o go guarda os binarios que 'go install' compila
```

# GOPATH vs GOROOT

### $GOROOT 
```sh
é onde fica guardado o código do go, literalmente. Codigo do compilador, das ferramentas e etc... é para la que o código 'go' aponta para executar os comandos 
do Golang, assim ele não é nosso código fonte. o $GOROOT geralmente fica em algo como /usr/local/go.
```

### $GOPATH  
```sh
geralmente é algo como $HOME/go (onde guardamos nosso código/dependencias)
```
# Pacotes 
```sh
Códigos go são organizados em pacotes. Um pacote representa todos os arquivos em um unico diretório no seu PC. Um diretório pode conter apenas alguns certos arquivos
de algum certo pacote. Os pacotes são guardados a baixo do diretório $GOPATH/src.
(Cada pasta é um pacote)
```
# Modules  
```sh
è uma coleção de pacotes GO guardados em um arquivo com nome de go.mod na raiz do projeto.
(lugar onde guarda todos pacotes) (dependencias requeridas no projeto) (caminho do module no repositório)
```
# Versão dos module  
```sh
utilize branches com versão ou tags com versão do github
```
### Estrutura do go.mod 
```sh
module<nome>
go<versao-do-go>
require<caminho-do-module> <versão-do-pacote> (pacotes que seu projeto precisam para executar)
replace<caminho-do-module> => /home/test/test-module (pacotes que voce quer substituir por alguma versão que voce tenha local daquele mesmo pacote
exclude<caminho-do-module> <versão-do-pacote> (pacotes/versão que voce nao quer que executem no seu projeto

no require podemos ter um '//indirect depois do modulo, o que é isso? -> dois casos que podem estar acontecendo:
primeiro: uma dependencia minha precisa desse pacote para executar
segundo: da um go get nessa dependencia mais nao usa no projeto
```
### Estrutura do go.sum
```sh
quando der um go.mod init e baixar as coisas vai aparece ese arquivo chamado go.sum (ele é auto gerado)
ele aponda os hashes do comite, versão e nomes(nao altere ele)
```

# Utilizando uma versão especifica 
```sh
go get github.com/gin-gonic/gin@v1.7.2
```

# Diferença entre Go install e go get -u -v
```sh
go get -u -v -> sai baixando e atualizando versão de tudo
se voce quer que atualize utilizae go get -u -v
se voce quer baixar e manter as versoes utilize go install
```
