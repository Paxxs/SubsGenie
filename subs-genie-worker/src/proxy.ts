export default {
	async fetch(proxyUrl: string): Promise<Response> {
		if (!proxyUrl) {
			return new Response('Bad request: Missing `proxyUrl`', { status: 400 });
		}
		let res = await fetch(proxyUrl);
		return res;
	},
};
