# Encryption Helpers → Domain Experiments

## 18. Simple AES Encryption
**Signal**: Basic data security needs  
**Domain experiment**: Add `EncryptAES(plaintext string, key string) (string, error)` for mock AES encryption  
**Next failing test**: EncryptAES("hello", "key") = "encrypted", want "encrypted" (pure function, no I/O)

##