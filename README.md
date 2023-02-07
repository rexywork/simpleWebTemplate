<h1>Overview</h1>
This project is intended to be used as web server template development.

The work currently is in progress to create more organized template

<h1>Prerequisites</h1>
- Go Language Installed

<h2>How to run the services</h2>

Windows
````
- go build cmd\main.go
- go run cmd\main.go
````

Linux
````
- go build cmd/main.go
- go run cmd/main.go
````

<h2> How to access Swagger </h2>

URL: http://localhost:8080/swagger/<version>/index.html

<h2> How to renew Swagger description </h2>
Usually the swagger would be updated to latest.
However if you need to do some development and changing the API,
you need to update the swagger again.
Make sure you already have <b>swag</b> by using 

````
go install github.com/swaggo/swag/cmd/swag@latest
````

and then you can run:

Windows
- `swag i -g pkg\api\api_v1.go  --instanceName v1`

Linux
- `swag i -g pkg/api/api_v1.go  --instanceName v1`