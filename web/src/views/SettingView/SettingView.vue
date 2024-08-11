<script setup lang="ts">
import SideBar from '@/components/SideBar.vue'
</script>

<template>
  <div flex>
    <SideBar m="t-20 r-5" p="5" />
    <div w-231>
      <RouterView />
    </div>
  </div>
</template>
