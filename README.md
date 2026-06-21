# Go Learning Roadmap

Current level: Junior 4
Next task: jun4_1

## Completed

- [x] jun0_1 CountVowels
  - strings
  - rune
  - loops

- [x] jun0_2 IsPalindrome
  - strings
  - rune
  - indices

- [x] jun0_3 FirstUniqueChar
  - map[rune]int
  - frequency counting
  - unicode

- [x] jun0_4 RemoveDuplicates
  - slice
  - append
  - map as set

- [x] jun0_5 MostFrequentChar
  - map
  - frequency counting
  - first winner logic

- [x] jun0_6 RotateRight
  - modulo
  - slices
  - indexing
  - edge cases

- [x] jun0_7 RemoveAt
  - removing elements from slice
  - backing array
  - slice copying
  - off-by-one

- [x] jun0_8 ReverseWords
  - strings.Fields
  - strings.Join
  - reverse slice
  - string immutability

- [x] jun0_9 Counter
  - struct
  - methods
  - value receiver
  - pointer receiver
  - struct copy
  - pointer semantics

- [x] jun0_10 Swap
  - pointers
  - dereference
  - address operator
  - nil pointers
  - multiple assignment

- [x] jun1_1 Divide
  - error
  - errors.New
  - multiple return values
  - error handling

- [x] jun1_2 Custom Validation Error
  - custom errors
  - error interface
  - implicit interface implementation
  - method sets

- [x] jun1_3 Interfaces and Polymorphism
  - interfaces
  - polymorphism
  - implicit interface implementation
  - interface values

- [x] jun1_4 Interface Nil Trap
  - nil interface
  - interface value
  - pointer nil
  - interface comparison

- [x] jun1_5 Pointer Receiver Interface
  - method set
  - pointer receiver
  - interface implementation

- [x] jun1_6 Interfaces + Polymorphism (Shape)
  - Interface implementation
  - Method Set basics
  - Polymorphism in Go
  - Struct vs interface behavior

  - [x] jun1_7 User Repository
  - custom errors
  - interfaces
  - interface values
  - repository pattern basics
  - map lookup

- [x] jun1_8 Notification Service
  - composition
  - dependency injection basics
  - interfaces as dependencies
  - custom errors

- [x] jun1_9 Logger Embedding
  - embedding
  - method promotion
  - composition vs embedding
  - shadowing

- [x] jun1_10 Auditor Embedding
  - embedding
  - method promotion
  - shadowing
  - method resolution

- [x] jun2_1 Packages and Visibility
  - packages
  - imports
  - exported identifiers
  - unexported identifiers
  - constructors

- [x] jun2_2 Package API Design
  - encapsulation
  - exported methods
  - unexported fields
  - package api

- [x] jun2_3 Config Constructor Validation
  - constructors
  - validation
  - package boundaries
  - custom errors

- [x] jun2_4 Modules and Package Imports
  - go.mod
  - module path
  - package imports
  - package organization

- [x] jun2_5 Sentinel Errors
  - errors.New
  - sentinel errors
  - package variables
  - error comparison

  - [x] jun2_6 Error Wrapping and errors.Is
  - error wrapping
  - fmt.Errorf
  - %w
  - errors.Is

  - [x] jun2_7 errors.As and Custom Errors
  - errors.As
  - custom error types
  - validation errors
  - type assertions via errors.As

  - [x] jun2_8 Error Polymorphism
  - errors.As
  - interface polymorphism
  - custom error types
  - error dispatching

  - [x] jun2_9 Dependency Injection Basics
  - dependency injection
  - service layer basics
  - constructor injection
  - interface dependencies

  - [x] jun3_1 Composition vs Embedding
  - composition
  - embedding
  - method promotion
  - shadowing

  - [x] jun3_2 JSON Basics
  - encoding/json
  - json.Marshal
  - json.Unmarshal
  - struct tags

  - [x] jun3_3 JSON Response Design
  - nested JSON
  - API responses
  - response DTO
  - json tags

  - [x] jun3_4 First HTTP Server
  - net/http
  - http.HandlerFunc
  - ResponseWriter
  - Request
  - HTTP methods

  - [x] jun3_5 HTTP JSON Endpoint
  - json.NewEncoder
  - Content-Type
  - HTTP Response
  - ServeMux
  - JSON API basics

  - [x] jun3_6 POST JSON Endpoint
  - POST method
  - Request Body
  - json.NewDecoder
  - json.NewEncoder
  - HTTP status codes
  - Bad Request handling

  - [x] jun3_7 User Validation API
  - HTTP POST
  - JSON Decoder
  - Validation Layer
  - Error Response
  - HTTP Status Codes

  - [x] jun3_8 Query Parameters API
  - Query Parameters
  - URL Values
  - GET Endpoint
  - JSON Response Helper

  - [x] jun3_9 Path Parameters API
    - Path Parameters
    - Request Routing
    - strconv.Atoi
    - URL Parameter Validation
    - Service Layer
    - Repository Layer
    - Dependency Injection
    - HTTP 400 vs 404 vs 500
    - ServeMux Route Patterns

  - [x] jun3_10 User Create API
    - HTTP POST
    - DTO
    - Validation
    - Layered Architecture
    - Dependency Injection
    - In-Memory Repository
    - HTTP 201 Created
    - Error Response Pattern

  - [x] jun3_10 HTTP Middleware Basics
    - Middleware
    - Middleware Chain
    - Request ID
    - Logging Middleware
    - ResponseWriter Wrapping
    - Request Lifecycle
    - Cross-Cutting Concerns

