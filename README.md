
# MiniEconomy

### ES
MiniEconomy es un proyecto personal cuyo objetivo es eminentemente educativo. MiniEconomy consiste en un sistema de simulación bancario diseñado para asistir eventos pequeños en los que se requiere una aproximación al sistema monetario real.

### EN
MiniEconomy is a personal project whose objective is mainly educative. MiniEconomy consist of a bank simulation system designed to aid small events where it's required a real monetary approach  

## Stack tecnológico

- **Golang - Backend**  
   **EN**  
   Internally developed with Gin to handle HTTP request, configured with JWT Authentication, Gorm as ORM with Automigrations enabled and Hexagonal Architecture with vertical slicing  
   **ES**  
   Desarrollado a nivel interno con Gin para manejar peticiones HTTP, con autenticación mediante JWT, Gorm como ORM con automigraciones y arquitectura hexagonal con vertical slicing

- **Flutter - Mobile** (TODO)  
  **EN**  
  Developed as user main interaction method to manage their accounts and make transactions


## Roadmap (Golang)
 - Architecture details should be covered, currently I'm trying to handle the Database connection with Singleton Pattern but in the current file structure the database connection is hanled inside the authentication module
 - The monetary transactions module should use Gorm transactions at repository level to enable rollbacks in error cases
 - Implement transactions rollback
- SMTP Module to send emails to users
- Generate PDF with transaction details (Probably just to educative purposes this will be developed with Spring Boot with Jasper Reports, that microservice would receive the transaction details through an AMQP messaging queue and then send it to email through SMTP, PDT: Yes, probably the SMTP module itself would be better held here)  
- Create a service to get a single account
  
