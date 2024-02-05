# Курс валют сервис

## Установка и запуск

1. Склонируйте репозиторий: `git clone https://github.com/yourusername/yourproject.git  `
2. Перейдите в директорию проекта: `cd yourproject`
3. Запустите приложение с помощью Docker Compose: `docker-compose up -d`

...

## API Endpoints

### GET /rates

Получить курсы всех валют.

### GET /rates/{cryptocurrency}

Получить курс конкретной криптовалюты.

...

## Telegram Бот

### /start

Начать взаимодействие с ботом.

### /rates

Получить текущие курсы валют.

### /rates {cryptocurrency}

Получить курс конкретной криптовалюты.

### /start-auto {minutes_count}

Включить автоматическую отправку данных каждые {minutes_count} минут.

### /stop-auto

Выключить автоматическую отправку данных.

...




