### This is a golang rest api made as an exercise.
By default this api runs on port 8080, this can be changed by editing line 65 in the "main.go" to whatever port you like.

## Setup
1. Run `go get github.com/gin-gonic/gin`
2. Run `go get github.com/jinzhu/gorm`
3. Run `go run ./main.go`

## Endpoints
Endpoints are based on the default port of 8080.  
> POST localhost:8080/api/v1/movie  
Creates a movie, has to include json object to create in db example is posted below:  

```
{  
	"ImdbId" : "tt0816692",  
	"Name" : "Interstellar",  
	"Year" : 2014,  
	"ImdbScore" : 8.7  
}
```

> GET localhost:8080/api/v1/movie  
Lists all movies that are currently in the db instance.  

> GET localhost:080/api/v1/movie/imdbid  
Lists a specific movie that is indexed by their IMDBID example request would be: localhost:8080/api/v1/movie/tt0816692 for Interstellar.  

