# Cookie Jar

This is a backend API for demonstrating the use of cookies in a web application. It allows users to set a cookie, destroys the cookie and view a protected API.

## Prerequisites
- Go 1.20 or higher

## Getting Started
To run the application, you will need to have Go installed. You can download it [here](https://go.dev/dl/).
```bash
go run main.go
```

## Endpoints

The following endpoints are available:
```
GET /health # Health check
GET /api/auth/signin # Sets a cookie
GET /api/auth/signout # Destroys the cookie
GET /api/secure/resource # Protected resource
```

## Extras
This is to pair with the [Cookie Monster](https://github.com/jintaokoong/cookie-monster) frontend application.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.