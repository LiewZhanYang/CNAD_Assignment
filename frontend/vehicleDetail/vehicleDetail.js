document.addEventListener("DOMContentLoaded", () => {
  const urlParams = new URLSearchParams(window.location.search);
  const vehicleId = urlParams.get("id");

  if (!vehicleId) {
    alert("Vehicle ID is missing!");
    return;
  }

  // Fetch vehicle details
  fetch(`http://localhost:8081/vehicles/${vehicleId}`)
    .then((response) => {
      if (!response.ok) {
        throw new Error("Failed to fetch vehicle details");
      }
      return response.json();
    })
    .then((data) => {
      const vehicle = data;

      // Populate vehicle details
      document.getElementById(
        "vehicle-name"
      ).textContent = `${vehicle.brand} ${vehicle.model}`;
      document.getElementById("vehicle-description").textContent =
        vehicle.description || "Description";
      document.getElementById(
        "vehicle-price"
      ).textContent = `$${vehicle.price}/day`;
      document.getElementById("main-image").src = vehicle.image_url;

      // Populate specifications
      const specs = document.getElementById("vehicle-specs");
      specs.innerHTML = `
          <div>Type: ${vehicle.vehicleType}</div>
          <div>Capacity: ${vehicle.personCapacity} People</div>
          <div>Location: ${vehicle.area}</div>
        `;

      // Populate image gallery
      const gallery = document.getElementById("image-gallery");
      gallery.innerHTML = `
          <img src="${vehicle.image_url}" alt="Gallery Image" />
        `;

      // Rent Now button action
      const rentNowBtn = document.getElementById("rent-now-btn");
      rentNowBtn.href = `rentNow.html?id=${vehicleId}`;
    })
    .catch((error) => {
      console.error(error);
      alert("Error fetching vehicle details.");
    });
});
