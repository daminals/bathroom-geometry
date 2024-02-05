<!-- Tooltip.svelte -->

<script>
  export let text = '';
</script>

<div class="tooltip">
  {text}
</div>

<style>
  .tooltip {
    position: absolute;
    background-color: rgba(51, 51, 51, 0.8); /* Adjust the alpha value (0.8) for transparency */
    color: #fff;
    padding: 8px;
    border-radius: 4px;
    z-index: 1;
    visibility: visible;
    /* Set top and left to 0 initially */
    top: -375%;
    left: 25%;
    right: 25%;
    /* Adjust as needed */
  }
</style>
