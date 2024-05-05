<script lang="ts">
    import Welcome from './Welcome.svelte'; 
    import '../main.css';

</script>

<div class="auth">
    <a href="/signin"><button type="button" class="btn variant-filled">Sign In</button></a>
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
        font-family: 'jbm';
        color: #f3f3f3;
    }

</style>
