<h1 align="center">💫 Operacion Quasar Fire 🚀</h1>

## Requirements ⏬
- Go 1.16 or higher

## Run development environment 🛠️
To install this application for development you must meet the above requirements.

1. Run `go run cmd/server/main.go` (If you have docker you can use `docker-compose -f ./deployments/docker-compose.yml up --build`)
3. Go to [127.0.0.1:8000](http://127.0.0.1:8000/)

## Services 🌐
- **POST** -> `/topsecret/`
  - Payload Schema
    ```
    {
      "satellites": [
        {
          "name": "s1",
          "distance": 10.0,
          "message": ["hello", "", "!!!"]
       },
       {
          "name": "s2",
          "distance": 8.0,
          "message": ["hello", "world", ""]
       },
       {
          "name": "s3",
          "distance": 3.0,
          "message": ["", "", "!!!"]
       }
      ]
    }
    ```
- **POST** -> `/topsecret_split/{satellite_name}/`
  - Payload Schema
    ```
    {
      "distance": 80.0,
      "message": ["", "world", "everybody"]
    }
    ```
- **GET** -> `/topsecret_split/`