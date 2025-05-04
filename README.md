# CEP Weather API

Este projeto é uma API escrita em Go que permite consultar a temperatura de uma localização brasileira com base em um CEP válido. A aplicação realiza a conversão da temperatura em Celsius para Fahrenheit e Kelvin, respondendo adequadamente para diferentes cenários de sucesso e erro.

## Funcionalidade

A API possui um único endpoint principal que aceita um CEP de 8 dígitos e retorna as temperaturas na localidade correspondente nos seguintes formatos:

- Celsius (`temp_C`)
- Fahrenheit (`temp_F`)
- Kelvin (`temp_K`)

A aplicação realiza duas integrações externas:

- **[ViaCEP](https://viacep.com.br/)**: para obter a localidade correspondente ao CEP.
- **[WeatherAPI](https://www.weatherapi.com/)**: para obter a temperatura em Celsius da localidade encontrada.

## Conversões realizadas

- **Fahrenheit:** `F = C * 1.8 + 32`
- **Kelvin:** `K = C + 273`

## Rotas disponíveis

- `GET /appstatus`  
  Retorna o status atual da aplicação.

- `GET /cep/:cep`  
  Recebe um CEP como parâmetro de URL e retorna a temperatura convertida nas três unidades.

- `GET /swagger/index.html`  
  Interface de documentação interativa da API gerada automaticamente com Swagger.


## Respostas da API /cep

- **200 OK** – Quando o CEP é válido e a temperatura foi localizada com sucesso.  
  Exemplo de retorno:  
  ```json
  {
    "temp_C": 28.5,
    "temp_F": 83.3,
    "temp_K": 301.5
  }
  ```

- **422 Unprocessable Entity** – Quando o CEP possui 8 dígitos, mas não é um CEP válido.  
  Mensagem: `invalid zipcode`

- **404 Not Found** – Quando o CEP não é encontrado no ViaCEP.  
  Mensagem: `can not find zipcode`

## Deploy

A aplicação está disponível publicamente no Google Cloud Run:

> https://golang-cloudrun-57107388905.us-central1.run.app/

## Como executar localmente

### Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- Uma conta gratuita na [WeatherAPI](https://www.weatherapi.com/) para obter sua chave de API.

### Passos para execução

1. **Clone este repositório:**

   ```bash
   git clone https://github.com/hydde7/goexpert-challenge-1
   cd seu-repo
   ```

2. **Adicione sua chave da WeatherAPI no `docker-compose.yml`:**

   No campo `FREEWEATHER_API_KEY`, substitua `COLOCAR API KEY AQUI` pela sua chave real.

3. **Execute a aplicação com Docker Compose:**

   ```bash
   docker-compose up
   ```

4. **Acesse a aplicação:**

   Ela estará disponível em:

   ```
   http://localhost:8080
   ```

   Você também pode acessar a documentação interativa Swagger em:

   ```
   http://localhost:8080/swagger/index.html
   ```

## Documentação Swagger

A documentação da API é gerada automaticamente com [Swaggo](https://github.com/swaggo/swag) e está acessível no endpoint `/swagger/index.html`.
