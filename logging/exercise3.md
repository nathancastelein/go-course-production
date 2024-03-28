# Exercise 3

Let's now use our logger from context!

## Log user from query param

Open [handler.go](./handler.go) file:
- Get logger from context (current context is in `c.Request().Context()`)
- Print a log `saying hello` with an attribute `user` with the given user in query param

Test your code:
```
go test -run 'TestLoggerMiddlewareSuite/TestLoggerFromContext'
```

## Warn if no user

Open [handler.go](./handler.go) file:
- Add a WARN log `empty user` if user is empty

Test your code:
```
go test -run 'TestLoggerMiddlewareSuite/TestWarningWithNoQueryParam'
```