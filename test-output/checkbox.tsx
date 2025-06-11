import { component$, useStyles$ } from "@builder.io/qwik";
import { Checkbox } from "@kunai-consulting/qwik";
import styles from './checkbox.module.css';


export default component$(() => {
  return (
    <Checkbox.Root>
      <Checkbox.HiddenInput />
      <div class="{styles.node0}">
        <Checkbox.Trigger
          class="{styles.trigger}"
        >
          <Checkbox.Indicator
            class="{styles.indicator}"
          >
            <LuCheck />
          </Checkbox.Indicator>
        </Checkbox.Trigger>
        <Checkbox.Label class="{styles.label}">
          This is a trusted device, don't ask again
        </Checkbox.Label>
      </div>
    </Checkbox.Root>
  );
}); 