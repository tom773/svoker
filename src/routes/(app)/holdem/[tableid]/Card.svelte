<script lang="ts">
    export let drawn: any;
</script>

<div class="usercard">
    <img src="../{drawn['Rank']+drawn['Suit']}.png" alt="card">
</div>

<style>
    .usercard{
        position: relative;
        margin-bottom: 30px;
        object-position: top center;
        width: 100px; 
        height: 150px;
        rotate: var(--rot);
    }
</style>
