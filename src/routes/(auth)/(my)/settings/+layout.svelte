<script lang="ts">
    import { Button } from "$lib/components/ui/button/index";
    import '../../../../main.css';
    import { pagey } from '$lib/stores/user';
</script>
<div>
  <main
    class="flex h-full flex-1 flex-col gap-4 p-4 md:gap-8 md:p-10"
  >
    <div class="mx-auto grid w-full max-w-6xl gap-2">
      <h1 class="text-3xl font-semibold">Settings</h1>
    </div>
    <div
      class="mx-auto grid w-full max-w-6xl items-start gap-6 md:grid-cols-[180px_1fr] lg:grid-cols-[250px_1fr]"
    >
      <nav class="grid gap-4 w-8 text-sm text-muted-foreground"
        data-x-chunk-container="chunk-container after:right-0">
        <Button on:click={() => pagey.set("profile")} class="font-semibold text-white"> General </Button>
        <Button on:click={() => pagey.set("security")} class="font-semibold text-white"> Security </Button>
        <Button on:click={() => pagey.set("payment")} class="font-semibold text-white"> Payment </Button>
        <Button on:click={() => pagey.set("support")} class="font-semibold text-white"> Support </Button>
      </nav>
      <div class="grid gap-6">
        <slot />
      </div>
    </div>
  </main>
</div>

