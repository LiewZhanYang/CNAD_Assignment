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

### 1.5. Security
- **JWT-based authentication** is implemented to ensure secure communication between services and clients.
- Sensitive information, such as passwords and email credentials, is stored securely and accessed through **environment variables**, reducing the risk of exposure.

## 2. Architecture Diagram
![Screenshot 2024-12-09 220541](https://github.com/user-attachments/assets/3ec51aea-6f84-4e26-917e-287d7c62b9e6)


This diagram illustrates the architecture of the system, showcasing the interactions between the client, frontend, microservices, and databases.

### Components:

1. **Client**: 
   - Users access the application through a web or mobile interface to interact with the services.

2. **Frontend**:
   - The user interface of the application.
   - Sends HTTP requests to the backend services to fetch or modify data.
   - Displays responses and updates dynamically based on user actions.

3. **API Gateway**:
   - Acts as a central routing point for client requests.
   - Handles load balancing, authentication, and routing requests to the appropriate microservices.

4. **Microservices**:
   - **User Service**:
     - Handles user authentication, registration, and profile management.
     - Includes endpoints such as `GetUserById`, `Register User`, `Login`, `Modify User Details`, and `GetAllUsers`.

   - **Vehicle Service**:
     - Manages vehicle listings, availability, and details.
     - Includes endpoints such as `Display All Vehicles` and `Get Vehicles Detail By Id`.

   - **Booking Service**:
     - Manages bookings, billing, and invoicing.
     - Includes endpoints such as:
       - `Create Booking`, `Modify Booking`, `Delete Booking`
       - `Get Billing By Id`, `Create Billing`
       - `Generate Invoice & Send`
       - `Get Booking By Id`, `Get Booking By UserId`
     - Integrates with a third-party payment service for transaction handling.

5. **Third-Party Payment**:
   - Processes user payments securely and integrates with the booking service for financial transactions.

### Key Features:
- **Modularity**: Each microservice focuses on a specific domain.
- **Scalability**: Microservices can scale independently based on demand.
- **Fault Isolation**: Issues in one service do not impact others.
- **Security**: API Gateway enforces authentication and authorization.
- **Efficiency**: RESTful APIs ensure seamless communication between components.


## 3. Instructions for Setting Up and Running Microservices

Follow these steps to set up and run the microservices:

### 1. Setting up the Database
- Install **MySQL** or any compatible database service.
- Create the required databases for each microservice:
  - `user_service`
  - `payment_service`
  - `vehicle_service`
- Import the SQL schema for each microservice into the respective database:
  - Use `.sql` files (e.g., `user_service.sql`, `payment_service.sql`, `vehicle_service.sql`) to create tables and relationships.
- Update the database connection details (host, port, username, password) in each microservice's configuration file (e.g., `.env`).

### 2. Loading the Microservices from GitHub
- Clone the project repository:
  ```bash
  git clone https://github.com/your-repo/microservices-project.git
  cd microservices-project
  ```
- Navigate to each microservice directory:
  - `user_service/`
  - `payment_service/`
  - `vehicle_service/`

### 3. Running the Microservices
- Install dependencies for each microservice:
  ```bash
  go mod tidy
  ```
- Run each service:
  ```bash
  go run main.go
  ```
- Verify that the services are running on their respective ports:
  - User Service: `http://localhost:8080`
  - Payment Service: `http://localhost:8082`
  - Vehicle Service: `http://localhost:8081`

### 4. Testing the Microservices with Postman
- Import the Postman collection provided in the repository (`postman_collection.json`).
- Use Postman to:
  - Test the endpoints for each service (e.g., login, create booking, fetch vehicle details).
  - Verify responses and database updates.
  - Example endpoints:
    - `GET http://localhost:8080/users/{id}`
    - `POST http://localhost:8082/payments/bookings`
    - `GET http://localhost:8081/vehicles`

### 5. Running the Frontend
- Open the frontend project directory:
  ```bash
  cd frontend/
  ```
- Start a live server to run the frontend (e.g., using Visual Studio Code's Live Server extension).
- Open the `signin.html` page in the browser:
  ```bash
  http://127.0.0.1:5500/frontend/signin.html
  ```
- Use the UI to test the integration of frontend with backend services:
  - Login to the system.
  - View and book vehicles.
  - Verify payment and invoicing processes.

### Additional Notes
- Ensure that all services are running before interacting with the frontend.
- Check `.env` files for configuration and update as required.
- For troubleshooting:
  - Check service logs for errors.
  - Verify database connections and configurations.



