package main

import (
	"errors"
	"fmt"
)

//objeto de erro com os campos que eu quero para eu poder montar uma string a partir dele
type ErrorTest struct {
	//Code int -> referencia qual o código de erro
	Code    int
	Message string
}

//para eu satisfazer a interface Error, eu preciso implementar meu 'ErrorTest' com a interface Error
//isso seria basicamente um 'this'
func (e ErrorTest) Error() string {
	return fmt.Sprintf("Erro ao processar dados, code = %d, message = %s",
		e.Code,
		e.Message,
	)
}

func main() {
	// Errror-> uma interface

	// "_" -> estou ignorando o erro
	//res, _ := test()

	//sempre veremos dessa forma, atribuindo a função a uma variavel e fazendo um if para tratar o erro
	//se eu nao tratar dessa forma com o if, vai vim sempre um valor nulo dentro desse erro
	_, err := test()
	if err != nil {
		fmt.Println(err)
	}

	//posso fazer o metodo acima dessa forma tambem:
	// _, err := test() err != nil {
	// 	fmt.Println(err)
	// }

	//utilizando nossa struct de error
	if err := teste1(); err != nil {
		fmt.Println(err)
	}

	//utilizando nossa struct de error pegando somente os campos code e message sem a mensagem anterior a eles
	if err := teste2(); err != nil {
		fmt.Println(err.Code, err.Message)
	}

	//printando o objeto inteiro e o arquivo que esta
	if err := teste2(); err != nil {
		fmt.Printf("%#v", err)
	}
}

//Boa pratica do go sempre deixar o ultimo paramtro caso ouver multiplos o parametro de erro
func test() (string, error) {
	return "", errors.New("erro teste")
}

//utilizando nossa struct de erro
func teste1() error {
	err := ErrorTest{
		Code:    500,
		Message: "Deu erro aqui",
	}
	return err
}

func teste2() *ErrorTest {
	err := &ErrorTest{
		Code:    500,
		Message: "Deu erro aqui",
	}
	return err
}
func teste3() *ErrorTest {
	err := &ErrorTest{
		Code:    500,
		Message: "Deu erro aqui",
	}
	return err
}
