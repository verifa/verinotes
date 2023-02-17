<script lang="ts">
	import { createHttpStore } from '$lib/http/store';
	import { onMount } from 'svelte';

	type NotesResp = [{
		id: string;
		data: string;
	}]

	const notesStore = createHttpStore<NotesResp>();
	onMount(async () => {
		console.log("get notes!")
		notesStore.get('/notes');
	})

</script>

<h1>Welcome to VeriNotes!</h1>
<!-- <p>Visit <a href="https://verifa.io">verifa.io</a> to learn more about Verifa</p> -->

{#if $notesStore.fetching}
	<h2>Loading projects</h2>
{:else if $notesStore.ok && $notesStore.data}
	<div class="overflow-x-auto w-full">
		<table class="table table-auto w-full">
			<tbody>
				{#each $notesStore.data as note}
					<tr>
						<div class="card w-96 bg-base-100 shadow-xl my-4 border border-gray-200 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700">
							<div class="card-body">
								<h2 class="card-title">Note!</h2>
								<p>{note.data}</p>
								<div class="card-actions justify-end">
								<button class="btn btn-primary">Click!</button>
								</div>
							</div>
						</div>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
{/if}


<!-- TODO: button to add notes <a href="{base}/projects/new" class="btn btn-primary">New Project</a> -->