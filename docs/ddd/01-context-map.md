# Context Map

## Bounded Contexts
- **String Manipulation**: Handles string splitting, reversing, trimming, joining, and replacement.  
- **Encoding/Decoding**: Provides base64 and hex encoding/decoding utilities.  
- **Data Analysis**: Includes numeric range generation and average calculation.  
- **Logging/Serialization**: Offers mock logging and JSON/YAML serialization.  
- **Security**: Contains AES encryption mock functionality.  
- **Unicode Handling**: Specialized for UTF-8 byte splitting, surrogate pairs, and emoji processing.  

## Aggregates
- **String**: Core domain object for all string manipulation operations.  
- **Data**: Aggregate for encoding/decoding operations.  
- **NumericRange**: Represents a range of integers for analysis.  
- **LogEntry**: Represents a log message in the logging context.  
- **SerializedData**: Represents data