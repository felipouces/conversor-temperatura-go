package converter

func CelsiusParaFahrenheit(c float64) float64 {
    return (c * 9 / 5) + 32
}

func FahrenheitParaCelsius(f float64) float64 {
    return (f - 32) * 5 / 9
}
