<script lang="ts">
  import type { Snippet } from 'svelte';
  import Button from './Button.svelte';

  let {
    open = $bindable(false),
    title,
    children
  }: { open?: boolean; title: string; children: Snippet } = $props();

  let dialog: HTMLDialogElement;

  // Sync the `open` prop with the native dialog.
  // Escape and the close button both end with open = false.
  $effect(() => {
    if (open && !dialog.open) dialog.showModal();
    if (!open && dialog.open) dialog.close();
  });
</script>

<dialog bind:this={dialog} aria-label={title} onclose={() => (open = false)}>
  <header>
    <h2>{title}</h2>
    <Button variant="ghost" aria-label="Close" onclick={() => (open = false)}>
      ×
    </Button>
  </header>
  <div class="body">
    {@render children()}
  </div>
</dialog>

<style>
  dialog {
    width: min(480px, 100%);
    padding: var(--space-6);
    border: 1px solid var(--color-border);
    border-radius: var(--radius-lg);
    background: var(--color-surface);
    color: var(--color-text);
  }

  dialog::backdrop {
    background: rgb(0 0 0 / 0.5);
  }

  header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: var(--space-4);
    margin-bottom: var(--space-4);
  }
</style>
