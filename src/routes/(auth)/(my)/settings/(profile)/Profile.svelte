<script lang="ts">
    import { Button } from "$lib/components/ui/button/index";
    import * as Card from "$lib/components/ui/card/index";
    import { Checkbox } from "$lib/components/ui/checkbox/index";
    import { Input } from "$lib/components/ui/input/index";

</script>
<div class="grid grid-cols-2 gap-4">
    <Card.Root>
      <Card.Header>
        <Card.Title>Display Name</Card.Title>
        <Card.Description>
          Used to identify you by an alias such as "TheAssMan"
        </Card.Description>
      </Card.Header>
      <Card.Content>
        <form>
          <Input placeholder="User Name" />
        </form>
      </Card.Content>
      <Card.Footer class="border-t px-6 py-4">
        <Button>Save</Button>
      </Card.Footer>
    </Card.Root>
    <Card.Root>
      <Card.Header>
        <Card.Title>Update Email</Card.Title>
        <Card.Description>
          Update your email address.
        </Card.Description>
      </Card.Header>
      <Card.Content>
        <form class="flex flex-col gap-4">
          <Input placeholder="assman@kramerica.com" />
          <div class="flex items-center space-x-2">
            <Checkbox id="include" checked={true} />
            <label
              for="include"
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
            >
              Make Visible?
            </label>
          </div>
        </form>
      </Card.Content>
      <Card.Footer class="border-t px-6 py-4">
        <Button>Save</Button>
      </Card.Footer>
    </Card.Root>
</div>

