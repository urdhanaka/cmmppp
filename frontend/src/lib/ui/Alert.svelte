<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { HTMLAttributes } from 'svelte/elements';

  type Props = Omit<HTMLAttributes<HTMLDivElement>, 'class'> & {
    variant?: 'info' | 'success' | 'danger';
    class?: string;
    children: Snippet;
  };

  let {
    variant = 'info',
    class: className = '',
    children,
    ...rest
  }: Props = $props();
</script>

<div
  class="alert alert-{variant} {className}"
  role={variant === 'danger' ? 'alert' : 'status'}
  {...rest}
>
  {@render children()}
</div>

<style>
  .alert {
    padding: var(--space-3) var(--space-4);
    border: 1px solid;
    border-radius: var(--radius-md);
    font-size: var(--font-sm);
    color: var(--color-text);
  }

  .alert-info {
    border-color: var(--color-info);
    background: color-mix(in srgb, var(--color-info) 12%, transparent);
  }

  .alert-success {
    border-color: var(--color-success);
    background: color-mix(in srgb, var(--color-success) 12%, transparent);
  }

  .alert-danger {
    border-color: var(--color-danger);
    background: color-mix(in srgb, var(--color-danger) 12%, transparent);
  }
</style>
