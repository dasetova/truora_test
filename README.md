# Prueba practica Truora
El presente repositorio se creó para dar solución a la prueba práctica para Truora.

Restful API en Go.

## Environment Variables
To run this RestAPI, the followed environment variables are required:
DATABASE_TYPE = "postgres"
DATABASE_URL = "host=localhost port=26257 user=truora_test dbname=recipes_db sslmode=disable"
API_PORT = "6767"

## Executing the API
Database must be created but the tables are created by migrations, just execute de API with ```go run main.go``` or build and exec.
