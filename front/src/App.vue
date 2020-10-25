<template>
<div class="container"> 
  <form v-on:submit.prevent="submitForm(imdbID)"> 
  <input v-model="imdbID">
  <button>Submit</button>
  </form>
  <movie v-for="movie in movies" :key="movie.id" v-bind="movie">
  </movie>
  </div>
</template>

<script>
import axios from 'axios'
import movie from './Movie'


export default {
  mounted() {
    axios.get('http://localhost:8000/api/v1/movie')
    .then(response => {
      // JSON responses are automatically parsed.
      this.movies = response.data
      console.log(response.data)
    })
    .catch(e => {
      this.errors.push(e)
    })
  },
  methods : {
    submitForm: function (imdbID) {
      axios.post(`http://localhost:8000/api/v1/movie/${imdbID}`).then(response=> {
        console.log(response.data)
        this.movies.push(response.data.success)
      })
    }
  },
  data () {
    return {
      movies: null,
      imdbID: null
    } 
  },
  components: {
    movie
  }
}
</script>

<style>

</style>
