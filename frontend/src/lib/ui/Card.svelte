<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { HTMLAttributes } from 'svelte/elements';

  type Props = Omit<HTMLAttributes<HTMLElement>, 'class'> & {
    class?: string;
    children: Snippet;
  };

  let { class: className = '', children, ...rest }: Props = $props();
</script>

<section class="card {className}" {...rest}>
  {@render children()}
</section>

<style>
  .card {
    display: grid;
    gap: var(--space-4);
    padding: var(--space-8);
    border: 1px solid var(--color-border);
    border-radius: var(--radius-lg);
    background: var(--color-surface);
  }
</style>
