# Golang doc (Itens que estudei e entendi)
## Aqui terá alguns exemplos do código go para ser usado em projetos futuros (de como construir um código em go)

> "A persistencia supera o talento"

[![Linkedin Badge](https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https:/https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)](https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)
[![Instagram Badge](https://img.shields.io/badge/-Instagram-a43b9d?style=flat-square&logo=Instagram&logoColor=white&link=https://https://www.instagram.com/vihhbz/?hl=pt-br/)](https://www.instagram.com/vihhstx/?hl=pt-br/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:vitorbrussolo@gmail.com)](mailto:vitorbrussolo@gmail.com)

# Generics
-  São tipos a serem especificados posteriormente 

### Generics antes da versão v1.18.1
```sh
O tipo '...interface{}' assumia essa função de 'T'

E se quisermos apenas aceitar dois tipos em um metodo como fariamos,
na versão v1.18.1?
func userTest(user interface{}) {
	switch v := user.(type){
	case UserTest1:
	case UserTest2:
	default:
	 panic("Invalid type!")
    }
}
```

### Generics depois da versão v1.18.1
```sh
func Sorc[T interface { Less(y T) bool}](list []T)
```
