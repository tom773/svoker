<script lang="ts">
    import '../../main.css';
    import Hero from './Hero.svelte';
    import Balance from './Balance.svelte';
    import Table from './Table.svelte';
</script>

<div class="bal">
    <Balance />
</div>

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
    .bal{
        text-align: right;
        right: 0;
        padding: 1rem;
        color: #f3f3f3;
        position: absolute;
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
