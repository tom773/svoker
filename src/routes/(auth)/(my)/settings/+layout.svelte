<script lang="ts">
    import { Button } from "$lib/components/ui/button/index";
    import '../../../../main.css';
    import { pagey } from '$lib/stores/user';
    import { Separator } from '$lib/components/ui/separator/index';
    import { ArrowLeft } from 'lucide-svelte';
        
    export let data;
</script>
<div>
    <div class="back absolute top-3 left-3 z-10 w-10">
        <Button href="/" variant="ghost" class="text-white">
           <ArrowLeft class="w-6 h-6"/> 
        </Button>
    </div>
  <main
    class="flex h-full flex-1 flex-col gap-4 p-4 md:gap-8 md:p-10"
  >
    <div class="mx-auto grid w-full max-w-6xl gap-2">
        <p class="text-6xl font-semibold">Settings</p>
        <Separator class="bg-black my-4"/>
    </div>
    <div
      class="mx-auto grid w-full max-w-6xl items-start gap-6 md:grid-cols-[180px_1fr] lg:grid-cols-[250px_1fr]"
    >
      <nav class="grid gap-6 w-8 text-sm text-muted-foreground"
        data-x-chunk-container="chunk-container after:right-0">
        <Button on:click={() => pagey.set("profile")} class="font-semibold text-white"> General </Button>
        <Button on:click={() => pagey.set("security")} class="font-semibold text-white"> Security </Button>
        <Button on:click={() => pagey.set("payment")} class="font-semibold text-white"> Payment </Button>
        <Button on:click={() => pagey.set("support")} class="font-semibold text-white"> Support </Button>
      </nav>
      <div class="grid gap-6">
        <slot data={data}/>
      </div>
    </div>
  </main>
</div>

