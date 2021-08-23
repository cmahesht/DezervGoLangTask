# DezervGoLangTask

Task | Auth and save user info API

The task is to develop a basic version of user create and update details API.

Folder Structure

- api
    - constants - contains globel constants and error codes
    - model - User model
    - module
      - Login
        - loginUser.go    - user data access layer
        - loginRoute.go   - login api routes
        - loginService.go - login services
    - api.go  - registration of all modules with api groups
    - helpers - helper folder contains reusebale libraries for
      - configHelper     - helper to load configs
      - databasehelper   - helper for database operations
      - logginghelper    - helper for logging operations
      - validationhelper - helper for model validation 
    - config.json - app config file
    - go.mod - go dependency file
    - server.go - main file to run app

Steps to execute program

1. open code in vs code editor
2. open config.json file and change "MONGODSN" value to your localhost mongodb string
3. Go to terminal and execute "go run server.go" command
4. Check Api End Points.txt file in repo for api end points and its request body.  -- \DezervGoLangTask\Api End Points.txt
5. Once server gets started, go to postman and send request to api endpoints with its request body

