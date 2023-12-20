##  CRUD API's using GIN framework

Sample codebase to represent the CURD API's creation using Golang.

## Getting Started 🚀

To start using this, clone this repo to a new directory:

```bash
git clone https://github.com/Yatin-Uppal/Golang-REST-API.git
```
and install all the packages using 
``` bash
go mod download
```
dependencies required to run the project -
```
1. Mongodb (should be installed on your system)
2. create a dev.env file on the root of the project and insert the following key=value pairs
   - MONGO_URI = <Mongodb connection URI here>
   - APP_PORT= <PORT on which you want your app to run> (example : 3000)
   - DATABASE_NAME= <Name of your database> (example: sample-database)
```
and run the project using following command -
```
export ENV=dev && go run main.go
```
## License
MIT