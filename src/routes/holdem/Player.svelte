<script lang="ts">
    
    export let name: string;
    export let chips: number;
    export let avatar: string;
    import * as Avatar from '$lib/components/ui/avatar';
</script>

<div class="player1 flex flex-row">
    <div class="flex avatar">
        <Avatar.Root class="w-16 h-16">
            <Avatar.Image class="bg-white border-4 rounded-full border-gray-600 cursor-pointer hover:border-gray-400" src={avatar} alt="avatar" />
        </Avatar.Root>
    </div>
    <div class="flex text-left flex-col">
        <p>@{name}</p>
        <p>${chips}</p>
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

</style>
