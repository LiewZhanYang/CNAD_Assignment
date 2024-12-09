document.addEventListener("DOMContentLoaded", () => {
  // Extract car details from the URL query string
  const params = new URLSearchParams(window.location.search);
  const carId = params.get("id");

  if (!carId) {
    alert("Car details not available!");
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
      ).innerText = `Subtotal: $${data.price}/day`;
      document.getElementById("car-total").innerText = `Total: $${data.price}`;
    })
    .catch((error) => console.error("Error fetching car details:", error));

  // Handle rent now button click
  document.getElementById("rent-now-btn").addEventListener("click", () => {
    const pickupLocation = document.getElementById("pickup-location").value;
    const dropoffLocation = document.getElementById("dropoff-location").value;
    const pickupDate = document.getElementById("pickup-date").value;
    const dropoffDate = document.getElementById("dropoff-date").value;

    if (!pickupLocation || !dropoffLocation || !pickupDate || !dropoffDate) {
      alert("Please fill out all required fields.");
      return;
    }

    alert("Your rental has been successfully booked!");
    // Redirect or perform further actions
  });
});
