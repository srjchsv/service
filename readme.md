## Educational project- Auth-Service

- Authorization and authentification using JWT

Used as auth-service for simple-bank https://github.com/srjchsv/simplebank

## Makefile commands to test endpoints:

- `make run` - runs db in container and the auth service
- `make signup` - perform a user registration
- `make signin` - sign in as a registered user. Get a token. Set cookie.
- `make api` - check if a signed in user with jwt token can be authenticated at a secured endpoint.
- `make refresh` - refresh token.
- `make logout` - delete cookie with token. logout.

useful links:

- jwt, clean architecture, project structure, patterns, api https://www.youtube.com/watch?v=1LFbmWk7NLQ&list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8&index=1
