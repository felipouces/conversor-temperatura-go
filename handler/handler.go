package handler

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "HELLO/converter"
)

func Run() {
    reader := bufio.NewReader(os.Stdin)
    for {
        showMenu()
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        option, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("❌ Inválido. Digite um número.")
            continue
        }

        switch option {
        case 1:
            handleCelsiusToFahrenheit(reader)
        case 2:
            handleFahrenheitToCelsius(reader)
        case 0:
            fmt.Println("👋 Tchau...")
            return
        default:
            fmt.Println("❌ Inválido. Tente novamente.")
        }
    }
}

func showMenu() {
    fmt.Println("\n🔥 Conversor de Temperatura 🔥")
    fmt.Println("1 - Celsius para Fahrenheit")
    fmt.Println("2 - Fahrenheit para Celsius")
    fmt.Println("0 - Sair")
    fmt.Print("Escolha uma opção: ")
}

func handleCelsiusToFahrenheit(reader *bufio.Reader) {
    fmt.Print("Digite a temperatura em Celsius: ")
    temp := getFloatInput(reader)
    result := converter.CelsiusParaFahrenheit(temp)
    fmt.Printf("🌡️ %.2f °C = %.2f °F\n", temp, result)
}

func handleFahrenheitToCelsius(reader *bufio.Reader) {
    fmt.Print("Digite a temperatura em Fahrenheit: ")
    temp := getFloatInput(reader)
    result := converter.FahrenheitParaCelsius(temp)
    fmt.Printf("🌡️ %.2f °F = %.2f °C\n", temp, result)
}

func getFloatInput(reader *bufio.Reader) float64 {
    for {
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        value, err := strconv.ParseFloat(input, 64)
        if err != nil {
            fmt.Print("❌ Entrada inválida. Digite um número válido: ")
            continue
        }
        return value
    }
}
