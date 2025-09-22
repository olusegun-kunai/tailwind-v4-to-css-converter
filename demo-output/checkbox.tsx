import { component$, useStyles$ } from "@builder.io/qwik";
import { Checkbox } from "@kunai-consulting/qwik";
import styles from './checkbox.module.css';


export default component$(() => {
  return (
    <Checkbox.Root>
      <Checkbox.HiddenInput />
      <div className={styles.node0}>
        <Checkbox.Trigger
          className={styles.trigger}
        >
          <Checkbox.Indicator
            className={styles.indicator}
          >
            <LuCheck />
          </Checkbox.Indicator>
        </Checkbox.Trigger>
        <Checkbox.Label className={styles.label}>
          This is a trusted device, don't ask again
        </Checkbox.Label>
      </div>
    </Checkbox.Root>
  );
}); 