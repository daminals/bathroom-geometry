// ratingsStore.js

import { writable } from 'svelte/store';

// Initialize the stores for each category with initial values of 0
export const accessibilityRating = writable(0);
export const cleanRating = writable(0);
export const menstrualRating = writable(0);
export const overallRating = writable(0);

interface Comment {
    text: string;
    // You can add more properties as needed, such as timestamp, user info, etc.
}

// Create a writable store for comments
export const comments = writable<Comment[]>([]);

// Function to add a new comment to the store
export function addComment(comment: Comment) {
    comments.update(existingComments => [...existingComments, comment]);
}