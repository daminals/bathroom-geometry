<script lang="ts">
	import { onMount } from 'svelte';
	import { type Bathroom } from '$lib/types';
	import ViewBathroom from '$lib/viewer/ViewBathroom.svelte';
	import { Button, Input } from 'flowbite-svelte';
	import { PUBLIC_API_ADDRESS } from '$env/static/public';

	let container: HTMLDivElement;
	let map: google.maps.Map;

	// {"name":"Stony Brook","coordinates":[{"lat":40.909119,"lng":-73.1194032},{"lat":40.907119,"lng":-73.1214032}],"bathrooms":[]}
	let grid: number[][];
	let bathrooms: Map<number, Bathroom> = new Map();
	let mapName = '';

	// Handle map initialization
	onMount(async () => {
		map = new google.maps.Map(container, {
			zoom: 17
		});
	});

	type BathroomMap = {
		name: string;
		coordinates: [
			{
				lat: number;
				lng: number;
			},
			{
				lat: number;
				lng: number;
			}
		];
		grid: number[][];
		bathrooms: Bathroom[];
	};

	function hslToHex(h: number, s: number, l: number) {
		let r, g, b;

		if (s === 0) {
			r = g = b = l; // achromatic
		} else {
			const hue2rgb = (p: number, q: number, t: number) => {
				if (t < 0) t += 1;
				if (t > 1) t -= 1;
				if (t < 1 / 6) return p + (q - p) * 6 * t;
				if (t < 1 / 2) return q;
				if (t < 2 / 3) return p + (q - p) * (2 / 3 - t) * 6;
				return p;
			};

			const q = l < 0.5 ? l * (1 + s) : l + s - l * s;
			const p = 2 * l - q;
			r = hue2rgb(p, q, h + 1 / 3);
			g = hue2rgb(p, q, h);
			b = hue2rgb(p, q, h - 1 / 3);
		}

		// Convert RGB to Hex
		const toHex = (x: number) => {
			const hex = Math.round(x * 255).toString(16);
			return hex.length === 1 ? '0' + hex : hex;
		};

		const hexColor = `#${toHex(r)}${toHex(g)}${toHex(b)}`;
		return hexColor.toUpperCase();
	}

	// Handle ID input
	let markers: Map<number, google.maps.Marker> = new Map();
	let rectangles: google.maps.Rectangle[] = [];
	async function handleSubmit(event: Event) {
		event.preventDefault();
		const form = event.target as HTMLFormElement;
		const input = form.querySelector('input') as HTMLInputElement;
		const value = input.value;

		if (value && map) {
			const json = JSON.stringify({ ID: parseInt(value) });
			const res = await fetch(`${PUBLIC_API_ADDRESS}/bathroom/maps/id`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: json
			});
			const data = (await res.json()) as BathroomMap;
			mapName = data.name;
			grid = data.grid;
			bathrooms = new Map();
			data.bathrooms.forEach((bathroom, index) => {
				// Assign color to bathroom
				const color = hslToHex(index / data.bathrooms.length, 0.8, 0.6);
				bathrooms.set(bathroom.id, { ...bathroom, color });
			});

			// Center map
			const bounds = new google.maps.LatLngBounds();
			bounds.extend(data.coordinates[0]);
			bounds.extend(data.coordinates[1]);
			map.fitBounds(bounds);

			// Draw in grid
			let xCount = grid[0].length;
			let yCount = grid.length;

			// Draw grid within the rectangle
			const north = data.coordinates[0].lat;
			const south = data.coordinates[1].lat;
			const east = data.coordinates[0].lng;
			const west = data.coordinates[1].lng;
			const latStep = (north - south) / yCount;
			const lngStep = (east - west) / xCount;
			for (let i = 0; i < yCount; i++) {
				for (let j = 0; j < xCount; j++) {
					let color = 'white';
					if (grid[i][j] == -1) {
						color = 'black';
					} else if (grid[i][j] > 0) {
						const marker = new google.maps.Marker({
							position: {
								lat: north - i * latStep - latStep / 2,
								lng: west + j * lngStep + lngStep / 2
							},
							map,
							label: {
								text: grid[i][j].toString(),
								color: 'white'
							}
						});
						markers.set(grid[i][j], marker);
						color = bathrooms.get(grid[i][j])?.color || 'white';
					}
					const rectangle = new google.maps.Rectangle({
						map,
						bounds: {
							north: north - i * latStep,
							south: north - (i + 1) * latStep,
							east: west + (j + 1) * lngStep,
							west: west + j * lngStep
						},
						fillColor: color,
						fillOpacity: color == 'white' ? 0.0 : 0.8,
						strokeWeight: 1
					});
					rectangles.push(rectangle);
				}
			}
		}
	}

	// Handle compute geometry
	type VoronoiResponse = number[][]
	async function handleCompute() {
		if (map) {
			const json = JSON.stringify({ matrix: grid });
			const res = await fetch(`${PUBLIC_API_ADDRESS}/voronoi`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: json
			});
			const matrix = await res.json() as VoronoiResponse;
			console.log(matrix);
			// Update rectangles with new colors
			for (let i = 0; i < matrix.length; i++) {
				for (let j = 0; j < matrix[i].length; j++) {
					let id = matrix[i][j];
					let color = 'white';
					if (id > 0) {
						color = bathrooms.get(id)?.color || 'white';
					}
					if (id == -1) {
						color = 'black';
					}
					rectangles[i * matrix[i].length + j].setOptions({
						fillColor: color,
						fillOpacity: color == 'white' ? 0.0 : 0.8
					});
				}
			}
		}
	}
</script>

<div class="flex h-full w-screen flex-col">
	<div class="bg-primary-800 flex justify-between p-2">
		<form on:submit={handleSubmit}>
			<Input type="text" placeholder="Enter ID"></Input>
		</form>
	</div>
	<div class="flex h-0 w-full flex-grow">
		<div class="flex h-full w-3/5 flex-col">
			<div bind:this={container} class="h-0 w-full flex-grow" />
		</div>
		<div class="flex h-full w-2/5 flex-col bg-slate-200">
			<div class="flex h-0 flex-grow flex-col gap-2 overflow-y-scroll p-2">
				{#each Array.from(bathrooms.values()) as bathroom}
					<ViewBathroom {bathroom} />
				{/each}
			</div>
			<div>
				<Button on:click={handleCompute}>Compute Geometry</Button>
			</div>
		</div>
	</div>
</div>
