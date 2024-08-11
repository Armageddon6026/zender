<script setup lang="ts">
import { onMounted } from 'vue'
import { useServicesStore } from '@/stores/ServicesStore'
import { useEventStore } from '@/stores/EventStore'
import { useGroupStore } from '@/stores/GroupStore'

const ourProps = defineProps<{ id: string }>()
const ourServiceStore = useServicesStore()
const ourGroupStore = useGroupStore()
const ourEventStore = useEventStore()

onMounted(async () => {
  await ourServiceStore.GetServices()
  await ourGroupStore.GetGroups()
})

function isServiceLogin(item: ResponseModel.Service): boolean {
  return ourEventStore.LoginServices.some((ele) => ele.uuid === item.uuid)
}

function isGroupService(item: EventModel.LoginService): boolean {
  return (
    ourServiceStore.Services.has(item.uuid) &&
    ourServiceStore.Services.get(item.uuid)!.groupId == Number(ourProps.id)
  )
}

function getServiceStatusColor(item: EventModel.LoginService): string {
  return ourEventStore.ErrorServiceUuids.has(item.uuid) ? 'yellow' : 'var(--main-color)'
}
</script>

<template>
  <div flex="col" w="200" text="white">
    <h1 text="center">Service List {{ ourGroupStore.Groups.get(Number(ourProps.id))?.name }}</h1>
    <div overflow="y-scroll x-hidden" h="75vh">
      <div v-for="(item, i) in ourEventStore.LoginServices" :key="i">
        <div
          v-if="isGroupService(item)"
          bg="gray-8"
          m="3"
          h="20"
          p="10px"
          un-border="solid 2px rd-10px"
          hover:shadow="[0_0_5px_1px] main-normal op-50"
          flex="wrap"
          cursor="pointer"
          :style="{ borderColor: getServiceStatusColor(item) }"
        >
          <div m="t2" w="full" flex="~">
            <i m="1" i="ic-baseline-circle" :style="{ color: getServiceStatusColor(item) }"></i>
            <div text="size-20px" m="l2">[ {{ item.name }} ]</div>
            <div
              text="size-15px"
              m="l3"
              p="l5px r5px t2px"
              bg="accept-dark"
              un-border="solid 2 accept-normal rd-5px"
            >
              IP : {{ item.ip }}
            </div>
            <div
              text="size-15px"
              m="l3"
              p="l5px r5px t2px"
              bg="main-dark"
              un-border="solid 2 main-normal rd-5px"
            >
              UUID : {{ item.uuid }}
            </div>
          </div>
          <div text="size-15px" p="t3" m="t2 l8" w="full" flex>
            <div>Last Data Time :</div>
            <div m-l2>{{ ourEventStore.GetLastScanFileTime(item.scanFiles) }}</div>
            <div m-l10>Data Count On The Day :</div>
            <div m-l2>{{ ourEventStore.GetAllScanFilesDataCount(item.scanFiles) }}</div>
            <div m-l10>Applications Status :</div>
            <div m-l2>
              {{
                ourEventStore.GetAliveApplicationCount(item.scanApplications) +
                '/' +
                Object.values(item.scanApplications).length
              }}
            </div>
          </div>
        </div>
      </div>
      <div v-for="[key, value] in ourServiceStore.Services" :key="key">
        <div
          v-if="value.groupId == Number(ourProps.id) && !isServiceLogin(value)"
          text="coolGray"
          bg="gray-8"
          m="3"
          h="20"
          p="10px"
          un-border="solid 2px gray rd-10px"
          flex="wrap"
          cursor="not-allowed"
        >
          <div m-t2 w-full flex>
            <i m="1" i="ic-baseline-circle" c="deny-normal"></i>
            <div text="size-20px" m="l-2">[ {{ value.name }} ]</div>
            <div
              text="size-15px"
              m="l-3"
              p="l-5px r-5px t-2px"
              bg="accept-dark"
              un-border="solid 2 accept-normal rd-5px"
            >
              IP : -.-.-.-
            </div>
            <div
              text="size-15px"
              m="l-3"
              p="l-5px r-5px t-2px"
              bg="main-dark"
              un-border="solid 2 main-normal rd-5px"
            >
              UUID : {{ value.uuid }}
            </div>
          </div>
          <div text="size-15px" p="t-3" m="t2 l-8" w="full" flex>
            <div>Last DataTime :</div>
            <div m-l2>---- / -- / -- -- : --</div>
            <div m-l10>DataNum in Last hour :</div>
            <div m-l2>---</div>
            <div m-l10>Subsystem Alive :</div>
            <div m-l2>--</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
