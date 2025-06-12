import '../stylesheets/styles.css';
import Error from '../components/layouts/Error.svelte';

const app = new Error({
	target: document.body,
	props: {
		name: 'world'
	}
});

export default app;