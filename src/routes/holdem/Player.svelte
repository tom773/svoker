<script lang="ts">
    
    export let name: string;
    export let chips: number;
    export let avatar: string;

</script>

<div class="player1 flex flex-row">
    <div class="avatar">
        <img src={avatar} alt="avatar" />
    </div>
    <div class="flex flex-col">
        <p>{name}</p>
        <p>{chips}</p>
    </div>
</div>

<style>
    .player1{
        align-items: center;
        justify-content: center;
        margin: 10px;
        margin-left: 5px;
        font-size: 15px;
        text-align: center;
    }
    
    .avatar{
        margin-top: -10px;
        margin-bottom: 40px;
        width: 50px;
        height: 50px;
        border-radius: 50%;
    }

</style>
