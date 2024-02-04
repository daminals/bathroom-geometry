<script lang="ts">
    import Rate from './Rate.svelte';
    import { onMount } from 'svelte';
    import { comments, accessibilityRating, cleanRating, menstrualRating, overallRating } from './ratingsStore';

    let showForm = false;
    let currentIndex = 0;
    let commentsData: { text: string; user: string }[] = [];

    let images = [
        { name: 'Jasmine', src: 'jasmine.jpg' },
        { name: 'SAC', src: 'sac.jpg' }
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

    onMount(() => {
        const unsubscribe = comments.subscribe((value) => {
            commentsData = value;
        });
        return unsubscribe;
    });
</script>

<div class="gallery-container">
    {#if !showForm}
        <div class="gallery">
            <h1>{images[currentIndex].name}</h1>
            <div class="ratings-container">
                <h2>Ratings</h2>
                <p>Accessibility Rating: {$accessibilityRating.toFixed(0)}</p>
                <p>Clean Rating: {$cleanRating.toFixed(0)}</p>
                <p>Menstrual Rating: {$menstrualRating.toFixed(0)}</p>
                <p>Overall Rating: {$overallRating.toFixed(0)}</p>
            </div>
            {#each images as { src }, index}
                <img
                    {src}
                    alt={images[currentIndex].name}
                    class="gallery-image"
                    style={index === currentIndex ? 'display: block;' : 'display: none;'}
                />
            {/each}
            <div class="nav-buttons">
                <button class="prev-btn" on:click={prevImage}>&lt;</button>
                <button class="next-btn" on:click={nextImage}>&gt;</button>
                <button class="rate-btn" on:click={handleRate}>Rate</button>
            </div>
        </div>
        <div class="comments-section">
            <h3>Comments</h3>
            {#if commentsData.length > 0}
                {#each commentsData as comment}
                    <div class="comment">
                        <p>{comment.text}</p>
                    </div>
                {/each}
            {:else}
                <p>No comments available</p>
            {/if}
        </div>
    {/if}

    {#if showForm}
        <div class="rate-form">
            <Rate />
        </div>
    {/if}
</div>

<style>

h1 {
	font-size: 20px;
	font-weight: bold;
}

h2 {
	font-size: 18px;
}
    /* Modern and Stylish CSS Styles */
    .gallery-container {
        font-family: 'Roboto', sans-serif;
        position: relative;
        max-width: 600px;
        margin: 0 auto;
		
    }

    .gallery {
        display: flex;
        flex-direction: column;
        align-items: center;
        background-color: #f9f9f9;
        padding: 20px;
        border-radius: 10px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }

    .ratings-container {
        margin-bottom: 20px;
    }

    .gallery-image {
        width: 100%;
        max-width: 400px;
        height: auto;
        border-radius: 8px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        margin-bottom: 20px;
    }

    .nav-buttons {
        margin-top: 10px;
        display: flex;
        align-items: center;
    }

    .prev-btn,
    .next-btn,
    .rate-btn {
        background-color: #4CAF50;
        color: white;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        padding: 10px 20px;
        margin: 0 10px;
        transition: background-color 0.3s ease;
    }

    .prev-btn:hover,
    .next-btn:hover,
    .rate-btn:hover {
        background-color: #45a049;
    }

    .comments-section {
        margin-top: 20px;
    }

    .comment {
        background-color: #fff;
        border: 1px solid #ddd;
        border-radius: 8px;
        padding: 10px;
        margin-bottom: 10px;
    }

    .comment p {
        margin: 0;
    }
</style>
