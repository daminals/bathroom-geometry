<script lang="ts">
	import { Label, Input, Radio, Checkbox, Card } from 'flowbite-svelte';
	import { type Bathroom } from '$lib/types';

	const genderOptions = ['F', 'M', 'U'];
	export let bathroom: Bathroom;
</script>

<Card class="max-w-full text-grey-900 text-sm font-medium">
	<div class="text-sbu font-bold mb-2">Bathroom #{bathroom.id} - {bathroom.name}</div>
	<div>
		Gender
	</div>
	<div class="flex gap-3">
		{#each genderOptions as gender}
			<Radio bind:group={bathroom.gender} value={gender} disabled>{gender}</Radio>
		{/each}
	</div>
	<div>Accessible</div>
	<Checkbox bind:checked={bathroom.accessible} disabled />
	<div>Menstrual Products</div>
	<Checkbox bind:checked={bathroom.menstrualProducts} disabled />
</Card>
