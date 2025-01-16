# Go WebSocket Chat

Это проект для чата с использованием **WebSocket** на языке программирования **Go**. Он реализует чат-приложение с поддержкой многопользовательского взаимодействия в реальном времени, используя **Go**, **Gorilla WebSocket** и архитектуру **Чистой Архитектуры**.

## Структура проекта

Проект организован по принципам **Чистой Архитектуры**, что позволяет отделить бизнес-логику от деталей реализации (таких как взаимодействие с WebSocket или HTTP). Он разделен на несколько слоев:

1. **`cmd/`** — точка входа в приложение (например, запуск сервера).
2. **`internal/`** — основная логика приложения:
    - **`domain/`** — содержит доменные сущности (например, клиент, сообщения).
    - **`interactor/`** — слой приложения, который обрабатывает бизнес-логику.
    - **`interfaces/`** — внешние интерфейсы, такие как обработка WebSocket и HTTP.
3. **`go.mod`** — файл модулей Go.

## Особенности

- **WebSocket-соединения** для реального времени.
- **Чистая Архитектура** для четкого разделения логики и интерфейсов.
- Простая логика для добавления и удаления клиентов.
- Отправка сообщений всем подключенным клиентам.