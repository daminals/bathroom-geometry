<script>
    // Import the Header, Sidebar, and Gallery components
    import Header from '../../lib/Header.svelte';
    import SidebarEdit from '../../lib/SidebarEdit.svelte';

</script>


<main>
    <!-- Your main content goes here -->
  
    <Header />
	<nav>
		<a href="/">Gallery</a>
        <a href="/new">New</a>
        <a href="/edit">Edit</a>
        <a href="/rate">Rate</a>
    </nav>
    <div class="content">
        <div class="sidebar-content">
            <SidebarEdit />
        </div>
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

    /* Adjustments to make the content align properly */
    .content {
        display: flex;
        align-items: flex-start; /* Align items at the start of the cross axis */
    }
</style>
