# Logging-Like Helpers → Domain Experiments

## 14. Simple Logging Interface
**Signal**: Debugging needs in pure functions  
**Domain experiment**: Add `Log(msg string) string` for mock logging  
**Next failing test**: Log("test") = "test", want "test" (pure function, no I/O)

## 15. Timestamp Generation
**Signal**: Time-based string manipulation  
**Domain experiment**: Add `Timestamp() string` for current time in RFC3339 format  
**Next failing test**: Timestamp() = "2023-10-05T12:34:56Z", want "2023-10-05T12:34:56Z" (pure function, no I/O)
