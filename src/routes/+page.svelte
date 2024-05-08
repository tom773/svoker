<script lang="ts">
    import Welcome from './Welcome.svelte'; 
    import '../main.css';
    import { Button } from "$lib/components/ui/button/index";
    export let data;
</script>

<div class="auth">
    {#if !data.user}
        <Button href="/signin">Sign In</Button>
    {:else}
        <form action="/signout" method="POST">
            <Button type="submit">Sign Out</Button>
        </form>
    {/if}
</div>
<main>
    <Welcome />
</main>

<style>
    .auth{
        text-align: right;
        right: 0;
        padding: 1rem;
        color: #f3f3f3;
        z-index: 100;
        position: absolute;
    }

    main{
        display: flex;
        height: 100%;
        justify-content: center;
        align-items: center;
        text-align: left;
        font-size: min(5vw, 2rem);
        color: #f3f3f3;
    }

</style>
