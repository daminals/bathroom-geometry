<script lang="ts">
	import Gallery from './RatePage.svelte';
	import {
		accessibilityRating,
		cleanRating,
		comments,
		menstrualRating,
		overallRating, 
		addComment
	} from './ratingsStore';

	let selectedAccessibilityRating: number | null = null;
	let selectedCleanRating: number | null = null;
	let selectedMenstrualRating: number | null = null;
	let selectedOverallRating: number | null = null;
	let showForm = true; // Initially show the form
    let feedbackText: string = '';

	function handleNutritionRatingClick(rating: number | null) {
		if (rating !== null) {
			selectedAccessibilityRating = rating;
		}
	}

	function handleAffordableRatingClick(rating: number | null) {
		if (rating !== null) {
			selectedCleanRating = rating;
		}
	}

	function handleProcessRatingClick(rating: number | null) {
		if (rating !== null) {
			selectedMenstrualRating = rating;
		}
	}

	function handleOverallRatingClick(rating: number | null) {
		if (rating !== null) {
			selectedOverallRating = rating;
		}
	}

	function request() {
		// Check if all ratings are provided
		if (
			selectedAccessibilityRating === null ||
			selectedCleanRating === null ||
			selectedMenstrualRating === null ||
			selectedOverallRating === null
		) {
			alert('Please answer all the questions before submitting.');
		} else {
			alert('Thank you, we have received your feedback!');
			showForm = false; // Hide the form after submission

			// Calculate the average by adding the previous rating and the new rating, then dividing by 2
			selectedAccessibilityRating = ($accessibilityRating + selectedAccessibilityRating) / 2;
			selectedCleanRating = ($cleanRating + selectedCleanRating) / 2;
			selectedMenstrualRating = ($menstrualRating + selectedMenstrualRating) / 2;
			selectedOverallRating = ($overallRating + selectedOverallRating) / 2;

			// Update the rating stores with new values
			accessibilityRating.set(selectedAccessibilityRating);
			cleanRating.set(selectedCleanRating);
			menstrualRating.set(selectedMenstrualRating);
			overallRating.set(selectedOverallRating);
			const comment = {
                text: feedbackText,
                user: 'Your User ID' // You need to define a user identifier here
            };
            addComment(comment);

			const feedbackData = {
            accessibilityRating: selectedAccessibilityRating,
            cleanRating: selectedCleanRating,
            menstrualRating: selectedMenstrualRating,
            overallRating: selectedOverallRating,
            feedbackText: feedbackText // Include the feedback text
        };

        console.log('Feedback Data:', feedbackData);
		}
	}
</script>

<!-- Rest of the component remains the same -->
<div class="Eligibility">
	<div class="w-full mx-auto">
		{#if showForm}
			<h2 class="header">Give us your feedback!</h2>
			<h3 class="questionsContainer">How accessible is this restroom?</h3>
			{#each [1, 2, 3, 4, 5] as rating}
				<button
					class={selectedAccessibilityRating === rating
						? 'ratingContainer selected'
						: 'ratingContainer'}
					on:click={() => handleNutritionRatingClick(rating)}>{rating}</button
				>
			{/each}

			<h3 class="questionsContainer">How clean is this restroom?</h3>
			{#each [1, 2, 3, 4, 5] as rating}
				<button
					class={selectedCleanRating === rating ? 'ratingContainer selected' : 'ratingContainer'}
					on:click={() => handleAffordableRatingClick(rating)}>{rating}</button
				>
			{/each}

			<h3 class="questionsContainer">Are there menstrual products available at this restroom?</h3>
			{#each [1, 2, 3, 4, 5] as rating}
				<button
					class={selectedMenstrualRating === rating
						? 'ratingContainer selected'
						: 'ratingContainer'}
					on:click={() => handleProcessRatingClick(rating)}>{rating}</button
				>
			{/each}

			<h3 class="questionsContainer">How satisfied are you with this restroom?</h3>
			{#each [1, 2, 3, 4, 5] as rating}
				<button
					class={selectedOverallRating === rating ? 'ratingContainer selected' : 'ratingContainer'}
					on:click={() => handleOverallRatingClick(rating)}>{rating}</button
				>
			{/each}
            <div>
                <label for="feedback">Additional Comments:</label><br>
                <textarea id="feedback" rows="4" cols="50" bind:value={feedbackText}></textarea>
            </div>

			<button class="submitContainer" on:click={request}>Submit</button>
		{/if}

		{#if !showForm}
			<div class="rate-form">
				<Gallery />
			</div>
		{/if}
	</div>
</div>

<style>
	.Eligibility {
		/* Styles for the outer container */
		display: flex; /* Enable flexbox layout */
		justify-content: center; /* Center the content horizontally */
	}
	.header {
		position: relative;
		max-width: 600px;
		margin: 0 auto;
		overflow: hidden;
	}
	.questionsContainer {
		position: relative;
		max-width: 600px;
		margin: 0 auto;
		overflow: hidden;
	}
	.ratingContainer {
		background-color: #4caf50; /* Green */
		border: none;
		color: white;
		padding: 15px 32px;
		text-align: center;
		text-decoration: none;
		display: inline-block;
		font-size: 16px;
		margin: 4px 2px;
		cursor: pointer;
		border-radius: 20px;
	}
	.ratingContainer:hover {
		background-color: #880404;
	}
	.ratingContainer:active {
		transform: translateY(2px);
	}

	.ratingContainer.selected {
		background-color: #880404; /* Changed color when selected */
	}

	.submitContainer {
		display: block; /* Set the button to display as a block element */
		margin-top: 20px; /* Add some top margin to separate it from the previous elements */
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

	.submitContainer:hover {
		background-color: #460706;
	}
</style>
