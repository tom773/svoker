<script lang="ts">
    import '../../../../main.css';
    import Hero from './Hero.svelte';
    import Table from './Table.svelte';
    import { page } from '$app/stores';
    import { onMount } from 'svelte'; 
    import { initWS } from '$lib/stores/websocket';

    export let data: any;
    const tableid = $page.params.tableid;
    
    onMount(() => {
        initWS("ws://localhost:8080/ws?id="+data.user.id);
    });
</script>

<div class="hero">
    <Hero data={data} tableid={tableid}/>
</div>

<main>
    <Table data={data} tableid={tableid}/> 
</main>

<style>
    .hero{
        margin: auto;
        width: 100%;
        position: absolute;
        padding: 1rem;
        color: #f3f3f3;
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
