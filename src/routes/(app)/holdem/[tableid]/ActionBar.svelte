<script lang="ts">
    import { Slider } from '$lib/components/ui/slider';
    import { Button } from '$lib/components/ui/button/index';
    import Card from './Card.svelte';
    //import { hand_store } from '$lib/stores/dealt';
    import { fade } from 'svelte/transition';
    import {dealCards, reset, betfunc} from '$lib/utils/game';
    import YourChips from './YourChips.svelte';
    import Result from './\(results)/Result.svelte';
    import { onMount, onDestroy } from 'svelte';
    import { nextaction, setFlop, setTurn, setRiver } from '$lib/utils/game';
    import PocketBase from 'pocketbase';
    const pb = new PocketBase("http://localhost:8090");

    export let currentPhase_: any;
    export let data: any;
    export let tableid: any;

    let socket;
    let gameState = {};
    
    let ttnh: string;
    let now = new Date().getTime();
    let countDownDate = new Date(now + 32000).getTime();
    let x = setInterval(function() {
        let now = new Date().getTime();
        let distance = countDownDate - now;
        let seconds = Math.floor((distance % (1000 * 60)) / 1000);
        ttnh = seconds + "s";
        if (distance < 0) {
            clearInterval(x);
            ttnh = "0s";
        }
    }, 1000);
    
    $: dealt = null;
    $: message = "";
    async function getCards(tableid: string) {
         
        await fetch("http://localhost:8090/api/hand",{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({"userid": data.user.id, "tableid": tableid})
        }).then(res => res.json()).then(data => {
            dealt = data;
            dealt = JSON.parse(dealt["cards"])
        });
        dealt = dealt;
        return dealt;
    }

    pb.collection("gametable").subscribe("*", async () => {
        dealt = await getCards(tableid);
    });
    
    onMount(async () => {
        socket = new WebSocket("ws://localhost:8090/ws");
        socket.onopen = () => {
            console.log("Connected to server");
        }
        socket.onmessage = (e: any) => {
            let data = JSON.parse(e.data);
            message = data.message;
            console.log(data);
        }
        socket.onerror = (e: any) => {
            console.log("Error: ", e);
        }
        dealt = await getCards(tableid);
    });

    onDestroy(() => {
        pb.collection("gametable").unsubscribe('*');
    });
    
    $: if (currentPhase_ == 1) {
        setFlop(tableid);
    } else if (currentPhase_ == 2) {
        setTurn(tableid);
    } else if (currentPhase_ == 3) {
        setRiver(tableid);
    }
</script>
<div class="bar"> 
{#if data.user}
    <div class="actions">
        <div class="cardset_ m-auto flex flex-row justify-evenly">
            <div class="flex items-center actionprof">
                <YourChips data={data}/>
            </div>
            <div style="font-size: 18px;" class="deal m-auto justify-start mx-5 items-center flex flex-row">
                {#if currentPhase_>= 0}
                    <div class="flex flex-col items-center m-auto justify-center">
                        <form class="flex flex-col items-center justify-center mx-5 m-auto">
                            <Slider class="w-32 my-2" max={data.user.balance} />
                            <div class="flex flex-row items-center justify-center">
                                <Button style="font-size: 18px;" variant="ghost" type="button" on:click={betfunc} class="btn my-2 m-auto variant-filled">Bet</Button>
                                <p style="font-size: 18px;">$0</p>
                            </div>
                        </form>
                    </div>
                    <Button style="font-size: 18px;" variant="ghost" on:click={()=>nextaction(tableid)} type="button" class="btn mx-2 my-5 variant-filled">Check</Button>
                    <Button style="font-size: 18px;" variant="ghost" on:click={()=>reset(tableid, data.user.id)} type="button" class="btn mx-2 my-5 variant-filled">Fold</Button>
                {:else if ttnh}
                    <div style="width: 500px" class="countdown inline-flex flex items-center flex-row">
                        <Button style="font-size: 18px;" variant="ghost" on:click={()=>reset(tableid, data.user.id)} type="button" class="btn mx-2 my-5 variant-filled">[Debug]</Button>
                        <h1 class="text-white">Next Hand In: {ttnh}</h1>
                    </div>
                {/if}
            </div>
            <div class="cardset justify-center items-center flex flex-row">
                <div class="cs items-center justify-center flex flex-row" in:fade={{delay: 1000, duration: 200}}>
                    {#if dealt != null}
                        {#if dealt['c1']}
                            <div class="cs items-center justify-center flex flex-row" in:fade={{delay: 1000, duration: 200}}>
                                <Card drawn={dealt['c1']} --rot="-10deg"/>
                                <Card drawn={dealt['c2']} --rot="10deg"/>
                            </div>
                        {/if}
                    {/if}
                </div>
                <div class="mx-20">
                    <Result />
                </div>
            </div>
            <div class="flex items-center nexthand">
                {#if currentPhase_ == -1}
                    <Button variant="ghost" on:click={async ()=> await dealCards(data.user.id, tableid)} type="button" style="display: inline-flex; align-items: center;" class="btn my-5 w-40 variant-filled">
                        <img width="32" src="../next.png" alt="next"/>&nbsp;Deal</Button>
                {:else}
                    <Button variant="ghost" on:click={()=>reset(tableid, data.user.id)} type="button" style="display: inline-flex; align-items: center;" class="btn my-5 w-40 variant-filled">
                        <img width="32" src="../next.png" alt="next"/>&nbsp;Skip Hand & {message}</Button>
                {/if}
            </div>
        </div>
    </div>
{:else}
    <div class="actions">
        <div class="cardset_ m-auto flex flex-row justify-evenly">
            <Button href="/signin">Sign In</Button>
        </div>
    </div>
{/if}
</div>

<style>
.cs {
   bottom: 10px; 
}
.cardset {
    object-fit: contain;
    display: flex;
    justify-content: start;
    align-items: center;
    width: 100%;
    height: 100%;
}
.cardset_ {
    height: 100%;
    width: 100%;
}
.bar {
    width: 100%;
    height: 10%;
    position: fixed;
    bottom: 0;
    padding: 0 1rem;
    background-color: #0F0F0F;
}
.actions {
    display: flex;
    width: 100%;
    height: 100%;
}
</style>

