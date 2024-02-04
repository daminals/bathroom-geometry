<!-- Logout.svelte -->
<script>
    import { writable } from 'svelte/store';
    import {logoutStatus, usernameStore, passwordStore} from './ratingsStore';
    let logoutSuccessful = writable(false);
    function handleLogout() {
        console.log("Logout successful");
        logoutSuccessful.set(true);
        logoutStatus.set(true);
        usernameStore.set("")
        passwordStore.set("")
   
    }

</script>

<style>
 .b-container {
        display: flex;
        justify-content: center; /* Horizontally center the button */
        margin-top: 20px; /* Add margin for spacing */
    }
    .b {
        background-color: #880404;
        color: #fff;
        padding: 10px 15px;
        border: none;
        border-radius: 3px;
        cursor: pointer;
        font-size: 16px;
    }

    .b:hover {
        background-color: #460706;
    }
    .login-message {
        display: flex;
        flex-direction: column;
        align-items: center; /* Center items horizontally */
        height: 100vh; 
        color: #4CAF50; /* Green color for success message */
        font-size: 18px;
        font-weight: bold;
        margin-top: 10px;
    }

   
</style>

{#if $logoutSuccessful}
    <!-- Show the success message only when logout is successful -->
    <p class="login-message">Logout successful!</p>
{:else}
    <!-- Show the logout button and form when not logged out -->
    <div class="b-container">
        <button class="b" on:click={handleLogout}>Logout</button>
    </div>
{/if}
