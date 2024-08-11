<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'

const ourModal = ref<HTMLDialogElement>()
const ourModalValue = defineModel<Dialog.GroupInputMessage>({ required: true })
const ourNameInput = ref<string>()
const ourNoteInput = ref<string>('')
const ourEmit = defineEmits<{ onConfirm: [RequestModel.Group, ...any] }>()

watch(ourModalValue.value, (message) => {
  if (message.visible) {
    ourModal.value?.showModal()
    ourNameInput.value = ourModalValue.value.group.name
    ourNoteInput.value = ourModalValue.value.group.note
  }
})

onMounted(() => {
  if (ourModalValue.value.visible) {
    ourModal.value?.showModal()
  }
})

function handleConfirm() {
  if (!ourNameInput.value || !ourModalValue.value) return

  ourModalValue.value.visible = false
  ourEmit(
    'onConfirm',
    { name: ourNameInput.value, note: ourNoteInput.value },
    ...ourModalValue.value.args
  )
}

function handleCancel() {
  ourModalValue.value.visible = false
}

function afterLeave() {
  ourModal.value?.close()
}
</script>

<template>
  <transition @after-leave="afterLeave">
    <dialog
      v-show="ourModalValue.visible"
      ref="ourModal"
      zender-form
      bg="gray-8"
      flex="~ col"
      gap="16px"
      p="20px"
      h="340px"
      w="full"
      max-w="300px"
      @close="ourModalValue.visible = false"
    >
      <div c="white" font="bold size-20px" text="uppercase">
        {{ ourModalValue.header }}
      </div>
      <div>
        <input
          v-model="ourNameInput"
          c="main-normal"
          font="size-20px"
          bg="gray-9"
          un-border="none rd-2"
          p="10px"
          w="250px"
          type="text"
          placeholder="GroupName"
        />
      </div>
      <div>
        <textarea
          v-model="ourNoteInput"
          c="main-normal"
          font="size-20px"
          bg="gray-9"
          un-border="none rd-2"
          p="10px"
          h="160px"
          w="250px"
          type="text"
          placeholder="Note"
        ></textarea>
      </div>
      <div flex m="t-a" grid="justify-end">
        <slot name="footer">
          <button
            bg="transparent"
            un-border="none"
            w="90px"
            h="34px"
            font="size-16px bold"
            c="red-7"
            cursor="pointer"
            @click="handleCancel"
          >
            CANCEL
          </button>
          <button
            :disabled="!ourNameInput"
            bg="main-normal"
            un-border="none rd-5px"
            w="90px"
            h="34px"
            m="l-10px"
            font="size-16px bold"
            c="white"
            cursor="pointer"
            hover-disabled="cursor-not-allowed "
            @click="handleConfirm"
          >
            OK
          </button>
        </slot>
      </div>
    </dialog>
  </transition>
</template>

<style scoped>
.v-enter-from,
.v-leave-to {
  opacity: 0;
}

.v-enter-active {
  transition-timing-function: 'ease-out';
}

.v-leave-active {
  transition-timing-function: 'ease-in';
}

.v-enter-active,
.v-leave-active {
  transition-duration: 200ms;
}

.v-enter-from,
.v-leave-to {
  transform: scale(0.9) translateY(-2rem);
}

::backdrop {
  background-color: rgba(0, 0, 0, 0.7);
}
</style>
