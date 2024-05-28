<script lang="ts">
    import Profile from './(profile)/Profile.svelte';
    import Money from './(money)/Money.svelte';
    import Security from './(security)/Security.svelte';
    import Support from './(support)/Support.svelte';
    import { pagey } from '$lib/stores/user';
    export let data; 
</script>

<div>

    {#if $pagey === "profile"}
        <Profile data={data}/>
    {:else if $pagey === "security"}
        <Security data={data} />
    {:else if $pagey === "payment"}
        <Money data={data} /> 
    {:else if $pagey === "support"}
        <Support data={data} />
    {:else}
        <Profile data={data}/>
    {/if}

</div>
