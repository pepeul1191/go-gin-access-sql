import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';
import 'font-awesome/css/font-awesome.min.css';
import '../stylesheets/styles.css';
import Error from '../components/layouts/Error.svelte';

const app = new Error({
	target: document.body,
	props: {
		name: 'world'
	}
});

export default app;