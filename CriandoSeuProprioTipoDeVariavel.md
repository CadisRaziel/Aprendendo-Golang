# Golang doc (Itens que estudei e entendi)
## Aqui terá alguns exemplos do código go para ser usado em projetos futuros (de como construir um código em go)

> "A persistencia supera o talento"

[![Linkedin Badge](https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https:/https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)](https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)
[![Instagram Badge](https://img.shields.io/badge/-Instagram-a43b9d?style=flat-square&logo=Instagram&logoColor=white&link=https://https://www.instagram.com/vihhbz/?hl=pt-br/)](https://www.instagram.com/vihhstx/?hl=pt-br/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:vitorbrussolo@gmail.com)](mailto:vitorbrussolo@gmail.com)

# Criando seu propio tipo de variavel 


### Lembre-se se eu declarar uma variavel com algum tipo exemplo: 'int', 'string' etc... ela vai ser esse tipo até morrer

```sh
//*O tipo base do hotdog aqui é int
type hotdog int

var c hotdog

//Porém porque criar um tipo prórpio nosso do que usar simplesmente o basico
//A resposta é porque quando criamos nosso propio tipo podemos fazer muitas coisas com ele que a gente não pode fazer com o tipo padrao 'var' ou ':='

func main() {
	fmt.Printf("%T", c)
}
```
