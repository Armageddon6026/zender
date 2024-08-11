<script setup lang="ts">
import router from '@/router'
import { useEventStore } from '@/stores/EventStore'
import { useGroupStore } from '@/stores/GroupStore'
import { useServicesStore } from '@/stores/ServicesStore'
import { onMounted } from 'vue'

const ourGroupStore = useGroupStore()
const ourServicesStore = useServicesStore()
const ourEventStore = useEventStore()

onMounted(() => {
  ourGroupStore.GetGroups()
})

function showServiceList(groupId: number) {
  router.push({ path: `/service/group/${groupId}` })
}

function getGroupStatusColor(groupId: number): string {
  for (const [, service] of ourServicesStore.Services) {
    if (service.groupId != groupId) continue

    if (ourEventStore.ErrorServiceUuids.has(service.uuid)) {
      return 'yellow'
    }
  }
  return 'var(--main-color)'
}
</script>
<template>
  <div flex="~ col" c="white" items="center">
    <h1>Groups page</h1>
    <div flex="~ wrap " content="start" overflow="y-scroll x-hidden" h="75vh" w="204">
      <div
        v-for="[key, value] in ourGroupStore.Groups"
        :key="key"
        zender="form"
        m="l6 t6"
        h="30"
        w="60"
        cursor="pointer"
        @click="showServiceList(value.id)"
        bg="gray-8"
        hover-shadow="[0_0px_13px] cyan"
        :style="{ borderColor: getGroupStatusColor(value.id) }"
      >
        <div flex="~" m="3">
          <i i="ic-baseline-circle" m="r1" :style="{ color: getGroupStatusColor(value.id) }">X</i>
          <div text="ellipsis">Group : {{ value.name }}</div>
        </div>
        <div break="words" overflow="y-scroll" m="3" p="1" h="14" bg="gray-9" un-border="rd-2">
          {{ value.note }}
        </div>
      </div>
    </div>
  </div>
</template>
