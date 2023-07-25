# README

## Running
- For convenience, while the project lacks configuration, it can be built and run with docker-compose
- Use the make command 
```
make docker_compose
```
- The correct context and paths and clean up will be handled
- Users registered with the domain @jackboxgames.com are automatically admins
- All other users are not

## Access
- Service runs on 8080
- Swagger docs can be found out http://localhost:8080/swagger
- Frontend app is served at http://localhost:8080
- API is accessible at http://localhost:8080/api/v1

## General considerations
- Add actual configuration
- Add tests
- Hook up monitoring
- Separate some functionality out
- Styling frontend
- Handle state and store jwt differently
- Use open api generator for frontend client
- lots of stuff...