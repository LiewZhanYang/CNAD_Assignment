const API_BASE_URL = "http://localhost:8080";

// Handle Sign In
document
  .getElementById("sign-in-button")
  .addEventListener("click", async () => {
    const email = document.getElementById("sign-in-user").value.trim();
    const password = document.getElementById("sign-in-pass").value.trim();

    if (!email || !password) {
      alert("Please fill in both email and password.");
      return;
    }

    try {
      const response = await fetch(`${API_BASE_URL}/signin`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      if (response.ok) {
        const data = await response.json();
        alert("Login successful!");

        // Save token to localStorage (optional)
        localStorage.setItem("authToken", data.token);

        // Redirect to home.html
        window.location.href = "../home.html"; // Adjust the path if necessary
      } else {
        const errorData = await response.json();
        alert(`Login failed: ${errorData.error}`);
      }
    } catch (error) {
      alert("An error occurred during login. Please try again.");
    }
  });

// Handle Sign Up
document
  .getElementById("sign-up-button")
  .addEventListener("click", async () => {
    const username = document.getElementById("sign-up-user").value.trim();
    const email = document.getElementById("sign-up-email").value.trim();
    const password = document.getElementById("sign-up-pass").value.trim();
    const repeatPassword = document
      .getElementById("sign-up-repeat-pass")
      .value.trim();

    if (!username || !email || !password || !repeatPassword) {
      alert("Please fill in all fields.");
      return;
    }

    if (password !== repeatPassword) {
      alert("Passwords do not match.");
      return;
    }

    try {
      const response = await fetch(`${API_BASE_URL}/signup`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name: username, email, password }),
      });

      if (response.ok) {
        const data = await response.json();
        alert(data.message); // Success message
      } else {
        const errorData = await response.json();
        alert(`Sign up failed: ${errorData.error}`);
      }
    } catch (error) {
      alert("An error occurred during sign up. Please try again.");
    }
  });
