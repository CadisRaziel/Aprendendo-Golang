# Golang doc (Itens que estudei e entendi)
## Aqui terá alguns exemplos do código go para ser usado em projetos futuros (de como construir um código em go)

> "A persistencia supera o talento"

[![Linkedin Badge](https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https:/https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)](https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)
[![Instagram Badge](https://img.shields.io/badge/-Instagram-a43b9d?style=flat-square&logo=Instagram&logoColor=white&link=https://https://www.instagram.com/vihhbz/?hl=pt-br/)](https://www.instagram.com/vihhstx/?hl=pt-br/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:vitorbrussolo@gmail.com)](mailto:vitorbrussolo@gmail.com)

### Utilizando net/http, o http default do golang ()
```sh
package main

import (
	"log"
	"net/http"
)

func main() {

	//Obs: problemas do pacote http (default do golang)
	//não conseguimos especificar se é um get/delet/post/put
	//para saber teriamos que por um switch
	//switch req.Method {
	//case "POST":
	//case "PUT":
	//case "GET":
	//case "DELETE":
	//default:
	//	fmt.Println("Rota inválida")
	//}
	//e se eu nao especificar a rota, independente do que o usuario chamar ele vai executar

	httpObj := NewHttpInterface()

	//3º -> aqui eu digo pro go que toda requisição(conexões) no /rota-que-eu-quero e toda vez que cair nesse cara, quero que redirecione para o metodo 'httpObj'
	http.Handle("/rota-que-eu-quero", httpObj)

	//4º -> uma maneira diferente de fazer o metodo acima e sem precisar do metodo server Http la em baixo
	//http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusOK)
	//	w.Write([]byte("123"))
	//	fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
	//})

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

type HttpStructImplementation struct {
}

//2º -> implementando o http.handler (e com essa função conseguimos fazer toda chamada na main
func NewHttpInterface() http.Handler {
	return &HttpStructImplementation{}
}

//1º server http -> nosso handler (vai tratar as chamadas do endpoint, exemplo /get /user etc.. Quem redireciona e trata esses endpoint é o handler)
func (h *HttpStructImplementation) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("server 8080 funcionando"))
	return
}


```


