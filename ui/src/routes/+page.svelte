<script lang="ts">
	import { createHttpStore } from '$lib/http/store';
	import { onMount } from 'svelte';
	import { base } from '$app/paths';

	let content = '';

	type NotesCreateResp = {
		id: string;
		data: string;
	};
	type NotesQueryResp = [NotesCreateResp];

	export const notesStore = createHttpStore<NotesQueryResp>();

	onMount(async () => {
		notesStore.get('/notes');
	});

	async function deleteNote(id: string) {
		requestAndUpdateNotes('DELETE', '/note/' + id, '', {});
	}

	function requestAndUpdateNotes(method: string, path: string, params?: string, data?: object) {
		let url = '/api/v1' + path;
		let body = JSON.stringify(data);
		const headers = {
			'Content-type': 'application/json'
		};

		fetch(url, {
			method,
			body,
			headers,
			// TODO: this could be same-origin when running on the same site
			credentials: 'include'
		})
			.then((response) => {
				if (response.ok) {
					response
						.json()
						.then((data) => {
							// bit of a hack, but fetch the notes again to update the store
							notesStore.get('/notes');
						})
						.catch((error) => {
							throw new Error('Error converting response to JSON');
						});
				} else {
					response
						.text()
						.then((text) => {
							notesStore.get('/notes');
						})
						.catch((error) => {
							throw new Error(error);
						});
				}
			})
			.catch((error) => {
				throw new Error(error);
			});
	}
</script>


	
<div class="text-center">
	<h1>Welcome to VeriNotes!</h1>
	<!-- <p>Visit <a href="https://verifa.io">verifa.io</a> to learn more about Verifa</p> -->

	<a href="{base}/notes/new" class="btn btn-primary w-96 text-white bg-black hover:bg-gray-800"
		>New note</a
	>
</div>

<div class="flex flex-col items-center text-left mt-4">
	{#if $notesStore.fetching}
		<h2>Loading notes</h2>
	{:else if $notesStore.ok && $notesStore.data}
			{#each $notesStore.data as note}
				<div
					class="card w-96 bg-base-100 shadow-xl my-4 border border-gray-200 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700"
				>
					<div class="card-body">
						<h2 class="card-title">{'#' + note.id}</h2>
						<p>{note.data}</p>
						<div class="card-actions justify-end">
							<button
								on:click={deleteNote(note.id)}
								class="btn btn-primary bg-red-400 hover:bg-red-500">Delete</button
							>
							<a
								href="{base}/notes/new/{note.id}"
								class="btn btn-primary bg-green-400 hover:bg-green-500">Edit</a
							>
						</div>
					</div>
				</div>
			{/each}
	{/if}
</div>




