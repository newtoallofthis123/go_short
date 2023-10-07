# Go Short

Simple URL shortener written in Go and HTMX.

## Usage

It is recommended to use the provided Dockerfile to run the application.
To use it, you must first fill in a `.env` file with the following variables:

```bash
DB=db_name
USER=postgres
PASSWORD=password
PORT=5432
HOST=database
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=db_name
```

I know the variables are repeated, I am working on it.

Then, you can run the application with the following command:

```bash
docker-compose up
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
