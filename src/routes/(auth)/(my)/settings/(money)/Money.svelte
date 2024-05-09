
<script lang="ts">
    import { Button } from "$lib/components/ui/button/index";
    import * as Card from "$lib/components/ui/card/index";
    import { Input } from "$lib/components/ui/input/index";
    import { DollarSign } from "lucide-svelte";
</script>
<div class="grid grid-cols-2 gap-4">
    <Card.Root>
      <Card.Header>
        <Card.Title>Deposit</Card.Title>
        <Card.Description>
            Deposit money into your account.
        </Card.Description>
      </Card.Header>
      <Card.Content>
        <form class="inline-flex items-center">
            <DollarSign class="mx-1" size="32" />
            <Input placeholder="42069" />
        </form>
      </Card.Content>
      <Card.Footer class="border-t px-6 py-4">
        <Button>Desposit</Button>
      </Card.Footer>
    </Card.Root>
</div>

