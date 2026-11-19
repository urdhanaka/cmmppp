import { createRouter } from 'sv-router';
import Login from './routes/Login.svelte'

export const { p, navigate, isActive, route } = createRouter({
  // Dev only: the lazy import is stripped from production builds.
  ...(import.meta.env.DEV && {
    '/styleguide': () => import('./routes/StyleGuide.svelte')
  }),

  '/login': Login,
})
