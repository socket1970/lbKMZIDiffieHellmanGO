# Лабораторная работа по предмету КМЗИ. Шифр Диффи — Хеллмана(протокол Диффи — Хеллмана)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
[![Telegram](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/funny_m4n)
***
## Описание
Алгоритм реализован на Go`lang.

Схема шифра:
```
    1. Alice:   Za = (Yb)^Xa mod p;
    2. Bob:     Zb = (Ya)^Xb mod p;   
при этом, 
    Za == Zb
```

В качестве дополнения к этой лабораторной работе можно посмотреть [декодирование из числового формата в символьный.](https://github.com/socket1970/KMZIcryptogramLB)
***
### Задание 1.
Заданы параметры шифра:
* сообщение;
* q - простое число;
* p - (q*2) + 1;
* g - параметр протокола Диффи — Хеллмана;
* Xa - ключ Алисы;
* Xb - ключ Боба;

Производится расчет недостающих ключей и переписки. Переписка выводится на экран.

### Задание 2.
Заданы параметры шифра:
* q - простое число;

Производится расчет недостающих ключей и переписки. Переписка выводится на экран.

***
### Компиляция.
```go
go build .\main.go
```
***