## Topics learned

- rune vs byte
- UTF-8
- len(string)
- len([]rune)
- nil map
- nil slice
- append
- backing array
- slice header
- len vs cap
- map lookup
- comma-ok idiom
- removing elements from slice
- backing array sharing
- append side effects
- off-by-one errors
- strings.Fields
- strings.Join
- string immutability
- string vs []byte
- string vs []rune
- Structs
- Methods
- Value Receiver
- Pointer Receiver
- Struct Copy Semantics
- Pointer Semantics
- Pointers
- Address Operator (&)
- Dereference Operator (*)
- Nil Pointers
- Multiple Assignment
- error interface
- errors.New
- explicit error handling
- Custom Errors
- Implicit Interface Implementation
- Method Sets
- Error Interface
- Interfaces
- Polymorphism
- Interface Values
- Nil Interface
- Interface Comparison
- Interface Internal Representation
- Method Set
- Pointer Receiver
- Interface Implementation Rules
- Interfaces (polymorphism)
- Interface method dispatch
- Struct method implementation rules
- Repository Pattern Basics
- Custom Error Types
- Error Interface Implementation
- Interface Dynamic Type
- Interface Dynamic Value
- Composition
- Dependency Injection Basics
- Interfaces as Dependencies
- Service Composition
- Embedding
- Method Promotion
- Shadowing
- Composition vs Embedding
- Method Resolution
- Embedded Method Access
- Shadowing vs Override
- Packages
- Imports
- Exported Identifiers
- Unexported Identifiers
- Package Visibility Rules
- Constructors (New pattern)
- Encapsulation
- Information Hiding
- Package API Design
- Exported Methods
- Unexported Fields
- Getter Pattern in Go
- Constructor Validation
- Package Boundaries
- Encapsulated Config Objects
- Validation Errors
- Go Modules
- Module Path
- Package Imports
- Package Organization
- Multi-Package Projects
- Sentinel Errors
- Package-Level Variables
- Error Identity
- Error Comparison
- Error Wrapping
- errors.Is
- Wrapped Errors
- Error Chains
- fmt.Errorf %w
- errors.As
- Custom Error Types
- Validation Errors
- Error Type Extraction
- Error Polymorphism
- Error Type Dispatching
- Multiple Error Types
- errors.As Deep Dive
- Dependency Injection
- Constructor Injection
- Service Layer Basics
- Dependency Inversion Basics
- Composition
- Embedding
- Method Promotion
- Shadowing
- Composition vs Embedding
- JSON
- encoding/json
- json.Marshal
- json.Unmarshal
- Struct Tags
- Reflection Basics
- DTO
- Response Model
- API Contract
- json.NewEncoder
- json:"-"
- API Response Design
- net/http
- http.Handler
- http.HandlerFunc
- http.ResponseWriter
- http.Request
- HTTP Methods
- ListenAndServe
- HTTP Status Codes
- json.NewEncoder
- HTTP JSON Response
- Content-Type
- ServeMux
- JSON API
- Header.Set vs Header.Add
- Request Body
- io.ReadCloser
- json.NewDecoder
- HTTP POST
- HTTP request lifecycle
- Request Body can be read only once
- HTTP POST
- JSON Decoder
- Validation Layer
- HTTP Status Codes
- Error Response Pattern
- Separation of Concerns
- Query Parameters
- URL Query
- url.Values
- Query().Get()
- RawQuery
- JSON Response Helper
- Path Parameters
- Request Path Variables
- r.PathValue()
- strconv.Atoi
- URL Parameter Validation
- Route Parameters
- ServeMux Pattern Routing
- Layered Architecture Basics
- Handler -> Service -> Repository
- Dependency Injection Through Layers
- HTTP 400 Bad Request
- HTTP 404 Not Found
- HTTP 500 Internal Server Error
- Error Mapping
- DTO
- Request DTO
- Create Endpoint
- In-Memory Repository
- HTTP 201 Created
- Service Layer Validation Flow
- Entity Creation Flow
- Middleware
- Middleware Chaining
- Cross-Cutting Concerns
- Request ID Pattern
- Logging Middleware
- ResponseWriter Wrapping
- Decorator Pattern
- HTTP Request Lifecycle
- WriteHeader Semantics
- Automatic Status Code 200