<script lang="ts">
    import Welcome from './Welcome.svelte'; 
    import '../main.css';

</script>

<main>
    <Welcome />
</main>

<style>

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
