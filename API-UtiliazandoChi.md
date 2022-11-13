# Golang doc (Itens que estudei e entendi)
## Aqui terá alguns exemplos do código go para ser usado em projetos futuros (de como construir um código em go)

> "A persistencia supera o talento"

[![Linkedin Badge](https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https:/https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)](https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)
[![Instagram Badge](https://img.shields.io/badge/-Instagram-a43b9d?style=flat-square&logo=Instagram&logoColor=white&link=https://https://www.instagram.com/vihhbz/?hl=pt-br/)](https://www.instagram.com/vihhstx/?hl=pt-br/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:vitorbrussolo@gmail.com)](mailto:vitorbrussolo@gmail.com)

### Utilizando Chi ()
```sh
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	//aqui eu ja posso espeficiar o tipo da rota mais tranquilo visualmente
	//o chi demora um pouco mais para fazer um 'reload'
	r.Get("/rota-que-eu-quero3", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Chi funcionou"))
	})
	http.ListenAndServe(":8080", r)
}

```

### nao se esqueça de no go.mod colocar a lib do Chi
```sh
module first_api_go_chi
go 1.18
require github.com/go-chi/chi/v5 v5.0.7


//pegamos a url do github
//para pegar a versão va em tags e coloque na mao la no go.mod
```

