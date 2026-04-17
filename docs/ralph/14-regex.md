# Regex-Like Pattern Matching → Domain Experiments

## 10. Simple Pattern Matching
**Signal**: Lightweight regex alternatives  
**Domain experiment**: Add `Match(pattern, s string) bool` for basic regex-like checks  
**Next failing test**: Match("^[a-z]+$", "abc123") = false, want false (test for simple regex syntax)

## 11. Wildcard Substring Search
**Signal**: Flexible string matching  
**Domain experiment**: Add `ContainsWildcard(s, pattern string) bool` for * and ? wildcards  
**Next failing test**: ContainsWildcard("hello", "h*l") = true, want true (test for wildcard matching)
