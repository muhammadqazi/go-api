# Student Information System (SIS) Clean Architecture


##### What is Gin?

Gin is a web framework written in Golang.
It features a Martini-like API, but with performance up to 40 times faster than Martini.
If you need performance and productivity, you will love Gin.


Clone the project

```
git clone git@github.com:muhammadqazi/SIS-Backend-Go.git
```

### Scripts

`make server` to run main.go without fast refresh.

`make fast` for the fast refresh.

`make func` to add a new function to application, it will make new file in handler,mapper,service and repository. Just write
` make func name=functionName `

`check-db` check dp will check if the database in docker it's working or not.

For `air` binary file 

`curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s` use this command it will create a `bin` directory in project copy air and move it to project root near make file.

##### Air Config
```
[build]
cmd = "go build -o ./tmp/main ./src/cmd/main.go"
```

### ENVIRONMENT VARIABLES

All the environment variables are in `.env.sample` file.
