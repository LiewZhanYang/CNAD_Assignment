document.addEventListener("DOMContentLoaded", () => {
  const notificationList = document.getElementById("notification-list");

  // Mock data (replace with backend fetch in production)
  const notifications = [
    {
      icon: "fas fa-car",
      title: "Booking Confirmed",
      message: "Your booking for Nissan GT-R has been successfully confirmed.",
      time: "10 minutes ago",
    },
    {
      icon: "fas fa-bell",
      title: "Promotion Alert",
      message: "Enjoy 10% off your next booking. Limited time only!",
      time: "1 hour ago",
    },
    {
      icon: "fas fa-check-circle",
      title: "Payment Received",
      message: "Your payment of $400.00 has been successfully processed.",
      time: "3 hours ago",
    },
  ];

  // Render notifications
  notifications.forEach((notification) => {
    const notificationItem = document.createElement("li");
    notificationItem.className = "notification-item";
    notificationItem.innerHTML = `
        <i class="${notification.icon} notification-icon"></i>
        <div class="notification-content">
          <h3>${notification.title}</h3>
          <p>${notification.message}</p>
          <p class="notification-time">${notification.time}</p>
        </div>
      `;
    notificationList.appendChild(notificationItem);
  });
});
