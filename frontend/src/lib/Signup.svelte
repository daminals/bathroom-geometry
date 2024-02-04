<script>
    import { usernameStore, passwordStore } from './ratingsStore';
    import { writable } from 'svelte/store';

    let username = '';
    let password = '';
    let confirmPassword = '';

    // Store for tracking sign-up status
    let signUpSuccessful = writable(false);

    function handleSignUp() {
        // Check if the username and password are valid (you would typically validate this on the server side)
        if (username !== '' && password !== '' && password === confirmPassword) {
            // Set username and password in the store
            usernameStore.set(username);
            passwordStore.set(password);

            // Set signUpSuccessful to true
            signUpSuccessful.set(true);
        } else {
            alert("Passwords do not match!");
        }
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

{#if $signUpSuccessful}
    <p class="login-message">Sign-up successful!</p>
{:else}
    <form on:submit|preventDefault={handleSignUp}>

        <div>
            <label for="username">Username:</label>
            <input type="text" id="username" bind:value={username} />
        </div>
        <div>
            <label for="password">Password:</label>
            <input type="password" id="password" bind:value={password} />
        </div>
        <div>
            <label for="confirmPassword">Confirm Password:</label>
            <input type="password" id="confirmPassword" bind:value={confirmPassword} />
        </div>
        <button type="submit">Sign Up</button>
    </form>
{/if}
