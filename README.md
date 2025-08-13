# Moxer [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Moxer** is a lightweight mock data generator that lets you define JSON templates with placeholders (e.g., `{{email}}`, `{{uuid}}`) and generates realistic fake data for development and testing.

Great for:

- :fast_forward: Seeding test databases
- :fast_forward: Generating mock API responses
- :fast_forward: Creating JSON datasets for frontend development

## :sparkles: Features

- :bust_in_silhouette: Register and login in as a user
- :page_facing_up: Create and save custom JSON templates with variables like `{{email}}`, `{{uuid}}`, etc.
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

## :closed_lock_with_key: API Endpoints `(api/v1)`

| Method |           Path            |           Description           |   Auth Required    |
| :----: | :-----------------------: | :-----------------------------: | :----------------: |
|  GET   |   `/swagger/index.html`   |     Swagger UI for the API      |        :x:         |
|  GET   |        `/healthz`         |  Check availability of the app  |        :x:         |
|  POST  |     `/auth/register`      |       Register a new user       |        :x:         |
|  POST  |       `/auth/login`       | Authenticate user and get token |        :x:         |
|  GET   |        `/auth/me`         |      Get current user info      | :white_check_mark: |
|  POST  |       `/templates`        |      Create a new template      | :white_check_mark: |
|  GET   |       `/templates`        |  Get current user's templates   | :white_check_mark: |
|  POST  | `/templates/:id/generate` |  Generate mocks for a template  | :white_check_mark: |

## :sparkles: Supported Template Variables

### Misc

| Variable |   Description    |                Example                 |
| :------: | :--------------: | :------------------------------------: |
| `{uuid}` | Random UUID (v4) | `682fb1f6-d93f-4a6d-91b4-f300ea71ae7d` |
| `{bool}` |  Random boolean  |           `true` or `false`            |

### Person

|    Variable     |         Description         |          Example           |
| :-------------: | :-------------------------: | :------------------------: |
|    `{name}`     |      Random full name       |      `Leopold Hansen`      |
| `{first_name}`  |      Random first name      |          `Jacey`           |
| `{middle_name}` |     Random middle name      |        `Guillermo`         |
|  `{last_name}`  |      Random last name       |         `Jacobson`         |
|    `{phone}`    |     Random phone number     |        `9164607404`        |
|    `{email}`    |    Random email address     | `athanrath@bechtelar.info` |
|  `{username}`   |       Random username       |       `Smitham7869`        |
|  `{password}`   | Random password of length 8 |         `VUpCsp3M`         |
|   `{gender}`    |        Random gender        |     `male` or `female`     |

### Address

|    Variable    |    Description    |       Example       |
| :------------: | :---------------: | :-----------------: |
|  `{country}`   |  Random country   |       `Spain`       |
|    `{city}`    |    Random city    |      `Chicago`      |
|   `{street}`   |   Random street   | `9961 West Rowland` |
|    `{zip}`     |  Random zip code  |       `95183`       |
|  `{latitude}`  |  Random latitude  |     `31.617583`     |
| `{longtitude}` | Random longtitude |     `75.031913`     |

### Words

|    Variable     |    Description     |                         Example                         |
| :-------------: | :----------------: | :-----------------------------------------------------: |
|    `{word}`     |    Random word     |                        `though`                         |
|    `{noun}`     |    Random noun     |                         `crowd`                         |
|    `{verb}`     |    Random verb     |                         `wash`                          |
|   `{adverb}`    |   Random adverb    |                        `clearly`                        |
| `{preposition}` | Random preposition |                      `opposite to`                      |
|  `{adjective}`  |  Random adjective  |                        `sparse`                         |
|   `{phrase}`    |   Random phrase    |                  `you had to be there`                  |
|  `{question}`   |  Random question   | `Kale chips mlkshk mustache butcher shabby chic tacos?` |

### Colors

| Variable  |       Description        |  Example  |
| :-------: | :----------------------: | :-------: |
| `{color}` |       Random color       |  `White`  |
|  `{hex}`  | Random hexadecimal color | `#680da6` |

### Internet

|     Variable      |         Description         |                  Example                  |
| :---------------: | :-------------------------: | :---------------------------------------: |
|      `{url}`      |         Random URL          | `http://www.productvirtual.name/generate` |
|  `{domain_name}`  |     Random domain name      |            `futureglobal.info`            |
| `{domain_suffix}` |    Random domain suffix     |                   `org`                   |
|     `{ipv4}`      | Random version 4 IP address |              `227.44.58.21`               |
|     `{ipv6}`      | Random version 6 IP address | `10da:dd53:10af:49a5:f108:cbcb:8df7:d672` |

### Date/Time

|    Variable     |                Description                |                           Example                            |
| :-------------: | :---------------------------------------: | :----------------------------------------------------------: |
|    `{date}`     |                Random date                |          `1915-09-16 08:38:40.340355206 +0000 UTC`           |
|  `{past_date}`  |             Random past date              | `2025-08-13 04:07:27.273437073 +0300 MSK m=-40484.952189785` |
| `{future_date}` |            Random future date             | `2025-08-13 18:07:27.273359687 +0300 MSK m=+9915.047733388`  |
|    `{year}`     | Random year between 1900 and current year |                            `1959`                            |
|    `{month}`    |        Random month ordinal number        |                             `8`                              |
|   `{weekday}`   |              Random weekday               |                           `Monday`                           |
|    `{hour}`     |                Random hour                |                             `6`                              |
|   `{minute}`    |               Random minute               |                             `16`                             |
|   `{second}`    |               Random second               |                             `18`                             |

## :mailbox_with_mail: Postman API Collection

This project includes a Postman collection and environment for testing the API easily.

- :bookmark_tabs: [Postman Collection](backend/postman/moxer.postman_collection.json)
- :world_map: [Postman Environment](backend/postman/moxer.postman_environment.json)

### :white_check_mark: How to use

1. Import both collection and environment files into Postman.
2. Send the register and then login requests. If successful, the `jwt-token` will be saved autimatically and used in the `Authorization` header for all other requests where it is needed.

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
