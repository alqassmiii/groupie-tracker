# Groupie Tracker

Groupie Tracker is a web application written in Go that displays information about music artists, their members, concert dates, and locations. It fetches data from the [Groupie Tracker API](https://groupietrackers.herokuapp.com/api/artists) and presents it in a user-friendly interface.

## Features
- Home page listing all artists with images
- Artist detail page showing members, creation date, first album, concert dates, and locations
- Error pages for 400, 404, and 500 HTTP errors
- Responsive design with custom CSS

## Project Structure
```
├── go.mod                # Go module definition
├── main.go               # Entry point, starts the web server
├── API/
│   ├── api.go            # Fetches and processes API data
│   ├── handel.go         # HTTP handlers for main and artist pages
│   └── struct.go         # Data structures for artists, locations, dates, relations
├── Static/
│   └── Style.css         # Styles for all pages
├── template/
│   ├── 400Error.html     # 400 Bad Request error page
│   ├── 404Error.html     # 404 Not Found error page
│   ├── 500Error.html     # 500 Internal Server Error page
│   ├── ArtistPage.html   # Artist details page template
│   └── HomePage.html     # Home page template
```

## How It Works
- The server runs on port 8080.
- Static files (CSS) are served from `/Static/`.
- `/` displays the home page with all artists.
- `/artistInfo?name=ARTIST_NAME` displays details for a selected artist.
- Error pages are rendered for invalid requests or server errors.

## Getting Started
1. **Install Go (>= 1.22)**
2. Clone this repository:
   ```sh
   git clone <repo-url>
   cd groupie-tracker
   ```
3. Run the server:
   ```sh
   go run main.go
   ```
4. Open [http://localhost:8080](http://localhost:8080) in your browser.

---


## 👨‍💻 Author

**Qassim Aljaafar**  
[LinkedIn](https://www.linkedin.com/in/qassim-aljaffer)  
📧 qassimhassan9@gmail.com  
📍 Manama, Bahrain

