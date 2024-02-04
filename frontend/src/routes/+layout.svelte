<script>
	import '../app.pcss';

	import { PUBLIC_MAPS_KEY } from '$env/static/public';
</script>

<svelte:head>
	<script
		defer
		async
		src="https://maps.googleapis.com/maps/api/js?key={PUBLIC_MAPS_KEY}&libraries=places,geometry,marker"
	>
	</script>
</svelte:head>

<slot />
