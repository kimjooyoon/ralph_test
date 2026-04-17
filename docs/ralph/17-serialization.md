# Serialization Helpers → Domain Experiments

## 16. JSON Serialization
**Signal**: Data format conversion needs  
**Domain experiment**: Add `SerializeJSON(data interface{}) (string, error)` for JSON encoding  
**Next failing test**: SerializeJSON(map[string]string{"a": "b"}) = "{\"a\":\"b\"}", want "{\"a\":\"b\"}" (pure function, no I/O)

## 17. YAML Serialization
**Signal**: Alternative data format needs  
**Domain experiment**: Add `SerializeYAML(data interface{}) (string, error)` for YAML encoding  
**Next failing test**: SerializeYAML(map[string]string{"a": "b"}) = "a: b\n", want "a: b\n" (pure function, no I/O)
