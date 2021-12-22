# Fibonacci Rest and gRPS server & client






### REST API
Чтобы получить все числа последовательности Фибоначчи с порядковыми номерами от first до last, необходимо отправить GET запрос вида:
```azure
http://localhost:8090/fibonacci? x=last y=first
```

### gRPC
Proto фаил(см. [fibonacci.proto](pkg/api/fibonacci.proto)) находится по пути:
```
./pkg/api/fibonacci.proto
```

Чтобы получить все значения от first до last нужно скомпилировать server(cmd make), а также клиент (cmd make mclient) для отправки значения от first(x) и last(y) 
