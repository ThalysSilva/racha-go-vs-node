package main

import (
    "fmt"
    "time"
)

func lerPedidos(pedidos chan<- int) {
    for i := 1; i <= 5; i++ {
        pedidos <- i
        time.Sleep(200 * time.Millisecond)
    }
    close(pedidos)
}

func calcularPrecos(pedidos <-chan int, precos chan<- float64) {
    for pedido := range pedidos {
        preco := float64(pedido) * 10.5
        precos <- preco
    }
    close(precos)
}

func salvarNoBanco(precos <-chan float64, resultado chan<- string) {
    for preco := range precos {
        time.Sleep(300 * time.Millisecond)
        resultado <- fmt.Sprintf("Salvo pedido com preço %.2f", preco)
    }
    close(resultado)
}

func main() {
    pedidos := make(chan int)
    precos := make(chan float64)
    resultado := make(chan string)

    go lerPedidos(pedidos)
    go calcularPrecos(pedidos, precos)
    go salvarNoBanco(precos, resultado)

    for msg := range resultado {
        fmt.Println(msg)
    }
}
