<script lang="ts">
    export let drawn: string;
</script>

<div class="usercard">
    <img src="../{drawn}.png" alt="card">
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
