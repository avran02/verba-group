# Verba group test task

### Запуск

```
# Скачиваем
git clone https://github.com/avran02/verba-group

# Заходим в папку проекта
cd verba-group

# Создаём .env из шаблона
cp example.env .env
```

___Запуск в docker (рекомендуется)___
```
docker-compose up -d --build
```

___Запуск локально___
- Изменить конфиги в .env файле
- Выполнить миграции
- Запустить: `go run main.go` 

### Задание

__Проект__: Разработка REST API для управления задачами (To-Do List)

__Роль__: Junior Backend Developer (Golang)

__Цель__:
Разработать REST API для системы управления задачами, которая позволяет пользователям создавать, просматривать, обновлять и удалять задачи.

__Требования к функционалу__:

_Создание задачи_

_Метод:_ __POST__ /tasks
Описание: Создать новую задачу.
Запрос:
 Заголовки:
  Content-Type: application/json
 Тело:
 ```
  {
    "title": "string",
    "description": "string",
    "due_date": "string (RFC3339 format)"
  }
  ```
Ответ:
 Успех (201 Created):
 ```
  {
    "id": "int",
    "title": "string",
    "description": "string",
    "due_date": "string (RFC3339 format)",
    "created_at": "string (RFC3339 format)",
    "updated_at": "string (RFC3339 format)"
  }
  ```
Ошибка (400 Bad Request): Неправильный формат данных.
Ошибка (500 Internal Server Error): Проблема на сервере.


_Просмотр списка задач_

_Метод:_ __GET__ /tasks
Описание: Получить список всех задач.
Запрос:
 Заголовки:
 Content-Type: application/json
 Ответ:
 Успех (200 OK):
 ```
 [
   {
     "id": "int",
     "title": "string",
     "description": "string",
    "due_date": "string (RFC3339 format)",
     "created_at": "string (RFC3339 format)",
     "updated_at": "string (RFC3339 format)"
   }
 ]
 ```
Ошибка (500 Internal Server Error): Проблема на сервере.

_Просмотр задачи_

_Метод:_ __GET__ /tasks/{id}
Описание: Получить задачу по ID.
Запрос:
 Параметры пути:
  id: ID задачи (int)
 Заголовки:
  Content-Type: application/json
 Ответ:
 Успех (200 OK):
 ```
 {
   "id": "int",
   "title": "string",
   "description": "string",
   "due_date": "string (RFC3339 format)",
   "created_at": "string (RFC3339 format)",
   "updated_at": "string (RFC3339 format)"
 }
 ```
Ошибка (404 Not Found): Задача не найдена.
Ошибка (500 Internal Server Error): Проблема на сервере.


_Обновление задачи_

_Метод:_ __PUT__ /tasks/{id}
Описание: Обновить задачу по ID.
Запрос:
 Параметры пути:
  id: ID задачи (int)
 Заголовки:
  Content-Type: application/json
 Тело:
 ```
 {
   "title": "string",
   "description": "string",
   "due_date": "string (RFC3339 format)"
 }
 ```
Ответ:
Успех (200 OK):
```
{
  "id": "int",
  "title": "string",
  "description": "string",
  "due_date": "string (RFC3339 format)",
  "created_at": "string (RFC3339 format)",
  "updated_at": "string (RFC3339 format)"
}
```
Ошибка (400 Bad Request): Неправильный формат данных.
Ошибка (404 Not Found): Задача не найдена.
Ошибка (500 Internal Server Error): Проблема на сервере.
_
Удаление задачи_

_Метод:_ __DELETE__ /tasks/{id}
Описание: Удалить задачу по ID.
Запрос:
 Параметры пути:
  id: ID задачи (int)
 Заголовки:
  Content-Type: application/json
Ответ:
Успех (204 No Content): Задача удалена.
Ошибка (404 Not Found): Задача не найдена.
Ошибка (500 Internal Server Error): Проблема на сервере.

Дополнительные требования:
Вся информация должна сохраняться в базу данных (используйте Postgres).
Использовать Go modules для управления зависимостями.

Критерии приемки:
Корректная работа всех конечных точек API.
Код соответствует стандартам Go (gofmt, golint).
Проект легко запускается и настраивается с помощью команды go run.
Ожидаемый результат:
Исходный код проекта, расположенный в репозитории на GitHub или любой другой системе контроля версий, с подробным описанием в README.md о том, как запустить проект.