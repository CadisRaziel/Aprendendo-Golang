# Golang doc (Itens que estudei e entendi)
## Aqui terá alguns exemplos do código go para ser usado em projetos futuros (de como construir um código em go)

> "A persistencia supera o talento"

[![Linkedin Badge](https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https:/https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)](https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)
[![Instagram Badge](https://img.shields.io/badge/-Instagram-a43b9d?style=flat-square&logo=Instagram&logoColor=white&link=https://https://www.instagram.com/vihhbz/?hl=pt-br/)](https://www.instagram.com/vihhstx/?hl=pt-br/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:vitorbrussolo@gmail.com)](mailto:vitorbrussolo@gmail.com)

### Utilizando GorilaMux ()
```sh
package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	//toda vez que chamar a rota "/rota-que-eu-quero2" eu quero que caia dentro da função 'func'
	r.HandleFunc("/rota-que-eu-quero2", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("funcionando gorilaMux"))
		return
		//Methods("GET", "POST") -> aqui estou dizendo para que só haja requisições nessa rota se o metodo for get ou post
	}).Methods("GET", "POST")
	http.ListenAndServe(":8080", r)
}
```

### nao se esqueça de no go.mod colocar a lib do gorilamux
```sh
module fisrt_api_gorila_mux
go 1.18
require github.com/gorilla/mux v1.8.0

//pegamos a url do github
//para pegar a versão va em tags e coloque na mao la no go.mod
```

