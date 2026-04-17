# Numeric Helpers → Domain Experiments

## 4. Range Generation
**Signal**: Data analysis simplicity  
**Domain experiment**: Add `Range(start, end int) []int` that returns inclusive range  
**Next failing test**: Range(1, 5) = [1,2,3,4,5], want [1,2,3,4,5] (pure function, no I/O)

## 5. String to Integer Conversion
**Signal**: Basic data manipulation needs  
**Domain experiment**: Add `ParseInt(s string) (int, error)` for simple conversions  
**Next failing test**: ParseInt("123") = 123, want 123 (handle invalid inputs as failing tests)
