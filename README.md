## Оглавление
- [Требования](#требования)
- [Установка](#установка)
- [Структура проекта](#структура-проекта)
- [Структура базы данных](#структура-базы-данных)
- [Запуск проекта](#запуск-проекта)

## Требования
- Go версии 1.16 и выше.
- Установленный MySQL.
- Установленные зависимости.

## Установка

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/ZakharSol/Currency-Service.git
   cd currency_service

2. Установите зависимости:

   ```bash
   go mod tidy

## Структура проекта

```
currency_service/
├── main.go               # Главный файл приложения
├── db/                   # Файлы работы с БД
│   └── db.go             # Инициализация и подключение к БД
├── api/                  # Файлы для работы с HTTP API
│   ├── handlers.go       # Обработчики запросов
│   └── router.go         # Настройка маршрутизации
├── models/               # Структуры данных и модели
│   └── RatesModel.go     # Модель данных валют
└── utils/                # Утилиты и вспомогательные функции
    └── fetch.go          # Логика для извлечения данных валют
```

## Структура базы данных
### База данных currencyDB
### Таблица rates
- cur_id (INT, PRIMARY KEY)
- abbreviation (VARCHAR(10))
- scale (INT)
- official_rate (DECIMAL(10, 4))
- date (DATE)

## Запуск проекта

1. Проверьте все представленные ранее пункты.
2. Запустите проект:

   ```bash
   go run main.go

