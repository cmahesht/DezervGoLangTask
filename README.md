# DezervGoLangTask

Task | Auth and save user info API

The task is to develop a basic version of user create and update details API.

Folder Structure

- api
    - constants - contains globel constants and error codes
    - model - User models
    - module
      - Login
        - loginUser.go     - user data access layer
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

1. 
2. open config.json file and change "MONGODSN" Ip value to your system IP
3. run command at app root directory   "docker-compose up -d"

Check Api End Points.txt file in repo for api end points. \DezervGoLangTask\Api End Points.txt
