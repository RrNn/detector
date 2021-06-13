## Must haves on your system

- [golang](https://golang.org/)
- [node](https://nodejs.org/en/)
- [postgres](https://www.postgresql.org/)

### API - golang

#### `Routes`

- GET `/`
- POST `/register`

```js
{
    "name": "Ricardo Banks",
    "email": "nabrrikk@yahoo.com",
    "address": "Kabale, Uganda",
    "contact": "+9887765987"
}
```

- POST `/auth/login`

```js
{
    "email": "nabrrikk@yahoo.com",
    "password": "lalala"
}
```

- POST `/auth/url`

```js
  {
  "link": "https://play.golang.org/"
  }
```

- GET `/auth/urls` - querystrings - `(limit int, offset int)`
- GET `/auth/urls/:id` - querystrings - `(withpings bool, limit int, offset int)`
- DELETE `/auth/urls/:id`

#### `Commands`

### UI - vuejs

- `npm install`
- `npm run serve`

To run the app, clone it, of course, and;
In the `/api` dircetory, create your `.env` file with the following environment variables

- **JWT_STRING**=`some_random_string`
- **DB_USER**=`your_database_user_name`
- **DB_PASSWORD**=`your_database_password`
- **DB_NAME**=`your_database_name`
- **DB_PORT**=`your_database_port` - usually 5432 for postgres, ie, the default
- **DB_SSL_MODE**=`disable_or_enable`
- **DB_HOST**=`your_database_host`

  After that;

- In one terminal window, `cd api && go run main.go`
- In another terminal window, `cd ui && npm install && npm run serve`
