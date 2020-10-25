<template>
<div class="box">
  <h1 class="title">{{name}}</h1>
  <p>IMDBID: {{IMDBID}}</p>
  <p>Plot: {{plot}}</p>
  <p>Score: {{score}}</p>
  <p>Year: {{year}}</p>
</div>
</template>

<script>
export default {
    props: {
        name: String,
        IMDBID: String,
        plot: String,
        score: Number,
        year: Number
    }
}
</script>

<style>

</style>