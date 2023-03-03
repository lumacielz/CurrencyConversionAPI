# API de Conversão Monetária

API Rest que lida com conversão monetária. Suporta moedas reais e customizadas, criadas pelo usuário.

## Documentation
A documentação completa da API está no arquivo <b>openApi.yaml</b>, que pode ser aberto no editor <a>editor.swagger.io</a>.

A aplicação foi desenvolvida em Go, utilizando conceitos de Clean Architecture. Foram utilizadas algumas bibliotecas como o router <b>chi</b> e alguns utilitários como o <b>viper</b> para arquivos de config e o <b>logrus</b> para logs.

Os dados de cotação são consumidos da API pública awesome api e salvos em uma collection no mongoDB Atlas. Segundo a
documentação da API, os dados são atualizados a cada 30 segundos, dessa forma, a cada query de busca por uma moeda,
a aplicação verifica se os dados de cotação estão atualizados, e os atualiza caso negativo.

Ao solicitar uma conversão de/para uma moeda que não exista no banco, se esta for contemplada pela API,
será inserida automaticamente de forma que <b>todas as moedas suportadas pela awesome API são suportadas aqui</b>.

É possível ainda inserir moedas customizadas, para isso são feitas algumas validações como se o
código da moeda é único no banco, e se a taxa de conversão
é maior que zero de modo a evitar uma divisão por zero.

## Running

- <b>Makefile:</b> - necessita uma versão do Go instalada
o
```bash
$ make run
 ```

- <b>Docker</b>

```bash
$ docker build -t challenge-bravo .
$ docker run -p 8080:8080 challenge-bravo
 ```

O servidor será iniciado no endereço <b>localhost:8080</b>.

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

https://editor.swagger.io/?_ga=2.212483857.425198988.1677796467-1675336891.1677796467
