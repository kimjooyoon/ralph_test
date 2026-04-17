# Encoding/Decoding Helpers → Domain Experiments

## 12. Base64 Encoding
**Signal**: Data encoding needs  
**Domain experiment**: Add `EncodeBase64(data []byte) string` for simple encoding  
**Next failing test**: EncodeBase64([]byte("hello")) = "aGVsbG8=", want "aGVsbG8=" (pure function, no I/O)

## 13. Hexadecimal Encoding
**Signal**: Low-level data manipulation  
**Domain experiment**: Add `EncodeHex(data []byte) string` for hex representation  
**Next failing test**: EncodeHex([]byte("abc")) = "616263", want "616263" (pure function, no I/O)
