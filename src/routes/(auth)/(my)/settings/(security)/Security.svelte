<script lang="ts">
    import { Button } from "$lib/components/ui/button/index";
    import * as Card from "$lib/components/ui/card/index";
    import { Input } from "$lib/components/ui/input/index";
    export let data: any;
</script>
<div class="grid grid-cols-2 gap-4">
    <Card.Root>
      <Card.Header>
        <Card.Title>Change Your Password</Card.Title>
        <Card.Description>
            Forget your password again dory? You are a fish. Change it here.
        </Card.Description>
      </Card.Header>
      <form method="" action="" class="items-center">
          <Card.Content class="grid-rows-2 border-t px-6 py-5 items-center">
                <div class="grid grid-cols-2 gap-4">
                    <div class="flex flex-col w-full">
                        <label for="Password" class="font-bold text-l mb-1">New</label>
                        <Input id="password" name="password" type="password"/>
                    </div>
                    <div class="flex flex-col w-full">
                        <label for="repeat" class="font-bold text-l mb-1">Repeat New</label>
                        <Input id="repeat" name="repeat" type="password"/>
                    </div>
                </div>
            </Card.Content>
          <Card.Footer class="border-t px-6 py-5">
            <Button type="submit">Submit</Button>
          </Card.Footer>
      </form>
    </Card.Root>
</div>


