<script lang="ts">
  import type { HTMLInputAttributes } from 'svelte/elements';

  type Props = Omit<HTMLInputAttributes, 'class' | 'value'> & {
    label: string;
    value?: string;
    error?: string;
    class?: string;
  };

  let {
    label,
    value = $bindable(''),
    error,
    class: className = '',
    ...rest
  }: Props = $props();

  const uid = $props.id();
  const errId = `${uid}-error`;
</script>

<div class="field {className}">
  <label for={uid}>{label}</label>
  <input
    id={uid}
    bind:value
    aria-invalid={error ? 'true' : undefined}
    aria-describedby={error ? errId : undefined}
    {...rest}
  />
  {#if error}
    <p id={errId} class="error">{error}</p>
  {/if}
</div>

<style>
  .field {
    display: grid;
    gap: var(--space-2);
    font-size: var(--font-sm);
  }

  input {
    padding: var(--space-3);
    border: 1px solid var(--color-border);
    border-radius: var(--radius-md);
    background: var(--color-input-bg);
    font-size: var(--font-md);
  }

  input[aria-invalid='true'] {
    border-color: var(--color-danger);
  }

  .error {
    color: var(--color-danger);
  }
</style>
