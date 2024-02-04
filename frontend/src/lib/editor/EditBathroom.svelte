<script lang="ts">
	import { Label, Input, Radio, Checkbox, Card } from 'flowbite-svelte';
	import { type Bathroom } from '$lib/types';

	const genderOptions = ['F', 'M', 'U'];
	export let bathroom: Bathroom;
</script>

<Card class="max-w-full">
	<div class="text-sbu font-bold mb-2">Bathroom #{bathroom.id}</div>
	<Label for="name">Name</Label>
	<Input type="text" id="name" bind:value={bathroom.name} />
	<Label for="Gender">Gender</Label>
	<div class="flex gap-3">
		{#each genderOptions as gender}
			<Radio bind:group={bathroom.gender} value={gender}>{gender}</Radio>
		{/each}
	</div>
	<Label for="accessible">Accessible</Label>
	<Checkbox bind:checked={bathroom.accessible} />
	<Label for="menstrualProducts">Menstrual Products</Label>
	<Checkbox bind:checked={bathroom.menstrualProducts} />
</Card>
