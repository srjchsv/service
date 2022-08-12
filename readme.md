## Educational project- Service

### Currently made:
- Authorization using JWT
- Load balancer and proxy nginx

### To be made:
-More feature

### To run use `docker compose up`

### Makefile commands to test authorization endpoint

- `make signup` - perform a user registration
- `make signin` - sign in as a registered user. Get a token.
- `make api` - check if a signed in user with jwt token can be authenticated at a secured endpoint. Add a token to a variable TOKEN before using this command.


useful links:
- jwt, project structure, patterns, api https://www.youtube.com/watch?v=1LFbmWk7NLQ&list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8&index=1
- loadbalancing, deploying https://codeburst.io/load-balancing-go-api-with-docker-nginx-digital-ocean-d7f05f7c9b31
