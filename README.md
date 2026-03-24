## Быстрый запуск (Docker)

Вся инфраструктура (База данных, миграции, Бэкенд и Фронтенд) поднимается одной командой:

```bash
docker-compose -f docker-compose.run.yml up -d --build
```

**После запуска:**

- **Frontend:** [http://localhost](http://localhost) (порт 80)
- **Backend API:** [http://localhost:9000/api/v1](http://localhost:9000/api/v1)
- **Swagger UI:** [http://localhost:9000/api/v1/swagger/index.html](http://localhost:9000/api/v1/swagger/index.html)

---

## Технологический стек

### Backend (Go)

- **Язык:** Go 1.24+
- **Framework:** [Chi v5](https://github.com/go-chi/chi) (Router)
- **ORM:** [GORM](https://gorm.io/) с использованием [GORM Gen](https://gorm.io/gen/) (Type-safe запросы)
- **База данных:** PostgreSQL 15
- **Миграции:** [golang-migrate](https://github.com/golang-migrate/migrate)
- **Логирование:** [Uber-Go Zap](https://github.com/uber-go/zap) (структурированные логи)
- **Документация:** Swagger (swaggo)

### Frontend (React & TypeScript)

- **UI Kit:** [Ant Design 5](https://ant.design/)
- **Сборка:** Vite
- **HTTP Клиент:** Axios

---

## 📡 Примеры запросов к API

### 1. Получение списка устройств с фильтрацией

**Запрос:**

```bash
curl -X GET "http://localhost:9000/api/v1/devices?hostname=srv&is_active=true&page=1&limit=10"
```

**Описание:** Поиск активных устройств, у которых в hostname есть подстрока "srv".

### 2. Создание нового устройства

**Запрос:**

```bash
curl -X POST http://localhost:9000/api/v1/devices \
     -H "Content-Type: application/json" \
     -d '{
           "hostname": "new-server-01",
           "ip": "192.168.10.5",
           "platform_code": "LINUX",
           "is_active": true
         }'
```

### 3. Обновление устройства

**Запрос:**

```bash
curl -X PUT http://localhost:9000/api/v1/devices/{UUID} \
     -H "Content-Type: application/json" \
     -d '{
           "hostname": "updated-hostname",
           "is_active": false
         }'
```

### 4. Просмотр журнала аудита

**Запрос:**

```bash
curl -X GET "http://localhost:9000/api/v1/audit?page=1&limit=5"
```

**Описание:** Возвращает последние 5 действий в системе (кто, когда и что изменил).

---
