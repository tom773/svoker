<script lang="ts">
    import '../../../../main.css';
    import Hero from './Hero.svelte';
    import Table from './Table.svelte';
    export let data;    
</script>

<div class="hero">
    <Hero />
</div>

<main>
    <Table data={data}/> 
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
