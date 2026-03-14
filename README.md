# 🔗 Go URL Shortener

A simple and fast **URL Shortener** built with **Golang, PostgreSQL, and a minimal frontend UI**.  
This project allows users to convert long URLs into short, shareable links and automatically redirects users to the original destination.

🌐 **Live Demo:**  
https://anu-dddi.onrender.com/

---

# 📌 Features

- Shorten long URLs instantly
- Redirect short links to the original URL
- PostgreSQL database for persistent storage
- Simple clean frontend UI
- Copy-to-clipboard functionality
- Loading spinner & error handling
- Production deployment on cloud

---

# 🧠 How It Works

1. User enters a **long URL**
2. The backend stores the URL in PostgreSQL
3. The database generates a unique **ID**
4. The ID is converted into a **Base62 encoded short code**
5. The short URL is returned to the user

Example:

Original URL
