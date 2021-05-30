# Getir-Challenge

Application deployed on heroku [fufu-getir-challange.herokuapp.com](https://fufu-getir-challange.herokuapp.com/)
[Documentation](https://fufu-getir-challange.herokuapp.com)

## Run Application

Use docker to run the application

```bash
make build-docker
```

## Running Tests

Use docker to run tests (unit and integration)

```bash
make run-tests
```

## Endpoints

| Method        | Url                     | Parameters |Description | 
| ------------- | ------------------------| -----------|-------------
| POST          | /records                | `min_count`, `max_count`, `start_date`, `end_date`           |List mongo db records    |
| POST          | /in-memory              | `key`, `value`           |Create in-memory record  |
| GET           | /in-memory?key=         | `key`                    |Retrieve in-memory record|

