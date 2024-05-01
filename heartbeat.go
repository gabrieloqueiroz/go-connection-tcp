package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Endereço e porta do servidor
	// serverAddr := "vpce-064dd5b1910b0e601-eyhocq2e.vpce-svc-0b36720ece5381654.us-east-1.vpce.amazonaws.com:21000"
	serverAddr := "10.101.5.186:21000"

	// Estabelece conexão TCP
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Erro ao conectar:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Conexão estabelecida com", serverAddr)

	// Loop para receber dados continuamente
	for {
		// Buffer para armazenar a mensagem recebida
		data := make([]byte, 1024)

		// Lê a mensagem recebida
		time := time.Now().Add(180 * time.Second)
		conn.SetReadDeadline(time)
		_, err := conn.Read(data)
		if err != nil {
			fmt.Println("Erro ao ler mensagem:", err)
			break
		}

		// Exibe a mensagem recebida
		fmt.Println("Mensagem recebida:", string(data))

		// Verifica se a mensagem indica que a conexão deve ser encerrada
		if string(data) == "FIM" {
			fmt.Println("Recebido comando para encerrar a conexão.")
			break
		}
	}
}

// func main() {
// 	// Endereço e porta que o servidor irá escutar
// 	host := "vpce-064dd5b1910b0e601-eyhocq2e.vpce-svc-0b36720ece5381654.us-east-1.vpce.amazonaws.com"
// 	port := "21000"

// 	// Cria o endereço do servidor
// 	serverAddr := host + ":" + port

// 	// Abre o socket TCP
// 	listener, err := net.Listen("tcp", serverAddr)
// 	if err != nil {
// 		fmt.Println("Erro ao abrir o socket:", err)
// 		os.Exit(1)
// 	}
// 	defer listener.Close()

// 	fmt.Println("Servidor escutando em", serverAddr)

// 	// Loop infinito para aceitar conexões e escutar mensagens
// 	for {
// 		// Aceita uma conexão de cliente
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Erro ao aceitar conexão:", err)
// 			continue
// 		}

// 		// Trata a conexão em uma goroutine
// 		go handleConnection(conn)
// 	}
// }

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()

// 	fmt.Println("Nova conexão estabelecida:", conn.RemoteAddr())

// 	// Buffer para armazenar a mensagem recebida
// 	message := make([]byte, 1024)

// 	// Loop para ler mensagens recebidas
// 	for {
// 		// Lê a mensagem recebida
// 		n, err := conn.Read(message)
// 		if err != nil {
// 			fmt.Println("Erro ao ler mensagem:", err)
// 			break
// 		}

// 		// Exibe a mensagem recebida
// 		fmt.Printf("Mensagem recebida de %s: %s\n", conn.RemoteAddr(), string(message[:n]))
// 	}
// }

// func main() {
// 	// Endereço e porta que o servidor irá escutar
// 	host := "vpce-064dd5b1910b0e601-eyhocq2e.vpce-svc-0b36720ece5381654.us-east-1.vpce.amazonaws.com"
// 	port := "21000"

// 	// Cria o endereço do servidor
// 	serverAddr := host + ":" + port

// 	// Abre o socket TCP
// 	listener, err := net.Listen("tcp", serverAddr)
// 	if err != nil {
// 		fmt.Println("Erro ao abrir o socket:", err)
// 		os.Exit(1)
// 	}
// 	defer listener.Close()

// 	fmt.Println("Servidor escutando em", serverAddr)

// 	// Loop infinito para aceitar conexões e escutar mensagens
// 	for {
// 		// Aceita uma conexão de cliente
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Erro ao aceitar conexão:", err)
// 			continue
// 		}

// 		// Trata a conexão em uma goroutine
// 		go handleConnection(conn)
// 	}
// }

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()

// 	fmt.Println("Nova conexão estabelecida:", conn.RemoteAddr())

// 	// Buffer para armazenar a mensagem recebida
// 	message := make([]byte, 1024)

// 	// Loop para ler mensagens recebidas
// 	for {
// 		// Lê a mensagem recebida
// 		n, err := conn.Read(message)
// 		if err != nil {
// 			fmt.Println("Erro ao ler mensagem:", err)
// 			break
// 		}

// 		// Exibe a mensagem recebida
// 		fmt.Printf("Mensagem recebida de %s: %s\n", conn.RemoteAddr(), string(message[:n]))
// 	}
// }
