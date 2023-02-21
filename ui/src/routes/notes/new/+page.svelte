<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { createHttpStore } from '$lib/http/store';
	import { writable } from 'svelte/store';

	type Note = {
		data: '';
	}
	type NewNote = {
		data: '';
	}


	const createNoteStore = createHttpStore<Note>();

	const store = writable<NewNote>({
		data: '',
	});


	createNoteStore.subscribe((value) => {
		if (value.ok) {
			goto(`${base}/`); // back to home when note is created
		} else {
			// TODO: something went wrong
		}
	});

	function handleOnSubmit() {
		createNoteStore.post('/note', {}, $store);
	}

</script>

<h1>Create note</h1>

<div class="card w-96 bg-base-100 shadow-xl my-4 border border-gray-200 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700">
	<div class="card-body">
		<input type="text" bind:value={$store.data} placeholder="Type here" class="input w-full max-w-xs border border-gray-500" />
		<button on:click={handleOnSubmit} class="btn btn-secondary">Add</button>
	</div>
</div>