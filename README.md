# Менеджер задач TODO

Пет проект на языке Go, для реализации использовал фреймворк gin

## Миграции

Для миграций использовал утилиту migrate

Перед самой миграцией необходимо запустить в докере базу данных, я использовал postgres
```
docker run --name to-do -e POSTGRES_PASSWORD=admin -e POSTGRES_USER=admin -p 5432:5432 -d postgres
```
Далее при помощи migrate сделаем инициализацию
```
migrate create -ext sql -dir ./schema -seq init
```
создадутся файлы 000001_init.down.sql и 000001_init.up.sql в которых необходимо прописать миграцию БД, при помощи следующей команды применить миграцию к созданной базе данных
```
migrate -path ./schema -database 'postgres://admin:admin@localhost:5432/postgres?sslmode=disable' up
```
Откат БД происходит делается следующим образом
```
migrate -path ./schema -database 'postgres://admin:admin@localhost:5432/postgres?sslmode=disable' down
```