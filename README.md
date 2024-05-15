# CCAS &mdash; Computer Club Accounting System

CCAS &mdash; это система учета компьютерного клуба, написанная на языке Go.


## Запуск на Linux

```
git clone git@github.com:l-golofastov/ccas.git
```
```
cd ccas
```

Программа обернута в Docker-контейнер, поэтому предварительно понадобится установить [Docker](https://docs.docker.com/get-docker/).

### Сборка

```
sudo docker build . -t ccas
```


### Запсук

Запустить программу можно на тестовом примере, который доступен сразу:
```
sudo docker run ccas data.txt
```

Также можно запустить программу на Ваших входных данных:
```
sudo docker run -v /home/user:/home/user ccas ~/test.txt
```
где ```~/test.txt``` &mdash; это полный путь до файла из домашней директории.


### Тесты

Запустить тесты можно из корня проекта командой:

```
go test
```
