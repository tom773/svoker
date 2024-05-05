<script lang="ts">
    import '../../main.css';
    import Hero from './Hero.svelte';
    import Table from './Table.svelte';
</script>

<div class="hero">
    <Hero />
</div>

<main>
    <Table /> 
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
        font-family: 'jbm';
        color: #f3f3f3;
    }
</style>
