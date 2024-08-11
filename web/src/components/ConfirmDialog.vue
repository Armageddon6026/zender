<script setup lang="ts">
import { ref, onMounted, watch, computed } from 'vue'

const ourModal = ref<HTMLDialogElement>()
const ourModalValue = defineModel<Dialog.ConfirmMessage>({ required: true })
const ourEmit = defineEmits<{ onConfirm: [...any] }>()
const ourHeaderColor = computed(() => {
  switch (ourModalValue.value.level) {
    case 'Normal':
      return 'white'
    case 'Critic':
      return 'red'
    case 'Warning':
      return 'yellow'
    default:
      return 'white'
  }
})

watch(ourModalValue.value, (message) => {
  if (message.visible) ourModal.value?.showModal()
})

onMounted(() => {
  if (ourModalValue.value.visible) ourModal.value?.showModal()
})

function handleConfirm() {
  if (!ourModalValue.value) return

  ourModalValue.value.visible = false
  ourEmit('onConfirm', ...ourModalValue.value.args)
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
      h="160px"
      w="full"
      max-w="300px"
      @close="ourModalValue.visible = false"
    >
      <div :style="{ color: ourHeaderColor }" font="bold size-20px" text="uppercase">
        <span v-if="ourModalValue.level == 'Critic'" i="ic-twotone-warning-amber">Ｘ</span>
        <span v-else-if="ourModalValue.level == 'Normal'" i="ic-baseline-check-circle-outline"
          >Ｖ</span
        >
        <span v-else-if="ourModalValue.level == 'Warning'" i="ic-round-notification-important"
          >ｗ</span
        >
        {{ ourModalValue.header }}
      </div>
      <div c="white" w-a break-all>
        {{ ourModalValue.body }}
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
            bg="main-normal"
            un-border="none rd-5px"
            w="90px"
            h="34px"
            font="size-16px bold"
            c="white"
            cursor="pointer"
            m="l-10px"
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
