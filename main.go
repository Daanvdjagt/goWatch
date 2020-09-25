package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Movies : Struct that defines a movie model
type Movies struct {
	ID        int     `gorm:"AUTO INCREMENT" form:"id" json:"id"`
	ImdbID    string  `gorm:"not null" form:"imdbid" json:"imdbid"`
	Name      string  `gorm:"not null" form:"name" json:"name"`
	Year      int     `gorm:"not null" form:"year" json:"year"`
	ImdbScore float32 `gorm:"not null" form:"imdbscore" json:"imdbscore"`
}

// InitDb InitsDb and creates if does not exist
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Movies{}) {
		db.CreateTable(&Movies{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Movies{})
	}

	return db
}

//Cors Define cors headers for api
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{

		v1.POST("/movies", PostMovie)
		v1.GET("/movies", GetMovies)
		v1.GET("/movies/:id", GetMovie)
	}

	r.Run(":8080")
}

// PostMovie Creates a new movie
func PostMovie(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var movie Movies
	c.Bind(&movie)

	if movie.ImdbID != "" && movie.Name != "" && movie.Year != 0 && movie.ImdbScore != 0 {
		// 	INSERT INTO "movies" (ImdbId) values (movie.ImdbId, movie.Name, movie.Year, movie.ImdbScore)
		db.Create(&movie)
		// Display error
		c.JSON(201, gin.H{"success": movie})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"imdbid\": \"tt0816692\", \"name\": \"InterStellar\", \"year\": \"2014\", \"imdbscore\": \"8,6\"  }" http://localhost:8080/api/v1/movies
}

// ParsWatchList Parses a watchlist and adds it to the db
func ParseWatchList(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	// Open the file
	csvfile, err := os.Open("watchlist.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ImdbID: %s Name: %s Year: %s ImdbScore: %s\n", record[1], record[5], record[10], record[8])

	}
}

// GetMovies Gets all movies in db
func GetMovies(c *gin.Context) {
	//Connection to db
	db := InitDb()
	//Close connection database
	defer db.Close()

	var movies []Movies
	// SELECT * FROM movies
	db.Find(&movies)

	//Display JSON result
	c.JSON(200, movies)

	// curl -i http://localhost:8080/api/v1/movies
}

// GetMovie Get specific movie by imdbid
func GetMovie(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	imdbid := c.Params.ByName("imdbid")
	var movie Movies
	// SELECT * FROM movies WHERE imdb = tt0816692;
	db.First(&movie, imdbid)

	if movie.ImdbID != "" {
		// Display JSON result
		c.JSON(200, movie)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Movie not found"})
	}

	// curl -i http://localhost:8080/api/v1/movies/tt0816692

}

// OptionsMovie Define options headers for api
func OptionsMovie(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
