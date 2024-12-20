# Схема приложения

![image](https://github.com/user-attachments/assets/ae797881-aaf6-41c9-bd4f-14c87c7e6c29)

# Описание сервисов

## API Gateway
API Gateway служит единой точкой доступа для всех клиентских приложений. Он обрабатывает входящие запросы и перенаправляет их к соответствующим микросервисам. Также он отвечает за аутентификацию, авторизацию и маршрутизацию.

## Auth
Сервис аутентификации отвечает за регистрацию и вход пользователей. Он обеспечивает:
- Регистрацию пользователя по почте и паролю, а также через OAuth.
- Авторизацию пользователей с использованием учетных данных или OAuth.
- Хранение информации о пользователях и управление токенами доступа.
  
**Хранимые данные:**
- Учетные данные пользователей (хэшированные пароли).
- Токены доступа.

**Взаимодействие с сервисами:**
- API Gateway (при получении запросов на вход и регистрацию).
- Profile (для обновления информации о пользователе).

**Методы/API:**
- `POST /auth/register`: регистрация пользователя.
- `POST /auth/login`: вход пользователя.
- `GET /auth/logout`: выход пользователя.

## Chats
Сервис чатов отвечает за обмен сообщениями между пользователями. Его функции включают:
- Отправку и получение сообщений.
- Хранение истории сообщений.

**Хранимые данные:**
- История сообщений (отправитель, получатель, текст сообщения, временные метки).
- Чаты (идентификаторы участников).

**Взаимодействие с сервисами:**
- API Gateway (при получении запросов на отправку и получение сообщений).
- Notifications (для отправки уведомлений о новых сообщениях).

**Методы/API:**
- `POST /chats/send`: отправка сообщения.
- `GET /chats/:userId`: получение истории сообщений с пользователем.

## Profile
Сервис профилей управляет пользовательской информацией. Он предоставляет функционал для:
- Редактирования профиля (никнейм, информация о себе, аватарка).
- Получения информации о пользователе по его идентификатору.

**Хранимые данные:**
- Уникальные никнеймы, информация о себе, аватарки пользователей.

**Взаимодействие с сервисами:**
- API Gateway (при запросах на редактирование профиля).
- Auth (для проверки прав доступа).

**Методы/API:**
- `GET /profile/:userId`: получение информации о пользователе.
- `PUT /profile/:userId`: обновление информации о профиле.

## Notifications
Сервис уведомлений отвечает за отправку уведомлений пользователям о событиях, таких как новые сообщения, запросы на дружбу и другие действия.

**Хранимые данные:**
- Уведомления (идентификатор пользователя, текст уведомления, статус прочтения).

**Взаимодействие с сервисами:**
- API Gateway (для получения уведомлений).
- Auth и Chats (для получения событий о входе и новых сообщениях).

**Методы/API:**
- `GET /notifications/:userId`: получение списка уведомлений.
- `POST /notifications/send`: отправка уведомления.

## Friends
Сервис друзей управляет отношениями между пользователями. Он обеспечивает:
- Поиск пользователей по никнейму.
- Добавление и удаление пользователей из друзей.
- Подтверждение или отклонение запросов на дружбу.
- Просмотр списка друзей.

**Хранимые данные:**
- Список друзей пользователей (идентификаторы друзей, статусы дружбы).

**Взаимодействие с сервисами:**
- API Gateway (для получения и обновления списка друзей).
- Notifications (для отправки уведомлений о запросах на дружбу).

**Методы/API:**
- `POST /friends/add`: добавление пользователя в друзья.
- `POST /friends/remove`: удаление пользователя из друзей.
- `GET /friends/:userId`: получение списка друзей.
