
<!-- Add this link tag in your HTML head to import the selected font -->
<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap">


<script>
  import { averageRating } from './ratingsStore.js';
  
  let currentAverageRating = 0;
  
  // Subscribe to the average rating store
  averageRating.subscribe(value => {
    currentAverageRating = value;
  });
</script>

<aside>
  <h2>Ratings</h2>
  <p>Average Rating: {currentAverageRating}</p>
</aside>

<style>
  /* Header component styles */
  aside {
    width: 200px;
    background-color: #f0f0f0;
    position: absolute;
    top: 100px;
    bottom: 0;
    text-align: center;
    font-family: 'Roboto', sans-serif;
  }
</style>
