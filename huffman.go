package main


import (
	"log"
	"os"
	"io/ioutil"
	"strings"
)

func main(){

	if len(os.Args) > 2 || len(os.Args) <= 1 {
		log.Fatal("É necessário apenas um argumento( o caminho do arquivo ), Conforme o comando :'huffman [path]'")
	}

	path := os.Args[1]

	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Não foi encontrado o arquivo especificado no informado PATH : %s ,  Error: %v", path, err);
	}
	content := string(b)
	var diffContentChars string
	for _, r := range content {
		if diffContentChars!= ""{
			if !strings.ContainsAny(diffContentChars,string(r)){
				diffContentChars+=string(r)
			}
		}else{
			diffContentChars = string(r)
		}
	}
	frequencyTable := make(map[string]int)
	for _, r := range diffContentChars {
		frequencyTable[string(r)] = strings.Count(content, string(r))
	}

	for key, val := range frequencyTable{
		log.Printf("value de %s tem um total de : %d",key ,val)
	}
}
		
