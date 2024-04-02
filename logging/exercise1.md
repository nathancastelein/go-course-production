# Exercise 1

Open [middleware.go](./middleware.go) file. A boilerplate middleware is already written.

## Add simple logs

Using [log/slog](https://pkg.go.dev/log/slog):
- Add a new log before each request, containing the message `handle request`.
- Add a new log after each request, containing the message `request handled`.

Test your code:
```
go test -run 'TestLoggerMiddlewareSuite/TestSimpleLog'
```

## Improve your logs with attributes

Add some attributes on your logs `handle request` and `request handled`:
- An attribute `http_method` with the request's method
- An attribute `http_uri` with the request's URI

Test your code:
```
go test -run 'TestLoggerMiddlewareSuite/TestLogWithAttributes'
```

## Add the request id in your logs

Add some attributes on your logs `handle request` and `request handled`:
- From the incoming http request, extract the `X-Request-Id` header
- If request id is empty, generate a new one using `uuid.New().String()` from [google/uuid](https://github.com/google/uuid) package
- Add the request id in your logs as an attribute named `request_id`

Test your code:
```
go test -run 'TestLoggerMiddlewareSuite/TestLogWithRequestID'
```

## Add elapsed time in your logs

Using time package and [time.Since](https://pkg.go.dev/time#Since) function, add a new field `elapsed_time` in your log to print the elapsed time in milliseconds on the `request handled` log.

Test your code:
```
go test -run 'TestLoggerMiddlewareSuite/TestLogWithElapsedTime'
```