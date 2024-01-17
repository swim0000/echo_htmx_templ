# echo_htmx_templ

### A golang project using 

* [Templ](https://templ.guide/) as the html template language, which generates go code from .templ(html with go code) files.

* [Echo](https://echo.labstack.com/docs) as the web framework for routing, middleware and handling requests.

* Postgres to store data.

* [Htmx-go](https://github.com/angelofallars/htmx-go) to assist with htmx headers.

### Prerequisites

* go installs
```
go install github.com/a-h/templ/cmd/templ@latest 
```

* go gets
```
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
go get github.com/angelofallars/htmx-go
go get github.com/a-h/templ
go get github.com/lib/pq
go get github.com/joho/godotenv
```

* Add your own .env file to your directory and paste in the line below, changing values to match your db.

```.env
DATABASE_URL=postgres://yourusername:yourpassword@localhost:5432/yourdbname?sslmode=disable
```

* Db schema needs to be inputed with sql tools or with pgadmin.

```sql
CREATE TABLE IF NOT EXISTS signup (
    id SERIAL PRIMARY KEY,
    user_email VARCHAR(255) NOT NULL,
    user_password VARCHAR(50) NOT NULL,
);
```

