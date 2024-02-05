<script>
  export let explanation;
</script>

<style>
  .info-box {
    background-color: #f4f4f4;
    border: 1px solid #ddd;
    border-radius: 8px;
    padding: 20px;
    margin: 20px 0;
    font-size: 16px;
    line-height: 1.5;
    text-align: center; /* Added this line */
  }

  /* You can add more styles as per your design preferences */
</style>

<div style="display: flex; justify-content: center; align-items: center;">
  <div style="width: 75%" class="info-box">
    <p>{explanation}</p>
  </div>
</div>