<!-- Add this link tag in your HTML head to import the selected font -->
<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap">

<script>
import { accessibilityRating, cleanRating, menstrualRating, overallRating } from './ratingsStore';

    
</script>

<aside>
  <h2>Ratings</h2>
  <p>Accessibility Rating: {$accessibilityRating.toFixed(0)}</p>
  <p>Clean Rating: {$cleanRating.toFixed(0)}</p>
  <p>Menstrual Rating: {$menstrualRating.toFixed(0)}</p>
  <p>Overall Rating: {$overallRating.toFixed(0)}</p>
</aside>


<style>
    /* Header component styles */
    aside {
        width: 200px;
        background-color: #333;
        position: absolute;
        top: 125px;
        bottom: 0;
        text-align: center;
        font-family: 'Roboto', sans-serif;
        color: #fff;
    }
</style>
