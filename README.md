# Moxer [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Moxer** is a lightweight mock data generator that lets you define JSON templates with placeholders (e.g., `{{email}}`, `{{uuid}}`) and generates realistic fake data for development and testing.

Great for:

- :fast_forward: Seeding test databases
- :fast_forward: Generating mock API responses
- :fast_forward: Creating JSON datasets for frontend development

## :sparkles: Features

- :bust_in_silhouette: Register and login in as a user
- :page_facing_up: Create and save custom JSON templates with variables like `{{email}}`,   `{{uuid}}`, etc.
- :books: Generate mock data from saved templates
- :open_file_folder: View your saved templates

## :hammer_and_wrench: Development Highlights

- :closed_lock_with_key: JWT-based authentication middleware
- :bricks: Clean 3-layer architecture: `handler → service → repository`
- :package: DTOs and mappers for clean data transformation
- :card_file_box: PostgreSQL with [GORM](https://gorm.io/index.html)
- :test_tube: Unit tests using [Testify](https://github.com/stretchr/testify)
- :books: API documentation powered by [swaggo/swag](https://github.com/swaggo/swag)
- :game_die: Mock data generation via [Gofakeit](https://github.com/brianvoe/gofakeit)

## :arrow_heading_down: Example Usage

Given the following input template:

```json
{
    "id": "{{uuid}}",
    "name": "{{name}}",
    "auth_data": {
        "email": "{{email}}",
        "password": "pass_{{password}}"
    }
}
```

Moxer will output:
```json
{
    "id": "ed725ce4-40cd-44f7-8b0d-ff0ba0265a1c",
    "name": "Estelle Strosin",
    "auth_data": {
        "email": "marisoltromp@mclaughlin.io",
        "password": "pass_@P7Laih1"
    }
}
```

## :closed_lock_with_key: API endpoints ```(api/v1)```

| Method | Path                      | Description                     | Auth Required |
|:------:|:-------------------------:|:-------------------------------:|:------------------:|
| GET    | `/swagger/index.html`     | Swagger UI for the API          | :x:                |
| GET    | `/healthz`                | Check availability of the app   | :x:                |
| POST   | `/auth/register`          | Register a new user             | :x:                |
| POST   | `/auth/login`             | Authenticate user and get token | :x:                |
| GET    | `/auth/me`                | Get current user info           | :white_check_mark: |
| POST   | `/templates`              | Create a new template           | :white_check_mark: |
| GET    | `/templates`              | Get current user's templates    | :white_check_mark: |
| POST   | `/templates/:id/generate` | Generate mocks for a template   | :white_check_mark: |

## :package: Dependencies
- [Echo](https://echo.labstack.com/)
- [GORM](https://gorm.io/index.html)
- [Gofakeit](https://github.com/brianvoe/gofakeit)
- [jwt-go](https://golang-jwt.github.io/jwt/)
- [Google UUID](https://github.com/google/uuid)
- [Testify](https://github.com/stretchr/testify)
- [swaggo/swag](https://github.com/swaggo/swag)
- [crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

## :page_facing_up: License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) for details.