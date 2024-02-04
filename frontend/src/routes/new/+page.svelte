<script lang="ts">
	import { PUBLIC_MAPS_KEY } from '$env/static/public';
	import { onMount } from 'svelte';

	let container: HTMLDivElement;
	let map: google.maps.Map;
	let zoom = 17;
	let marker1: google.maps.Marker;
	let marker2: google.maps.Marker;
	let rect: google.maps.Rectangle;

	let mode = 'Search';
	let nextId = 1;

	// Handle map initialization
	onMount(async () => {
		map = new google.maps.Map(container, {
			zoom
		});

		// Create draggable markers
		marker1 = new google.maps.Marker({
			position: { lat: 0, lng: 0 },
			map,
			draggable: true
		});
		marker2 = new google.maps.Marker({
			position: { lat: 0, lng: 0 },
			map,
			draggable: true
		});

		// Draw rectangle using the markers
		rect = new google.maps.Rectangle({
			map,
			bounds: {
				north: 0,
				south: 0,
				east: 0,
				west: 0
			}
		});

		// Add event listener to markers
		marker1.addListener('dragend', updateRect);
		marker2.addListener('dragend', updateRect);
	});

	// Handle location search
	function handleSubmit(event: Event) {
		event.preventDefault();
		const form = event.target as HTMLFormElement;
		const input = form.querySelector('input') as HTMLInputElement;
		const value = input.value;

		if (value && map) {
			// Search for the place
			var service = new google.maps.places.PlacesService(map);
			service.textSearch(
				{
					query: value
				},
				(results, status) => {
					if (status === google.maps.places.PlacesServiceStatus.OK && results) {
						const place = results[0];
						const location = place.geometry?.location;
						if (location) {
							map.setCenter(location);
							marker1.setPosition({ lat: location.lat() + 0.001, lng: location.lng() - 0.001 });
							marker2.setPosition({ lat: location.lat() - 0.001, lng: location.lng() + 0.001 });
							updateRect();
						}
					}
				}
			);
		}
	}

	// Update rectangle bounds
	function updateRect() {
		const pos1 = marker1.getPosition();
		const pos2 = marker2.getPosition();
		if (pos1 && pos2) {
			rect.setBounds({
				north: Math.max(pos1.lat(), pos2.lat()),
				south: Math.min(pos1.lat(), pos2.lat()),
				east: Math.max(pos1.lng(), pos2.lng()),
				west: Math.min(pos1.lng(), pos2.lng())
			});
		}
	}

	type Bathroom = {
		id: number;
		name: string;
	};

	// Handle go button click (draw grid)
	let grid: number[][];
	let bathrooms: Map<number, Bathroom> = new Map();
	let rectangles: google.maps.Rectangle[] = [];
	let markers: Map<number, google.maps.Marker> = new Map();
	function handleGo() {
		mode = 'Draw';
		let bounds = rect.getBounds();
		if (bounds) {
			// Calculate number of recetagles (10 meters = 1 recentagle)
			let distanceX = google.maps.geometry.spherical.computeDistanceBetween(
				new google.maps.LatLng(bounds.getNorthEast().lat(), bounds.getNorthEast().lng()),
				new google.maps.LatLng(bounds.getNorthEast().lat(), bounds.getSouthWest().lng())
			);
			let distanceY = google.maps.geometry.spherical.computeDistanceBetween(
				new google.maps.LatLng(bounds.getNorthEast().lat(), bounds.getNorthEast().lng()),
				new google.maps.LatLng(bounds.getSouthWest().lat(), bounds.getNorthEast().lng())
			);
			let xCount = Math.ceil(distanceX / 10);
			let yCount = Math.ceil(distanceY / 10);

			// Draw grid within the rectangle
			const north = bounds.getNorthEast().lat();
			const south = bounds.getSouthWest().lat();
			const east = bounds.getNorthEast().lng();
			const west = bounds.getSouthWest().lng();
			const latStep = (north - south) / yCount;
			const lngStep = (east - west) / xCount;

			// Hide markers + rectangle
			marker1.setMap(null);
			marker2.setMap(null);
			rect.setMap(null);

			// Initialize grid
			grid = Array(xCount)
				.fill(0)
				.map(() => Array(yCount).fill(0));
			for (let i = 0; i < yCount; i++) {
				for (let j = 0; j < xCount; j++) {
					const rectangle = new google.maps.Rectangle({
						map,
						bounds: {
							north: north - i * latStep,
							south: north - (i + 1) * latStep,
							east: west + (j + 1) * lngStep,
							west: west + j * lngStep
						},
						fillColor: 'white',
						fillOpacity: 0.1,
						strokeWeight: 1
					});
					rectangle.addListener('click', () => {
						handleRectClick(j, i);
					});
					rectangles.push(rectangle);
				}
			}
		}
	}

	// Handle rectangle click
	function handleRectClick(x: number, y: number) {
		// Get the rectangle
		const rectangle = rectangles[y * grid.length + x];

		// Check if in add mode
		if (mode === 'Add') {
			// Delete bathroom if exists
			let id = grid[x][y];
			if (id > 0) {
				// Remove bathroom
				bathrooms.delete(id);
				bathrooms = bathrooms;

				// Remove marker
				const marker = markers.get(id);
				if (marker) {
					marker.setMap(null);
					markers.delete(id);
				}

				// Update grid
				grid[x][y] = 0;
				return;
			}

			// Add bathroom
			id = nextId++;
			bathrooms.set(id, {
				id,
				name: `Bathroom ${id}`
			});
			bathrooms = bathrooms;
			grid[x][y] = id;

			// Create marker
			const bounds = rectangle.getBounds();
			if (bounds) {
				const center = bounds.getCenter();
				const marker = new google.maps.Marker({
					position: center,
					map,
					label: id.toString()
				});
				markers.set(id, marker);
			}
			return;
		}

		// Get current state
		const filled = grid[x][y] === -1;

		if (filled) {
			// Unfill the rectangle
			rectangle.setOptions({
				fillColor: 'white',
				fillOpacity: 0.1
			});
			grid[x][y] = 0;
		} else {
			// Fill the rectangle
			rectangle.setOptions({
				fillColor: 'black',
				fillOpacity: 1
			});
			grid[x][y] = -1;
		}
	}
</script>

<svelte:head>
	<script
		defer
		async
		src="https://maps.googleapis.com/maps/api/js?key={PUBLIC_MAPS_KEY}&libraries=places"
	>
	</script>
</svelte:head>

<div class="container">
	<div class="header">
		<form on:submit={handleSubmit}>
			<input type="text" placeholder="Search for a place" />
			<button type="submit">Search</button>
		</form>
		<div>Mode: {mode}</div>
		{#if mode === 'Search'}
			<button on:click={handleGo}>Go</button>
		{:else}
			<button on:click={() => (mode = mode === 'Draw' ? 'Add' : 'Draw')}>
				Toggle Draw / Add Mode
			</button>
		{/if}
	</div>
	<div bind:this={container} class="map"></div>
	<div>
		<div>
			{#each Array.from(bathrooms.values()) as bathroom}
				<div>{bathroom.name}</div>
			{/each}
		</div>
	</div>
</div>

<style>
	:global(body) {
		margin: 0;
	}
	.header {
		display: flex;
		justify-content: center;
	}
	.container {
		width: 100vw;
		height: 100vh;
		display: flex;
		flex-direction: column;
	}
	.map {
		width: 100%;
		height: 0%;
		flex-grow: 1;
	}
</style>
