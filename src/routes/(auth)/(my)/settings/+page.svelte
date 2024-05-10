<script lang="ts">
    import Profile from './(profile)/Profile.svelte';
    import Money from './(money)/Money.svelte';
    import { pagey } from '$lib/stores/user';
    export let data; 
</script>

<div>

    {#if $pagey === "profile"}
        <Profile data={data}/>
    {:else if $pagey === "security"}
        <p>Page not found</p>
    {:else if $pagey === "payment"}
        <Money data={data} /> 
    {:else if $pagey === "support"}
        <p>Page not found</p>
    {:else}
        <Profile data={data}/>
    {/if}

</div>
