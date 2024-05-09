<script lang="ts">
    
    import Profile from './(profile)/Profile.svelte';
    import { pagey } from '$lib/stores/user';
    
</script>

<div>

    {#if $pagey === "profile"}
        <Profile />
    {:else if $pagey === "security"}
        <p>Page not found</p>
    {:else if $pagey === "payment"}
        <p>Page not found</p>
    {:else if $pagey === "support"}
        <p>Page not found</p>
    {:else}
        <Profile />
    {/if}

</div>
