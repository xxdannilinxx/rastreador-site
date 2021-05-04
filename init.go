package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const QUANTITY = 3
const DELAY = 5

var LOGS []string

func main() {
	getIntroducao()
	/**
	*
	 */
	for {
		switch getOpcao() {
		case 0:
			sair()
		case 1:
			verLogs()
		case 2:
			initMonitoramento()
		default:
			fmt.Println("Opção não encontrada.")
			os.Exit(-1)
		}
	}
}

func getIntroducao() {
	var versao float32 = 1.0
	fmt.Println("Seja bem-vindo ao rastreador, versão", versao)
}

func getOpcao() int {
	var opcao int
	fmt.Println("[0] Sair do programa")
	fmt.Println("[1] Verificar arquivos de logs")
	fmt.Println("[2] Iniciar monitoramento")
	fmt.Scan(&opcao)
	fmt.Println("A opção escolhida foi", opcao, "em", &opcao)
	return opcao
}

func sair() {
	fmt.Println("Fechando...")
	os.Exit(0)
}

func initMonitoramento() {
	fmt.Println("Monitorando...")
	sites := lerSitesArquivo()
	for i := 0; i < QUANTITY; i++ {
		for _, site := range sites {
			monitorar(site)
		}
		time.Sleep(DELAY * time.Second)
	}
	fmt.Println("Monitoramento encerrado...")
}

func verLogs() {
	fmt.Println("Verificando logs...")
	fmt.Println(LOGS)
}

func monitorar(site string) {
	var newLog string
	res, _ := http.Get(site)
	switch res.StatusCode {
	case 200:
		newLog = "- O site" + site + "está OK."
	default:
		newLog = " - O site" + site + "está fora do ar."
	}
	LOGS = append(LOGS, newLog)
	fmt.Println(newLog)
}

func lerSitesArquivo() []string {
	var sites []string

	arquivo, error := os.Open("sites.ini")

	if error != nil {
		fmt.Println("Ocorreu um error: ", error)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, error := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if error == io.EOF {
			fmt.Println("Ocorreu um error: ", error)
			break
		}
	}

	return sites

}
