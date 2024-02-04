<script>
    import { usernameStore, passwordStore } from './ratingsStore';
    import { writable } from 'svelte/store';

    let username = '';
    let password = '';

    // Store for tracking login status
    let loginSuccessful = writable(false);

    function handleSubmit() {
        // Check if the username and password are correct (you would typically validate this on the server side)
       
            usernameStore.set(username);
            passwordStore.set(password);

            // Set loginSuccessful to true
            loginSuccessful.set(true);
        
    }
</script>

<style>
    form {
        max-width: 300px;
        margin: 0 auto;
        padding: 20px;
        border: 1px solid #ccc;
        border-radius: 5px;
        background-color: #f9f9f9;
    }

    label {
        display: block;
        margin-bottom: 10px;
    }

    input[type="text"],
    input[type="password"] {
        width: 100%;
        padding: 10px;
        margin-bottom: 15px;
        border: 1px solid #ccc;
        border-radius: 3px;
        box-sizing: border-box;
    }

    button[type="submit"] {
        background-color: #880404;
        color: #fff;
        padding: 10px 15px;
        border: none;
        border-radius: 3px;
        cursor: pointer;
        font-size: 16px;
    }

    button[type="submit"]:hover {
        background-color: #460706;
    }

    /* Hide the form when loginSuccessful is true */
    form.hidden {
        display: none;
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

    /* Define a 'hidden' class to hide elements */
    .hidden {
        display: none;
    }
</style>

{#if $loginSuccessful}
<p class="login-message">Login successful!</p>
{:else}
<form on:submit|preventDefault={handleSubmit} class={$loginSuccessful ? 'hidden' : ''}>

        <div>
            <label for="username">Username:</label>
            <input type="text" id="username" bind:value={username} />
        </div>
        <div>
            <label for="password">Password:</label>
            <input type="password" id="password" bind:value={password} />
        </div>
        <button type="submit">Login</button>
    </form>
{/if}
