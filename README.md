# Microservices Architecture

## 1. Design Consideration of Microservices

### 1.1. Modularity
Each microservice is designed to handle a specific business function:
- **User Service**: Manages user authentication, profile, and JWT-based authorization.
- **Payment Service**: Handles bookings, billing, invoicing, and payment processing.
- **Vehicle Service**: Manages vehicle listings, availability, and details.

### 1.2. Scalability
The microservices architecture ensures scalability by enabling independent scaling of each service based on load. For example:
- The **Payment Service** can scale independently during high transaction volumes.
- The **Vehicle Service** can scale during browsing-heavy periods.

### 1.3. Fault Isolation
Failures in one microservice do not impact the entire system. For instance:
- If the **Payment Service** encounters an issue, the **User Service** and **Vehicle Service** continue to operate independently.

### 1.4. Communication
- Microservices communicate via **RESTful APIs** using **JSON** as the data exchange format.
- Each service is **decoupled**, reducing interdependencies and making the system more resilient and flexible.

### 1.5. Database Design
Each microservice maintains its own database to ensure loose coupling:
- **User Service Database**: Stores user credentials, profiles, and tokens.
- **Payment Service Database**: Manages bookings, billing, and invoice information.
- **Vehicle Service Database**: Stores vehicle details, availability, and booking records.

### 1.6. Security
- **JWT-based authentication** is implemented to ensure secure communication between services and clients.
- Sensitive information, such as passwords and email credentials, is stored securely and accessed through **environment variables**, reducing the risk of exposure.

