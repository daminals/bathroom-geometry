<script lang="ts">
	// Import the Header, Sidebar, and Gallery components
	import Header from '../../lib/Header.svelte';
	import Viewer from '$lib/viewer/Viewer.svelte';

	import { usernameStore } from '$lib/ratingsStore';
	import { viewStore } from '$lib/viewStore';

    let username: string | null = null;
	usernameStore.subscribe(value => {
        username = value;
    });

	let view: number = -1;
	viewStore.subscribe(value => {
		view = value;
	});
</script>

<main class="h-screen w-screen flex flex-col">
	<!-- Your main content goes here -->

	<Header />
	<nav>
		<a href="/">Gallery</a>
		<a href="/viewer">Viewer</a>
    <a href="/editor">Editor</a>
		{#if username !== ''}
			<a href="/logout">Logout</a>
			<a href="/rate">Rate</a>
		{:else}
			<a href="/login">Login</a>
			<a href="/signup">Sign Up</a>
		{/if}
	</nav>
	<div class="h-0 flex-grow w-full flex">
		{#if view !== -1}
			<Viewer id={view} />
		{:else}
			<p>Nothing to see here</p>
		{/if}
	</div>
</main>

<style>
	:global(body) {
		margin: 0;
	}
	/* Style for the navigation links */
	nav {
		background-color: #333; /* Dark background color */
		padding: 10px 0; /* Add padding to the top and bottom */
	}

	nav a {
		color: #fff; /* White text color */
		text-decoration: none; /* Remove underline */
		padding: 10px 20px; /* Add padding to the links */
		margin-right: 10px; /* Add margin between links */
		border-radius: 5px; /* Add border radius for rounded corners */
		transition: background-color 0.3s ease; /* Smooth transition on hover */
	}

	nav a:hover {
		background-color: #555; /* Darker background color on hover */
	}
</style>
