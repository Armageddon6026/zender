<script setup lang="ts">
import { useUserStore } from '@/stores/UserStore'
import NavigationBar from '@/components/NavigationBar.vue'
import SubNavigationBar from '@/components/SubNavigationBar.vue'
// Init Login
const ourUserStore = useUserStore()
</script>

<template>
  <NavigationBar />
  <SubNavigationBar v-if="ourUserStore.IsUserLogin" />
  <RouterView />
</template>
