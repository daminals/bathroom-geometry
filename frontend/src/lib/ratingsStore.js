// ratingsStore.js

import { writable } from 'svelte/store';

// Initialize the average rating store with an initial value of 0
export const averageRating = writable(0);
