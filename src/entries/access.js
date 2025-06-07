import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';
import 'font-awesome/css/font-awesome.min.css';
import '../stylesheets/styles.css';
import Access from '../components/layouts/Access.svelte';

const app = new Access({
	target: document.body,
	props: {}
});

export default app;