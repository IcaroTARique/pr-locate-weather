# Locate Weather

Locate weather é uma aplicação que tem como objetivo retornar a temperatura de uma
determinada localidade a partir do CEP informado.


## Cloud URL
A URL da cloud é a seguinte:
https://pr-locate-weather-l5ol6kq5cq-uc.a.run.app/temperature/{CEP}
Aqui vemos que para solicitar a temperatura de um determinado local, utilizamos o
path `/temperature/{CEP}` onde `{CEP}` é o CEP da localidade que desejamos saber a
temperatura.

**_Atente-se para a utilização do CEP sem hífen._**

<h3 style="color:green"><b>Exemplo 200:</b></h3>
- https://pr-locate-weather-l5ol6kq5cq-uc.a.run.app/temperature/58046700
- https://pr-locate-weather-l5ol6kq5cq-uc.a.run.app/temperature/58704090

<h3 style="color:orange"><b>Exemplo 404:</b></h3>
- https://pr-locate-weather-l5ol6kq5cq-uc.a.run.app/temperature/58046701
- https://pr-locate-weather-l5ol6kq5cq-uc.a.run.app/temperature/58704091

<h3 style="color:red"><b>Exemplo 422:</b></h3>
- https://pr-locate-weather-l5ol6kq5cq-uc.a.run.app/temperature/5804670
- https://pr-locate-weather-l5ol6kq5cq-uc.a.run.app/temperature/8704091