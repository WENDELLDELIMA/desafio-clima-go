# Desafio Clima GO ☀️🌦️

Este projeto é um serviço em Go que recebe um **CEP** como entrada e retorna a **temperatura atual** da região correspondente, nos formatos **Celsius**, **Fahrenheit** e **Kelvin**.

---

## 📌 Objetivo

Criar uma API que:
- Recebe um CEP válido de 8 dígitos via HTTP.
- Consulta o endereço através da API [ViaCEP](https://viacep.com.br/).
- Consulta a temperatura da cidade usando a [WeatherAPI](https://www.weatherapi.com/).
- Retorna os dados em JSON com temperaturas convertidas para diferentes escalas.

---

## 📥 Requisição

### Endpoint:
```
GET /weather/{cep}
```

### Exemplos:
```
GET /weather/01001000
```

### Resposta de sucesso (HTTP 200):
```json
{
  "temp_C": 25.0,
  "temp_F": 77.0,
  "temp_K": 298.0
}
```

---

## 🧪 Tratamento de erros

- **CEP inválido (não 8 dígitos):**
  - Código: `422`
  - Mensagem: `invalid zipcode`

- **CEP não encontrado na API ViaCEP:**
  - Código: `404`
  - Mensagem: `can not find zipcode`

---

## 🧰 Como executar localmente

### Pré-requisitos:
- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/)

### Rodando com Docker Compose
```bash
docker-compose up --build
```

A aplicação estará acessível em: [http://localhost:8080/weather/01001000](http://localhost:8080/weather/01001000)

> ⚠️ Substitua `WEATHER_API_KEY` no docker-compose pela sua chave da [WeatherAPI](https://www.weatherapi.com/).

---

## 🚀 Deploy no Google Cloud Run

1. Compile o binário com Docker:
   ```bash
   docker build -t gcr.io/SEU_PROJETO/weather-service .
   ```
2. Publique com:
   ```bash
   gcloud run deploy --image gcr.io/SEU_PROJETO/weather-service --platform managed
   ```

---

## ✅ Testes

Rodar testes unitários:
```bash
go test ./...
```

---

## 📂 Estrutura de arquivos
```
├── main.go              # Inicialização do servidor
├── handler.go           # Handler do endpoint /weather/
├── viacep.go            # Consulta à API ViaCEP
├── weather.go           # Consulta à API WeatherAPI
├── models.go            # Structs de resposta
├── utils.go             # Validação de CEP
├── main_test.go         # Testes
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

## 👨‍💻 Desenvolvido por
[Wendell Lima](https://github.com/WENDELLDELIMA)