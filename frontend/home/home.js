const carListElement = document.getElementById("car-list");

// Fetch vehicles from the API
fetch("http://localhost:8081/vehicles")
  .then((response) => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
  })
  .then((data) => {
    const vehicles = data.vehicles;

    // Populate the car list dynamically
    vehicles.forEach((vehicle) => {
      const carCard = document.createElement("div");
      carCard.classList.add("car-card");
      carCard.innerHTML = `
        <img src="${vehicle.image_url}" alt="${vehicle.make} ${
        vehicle.model
      }" />
        <h3>${vehicle.make} ${vehicle.model}</h3>
        <p>${vehicle.vehicle_type}</p>
        <div class="specs">
          <span><i class="fas fa-map-marker-alt"></i> ${vehicle.location}</span>
          <span><i class="fas fa-user"></i> ${vehicle.capacity} People</span>
        </div>
        <div class="price">$${vehicle.price.toFixed(2)}/day</div>
        <button>Rent Now</button>
      `;
      carListElement.appendChild(carCard);
    });
  })
  .catch((error) => {
    console.error("Error fetching vehicles:", error);
  });
