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


---


## 👨‍💻 Author

**Qassim Aljaafar**  
[LinkedIn](https://www.linkedin.com/in/qassim-aljaffer)  
📧 qassimhassan9@gmail.com  
📍 Manama, Bahrain

