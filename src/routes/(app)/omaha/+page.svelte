<script lang="ts">
    import { Button } from '$lib/components/ui/button';
    export let data;
    
    $: userdata = null;
    function handleClick() {
        fetch('http://127.0.0.1:8090/api/basicuser', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "userid": data.user.id
                })
            })
            .then(response => response.json())
            .then(data => {
                console.log(data);
                userdata = JSON.parse(JSON.stringify(data));
            }); 
    }


</script>   
<div class="text-xl text-white justify-center w-screen h-screen m-auto items-center">
    <div class="m-auto">
        <Button on:click={handleClick}>Render some datum</Button>
        <div class="text-white">
            {#if userdata}
                {userdata["username"]}  
            {/if}
        </div>
    </div>
</div>
