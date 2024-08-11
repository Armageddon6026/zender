<script setup lang="ts"></script>
<template>
  <ul flex="~ col" text="size-20px">
    <RouterLink c="gray-4" p="2" decoration="none" hover="c-white" to="user-setting"
      ><i i="ic-sharp-manage-accounts">XX</i>User</RouterLink
    >
    <RouterLink c="gray-4" p="2" decoration="none" hover="c-white" to="service-setting"
      ><i i="ic-baseline-display-settings">XX</i>Service</RouterLink
    >
    <RouterLink c="gray-4" p="2" decoration="none" hover="c-white" to="group-setting"
      ><i i="ic-round-group">XX</i>Group</RouterLink
    >
    <RouterLink c="gray-4" p="2" decoration="none" hover="c-white" to="authorization-setting"
      ><i i="ic-outline-security">XX</i>Autherization</RouterLink
    >
    <RouterLink c="gray-4" p="2" decoration="none" hover="c-white" to="system-setting"
      ><i i="ic-sharp-settings-power">XX</i>System</RouterLink
    >
  </ul>
</template>

<style scoped>
.router-link-active {
  color: white;
}
</style>
