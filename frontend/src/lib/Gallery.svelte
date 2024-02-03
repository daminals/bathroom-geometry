<script>
    import Rate from "./Rate.svelte";
    
    let showForm = false;
    let currentIndex = 0;
    let images = [
        { name: "Image 1", src: "grid.jpg" },
        { name: "Image 2", src: "grid.jpg" },
        { name: "Image 3", src: "grid.jpg" }
    ];

    function prevImage() {
        currentIndex = (currentIndex - 1 + images.length) % images.length;
    }

    function nextImage() {
        currentIndex = (currentIndex + 1) % images.length;
    }

    function handleRate() {
        showForm = !showForm;
    }
</script>

<div class="gallery-container">
    {#if !showForm}
        <div class="gallery">
            <h2>{images[currentIndex].name}</h2>
            {#each images as { src }, index}
                <img src={src} alt={images[currentIndex].name} class="gallery-image" style="{index === currentIndex ? 'display: block;' : 'display: none;'}">
            {/each}
            <div class="nav-buttons">
                <button class="prev-btn" on:click={prevImage}>&lt;</button>
                <button class="next-btn" on:click={nextImage}>&gt;</button>
                <button class='b' on:click={handleRate}>Rate</button>
            </div>
        </div>
    {/if}

    <!-- Rating Form -->
    {#if showForm}
        <div class="rate-form">
            <Rate />
        </div>
    {/if}
</div>

<style>
    .gallery-container {
        font-family: 'Roboto', sans-serif;
        position: relative;
        max-width: 600px;
        margin: 0 auto;
        overflow: hidden; /* Hide overflow content */
    }

    .gallery {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .gallery-image {
        width: 300px;
        height: 200px;
        object-fit: cover;
        margin-bottom: 20px;
    }

    .nav-buttons {
        margin-top: 10px;
        text-align: center;
    }

    .prev-btn,
    .next-btn {
        background: transparent;
        border: none;
        font-size: 20px;
        color: #000;
        cursor: pointer;
        outline: none;
    }

    .prev-btn {
        margin-right: 10px;
    }

    .next-btn {
        margin-left: 10px;
    }

    .rate-form {
        margin-top: 20px;
    }
    .b {
        display: inline-block;
    padding: 10px 20px;
    background-color: #880404;
    color: #fff;
    border: none;
    border-radius: 20px;
    font-size: 16px;
    font-weight: bold;
    text-transform: uppercase;
    cursor: pointer;
    transition: background-color 0.3s ease;
    }
    .b:hover {
        background-color: #460706;
    }
    .b:active{
        transform: translateY(2px);
    }
</style>
