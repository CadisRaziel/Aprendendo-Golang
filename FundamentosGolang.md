# Golang Fundamentos (Itens que estudei e entendi)
## Aqui terá alguns exemplos do código go para ser usado em projetos futuros (de como construir um código em go)

> "A persistencia supera o talento"

[![Linkedin Badge](https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https:/https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)](https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)
[![Instagram Badge](https://img.shields.io/badge/-Instagram-a43b9d?style=flat-square&logo=Instagram&logoColor=white&link=https://https://www.instagram.com/vihhbz/?hl=pt-br/)](https://www.instagram.com/vihhstx/?hl=pt-br/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:vitorbrussolo@gmail.com)](mailto:vitorbrussolo@gmail.com)

## Func Main
- É a porta de entrada e porta de saida dos programas em golang (Tipo o main.dart)

## Imports em Go
- Podemos fazer importes com alias, imports diretos
- 1 tipo "Alias" -> Podemos por um apelido a nosso imports.
```sh
import (
 apelido "fmt"
)
func main (){
    apelido.Println("Ola")
}
```
- 2 tipo "Direto" . ->  Podemos utilizar os atributos de nosso import diretamente sem precisar por o nome dele antes, utilizando um "ponto"
```sh
import (
 . "fmt"
)
func main (){
    Println("Ola")
}
```

## Go inferindo os tipos
- Podemos deixar que o Go nos infira o tipo das variaveis
```
variavelString := "ooioioio"
	//%T -> Printa o tipo da variavel
	fmt.Printf("%T", variavelString)
```

## Utilizando tipos 'dynamic' no Go (Utilizamos a palavra chave 'interface{}')
```
var varDinamico interface{}
	varDinamico = 1 //pode ser qualquer tipo aqui
	fmt.Println(varDinamico)
```

## Criando maps e utilizando interfacfe{}

```
testJson := map[string]interface{}{
		"saude":    1,
		"economia": "oi",
		"salario":  1.50,
		"feliz":    true,
	}
```

## Array e Slice
```
    //Array -> Lista com tamanho definido
	var arrayList [4]string = [4]string{"Test1", "Test2", "Test3", "Test4"}
	fmt.Println(arrayList)
	fmt.Println(cap(arrayList))
	fmt.Println(len(arrayList))

	//Slices -> Lista com tamanho indefinido (repare que no slice não definimos valor nas chaves "[]"
	//len -> tamanho que ele tem
	//cap -> tamanho que ele suporta
	var sliceList []string = []string{"Test1", "Test2", "Test3", "Test4", "Test5", "Test6", "Test7", "Test8", "Test9", "Test10"}
	fmt.Println(sliceList)
	fmt.Println(cap(sliceList))
	fmt.Println(len(sliceList))

    OBSERVAÇÕES PARA SLICE
	//o go aloca um tamanho fixo em memoria, porém se adicionamos mais itens dentro de um  slice ele dobra esse tamanho (para nao dar problema de memoria)
	//exemplo, tenho 2 itens no slice, o go adiciona o dobro desses itens então o slice tera 4 itens porém 2 esta vago
	//exemplo, tenho 4 itens no slice, o go adiciona o dobro desses itens então o slice tera 8 itens porém 4 esta vago
	//exemplo, tenho 8 itens no slice, o go adiciona o dobro desses itens então o slice tera 16 itens porém 8 esta vago
	//com isso o go não vai deixar o slice pesado.
	//ja em array, como definimos um valor pra ele, então o go não vai te essa preocupação de gerenciar tamanhos e dobra-lo etc..
```

## Função publica e privada
```
//função privada -> começa com letra MINUSCULA (essa função não poderá ser acessada por arquivos fora desse !!)
func funcPrivado() {}

//função Publica -> começa com letra MAIUSCULA (essa função poderá ser acessada por arquivos fora desse !!)
func FuncPublica() {}
```

## Função com retorno multiplo
```
func funTest() (string, int) {
	return "teste", 2
}

func funTest1() (string, error) {
	return "teste", nil
}

func funTest2() (string, error) {
	return "teste", errors.New("teste")
}
//exemplo de utilizaçao funTest2
value, err := funcTest2()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)

```
## Passando função como parametro
```
//passando funções como parametros

func retornoFun (value func(string, int))  {
	value("Vitor", 29)
}
//exemplo de utilização da retornoFun
 funcaoTeste := func(test s	tring, testInt int) {
	 fmt.Println(test, testInt)
}
retornoFun(funcaoTeste)
``` 

## Recebendo função como retorno
```
//recebendo funções como retorno
func retornoZun ( )  func(string, int){
	recebeFunc :=  func(valorString string, valorInt int){
		fmt.Println(valorString, valorInt)
	}
	return  recebeFunc
}

//exemplo de utilização da retornoZun
funcao := retornoZun()
funcao("Vitor", 29)
```

## Função com retorno nomeado
```
//retorno nomeado
func funTest3() (retornoString string, retornoError error) {
	// = -> como ela ja esta instanciada no parametro não precisamos utilizar a atribuição ":=" e sim somente esse operador "="=
	retornoError = nil
	retornoString = "test"

	//aqui só damos o retorno
	return
}
```

## Função anonima
```
//função anonima

func retornoBoom ( )  func(string, int){
	func(valorString string, valorInt int){
		fmt.Println(valorString, valorInt)
	} ("Vitor", 29)
	//repare que a função anonima eu coloquei após os "{}" da função
}

//exemplo de utilização de func anonima
retornoBoom()

//ou

func anonimato(){
	func() {
		fmt.Println(ValorString, valorInt)
	}()
}
```

## Recebendo N parametros para dentro de uma função
```
//recebendo N parametros dentro de uma função
//imagine que quero receber varias strings
//... -> o que os 3 pontinhos esta dizendo, go eu quero que na variavel "valoresString" aceite 1 ou N parametros de string (com isso posso mandar uma lista de string)
func muitos(valoresString ...string){
	for i, x := range valoresString{
		fmt.Println(i, x)
	}
}

//exemplo de utilização
muitos("Vitor")
muitos("")
muitos("Vitor", "a", "b", "c", "d")

//podemos enviar varias funções assim tambem
func muitosFun(variasFunc ...func()){
	for _, x := range variasFunc{
		x()
	}
}
//exemplo de utilização
func1 := func(){
	fmt.Println("test")
}
muitosFun(func1)


func2 := func(){
	fmt.Println("test")
}
func3 := func(){
	fmt.Println("test")
}
func4 := func(){
	fmt.Println("test")
}

muitosFun(func2, func3, func4)

```

## Criando uma Struct e instanciando ela
```
//Structs -> igual objeto(modelo) em outras linguagens
//Primeira letra do nome que voce der a sua struct for minuscula  = privada 
//Primeira letra do nome que voce der a sua struct for maiuscula  = Publica 
type User struct {
	name string
	age  int
}

//struct com campos privados
type UserDois struct {
	//usando a mesma metodologia de struc e funções
	//Primeira letra maiscula fica publico (visivel)
	//Primeira letra minuscula fica privado (não visivel na hora de instanciar)
	Email    string
	password string
}


//instanciando a struct
	var user User = User{
		//os parametros não precisam estar na ordem !!!
		name: "Vitor",
		age:  29,
	}
	fmt.Println(user)
```

## If else em Go
```
//no golang não temos "()" no if
	if 2 < 1 {
		fmt.Println("é verdade")
	} else if 2 == 1 {
		fmt.Println("é verdade novamente")
	} else {
		fmt.Println("é mentira")
	}
```

## If com init (muito utilizado)
```
//Criando a função
//repare que o retorno da função a gente nao poe antes de 'func' e sim depois dos '()'
func initIf1() string {
	return "oi"
}

//repare que podemos retornar um erro tambem
func initIf2() (string, error) {
	return "oi", errors.New("ocorreu um erro")
}

//Instanciando a funçao
//if com init (criamos uma variavel só para poder validar se ela é isso ou aquilo
	//exemplo se a função 'initIf()' retornar a palavra "oi" eu quero executar tal coisa
	//porém preciso instanciar uma variavel pra receber esse valor para ver se o retorno é realmente a palavra "oi"
	//abaixo estou atribuindo a função para a variavel retorno
	if retorno1 := initIf1(); retorno1 == "teste" {
		fmt.Println("se o retorno for igual a palavra OI ele vai printar isso aqui")
		fmt.Println(retorno1)
	}
	//lembre-se essa variavel "retorno1" é somente usada no escopo do if em que ela esta, com isso não podemos pegar ela fora desse if !!

	//retornando com o error
	//se o que tive de retorno for diferente do erro, ele executa a função
	//funcao muito usada !!
	if retorno2, err := initIf2(); err != nil {
		fmt.Println("se o retorno for igual a palavra OI ele vai printar isso aqui")
		fmt.Println(retorno2)
	}
```

## Switch
```
	//Switch
	//fallthrough -> seria tipo um "continue" (ele passa na condição que ele esta, e nas proximas mesma que nao for valida ele cai nelas tambem)
	//fallthrough -> ou seja independente da condição ele fica passando entre os cases e executando eles
	//break -> se não colocar o golang ja deixa por default, porém porque existe a palavra break?
	//break -> porque se eu por o fallthrough e quiser que ele pare em um case especifico eu coloco o break
	testeum := "test"

	switch testeum {
	case "teste", "novoParametro":
		fmt.Println("caiu na primeira condição")
		fallthrough
	case "teste2":
		fmt.Println("caiu na segunda condição")
		break
	default:
		fmt.Println("nao caiu em nenhum case")
	}
```

## For
```
//FOR -> existem varias formas de faze-lo no go
	//repare que não tem ()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
```

## While
```
//While -> no go ele não existe, porém podemos adaptar um for para isso
	variable := 0
	for variable <= 10 {
		fmt.Println("Valor", variable)
		variable++
	}
```

## Do while
```
//Do while -> no go ele não existe, porém podemos adaptar um for para isso
	anExpression := false
	//enquanto o variableY for igual a anExpression execute o que esta dentro da condição
	//porém como variableY é true e anExpression é false, o do while vai entrar pelo menos 1 vez na condição
	for variableY := true; variableY; variableY = anExpression {
		fmt.Println("Passou aqui")
	}
```

## For com range

```
//For com range
	//disponibiliza duas variaveis (o indice / o valor em si)
	variableZ := []string{"teste,teste,teste"}
	for i, value := range variableZ {
		fmt.Println(value, i)
	}

	//_ -> ignora o valor do indice
	for _, value := range variableZ {
		fmt.Println(value)
	}
```
