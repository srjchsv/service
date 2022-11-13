## Educational project- Auth-Service

- Authorization. And authentification using JWT

Used as auth-service for simple-bank https://github.com/srjchsv/simplebank

### To build and run containerized use `docker compose up --build`
### To run postgres db for dev and test use `docker compose up db`

## Makefile commands to test endpoints (not containerised):
###
#### for dev use:
- `make signup` - perform a user registration
- `make signin` - sign in as a registered user. Get a token. Set cookie.
- `make api` - check if a signed in user with jwt token can be authenticated at a secured endpoint.
- `make refresh` - refresh token.
- `make logout` - delete cookie with token. logout.
### - `make coverage` - to run test and show coverage in the default browser.
#### for running docker container use:
- `make signupDocker` - perform a user registration
- `make signinDocker` - sign in as a registered user. Get a token. Set cookie.
- `make apiDocker` - check if a signed in user with jwt token can be authenticated at a secured endpoint.
- `make refreshDocker` - refresh token.
- `make logoutDocker` - delete cookie with token. logout.

useful links:
- jwt, clean architecture, project structure, patterns, api https://www.youtube.com/watch?v=1LFbmWk7NLQ&list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8&index=1
