package converter

import "testing"

func TestCelsiusParaFahrenheit(t *testing.T) {
    tests := []struct {
        input    float64
        expected float64
    }{
        {0, 32},
        {100, 212},
        {-40, -40},
    }

    for _, test := range tests {
        result := CelsiusParaFahrenheit(test.input)
        if result != test.expected {
            t.Errorf("CelsiusParaFahrenheit(%.2f) = %.2f; esperado %.2f", test.input, result, test.expected)
        }
    }
}

func TestFahrenheitParaCelsius(t *testing.T) {
    tests := []struct {
        input    float64
        expected float64
    }{
        {32, 0},
        {212, 100},
        {-40, -40},
    }

    for _, test := range tests {
        result := FahrenheitParaCelsius(test.input)
        if result != test.expected {
            t.Errorf("FahrenheitParaCelsius(%.2f) = %.2f; esperado %.2f", test.input, result, test.expected)
        }
    }
}
