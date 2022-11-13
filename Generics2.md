# Golang doc (Itens que estudei e entendi)
## Aqui terá alguns exemplos do código go para ser usado em projetos futuros (de como construir um código em go)

> "A persistencia supera o talento"

[![Linkedin Badge](https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https:/https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)](https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)
[![Instagram Badge](https://img.shields.io/badge/-Instagram-a43b9d?style=flat-square&logo=Instagram&logoColor=white&link=https://https://www.instagram.com/vihhbz/?hl=pt-br/)](https://www.instagram.com/vihhstx/?hl=pt-br/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:vitorbrussolo@gmail.com)](mailto:vitorbrussolo@gmail.com)

# Como comparar o tipo 'any'
-  São tipos a serem especificados posteriormente, Se os dois são any, você não consegue comparar os dois, Pensando nisso foi criado um tipo chamado 'comparable' no qual tambem, é um any mas te permite fazer comparações quando necessario.



### Exemplo COM e SEM comparable
```sh

o comparable nos permite utilizar '==' e '!=='

func anyTest[T any](arg1 T, arg2 T) {
	fmt.Println(arg1 == arg2)
}
func main(){
   anyTest(1,2)
}

no terminal vai aparecer:
invalid operation: Cannot compare arg1 == arg2 (operator == not defined on T)

exemplo COM o comparable:
func anyTest[T comparable](arg1 T, arg2 T){
	fmt.Println(arg1 == arg2)
}
func main(){
   anyTest(1,2)
}

//no terminal vai aparecer:
true ou false dependendo da comparação
```

### Utilizando sinal de maior e menor (lembrando o comparable só aceita comparação de igualdade e diferença)
```sh
Mais ai queremos fazer uma comparação de maior e menos (com o comparable não daria)
para fazer isso fariamos o seguinte:

func anyTest[T NumberTests(arg1 T, arg2 T){
	if arg1 > arg2 {

	}
}
type NumberTests interface {
	int64 | float64 | float32
}
```

### Utilizando TOKEN em Go "~"

```sh
Listar um unico tipo é inutil por si só. Para satisfação de restrições, queremos poder dizer
não apenas int, mas "qualquer tipo cujo tipo subjacente seja int".[...]
Queremos que funcione não apenas para fatias dos tipos ordenados pré-declarados, mas também para 
tipos definidos por um programa.

package main
//AnyString -> corresponde a qualquer tipo cujo tipo subjacente seja string.
//isso inclui, entre outros, o próprio tipo string e tipo MyString.
type AnyString interface {
   ~string
}

//AppproximateMyString é invalida.
type AppproximateMyString interface{
   ~MyString //INVALIDO: o tipo subjacente de MyString não é MyString
}

//AppproximateParameter é invalida.
type AppproximateMyString interface{
   ~T //INVALIDO: T é um parametro de tipo
}

//Outros exemplos
//colocando "~" estou dizendo, eu aceito int64 e qualquer tipo que implemente int64
func test[T ~int64](arg1 T, arg2 T){}

type Vitor int64

func main(){	
	var vitorTest Vitor = 20
	test(vitorTest, vitorTest)
}
```

### Structs e interfaces com tipos parametrizadas
```sh
// E se eu quiser definir uma struct na qual o tipo de um dos campos dela
seja definido posteriormente por outra função

type Setter2[T any] interface {
	Set(string)
	*B
}
type ListHead[T any] struct {
	Head *ListElement[T]
}
type ListHead[T any] struct {
	Val T
}
func test() {
	listElementeObject := ListHead[string] {
           Head: &ListElement[string] {
		Val: "teste",
	    },
	  }
	  fmt.Println(listElementObject)
	}
	
	//===========
	
package main

type User[T any, B any] struct {
	name T
	age B
}

//Repare que usango generics eu posso tipar ela só na hora que eu vou utiliza-la
func main(){
	userTest := User[string, int64] {
	    name: "vitor",
	    age: 29,
	}
}
```
