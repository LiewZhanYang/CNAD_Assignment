document.addEventListener("DOMContentLoaded", () => {
  const vehicleList = document.getElementById("vehicle-list");

  let allVehicles = []; // Store all vehicles to filter dynamically

  // Fetch vehicle data from the GetAllVehicles endpoint
  fetch("http://localhost:8081/vehicles/")
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.json();
    })
    .then((data) => {
      allVehicles = data.vehicles; // Save all vehicles
      renderVehicles(allVehicles); // Initial rendering
    })
    .catch((error) => {
      console.error("Error fetching vehicles:", error);
    });

  // Function to render vehicles
  function renderVehicles(vehicles) {
    vehicleList.innerHTML = ""; // Clear previous list
    vehicles.forEach((vehicle) => {
      const vehicleCard = document.createElement("div");
      vehicleCard.classList.add("vehicle-card");
      vehicleCard.setAttribute("data-type", vehicle.vehicleType.toLowerCase());
      vehicleCard.setAttribute("data-location", vehicle.area.toLowerCase());
      vehicleCard.setAttribute("data-price", vehicle.price);

      vehicleCard.innerHTML = `
        <img src="${vehicle.image_url}" alt="${vehicle.brand} ${vehicle.model}" />
        <h3>${vehicle.brand} ${vehicle.model}</h3>
        <p>${vehicle.vehicleType}</p>
        <div class="specs">
          <span><i class="fas fa-user"></i> ${vehicle.personCapacity} People</span>
          <span><i class="fas fa-map-marker-alt"></i> ${vehicle.area}</span>
        </div>
        <div class="price">$${vehicle.price}/day</div>
        <button class="rent-now-btn" data-id="${vehicle.id}">Rent Now</button>
      `;
      vehicleList.appendChild(vehicleCard);
    });

    // Attach event listeners to the "Rent Now" buttons
    document.querySelectorAll(".rent-now-btn").forEach((button) => {
      button.addEventListener("click", (event) => {
        const vehicleId = event.target.getAttribute("data-id");
        // Redirect to vehicle details page with the selected vehicle ID as a query parameter
        window.location.href = `../vehicleDetail/vehicleDetail.html?id=${vehicleId}`;
      });
    });
  }

  // Update Price Display
  document.getElementById("price").addEventListener("input", () => {
    const price = document.getElementById("price").value;
    document.getElementById("max-price").innerText = `$${price}`;
    filterVehicles();
  });

  // Filter Vehicles
  function filterVehicles() {
    const selectedTypes = Array.from(
      document.querySelectorAll(".filters input[type='checkbox']:checked")
    )
      .map((checkbox) => checkbox.value.toLowerCase())
      .filter((value) =>
        [
          "mpv",
          "sedan",
          "hatchback",
          "west",
          "east",
          "north",
          "south",
        ].includes(value)
      );

    const maxPrice = document.getElementById("price").value;

    const filteredVehicles = allVehicles.filter((vehicle) => {
      const matchesType =
        selectedTypes.includes(vehicle.vehicleType.toLowerCase()) ||
        selectedTypes.includes(vehicle.area.toLowerCase());

      const matchesPrice = vehicle.price <= parseFloat(maxPrice);

      return matchesType && matchesPrice;
    });

    renderVehicles(filteredVehicles);
  }

  // Attach event listeners to filter checkboxes
  document
    .querySelectorAll(".filters input[type='checkbox']")
    .forEach((checkbox) => {
      checkbox.addEventListener("change", filterVehicles);
    });
});
