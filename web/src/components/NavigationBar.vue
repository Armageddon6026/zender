<script setup lang="ts">
import router from '@/router'
import { useUserStore } from '@/stores/UserStore'
const ourUserStore = useUserStore()

async function handleLogout() {
  await ourUserStore.Logout()
  router.push({ name: 'LoginView' })
}
</script>

<template>
  <nav h="80px" w="100%">
    <img src="/logo.png" alt="X" w="80px" m="l5 b--8" />
    <span m="l3" text="size-34px white " font="bold">Project Zender</span>
    <ul v-if="ourUserStore.IsUserLogin" hidden>
      <RouterLink to="/"> Home </RouterLink>
      <RouterLink to="/service"> Service </RouterLink>
      <RouterLink to="/Document"> Document </RouterLink>
      <RouterLink to="/setting"> Setting </RouterLink>
      <button class="logout-button" @click="handleLogout">Logout</button>
    </ul>
  </nav>
</template>

<style scoped>
nav ul a {
  color: white;
  font-size: 17px;
  text-transform: capitalize;
  padding: 12px 25px;
  font-weight: bold;
  text-decoration: none;
}

nav ul a:hover {
  border-bottom: 2px solid var(--main-color);
}

nav .login-button {
  width: 80px;
  height: 40px;
  border-radius: 6px;
  letter-spacing: 1px;
  font-size: 15px;
  outline: none;
  border: 2px solid var(--main-color);
  background: var(--main-color-dark);
  color: white;
  margin: 0 10px;
}

nav .login-button:hover {
  background: var(--main-color);
}

nav .logout-button {
  width: 80px;
  height: 40px;
  border-radius: 6px;
  letter-spacing: 1px;
  font-size: 15px;
  outline: none;
  border: 2px solid var(--main-deny-color);
  background: var(--main-deny-color-dark);
  color: white;
  margin: 0 10px;
}

nav .logout-button:hover {
  background: var(--main-deny-color);
}

@media (min-width: 1024px) {
  nav ul {
    display: inline;
    float: right;
    list-style: none;
    line-height: 80px;
    margin: 0 10px;
  }
}
</style>
