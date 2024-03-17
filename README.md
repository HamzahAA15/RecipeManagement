# Recipe Service

Recipe Service is a service that enables you to perform CRUD operations on recipes, add comments, and rate recipes. It uses an inMemory data source for user domain, except for the username Hamzah which hardcoded. using MongoDB as a database to store all recipe-related information.

**Note:** View this readme file using preview mode (Ctrl-K + V) or (Command-K + V) for better readability.

## Features

- bcrypt hash to secure user's password.
- Asynchronous email notification feature using Gmail SMTP server (using feature flag).
- Modular project structure with dependency injection on the repository & controller layers.
- Using Docker Compose to ease the experience of using this service

## Installation

To enable the email notification feature using Gmail SMTP server, follow these steps:
- Set `CONFIG_EMAIL_SERVICE` to `true`.
- Set `CONFIG_AUTH_PASSWORD` with the app secret of your Gmail account.

Set other `CONFIG` variables as required in `config.go`.

## Usage

Run the following command:

```bash
docker-compose build
docker-compose up
```

import the JSON collection of request from the attachment of email into API platform such as Postman

insert the Basic Authorization with below data
- username: hamzah
- password: kaskus12345

Or 
Create new User with API POST `host:port/api/user/`
 
## API LIST

### Recipes
- **POST** `host:port/api/recipes/`: Create a new recipe (authentication required)
- **PATCH** `host:port/api/recipes/`: Update a recipe (authentication required)
- **DELETE** `host:port/api/recipes/`: Delete a recipe (authentication required)
- **GET** `host:port/api/recipes/`: Get a recipe by ID
- **GET** `host:port/api/recipes/all`: Get all recipes
- **GET** `host:port/api/recipes/filter`: Get recipes by filter criteria

### Comments
- **PATCH** `host:port/api/recipes/comments/`: Add a comment to a recipe

### Ratings
- **PATCH** `host:port/api/recipes/ratings/`: Add a rating to a recipe

### User
- **POST** `host:port/api/user/`: Create a new user.
- **GET** `host:port/api/user/`: Get a user by name.
