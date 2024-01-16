# echo_htmx_templ

### A golang project using 

* [Templ](https://templ.guide/project-structure/project-structure) as the html component generator, which generates html(.templ) files in to go code.

* [Echo](https://echo.labstack.com/docs) as the web framework for routing, middleware and handling requests.

* Postgres to store data.

* [Squirrel](https://github.com/Masterminds/squirrel) to build safe sql queries.

* [Htmx-go](https://github.com/angelofallars/htmx-go) to assist with htmx headers.

Add your own .env file to your directory and paste in the line below, changing values to match your db.

```.env
DATABASE_URL=postgres://yourusername:yourpassword@localhost:5432/yourdbname?sslmode=disable
```

The simple db schema can be inputed with sql tools or with pgadmin.

```sql
CREATE TABLE IF NOT EXISTS signup (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    user_password VARCHAR(50) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

