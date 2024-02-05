<script lang="ts">
	// Import the Header, Sidebar, and Gallery components
	import Header from '../lib/Header.svelte';
  import InfoBox from '../lib/InfoBox.svelte';
	import { usernameStore } from '../lib/ratingsStore';
	import { onMount } from 'svelte';
	import { PUBLIC_API_ADDRESS } from '$env/static/public';
	import { Button, Card } from 'flowbite-svelte';
	import { ArrowRightOutline } from 'flowbite-svelte-icons';
  import { viewStore } from '$lib/viewStore';
  import { goto } from '$app/navigation';
  

	let username: string | null = null;

	type Map = {
		ID: number;
		name: string;
	};
	let maps: Map[] = [];

	onMount(async () => {
		// Get maps from the server
		const res = await fetch(`${PUBLIC_API_ADDRESS}/bathroom/maps`);
		const serverMaps = await res.json();
		maps = serverMaps;
        console.log(maps);
	});

    function handleView(id: number) {
        // Set the view id
        viewStore.set(id);

        // Redirect to the viewer
        goto('/viewer');
    }

	usernameStore.subscribe((value) => {
		username = value;
	});
</script>

<main class="flex h-screen w-screen flex-col">
	<Header />
	<nav>
		<a href="/">Gallery</a>
		<a href="/viewer">Viewer</a>
		{#if username !== ''}
			<a href="/logout">Logout</a>
			<a href="/editor">Editor</a>
			<a href="/rate">Rate</a>
		{:else}
			<a href="/login">Login</a>
			<a href="/signup">Sign Up</a>
		{/if}
	</nav>
  
  
  <InfoBox explanation="Welcome to the Bathroom Map Gallery! From this page, you can view any of the predefined maps, and have the option to calculate a voronoi approximation of them. What does this mean? Put simply, we attempt to calculate the closest bathroom from any given point on a map, and then color in all the points which are closest to that bathroom with the same color. If you'd like to create your own map, please log in and go to the editor, where you can build and save your own map of any location!" style="width: 50%;"></InfoBox>

  <div class="flex w-full flex-grow flex-col justify-center">
    {#each Array.from({ length: Math.ceil(maps.length / 3) }) as _, rowIndex}
      <div class="h-fit flex gap-4 justify-center">
        {#each maps.slice(rowIndex * 3, (rowIndex + 1) * 3) as map}
          <Card class="flex-shrink mb-4">
            <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              {map.name}
            </h5>
            <p class="mb-3 font-normal leading-tight text-gray-700 dark:text-gray-400">
              Bathroom map
            </p>
            <Button class="w-fit" on:click={() => handleView(map.ID)}>
              View <ArrowRightOutline class="ms-2 h-3.5 w-3.5 text-white" />
            </Button>
          </Card>
        {/each}
      </div>
    {/each}
  </div>
  


</main>

<style>
	/* Style for the navigation links */
	:global(body) {
		margin: 0;
	}

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

	nav a:hover,
	nav a.selectedLink {
		background-color: #555; /* Darker background color on hover */
	}
</style>
