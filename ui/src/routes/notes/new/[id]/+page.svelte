<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { createHttpStore } from '$lib/http/store';
	import { writable } from 'svelte/store';

	export let data: any;

	type NotesCreateResp = {
		id: string;
		data: string;
	}
	
	type NewNote = {
		note: string;
	}
	const createNoteStore = createHttpStore<NotesCreateResp>();

	const store = writable<NewNote>({
		note: data.note.data // value starts as the old data
	});

	createNoteStore.subscribe((value) => {
		if (value.ok) {
			goto(`${base}/`); // back to main screen
		} else {
			// TODO: something went wrong
		}
	});

	function handleOnSubmit() {
		createNoteStore.put(`/note/${data.note.id}`, {}, { data: $store.note });
	}
</script>

<div class="items-center content-center flex flex-col">
<h1>Edit note</h1>

<div class="card w-96 bg-base-100 shadow-xl my-4 border border-gray-200 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700">
	<div class="card-body">
		<input type="text" bind:value={$store.note} placeholder="Type here" class="input w-full max-w-xs border border-gray-500" />
		<button on:click={handleOnSubmit} class="btn btn-secondary {$createNoteStore.fetching ? 'loading' : ''} ">Submit</button>
	</div>
</div>
</div>