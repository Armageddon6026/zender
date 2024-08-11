<script setup lang="ts">
import { onMounted, onUnmounted, computed } from 'vue'
import { useServicesStore } from '@/stores/ServicesStore'
import { useEventStore } from '@/stores/EventStore'

const ctrl = new AbortController()
const ourServiceStore = useServicesStore()
const ourEventStore = useEventStore()

const servicesNum = computed(() => {
  return ourServiceStore.Services.size
})

const ErrorServiceNum = computed(() => {
  return ourEventStore.ErrorServiceUuids.size
})

const loginServicesNum = computed(() => {
  return ourEventStore.LoginServices.length - ourEventStore.ErrorServiceUuids.size
})

const logoutServicesNum = computed(() => {
  return ourServiceStore.Services.size - ourEventStore.LoginServices.length
})

onMounted(async () => {
  await ourServiceStore.GetServices()
  try {
    ourEventStore.StartLoginServiceListen(ctrl)
  } catch (err) {
    console.log(err)
  }
})

onUnmounted(() => {
  ctrl.abort()
})
</script>

<template>
  <div flex bg="gray-8" h="30px" w="full" items="center" justify="center">
    <div flex items="center" w="40">
      <div i="ic-round-workspaces" text="2xl gray-1"></div>
      <div ml-3 text="15px white">{{ servicesNum }}</div>
    </div>
    <div flex items="center" w="40">
      <div i="ic-baseline-play-arrow" text="3xl main-normal"></div>
      <div ml-3 text="15px white">{{ loginServicesNum }}</div>
    </div>
    <div flex items="center" w="40">
      <div i="ic-baseline-stop" text="3xl deny-normal"></div>
      <div ml-3 text="15px white">{{ logoutServicesNum }}</div>
    </div>
    <div flex items="center" w="40">
      <div i="ic-baseline-warning" text="2xl yellow-5"></div>
      <div ml-3 text="15px white">{{ ErrorServiceNum }}</div>
    </div>
  </div>
</template>
