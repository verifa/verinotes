export async function load({ fetch, params }) {
    const res = await fetch(`/api/v1/note/${params.id}`);
    const note = await res.json();
   
    return { note };
}