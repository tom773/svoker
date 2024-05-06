<script lang="ts">
    import { highest_card } from '$lib/pkg/pokerutil';
    import '../../main.css';
    import Hero from './Hero.svelte';
    import Table from './Table.svelte';
    
    function highCard(){
        let cards = ['2C', 'AS', '3C', 'KD'];
        let highcard = highest_card(cards);
        console.log(highcard);
    }
    
</script>

<button type="button" on:click={highCard} class="btn variant-filled">Click Me!</button>

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
