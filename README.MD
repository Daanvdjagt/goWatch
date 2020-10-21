### This is a golang rest api made as an exercise.
By default this api runs on port 8080, this can be changed by editing line 65 in the "main.go" to whatever port you like.

## Setup
1. Run `go get github.com/gin-gonic/gin`
2. Run `go get github.com/jinzhu/gorm`
3. Run `go run ./main.go`

## Endpoints
Endpoints are based on the default port of 8080.  
POST localhost:8080/api/v1/movie - Creates a movie, has to include json object to create in db example is posted below:  

```
{  
	"ImdbId" : "tt6723592",  
	"Name" : "Tenet",  
	"Year" : 2020,  
	"ImdbScore" : 7,8  
}
```

GET localhost:8080/api/v1/movie - Lists all movies that are currently in the db instance.  

GET localhost:080/api/v1/movie/imdbid - Lists a specific movie that is indexed by their IMDBID   
An example request would be: localhost:8080/api/v1/movie/tt0816692 for Interstellar.  

