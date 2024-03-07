import { Router } from 'itty-router';
import handleProxy from './proxy';

// now let's create a router (note the lack of "new")
const router = Router();

// GET /api/v1/github/id?name=fileName
router.get('/api/v1/github/:id', ({ query, params }) => {
	const { name, user = 'Paxxs' } = query;
	const { id } = params;

	if (!name)
		return new Response('Not Found!', {
			status: 404,
		});
	const targetURL = `https://gist.githubusercontent.com/${user}/${id}/raw/${name}`;
	return handleProxy.fetch(targetURL);
});

// GET item
router.get('/api/todos/:id', ({ params }) => new Response(`Todo #${params.id}`));

// POST to the collection (we'll use async here)
router.post('/api/todos', async (request) => {
	const content = await request.json();

	return new Response('Creating Todo: ' + JSON.stringify(content));
});

// 404 for everything else
router.all('*', () => new Response('Not Found.', { status: 404 }));

export default router;
