document.addEventListener("DOMContentLoaded", () => {
  const params = new URLSearchParams(window.location.search);
  const carId = params.get("id");
  const pricePerDay = parseFloat(params.get("price"));

  if (!carId || !pricePerDay) {
    alert("Car details not available!");
    return;
  }

  // Decode JWT token to extract user details
  const token = localStorage.getItem("token");
  let userId; // Add this variable
  if (token) {
    const decodedToken = parseJwt(token);
    if (decodedToken && decodedToken.name && decodedToken.email) {
      document.getElementById("name").value = decodedToken.name;
      document.getElementById("phone").value = decodedToken.email;
      userId = decodedToken.user_id; // Set userId here
    } else {
      alert("Invalid user token!");
      window.location.href = "../signin/signin.html"; // Redirect to login page
      return;
    }
  } else {
    alert("You are not logged in!");
    window.location.href = "../signin/signin.html"; // Redirect to login page
    return;
  }

  // Fetch car details using the carId
  fetch(`http://localhost:8081/vehicles/${carId}`)
    .then((response) => response.json())
    .then((data) => {
      document.getElementById(
        "car-name"
      ).innerText = `${data.brand} ${data.model}`;
      document.getElementById(
        "car-type"
      ).innerText = `Type: ${data.vehicleType}`;
      document.getElementById(
        "car-capacity"
      ).innerText = `Capacity: ${data.personCapacity} People`;
      document.getElementById(
        "car-subtotal"
      ).innerText = `Subtotal: $${pricePerDay}/day`;
      document.getElementById("car-total").innerText = `Total: $${pricePerDay}`;
      document.querySelector(".summary img").src = data.image_url;
    })
    .catch((error) => {
      console.error("Error fetching car details:", error);
      alert("Error fetching vehicle details.");
    });

  // Event listeners for calculating total price
  document
    .getElementById("pickup-date")
    .addEventListener("change", calculateTotalPrice);
  document
    .getElementById("dropoff-date")
    .addEventListener("change", calculateTotalPrice);

  function calculateTotalPrice() {
    const pickupDate = new Date(document.getElementById("pickup-date").value);
    const dropoffDate = new Date(document.getElementById("dropoff-date").value);

    if (pickupDate && dropoffDate && dropoffDate >= pickupDate) {
      const days = Math.max(
        1,
        Math.ceil((dropoffDate - pickupDate) / (1000 * 60 * 60 * 24))
      );
      const totalPrice = days * pricePerDay;
      document.getElementById("car-total").innerText = `Total: $${totalPrice}`;
      return totalPrice;
    }
    return pricePerDay; // Default to 1 day price if dates are not valid
  }

  // Handle Rent Now button click
  document.getElementById("rent-now-btn").addEventListener("click", () => {
    const name = document.getElementById("name").value;
    const email = document.getElementById("phone").value;
    const address = document.getElementById("address").value;
    const pickupLocation = document.getElementById("pickup-location").value;
    const dropoffLocation = document.getElementById("dropoff-location").value;
    const pickupDate = document.getElementById("pickup-date").value;
    const pickupTime = document.getElementById("pickup-time").value;
    const dropoffDate = document.getElementById("dropoff-date").value;
    const dropoffTime = document.getElementById("dropoff-time").value;
    const cardNumber = document.getElementById("card-number").value;

    // Validate required fields
    if (
      !name ||
      !email ||
      !address ||
      !pickupLocation ||
      !dropoffLocation ||
      !pickupDate ||
      !pickupTime ||
      !dropoffDate ||
      !dropoffTime ||
      !cardNumber
    ) {
      alert("Please fill out all required fields.");
      return;
    }

    // Prepare booking data
    const bookingData = {
      user_id: userId, // User ID from JWT token
      address,
      pickUpLocation: pickupLocation,
      pickUpDate: pickupDate,
      pickUpTime: pickupTime,
      dropOffLocation: dropoffLocation,
      dropOffDate: dropoffDate,
      dropOffTime: dropoffTime,
      creditCardNumber: cardNumber,
      vehicle_id: carId, // Car ID from URL params
    };

    // Calculate total amount for billing
    const totalAmount = calculateTotalPrice();
    const billingData = {
      booking_id: null, // Placeholder, will be updated after booking creation
      amount: totalAmount,
      status: "Pending",
    };

    console.log("Booking Data:", bookingData); // Debug log for booking data
    console.log("Billing Data:", billingData); // Debug log for billing data

    // Post Booking Data to Backend
    fetch("http://localhost:8082/payments/bookings", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(bookingData),
    })
      .then((response) => {
        console.log("Booking Response Status:", response.status);
        if (!response.ok) {
          return response.json().then((err) => {
            console.error("Booking Error Response:", err);
            throw new Error(err.error || "Failed to create booking.");
          });
        }
        return response.json();
      })
      .then((bookingResponse) => {
        console.log("Booking Response:", bookingResponse);
        billingData.booking_id = bookingResponse.booking_id; // Assign booking ID
        // Post Billing Data to Backend
        return fetch("http://localhost:8082/payments/billing", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(billingData),
        });
      })
      .then((billingResponse) => {
        console.log("Billing Response Status:", billingResponse.status);
        if (!billingResponse.ok) {
          return billingResponse.json().then((err) => {
            console.error("Billing Error Response:", err);
            throw new Error(err.error || "Failed to create billing entry.");
          });
        }
        return billingResponse.json();
      })
      .then(() => {
        alert("Your rental has been successfully booked!");
        window.location.href = "../notifications/notificationAfterPay.html";
      })
      .catch((error) => {
        console.error(
          "Error processing the booking and billing:",
          error.message || error
        );
        alert(
          "An error occurred while creating the booking and billing. Please try again."
        );
      });
  });
  // Function to decode JWT token
  function parseJwt(token) {
    try {
      const base64Url = token.split(".")[1];
      const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
      const jsonPayload = decodeURIComponent(
        atob(base64)
          .split("")
          .map(function (c) {
            return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
          })
          .join("")
      );
      return JSON.parse(jsonPayload);
    } catch (error) {
      console.error("Invalid token format:", error);
      return null;
    }
  }
});
