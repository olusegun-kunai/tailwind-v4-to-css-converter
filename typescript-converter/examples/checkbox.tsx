import { component$, useStyles$ } from "@builder.io/qwik";
import { Checkbox } from "@kunai-consulting/qwik";

export default component$(() => {
  return (
    <Checkbox.Root>
      <Checkbox.HiddenInput />
      <div class="flex items-center gap-2">
        <Checkbox.Trigger
          class="size-[25px] rounded-lg relative bg-gray-500 
                 focus-visible:outline focus-visible:outline-1 focus-visible:outline-white
                 disabled:opacity-50 bg-qwik-neutral-200 data-[checked]:bg-qwik-blue-800 focus-visible:ring-[3px] ring-qwik-blue-600"
        >
          <Checkbox.Indicator
            class="data-[checked]:flex justify-center items-center absolute inset-0"
          >
            <LuCheck />
          </Checkbox.Indicator>
        </Checkbox.Trigger>
        <Checkbox.Label class="text-sm">
          This is a trusted device, don't ask again
        </Checkbox.Label>
      </div>
    </Checkbox.Root>
  );
}); 