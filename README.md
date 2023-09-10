# API de Conversão Monetária

API Rest que lida com conversão monetária. Suporta moedas reais e customizadas, criadas pelo usuário.

## Documentation
A documentação detalhada da API está disponível em:

- <a>https://lumacielz.github.io/currency-conversionAPI-docs/ </a>

A aplicação foi desenvolvida em Go, utilizando conceitos de Clean Architecture. Foram utilizadas algumas bibliotecas como o router <b>chi</b> e alguns utilitários como o <b>viper</b> para arquivos de config e o <b>logrus</b> para logs.

Os dados de cotação são consumidos da API pública de cotação de moedas <a href="https://docs.awesomeapi.com.br/api-de-moedas">Awesome API</a>
e salvos em uma collection no <b>mongoDB Atlas</b>. Segundo a
documentação da API, os dados são atualizados a cada 30 segundos, dessa forma, a cada query de busca por uma moeda,
a aplicação verifica se os dados de cotação estão atualizados, e os atualiza caso negativo.

Ao solicitar uma conversão de/para uma moeda que não exista no banco, se esta for contemplada pela Awesome API,
será inserida automaticamente de forma que <b>todas as moedas suportadas pela Awesome API são suportadas aqui</b>.

É possível ainda inserir moedas customizadas, para isso são feitas algumas validações como se o
código da moeda é único no banco, e se a taxa de conversão
é maior que zero de modo a evitar uma divisão por zero.

## Running

- <b>Makefile:</b> - necessita uma versão do Go instalada

```bash
$ make run
 ```

- <b>Docker</b>

```bash
$ docker build -t c-api .
$ docker run -p 8080:8080 c-api
 ```

Após executar um dos comandos, o servidor será iniciado no endereço <b>localhost:8080</b>.

## Testing
- <b>Makefile</b>:
```bash
$ make tests
 ```

## Test Coverage
- <b>Makefile</b>:
```bash
$ make coverage
 ```

