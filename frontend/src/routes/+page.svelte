<script lang="ts">
    // Import the Header, Sidebar, and Gallery components
    import Header from '../lib/Header.svelte';
    import SidebarGallery from '../lib/SidebarGallery.svelte';
    import { usernameStore } from '../lib/ratingsStore';
    import { onMount } from 'svelte';

    let username: string | null = null;

    usernameStore.subscribe(value => {
        username = value;
    });
</script>

<main class='h-screen w-screen flex flex-col'>
    <Header />
    <nav>
        <a href="/">Gallery</a>
        {#if username !== null}
            <!-- If the user is logged in, show logout and other links -->
            <a href="/logout">Logout</a>
            <a href="/new">New</a>
            <a href="/edit">Edit</a>
            <a href="/rate">Rate</a>
        {:else}
            <!-- If the user is not logged in, show login and signup links -->
            <a href="/login">Login</a>
            <a href="/signup">Sign Up</a>
        {/if}
    </nav>
    <div class="h-0 flex-grow w-full flex">
        <div class="bg-[#333] text-white w-64">
            <SidebarGallery />
        </div>
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
