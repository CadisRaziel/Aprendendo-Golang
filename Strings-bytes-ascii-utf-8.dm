# Golang doc (Itens que estudei e entendi)
## Aqui terá alguns exemplos do código go para ser usado em projetos futuros (de como construir um código em go)

> "A persistencia supera o talento"

[![Linkedin Badge](https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https:/https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)](https://www.linkedin.com/in/vitor-brussolo-zerbato-474447176//)
[![Instagram Badge](https://img.shields.io/badge/-Instagram-a43b9d?style=flat-square&logo=Instagram&logoColor=white&link=https://https://www.instagram.com/vihhbz/?hl=pt-br/)](https://www.instagram.com/vihhstx/?hl=pt-br/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:vitorbrussolo@gmail.com)](mailto:vitorbrussolo@gmail.com)

# Criando sua propria variavel







# Valor e tipo da variavel 
```sh
%v %T -> format strings,
%v -> me da o valor da string,
%T -> me da o tipo da string'(pode ser usada no Printf)
#U, %#X -> format strings
```

# transforma em 'ascii
```sh
    Slice of bytes = []byte(x) -> conversão para slice de bytes
// cada caracter em uma string é um byte separado

	s := "ascii éâ"
	sb := []byte(s)
	fmt.Printf("%v\n%T", sb, sb)

	//*Range em uma string, ele da caracter por caracter (utf-8)
	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	//*se fazer o for in tradicional, ele da byte por byte (ascii)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
```


