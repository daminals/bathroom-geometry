<script>
    // Import the Header, Sidebar, and Gallery components
    import Header from '../lib/Header.svelte';
    import Sidebar from '../lib/Sidebar.svelte';
    import Gallery from '../lib/Gallery.svelte'; // Adjust the path accordingly
</script>

<main>
    <!-- Your main content goes here -->
    <Header />
    <div class="content">
        <nav class="navigation-links">
            <!-- Add more navigation links as needed -->
        </nav>
        <div class="sidebar-content">
            <Sidebar />
        </div>
        <Gallery /> <!-- Include the Gallery component -->
    </div>
</main>

<style>
    .content {
        display: flex;
        align-items: flex-start; /* Align items at the start of the cross axis */
    }

    .sidebar-content {
        display: flex;
    }


</style>
