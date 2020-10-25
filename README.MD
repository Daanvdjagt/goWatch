### This is a golang rest api made as an exercise.

## API Setup
`cd ./backend`

#### All the imports
`go get github.com/gin-gonic/gin`  
`go get github.com/jinzhu/gorm`  
`go get github.com/mattn/go-sqlite3`  
`go get github.com/tidwall/gjson`  

#### Start the server
`go run ./main.go`

## Frontend setup
`cd ..`  
`cd ./front`  
`npm install`  
`npm run serve`  
```
	Frontend is hosted at 127.0.0.1:8080  
	Input an IMDBID into the input value and press submit to fetch and add the movie. 
```


## API Endpoints
Endpoints are based on the default port of 8000.  

GET 127.0.0.1:8000/api/v1/movie - Lists all movies that are currently in the db instance.  

GET 127.0.0.1:8000/api/v1/movie/imdbid - Lists a specific movie that is indexed by their IMDBID   
An example request would be: localhost:8080/api/v1/movie/tt0816692 for Interstellar.  


POST 127.0.0.1:8000/api/v1/movie - Creates a movie, has to include json object to create in db example is posted below:  

```
{  
	"ImdbId" : "tt6723592",  
	"Name" : "Tenet",  
	"Year" : 2020,  
	"ImdbScore" : 7,8  
}
```

PATCH 127.0.0.1:8080/api/v1/movie - Concurrently fetches all plots for movies in database 
Example request: 127.0.0.1:8000/api/v1/movie

POST 127.0.0.1:8080/api/v1/movie:id - Fetches all info for an IMDBID and saves it in the database
Example request: 127.0.1.1:8000/api/v1/movie/tt6723592
