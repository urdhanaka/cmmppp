import { createRouter } from 'sv-router';
import Login from './routes/Login.svelte'

export const { p, navigate, isActive, route } = createRouter({
  '/login': Login
})
