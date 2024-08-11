<script setup lang="ts">
import { ref, onMounted } from 'vue'

import ServiceInputDialog from '@/components/ServiceInputDialog.vue'
import ConfirmDialog from '@/components/ConfirmDialog.vue'
import { useServicesStore } from '@/stores/ServicesStore'
import { useGroupStore } from '@/stores/GroupStore'

const ourServiceStore = useServicesStore()
const ourGroupStore = useGroupStore()
const ourCreateMsg = ref<Dialog.ServiceInputMessage>({
  visible: false,
  service: { name: '', groupId: -1 },
  args: []
})
const ourUpdateMsg = ref<Dialog.ServiceInputMessage>({
  visible: false,
  service: { name: '', groupId: -1 },
  args: []
})
const ourDeleteMsg = ref<Dialog.ConfirmMessage>({ visible: false, args: [] })

onMounted(async () => {
  await ourGroupStore.GetGroups()
  await ourServiceStore.GetServices()
})

function showDeleteDialog(serviceInfo: ResponseModel.Service) {
  ourDeleteMsg.value.header = `Warning!!`
  ourDeleteMsg.value.body = `Delete Service "${serviceInfo.name}" ?`
  ourDeleteMsg.value.level = 'Critic'
  ourDeleteMsg.value.visible = true
  ourDeleteMsg.value.args = [serviceInfo.uuid]
}

function showInputDialog() {
  ourCreateMsg.value.header = `Add Service`
  ourCreateMsg.value.visible = true
}

function showUpdateDialog(serviceInfo: ResponseModel.Service) {
  ourUpdateMsg.value.header = `Update [ ${serviceInfo.name} ]`
  ourUpdateMsg.value.visible = true
  ourUpdateMsg.value.service = { name: serviceInfo.name, groupId: serviceInfo.groupId }
  ourUpdateMsg.value.args = [serviceInfo.uuid]
}
</script>

<template>
  <ConfirmDialog v-model="ourDeleteMsg" @onConfirm="ourServiceStore.DeletService" />
  <ServiceInputDialog v-model="ourCreateMsg" @onConfirm="ourServiceStore.InsertService" />
  <ServiceInputDialog v-model="ourUpdateMsg" @onConfirm="ourServiceStore.UpdatetService" />

  <div flex="~ col" items="center">
    <h1 text="white">Service Setting</h1>
    <div m="b2 r15" w="full" flex="~ row" justify="end">
      <button
        un-border="sky-5 solid 1px rd-6px"
        h="8"
        w="26"
        bg="sky-9"
        text="size-18px white "
        hover="bg-sky-5"
        @click="showInputDialog"
      >
        <i i="ic-baseline-add-circle-outline">Ｘ</i>
        Add
      </button>
    </div>
    <table zender="form" bg="gray-8" p="4" cellpadding="0" cellspacing="0">
      <tr h-12 c-main-normal font-bold>
        <td w-30px tds>SN</td>
        <td w-60px tds>NAME</td>
        <td w-auto tds>UUID</td>
        <td w-auto tds>PRIVATEKEY</td>
        <td w-auto tds>GROUP</td>
        <td w-100px tds>ACTION</td>
      </tr>
      <tr
        v-for="([key, value], i) in ourServiceStore.Services"
        :key="key"
        h="12"
        c="white"
        hover="bg-gray-6"
        transition="200"
      >
        <td tds>{{ i }}</td>
        <td tds>{{ value.name }}</td>
        <td tds>{{ value.uuid }}</td>
        <td tds>{{ value.privateKey }}</td>
        <td tds>{{ ourGroupStore.Groups.get(value.groupId)?.name }}</td>
        <td tds flex="~ row" items="center" h="inherit" gap="10px">
          <i
            i="ic-outline-edit"
            text="size-25px"
            c="blue-4"
            cursor="pointer"
            hover="c-blue-2"
            @click="showUpdateDialog(value)"
          ></i>
          <i
            i="ic-baseline-delete-outline"
            text="size-25px"
            c="red-8"
            cursor="pointer"
            hover="c-red-5"
            @click="showDeleteDialog(value)"
          ></i>
        </td>
      </tr>
    </table>
  </div>
</template>
