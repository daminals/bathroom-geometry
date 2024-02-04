<script lang="ts">
    // Import the Header, Sidebar, and Gallery components
    import Header from '../../lib/Header.svelte';
    import Signup from '../../lib/Signup.svelte';
    import { usernameStore} from '../../lib/ratingsStore';
    let username: string | null = null; // 
  
usernameStore.subscribe(value => {
    username = value;
});
console.log(username)
</script>


<main>
    <!-- Your main content goes here -->
  
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
    <div style="margin-top: 20px;">

            <Signup />
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
    nav a.selectedLink {
        background-color: #555; /* Darker background color on hover */
    }
   
</style>
