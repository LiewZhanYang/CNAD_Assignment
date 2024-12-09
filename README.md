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
  git clone [https://github.com/your-repo/microservices-project.git](https://github.com/LiewZhanYang/CNAD_Assignment.git)
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

Here is the detailed test data and routes for each function to help a person test each endpoint using Postman or any other HTTP client.

---

### **User Service**
1. **Get All Users**
   - **Route**: `GET http://localhost:8080/users`
   - **Description**: Retrieves a list of all users.
   - **Expected Response**:
     ```json
     [
       {
         "id": 1,
         "name": "John Doe",
         "email": "johndoe@example.com"
       },
       {
         "id": 2,
         "name": "Jane Smith",
         "email": "janesmith@example.com"
       }
     ]
     ```

2. **Get User by ID**
   - **Route**: `GET http://localhost:8080/users/1`
   - **Description**: Retrieves details of the user with ID 1.
   - **Expected Response**:
     ```json
     {
       "id": 1,
       "name": "John Doe",
       "email": "johndoe@example.com",
       "phone_number": "1234567890"
     }
     ```

3. **Register User**
   - **Route**: `POST http://localhost:8080/users/signup`
   - **Body**:
     ```json
     {
       "name": "Alice Johnson",
       "email": "alice@example.com",
       "password": "securepassword123"
     }
     ```
   - **Expected Response**:
     ```json
     {
       "message": "User registered successfully."
     }
     ```

4. **Login User**
   - **Route**: `POST http://localhost:8080/users/signin`
   - **Body**:
     ```json
     {
       "email": "alice@example.com",
       "password": "securepassword123"
     }
     ```
   - **Expected Response**:
     ```json
     {
       "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
     }
     ```

5. **Update User Profile**
   - **Route**: `PUT http://localhost:8080/users/profile/1`
   - **Body**:
     ```json
     {
       "name": "Alice Updated",
       "email": "aliceupdated@example.com",
       "phone_number": "9876543210"
     }
     ```
   - **Expected Response**:
     ```json
     {
       "message": "Profile updated successfully."
     }
     ```

---

### **Vehicle Service**
1. **Get All Vehicles**
   - **Route**: `GET http://localhost:8081/vehicles`
   - **Description**: Retrieves a list of all vehicles.
   - **Expected Response**:
     ```json
     [
       {
         "id": 1,
         "brand": "Toyota",
         "model": "Corolla",
         "price": 50.0,
         "location": "West"
       },
       {
         "id": 2,
         "brand": "Nissan",
         "model": "GT-R",
         "price": 80.0,
         "location": "East"
       }
     ]
     ```

2. **Get Vehicle by ID**
   - **Route**: `GET http://localhost:8081/vehicles/1`
   - **Description**: Retrieves details of the vehicle with ID 1.
   - **Expected Response**:
     ```json
     {
       "id": 1,
       "brand": "Toyota",
       "model": "Corolla",
       "price": 50.0,
       "location": "West",
       "capacity": 4
     }
     ```

---

### **Booking Service**
1. **Create Booking**
   - **Route**: `POST http://localhost:8082/payments/bookings`
   - **Body**:
     ```json
     {
       "user_id": 1,
       "address": "123 Main St, Singapore",
       "pickUpLocation": "West",
       "pickUpDate": "2024-12-15",
       "pickUpTime": "10:00",
       "dropOffLocation": "East",
       "dropOffDate": "2024-12-20",
       "dropOffTime": "16:00",
       "creditCardNumber": "4111111111111111",
       "vehicle_id": 1
     }
     ```
   - **Expected Response**:
     ```json
     {
       "message": "Booking created successfully.",
       "booking_id": 1
     }
     ```

2. **Modify Booking**
   - **Route**: `PUT http://localhost:8082/payments/bookings/1`
   - **Body**:
     ```json
     {
       "address": "456 Another St, Singapore",
       "pickUpLocation": "North",
       "pickUpDate": "2024-12-16",
       "pickUpTime": "09:00",
       "dropOffLocation": "South",
       "dropOffDate": "2024-12-22",
       "dropOffTime": "15:00",
       "creditCardNumber": "5555555555554444",
       "vehicle_id": 2
     }
     ```
   - **Expected Response**:
     ```json
     {
       "message": "Booking updated successfully."
     }
     ```

3. **Get Booking by ID**
   - **Route**: `GET http://localhost:8082/payments/bookings/1`
   - **Description**: Retrieves booking details with ID 1.
   - **Expected Response**:
     ```json
     {
       "id": 1,
       "user_id": 1,
       "address": "123 Main St, Singapore",
       "pickUpLocation": "West",
       "pickUpDate": "2024-12-15",
       "dropOffLocation": "East",
       "dropOffDate": "2024-12-20",
       "amount": 250.0
     }
     ```

4. **Delete Booking**
   - **Route**: `DELETE http://localhost:8082/payments/bookings/1`
   - **Description**: Deletes booking with ID 1.
   - **Expected Response**:
     ```json
     {
       "message": "Booking deleted successfully."
     }
     ```

---

### **Invoice Service**
1. **Generate and Send Invoice**
   - **Route**: `POST http://localhost:8082/invoices/generate`
   - **Body**:
     ```json
     {
       "booking_id": 1,
       "user_id": 1,
       "amount": 250.0,
       "status": "Paid"
     }
     ```
   - **Expected Response**:
     ```json
     {
       "message": "Invoice generated and sent successfully.",
       "invoice": {
         "id": 1,
         "booking_id": 1,
         "user_id": 1,
         "amount": 250.0,
         "status": "Paid",
         "invoice_date": "2024-12-15T10:00:00Z"
       }
     }
     ```

---

### Instructions:
- Test each endpoint by copying the **Route** into Postman.
- Use the provided **Body** data for `POST` and `PUT` requests.
- Verify the **Expected Response** matches the actual response.
- Ensure the services are running and the database is populated with the required data.
  - Check service logs for errors.
  - Verify database connections and configurations.



