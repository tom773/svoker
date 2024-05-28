<script lang="ts">
    import { Button } from "$lib/components/ui/button/index";
    import * as Card from "$lib/components/ui/card/index";
    import { Input } from "$lib/components/ui/input/index";
    import { numberWithCommas } from "$lib/utils";
    import { DollarSign } from "lucide-svelte";
    export let data: any;
</script>
<div class="grid grid-cols-2 gap-4">
    <Card.Root>
      <Card.Header>
        <Card.Title>Deposit</Card.Title>
        <Card.Description>
            Deposit money into your account.<br><strong>Current Balance:</strong>&nbsp;${numberWithCommas(data?.user.user.balance)}
        </Card.Description>
      </Card.Header>
      <form method="POST" action="?/amt" class="inline-flex flex-row items-center">
          <Card.Content class="inline-flex py-5 items-center">
                <DollarSign class="mx-1" size="32" />
                <Input id="amt" name="amt" type="currency" placeholder="42069" />
          </Card.Content>
          <Card.Footer class="border-t px-6 py-5">
            <Button type="submit">Desposit</Button>
          </Card.Footer>
      </form>
    </Card.Root>
</div>

