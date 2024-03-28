# Exercise 2

Let's play with [context.Context](https://pkg.go.dev/context)!

## Store and get request id from context

Open [context.go](./context.go) file:
- Write the function NewContextWithRequestID to create a new context with the request id as value
- Write the function GetRequestID to get the request id from a context

Test your code:
```
go test -run TestRequestIDInContext
```

## Store and get logger from context

Open [context.go](./context.go) file:
- Write the function NewContextWithLogger to create a new context with the logger as value
- Write the function GetLogger to get the logger from a context

Test your code:
```
go test -run TestLoggerInContext
```