document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");
  if (!token) {
    alert("You are not logged in!");
    window.location.href = "signin.html"; // Redirect to login page
    return;
  }

  // Decode JWT token to prefill the form
  const base64Url = token.split(".")[1];
  const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
  const jsonPayload = decodeURIComponent(
    atob(base64)
      .split("")
      .map((c) => `%${("00" + c.charCodeAt(0).toString(16)).slice(-2)}`)
      .join("")
  );
  const user = JSON.parse(jsonPayload);

  // Populate the form fields
  document.getElementById("name").value = user.name || "";
  document.getElementById("email").value = user.email || "";
  document.getElementById("phone").value = user.phone_number || "";

  // Handle form submission
  document
    .getElementById("update-details-form")
    .addEventListener("submit", async (event) => {
      event.preventDefault();

      const updatedDetails = {
        name: document.getElementById("name").value,
        email: document.getElementById("email").value,
        phone_number: document.getElementById("phone").value, // Use "phone_number" as the key
      };

      try {
        const response = await fetch(`http://localhost:8080/users/profile/${user.id}`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
          body: JSON.stringify(updatedDetails),
        });

        if (response.ok) {
          alert("Profile updated successfully!");
          window.location.href = "profile.html"; // Redirect to profile page
        } else {
          const errorData = await response.json();
          alert(`Failed to update profile: ${errorData.error}`);
        }
      } catch (error) {
        console.error("Error updating profile:", error);
        alert("An error occurred. Please try again.");
      }
    });
});
