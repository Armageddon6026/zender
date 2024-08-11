<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useUserStore } from '@/stores/UserStore'

const ourUserStore = useUserStore()
const ourEditMode = ref<boolean>(false)
const ourUserUpdateInput: RequestModel.User & { confirmPassword: string } = reactive({
  name: '',
  password: '',
  confirmPassword: ''
})

function handleEdit() {
  ourEditMode.value = true
}

function handleCancel() {
  ourEditMode.value = false
}

async function handleUpdate() {
  if (
    ourUserUpdateInput.password &&
    ourUserUpdateInput.password != ourUserUpdateInput.confirmPassword
  ) {
    console.log('wrong password')
    return
  }

  await ourUserStore.UpdateUser(ourUserUpdateInput)
  ourEditMode.value = false
}
</script>
<template>
  <div text="white" flex="~ col" items="center">
    <h1 text="center white">User Setting</h1>
    <div w="110">
      <div flex="~ row" justify-end>
        <div v-if="!ourEditMode" m="b-2">
          <button
            un-border="sky-5 solid 1px rd-6px"
            h="8"
            w="26"
            bg="sky-9"
            text="size-18px white "
            hover="bg-sky-5"
            @click="handleEdit"
          >
            <i i="ic-baseline-edit">Ｘ</i>
            Edit
          </button>
        </div>
        <div v-else m="b-2">
          <button
            ml="2"
            un-border="sky-5 solid 1px rd-6px"
            h="8"
            w="26"
            bg="sky-9"
            text="size-18px white "
            hover="bg-sky-5"
            @click="handleUpdate"
          >
            <i i="ic-sharp-file-download-done">Ｘ</i>
            Update
          </button>
          <button
            ml-2
            un-border="deny-normal solid 1px rd-6px"
            h="8"
            w="26"
            bg="deny-dark"
            text="size-18px white "
            hover="bg-deny-normal"
            @click="handleCancel"
          >
            <i i="ic-baseline-cancel">Ｘ</i>
            Cancel
          </button>
        </div>
      </div>

      <div zender-form bg="gray-8" p="4" text="size-20px">
        <table cellpadding="5" cellspacing="0">
          <tr>
            <td w-120px>Account :</td>
            <td>{{ ourUserStore.UserInfo!.account }}</td>
          </tr>
          <tr>
            <td>Name :</td>
            <td v-if="!ourEditMode">{{ ourUserStore.UserInfo!.name }}</td>
            <td v-else>
              <input
                v-model="ourUserUpdateInput.name"
                type="text"
                placeholder="new name"
                data-space-bottom="0.5rem"
                bg="gray-9"
                un-border="none rd-6px"
                outline="none"
                p="t2 b2 l2"
                c="main-normal"
              />
            </td>
          </tr>
          <tr>
            <td>Password :</td>
            <td v-if="!ourEditMode">***************</td>
            <td v-else>
              <input
                v-model="ourUserUpdateInput.password"
                type="password"
                placeholder="password"
                data-space-bottom="0.5rem"
                bg="gray-9"
                un-border="none rd-6px"
                outline="none"
                p="t2 b2 l2"
                c="main-normal"
              />
            </td>
          </tr>
          <tr v-if="ourEditMode">
            <td>Confirm :</td>
            <td>
              <input
                v-model="ourUserUpdateInput.confirmPassword"
                type="password"
                placeholder="confirm password"
                data-space-bottom="0.5rem"
                bg="gray-9"
                un-border="none rd-6px"
                outline="none"
                p="t2 b2 l2"
                c="main-normal"
              />
            </td>
          </tr>
          <tr>
            <td>Join Date :</td>
            <td>{{ new Date(ourUserStore.UserInfo!.date).toLocaleString() }}</td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</template>
