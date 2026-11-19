<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { HTMLButtonAttributes } from 'svelte/elements';

  type Props = Omit<HTMLButtonAttributes, 'class'> & {
    variant?: 'primary' | 'ghost' | 'danger';
    class?: string;
    children: Snippet;
  };

  let {
    variant = 'primary',
    type = 'button',
    class: className = '',
    children,
    ...rest
  }: Props = $props();
</script>

<button class="btn btn-{variant} {className}" {type} {...rest}>
  {@render children()}
</button>

<style>
  .btn {
    padding: var(--space-3) var(--space-4);
    border: 1px solid transparent;
    border-radius: var(--radius-md);
    font-weight: 500;
    transition:
      background var(--transition-fast),
      opacity var(--transition-fast);
  }

  .btn-primary {
    background: var(--color-primary);
    color: var(--color-primary-text);
  }
  .btn-primary:hover:not(:disabled) {
    background: var(--color-primary-hover);
  }

  .btn-ghost {
    background: transparent;
    color: var(--color-text);
    border-color: var(--color-border);
  }
  .btn-ghost:hover:not(:disabled) {
    background: var(--color-surface);
  }

  .btn-danger {
    background: var(--color-danger);
    color: var(--color-bg);
  }

  .btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }
</style>
