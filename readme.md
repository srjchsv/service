## Educational project- Service

### Currently made:
- Authorization. And authentification using JWT
- Load balancer and proxy nginx

### To be made:
-More feature

### To run "prod" use `docker compose up --build`
### To run postgres db for dev and test use `docker compose up db`

## Makefile commands to test endpoints (not containerised):
###
#### for local dev use:
- `make signup` - perform a user registration
- `make signin` - sign in as a registered user. Get a token. Set cookie.
- `make api` - check if a signed in user with jwt token can be authenticated at a secured endpoint.
- `make refresh` - refresh token.
- `make logout` - delete cookie with token. logout.
### - `make coverage` - to run test and show coverage in the default browser.
#### for running docker container use:
- `make signupProd` - perform a user registration
- `make signinProd` - sign in as a registered user. Get a token. Set cookie.
- `make apiProd` - check if a signed in user with jwt token can be authenticated at a secured endpoint.
- `make refreshProd` - refresh token.
- `make logoutProd` - delete cookie with token. logout.

useful links:
- jwt, clean architecture, project structure, patterns, api https://www.youtube.com/watch?v=1LFbmWk7NLQ&list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8&index=1
- loadbalancing, deploying https://codeburst.io/load-balancing-go-api-with-docker-nginx-digital-ocean-d7f05f7c9b31
