package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    converterTemperatura()
}

func converterTemperatura() {
    reader := bufio.NewReader(os.Stdin)

    for {
        Menu()
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        option, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("❌ Inválido. Digite um número.")
            continue
        }

        switch option {
        case 1:
            celsiusFahrenheit(reader)
        case 2:
            fahrenheitCelsius(reader)
        case 0:
            fmt.Println("👋 Tchau...")
            return
        default:
            fmt.Println("❌ Inválido. Tente novamente.")
        }
    }
}

func Menu() {
    fmt.Println("\n🔥 Conversor de Temperatura 🔥")
    fmt.Println("1 - Celsius para Fahrenheit")
    fmt.Println("2 - Fahrenheit para Celsius")
    fmt.Println("0 - Sair")
    fmt.Print("Escolha uma opção: ")
}

// função
func celsiusFahrenheit(reader *bufio.Reader) {
    fmt.Print("Digite a temperatura em Celsius: ")
    temp := floatInput(reader)
    fahrenheit := converterCelsiusParaFahrenheit(temp)
    fmt.Printf("🌡️ %.2f °C = %.2f °F\n", temp, fahrenheit)
}

// função
func fahrenheitCelsius(reader *bufio.Reader) {
    fmt.Print("Digite a temperatura em Fahrenheit: ")
    temp := floatInput(reader)
    celsius := converterFahrenheitParaCelsius(temp)
    fmt.Printf("🌡️ %.2f °F = %.2f °C\n", temp, celsius)
}

// função
func floatInput(reader *bufio.Reader) float64 {
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

func converterCelsiusParaFahrenheit(c float64) float64 {
    return (c * 9 / 5) + 32
}

func converterFahrenheitParaCelsius(f float64) float64 {
    return (f - 32) * 5 / 9
}