const carListElement = document.getElementById("car-list");

// Fetch vehicles from the API
fetch("http://localhost:8081/vehicles/")
  .then((response) => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
  })
  .then((data) => {
    const vehicles = data.vehicles;

    // Clear existing hardcoded car cards
    carListElement.innerHTML = "";

    // Populate the car list dynamically
    vehicles.forEach((vehicle) => {
      const carCard = document.createElement("div");
      carCard.classList.add("car-card");
      carCard.innerHTML = `
        <img src="${vehicle.image_url}" alt="${vehicle.brand} ${
        vehicle.model
      }" />
        <h3>${vehicle.brand} ${vehicle.model}</h3>
        <p>${vehicle.vehicleType}</p>
        <div class="specs">
          <span><i class="fas fa-map-marker-alt"></i> ${vehicle.area}</span>
          <span><i class="fas fa-user"></i> ${
            vehicle.personCapacity
          } People</span>
        </div>
        <div class="price">$${vehicle.price.toFixed(2)}/day</div>
        <button class="rent-now-btn" data-id="${vehicle.id}">Rent Now</button>
      `;
      carListElement.appendChild(carCard);
    });
    // Attach event listeners to the "Rent Now" buttons
    document.querySelectorAll(".rent-now-btn").forEach((button) => {
      button.addEventListener("click", (event) => {
        const vehicleId = event.target.getAttribute("data-id");
        // Redirect to vehicle details page with the selected vehicle ID as a query parameter
        window.location.href = `../vehicleDetail/vehicleDetail.html?id=${vehicleId}`;
      });
    });
  })
  .catch((error) => {
    console.error("Error fetching vehicles:", error);
    carListElement.innerHTML =
      "<p>Failed to load popular cars. Please try again later.</p>";
  });
