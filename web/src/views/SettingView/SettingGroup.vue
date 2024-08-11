<script setup lang="ts">
import GroupInputDialog from '@/components/GroupInputDialog.vue'
import ConfirmDialog from '@/components/ConfirmDialog.vue'
import { useGroupStore } from '@/stores/GroupStore'
import { onMounted, ref } from 'vue'

const ourGroupStore = useGroupStore()
const ourCreateMsg = ref<Dialog.GroupInputMessage>({
  visible: false,
  group: { name: '', note: '' },
  args: []
})
const ourUpdateMsg = ref<Dialog.GroupInputMessage>({
  visible: false,
  group: { name: '', note: '' },
  args: []
})
const ourDeleteMsg = ref<Dialog.ConfirmMessage>({ visible: false, args: [] })

onMounted(async () => {
  await ourGroupStore.GetGroups()
})

function showDeleteDialog(group: ResponseModel.Group) {
  ourDeleteMsg.value.header = `Warning!!`
  ourDeleteMsg.value.body = `Delete Group "${group.name}" ?`
  ourDeleteMsg.value.level = 'Critic'
  ourDeleteMsg.value.visible = true
  ourDeleteMsg.value.args = [group.id]
}

function showCreateDialog() {
  ourCreateMsg.value.header = `Add Group`
  ourCreateMsg.value.visible = true
}

function showUpdateDialog(group: ResponseModel.Group) {
  ourUpdateMsg.value.header = `Update [ ${group.name} ]`
  ourUpdateMsg.value.visible = true
  ourUpdateMsg.value.group = { name: group.name, note: group.note }
  ourUpdateMsg.value.args = [group.id]
}
</script>
<template>
  <GroupInputDialog v-model="ourCreateMsg" @onConfirm="ourGroupStore.InsertGroup" />
  <GroupInputDialog v-model="ourUpdateMsg" @onConfirm="ourGroupStore.UpdatetGroup" />
  <ConfirmDialog v-model="ourDeleteMsg" @onConfirm="ourGroupStore.DeleteGroup" />

  <div flex="~ col" c="white" items="center">
    <h1>Group Setting</h1>
    <div m="b2 r15" w="146" flex="~ row" justify="end">
      <button
        un-border="sky-5 solid 1px rd-6px"
        h="8"
        w="26"
        bg="sky-9"
        text="size-18px white "
        hover="bg-sky-5"
        @click="showCreateDialog"
      >
        <i i="ic-baseline-add-circle-outline">Ｘ</i>
        Add
      </button>
    </div>
    <table zender="form" bg="gray-8" p="4" cellpadding="0" cellspacing="0" table="fixed" w="60">
      <tr h="12" c="main-normal" font="bold">
        <td w-30px tds>SN</td>
        <td w-60px tds>NAME</td>
        <td w-60 tds>NOTE</td>
        <td w-100px tds>ACTION</td>
      </tr>
      <tr
        v-for="([, group], i) in ourGroupStore.Groups"
        :key="i"
        h="9"
        c="white"
        hover="bg-gray-6"
        transition="200"
      >
        <td tds>{{ i }}</td>
        <td overflow="hidden" text="ellipsis" tds>{{ group.name }}</td>
        <td overflow="hidden" break="words" tds>{{ group.note }}</td>
        <td p="l4" items="center">
          <i
            i="ic-outline-edit"
            text="size-25px"
            c="blue-4"
            cursor="pointer"
            hover="c-blue-2"
            @click="showUpdateDialog(group)"
            >ｘ</i
          >
          <i
            i="ic-baseline-delete-outline"
            text="size-25px"
            c="red-8"
            cursor="pointer"
            hover="c-red-5"
            m="l3"
            @click="showDeleteDialog(group)"
            >ｘ</i
          >
        </td>
      </tr>
    </table>
  </div>
</template>
