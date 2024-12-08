document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");
  if (!token) {
    alert("You are not logged in!");
    window.location.href = "signin.html"; // Redirect to login page
    return;
  }

  // Decode JWT token
  const base64Url = token.split(".")[1];
  const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
  const jsonPayload = decodeURIComponent(
    atob(base64)
      .split("")
      .map((c) => `%${("00" + c.charCodeAt(0).toString(16)).slice(-2)}`)
      .join("")
  );
  const user = JSON.parse(jsonPayload);

  // Populate user profile details
  document.getElementById("user-name").textContent = user.name;
  document.getElementById("user-email").textContent = `Email: ${user.email}`;
  document.getElementById(
    "user-phone"
  ).textContent = `Phone: ${user.phone_number}`;
  document.getElementById(
    "user-membership"
  ).textContent = `Membership: ${user.membership}`;
  document.getElementById("user-tier").textContent = user.membership;

  // Fetch rental history
  fetch(`http://localhost:8081/rentals/${user.user_id}`)
    .then((response) => response.json())
    .then((data) => {
      const rentalHistory = document.getElementById("rental-history");
      rentalHistory.innerHTML = ""; // Clear the table

      data.rentals.forEach((rental) => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${rental.car}</td>
            <td>${rental.rental_date}</td>
            <td>${rental.return_date}</td>
            <td>${rental.status}</td>
            <td>$${rental.amount}</td>
          `;
        rentalHistory.appendChild(row);
      });
    })
    .catch((error) => {
      console.error("Error fetching rental history:", error);
    });
});
