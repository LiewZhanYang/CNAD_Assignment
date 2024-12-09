CREATE database DBCNAD_Assignment;

-- id,name,email,password,phoneNum,membership(basic,premium,vip)
CREATE TABLE  user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phoneNum VARCHAR(15) ,
    membership ENUM('basic', 'premium', 'vip') NOT NULL
);

-- sample
INSERT INTO user (name, email, password, phoneNum, membership) VALUES
('Alice Johnson', 'alice.johnson@example.com', 'password123', '1234567890', 'basic'),
('Bob Smith', 'bob.smith@example.com', 'mypassword', '2345678901', 'premium'),
('Charlie Brown', 'charlie.brown@example.com', 'securepass', '3456789012', 'vip'),
('Diana Prince', 'diana.prince@example.com', 'wonderwoman', '4567890123', 'premium'),
('Ethan Hunt', 'ethan.hunt@example.com', 'impossible', '5678901234', 'basic');

Select * from User;


CREATE TABLE vehicle (
    id INT AUTO_INCREMENT PRIMARY KEY,
    brand ENUM('toyota', 'honda', 'bmw') NOT NULL,
    model VARCHAR(100) NOT NULL,
    vehicleType ENUM('sedan', 'mpv', 'hatchback') NOT NULL,
    area ENUM('west', 'east', 'north', 'south') NOT NULL,
    personCapacity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    image_url VARCHAR(255) 
);

INSERT INTO vehicle (brand, model, vehicleType, area, personCapacity, price, image_url) VALUES
('toyota', 'Corolla', 'sedan', 'west', 4, 50.00, 'https://hips.hearstapps.com/hmg-prod/images/2025-toyota-corolla-fx-102-6674930515eb4.jpg'),
('honda', 'Civic', 'sedan', 'east', 4, 60.00, 'https://www.honda.com.sg/images/cars/2021_All-New_Civic/Launch/Civic_Model_Page_Banner_-_1920x888px.jpg'),
('bmw', '3 Series', 'sedan', 'north', 4, 70.00, 'https://i.i-sgcm.com/new_cars/cars/12406/12406_m.jpg'),
('toyota', 'Innova', 'mpv', 'south', 7, 80.00, 'https://www.johortrip.com/wp-content/uploads/JohorTrip-Toyota-Innova.jpg'),
('honda', 'Jazz', 'hatchback', 'west', 5, 55.00, 'https://imgcdn.oto.com.sg/large/gallery/exterior/2/18/honda-jazz-front-angle-low-view-830673.jpg');

Select * from vehicle;

CREATE TABLE booking (
    id INT AUTO_INCREMENT PRIMARY KEY,
    userid INT NOT NULL,
    address VARCHAR(255) NOT NULL,
    pickUpLocation ENUM('west', 'east', 'north', 'south') NOT NULL,
    pickUpDate DATE NOT NULL,
    pickUpTime TIME NOT NULL,
    dropOffLocation ENUM('west', 'east', 'north', 'south') NOT NULL,
    dropOffDate DATE NOT NULL,
    dropOffTime TIME NOT NULL,
    creditCardNumber VARCHAR(16) NOT NULL,
    vehicleid INT NOT NULL,
    FOREIGN KEY (userid) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (vehicleid) REFERENCES vehicle(id) ON DELETE CASCADE
);

INSERT INTO booking (userid, address, pickUpLocation, pickUpDate, pickUpTime, dropOffLocation, dropOffDate, dropOffTime, creditCardNumber, vehicleid)
VALUES
(1, '123 Ang Mo Kio Ave 3, Singapore 560123', 'West', '2024-12-10', '09:00', 'East', '2024-12-15', '17:00', '4111111111111111', 1),
(1, '456 Jurong West St 52, Singapore 640456', 'North', '2024-12-12', '10:30', 'South', '2024-12-18', '18:30', '4222222222222222', 1),
(1, '789 Tampines Ave 8, Singapore 520789', 'East', '2024-12-13', '08:45', 'West', '2024-12-20', '16:00', '4333333333333333', 1),
(1, '321 Clementi Rd, Singapore 129321', 'South', '2024-12-14', '11:15', 'North', '2024-12-22', '15:45', '4444444444444444', 1),
(1, '987 Toa Payoh Lor 2, Singapore 319987', 'West', '2024-12-16', '07:00', 'East', '2024-12-25', '12:30', '4555555555555555', 1);

-- Drop existing foreign keys
ALTER TABLE booking
DROP FOREIGN KEY booking_ibfk_1;

ALTER TABLE booking
DROP FOREIGN KEY booking_ibfk_2;

-- Rename columns to match the referenced table
ALTER TABLE booking
CHANGE COLUMN userid user_id INT NOT NULL;

ALTER TABLE booking
CHANGE COLUMN vehicleid vehicle_id INT NOT NULL;

-- Recreate foreign keys with correct names
ALTER TABLE booking
ADD CONSTRAINT booking_ibfk_1
FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE;

ALTER TABLE booking
ADD CONSTRAINT booking_ibfk_2
FOREIGN KEY (vehicle_id) REFERENCES vehicle(id) ON DELETE CASCADE;

CREATE TABLE billing (
    id INT AUTO_INCREMENT PRIMARY KEY,
    booking_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    status ENUM('pending', 'paid', 'failed') DEFAULT 'pending',
    FOREIGN KEY (booking_id) REFERENCES booking(id) ON DELETE CASCADE
);

-- Sample data for billing table
INSERT INTO billing (booking_id, amount, status)
VALUES
    (1, 150.00, 'pending'),
    (2, 200.00, 'paid'),
    (3, 175.50, 'failed'),
    (4, 125.00, 'pending');


CREATE TABLE invoice (
  id INT AUTO_INCREMENT PRIMARY KEY,
  booking_id INT NOT NULL,
  user_id INT NOT NULL,
  amount DECIMAL(10, 2) NOT NULL,
  invoice_date DATETIME DEFAULT CURRENT_TIMESTAMP,
  status ENUM('Paid', 'Pending') DEFAULT 'Paid',
  FOREIGN KEY (booking_id) REFERENCES booking(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);









