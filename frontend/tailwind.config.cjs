const config = {
	content: ['./src/**/*.{html,js,svelte,ts}', './node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'],

	plugins: [require('flowbite/plugin')],

	darkMode: 'class',

	theme: {
		extend: {
			colors: {
				// flowbite-svelte
				primary: {
					50: "#FFDBDB",
					100: "#FFB8B8",
					200: "#FF7070",
					300: "#FF2929",
					400: "#E00000",
					500: "#990000",
					600: "#7A0000",
					700: "#5C0000",
					800: "#3D0000",
					900: "#1F0000",
					950: "#0F0000"
				},

			}
		}
	}
};

module.exports = config;