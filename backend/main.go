package main

import (
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tidwall/gjson"
)

//Movie defines a movie
type Movie struct {
	ID     int     `json:"ID"`
	Name   string  `json:"name"`
	Year   int     `json:"year"`
	IMDBID string  `json:"IMDBID"`
	Score  float64 `json:"score"`
	Plot   string  `json:"plot"`
}

//Cors Define cors headers for api
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	InitDB()

	r := gin.Default()
	r.Use(Cors())

	api := r.Group("api/v1")
	{
		api.POST("/movie", PostMovie)
		api.POST("/movie/:id", PostAndFillMovie)
		api.GET("/movie", GetMovies)
		api.GET("/movie/:id", GetMovie)
		api.PATCH("/movie", UpdateMoviePlots)
	}

	r.Run("localhost:8000")
}

//InitDB Inits the db connection
func InitDB() {
	db := OpenDB()
	defer db.Close()

	if !db.HasTable(&Movie{}) {
		db.CreateTable(&Movie{})
		db.Set("gorm:table_options",
			"ENGINE=InnoDB").CreateTable(&Movie{})
	}
}

//OpenDB opens db
func OpenDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	return db
}

//PostMovie Creates a movie
func PostMovie(c *gin.Context) {
	db := OpenDB()
	defer db.Close()
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	var movie Movie
	c.Bind(&movie)

	if movie.Name != "" && movie.IMDBID != "" {
		db.Create(&movie)
		c.JSON(201, gin.H{"success": movie})
	} else {
		c.Data(422, "text/html; charset=utf-8", []byte(`
			<h1>422 - Empty name and/or IMDB_ID</h1>
		`))
	}
}

//GetMovies Gets all movies
func GetMovies(c *gin.Context) {
	db := OpenDB()
	defer db.Close()

	var movies []Movie
	db.Find(&movies)

	c.JSON(200, movies)
}

//GetMovie Gets one single movie by IMDBID
func GetMovie(c *gin.Context) {
	db := OpenDB()
	defer db.Close()

	IMDBID := c.Params.ByName("id")
	var movie Movie
	db.First(&movie, "imdb_id = ?", IMDBID)

	if movie.ID != 0 {
		c.JSON(200, movie)
	} else {
		c.Data(404, "text/html; charset=utf-8", []byte(`
			<h1>404 - Movie not found</h1>
		`))
	}
}

//UpdateMoviePlots Updates plots for all movies
func UpdateMoviePlots(c *gin.Context) {
	db := OpenDB()
	defer db.Close()

	var movies []Movie
	db.Find(&movies)

	var wg sync.WaitGroup

	for _, mov := range movies {
		if mov.Plot == "" {
			wg.Add(1)
			go func(movie Movie) {
				moviePlot := FetchMoviePlot(movie.IMDBID)
				updateMovie(moviePlot, movie.IMDBID)
				wg.Done()
			}(mov)
		}
	}
	wg.Wait()
}

//PostAndFillMovie Updates plots for all movies
func PostAndFillMovie(c *gin.Context) {
	db := OpenDB()
	defer db.Close()

	IMDBID := c.Params.ByName("id")
	movie := FetchMovie(IMDBID)
	c.Bind(&movie)

	if movie.Name != "" && movie.IMDBID != "" {
		db.Create(&movie)
		c.JSON(201, gin.H{"success": movie})
	} else {
		c.Data(422, "text/html; charset=utf-8", []byte(`
			<h1>422 - Empty name and/or IMDB_ID</h1>
		`))
	}

}

//FetchMoviePlot Fetches plot for movie from OMDBAPI
func FetchMoviePlot(IMDBID string) string {
	var moviePlot string
	res, err := http.Get("http://www.omdbapi.com/?i=" + IMDBID + "&apikey=c71f6e33")
	if err != nil {
		panic(err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fetchedPlot := gjson.Get(string(data), "Plot")
		moviePlot = fetchedPlot.String()
	}
	return moviePlot
}

//FetchMovie Fetches movie in OMDBAPI
func FetchMovie(IMDBID string) Movie {
	var movie Movie
	res, err := http.Get("http://www.omdbapi.com/?i=" + IMDBID + "&apikey=c71f6e33")
	if err != nil {
		panic(err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		movie.IMDBID = IMDBID
		movie.Name = gjson.Get(string(data), "Title").String()
		movie.Year = int(gjson.Get(string(data), "Year").Int())
		movie.Score = gjson.Get(string(data), "imdbRating").Float()
		movie.Plot = gjson.Get(string(data), "Plot").String()

	}
	return movie
}

//updateMovie Updates a movie with the newly fetched movie plot
func updateMovie(moviePlot, IMDBID string) {
	db := OpenDB()
	defer db.Close()

	var localMovie Movie

	db.First(&localMovie, "imdb_id = ?", IMDBID)

	if localMovie.ID != 0 {
		localMovie.Plot = moviePlot
		db.Save(&localMovie)
	}
}

// OptionsMovie Define options headers for api
func OptionsMovie(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
