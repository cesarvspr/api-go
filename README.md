## Quick backend using Fiber.

[Fiber](https://gofiber.io/) is an Express.js inspired web framework build on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for **fast** development with **zero memory allocation** and **performance** in mind.

## ⚡️ Quick start

   - [github.com/swaggo/swag](https://github.com/swaggo/swag) for auto-generating Swagger API docs
   - [github.com/securego/gosec](https://github.com/securego/gosec) for checking Go security issues
   - [github.com/go-critic/go-critic](https://github.com/go-critic/go-critic) for checking Go the best practice issues
   - [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint) for checking Go linter issues

Run project by this command:

Use vscode launcher and play with `Launch local`

_you may need to place an api-key on the launch.json file._

5. Go to API Docs page (Swagger): [127.0.0.1:4555/swagger/index.html](http://127.0.0.1:4555/swagger/index.html)

## 📦 Used packages

| Name                                                                  | Version    | Type       |
| --------------------------------------------------------------------- | ---------- | ---------- |
| [gofiber/fiber](https://github.com/gofiber/fiber)                     | `v2.34.0`  | core       |
| [arsmn/fiber-swagger](https://github.com/arsmn/fiber-swagger)         | `v2.31.1`  | middleware |
| [stretchr/testify](https://github.com/stretchr/testify)               | `v1.7.1`   | tests      |
| [joho/godotenv](https://github.com/joho/godotenv)                     | `v1.4.0`   | config     |
| [IBM/cloudant-go-sdk](https://github.com/IBM/cloudant-go-sdk)                       | `v0.2.1`   | database   |
| [swaggo/swag](https://github.com/swaggo/swag)                         | `v1.8.2`   | utils      |
| [google/uuid](https://github.com/google/uuid)                         | `v1.3.0`   | utils      |
| [go-playground/validator](https://github.com/go-playground/validator) | `v10.10.0` | utils      |

### ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/models` folder for describe business models and methods of your project
- `./app/queries` folder for describe queries for models of your project

### ./docs

**Folder with API Documentation**. This directory contains config files for auto-generated API Docs by Swagger.

### ./pkg

**Folder with project-specific functionality**. This directory contains all the project-specific code tailored only for your business use case, like _configs_, _middleware_, _routes_ or _utils_.

- `./pkg/configs` folder for configuration functions
- `./pkg/middleware` folder for add middleware (Fiber built-in and yours)
- `./pkg/repository` folder for describe `const` of your project
- `./pkg/routes` folder for describe routes of your project
- `./pkg/utils` folder with utility functions (server starter, error checker, etc)

### ./platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project, like _setting up the database_ or _cache server instance_ and _storing migrations_.

- `./platform/database` folder with database setup functions (by default, CLoudant)

## ⚙️ Configuration

```ini
#

# Stage status to start server:
#   - "dev", for start server without graceful shutdown
#   - "prod", for start server with graceful shutdown
STAGE_STATUS="dev"

# Server settings:
SERVER_HOST="0.0.0.0"
SERVER_PORT=5000
SERVER_READ_TIMEOUT=60

# AUTH settings:
API-KEY=2a855350-40b2-4598-b937-3ad5824f8ac4

# Database settings:
CLOUDANT_APIKEY=#create your db (or request to use mine?)
CLOUDANT_URL=https://a273a6f7-77f4-43eb-88ef-58fb17fbeca0-bluemix.cloudantnosqldb.appdomain.cloud

```
