---
name: Чат
sort: 1
---

# Чат

Это простая демонстрация двух способов, которыми можно реализовать Web IM приложение:

Используя long polling.
Используя WebSocket.

Оба способа, по умолчанию, сохраняют данные в памяти так, что каждый раз перезапуская приложение вы будете терять данные, но вы можете изменить настройки в файле `conf/app.conf`, чтобы включить использование базы данных.

Вот пример организации проекта:

```bash
WebIM/
    WebIM.go            # Главный пакет
    conf
        app.conf        # Конфигурационный файл
    controllers
        app.go          # Окно приветствия которое позволяет выбрать технологию и имя пользователя
        chatroom.go     # Функции для манипулирования данными
        longpolling.go  # Контроллер и методы для чата на long polling технологии
        websocket.go    # Контроллер и методы для чата на WebSocket технологии
    models
        archive.go      # Общии функции для работы с данным (модель). Общие для обоих технологий (WebSocket и long polling)
    views
        ...             # Файлы представления
    static
        ...             # JavaScript и CSS файлы
```

[Посмотреть код на GitHub](https://github.com/beego/samples/tree/master/WebIM)