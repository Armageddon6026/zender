import { defineStore, acceptHMRUpdate } from 'pinia'
import { ref } from 'vue'
import { $fetchV2 } from '@/common/http'

export const useServicesStore = defineStore('ourServicesStore', () => {
  // state:
  const Services = ref<Map<string, ResponseModel.Service>>(new Map())

  // getters:

  // actions:
  async function GetServices() {
    try {
      const res = await $fetchV2.Get<ResponseModel.Service[]>(`/services`)
      res.forEach((item) => {
        Services.value.set(item.uuid, item)
      })
    } catch (err) {
      Services.value.clear()
      console.log(err)
    }
  }

  async function InsertService(service: RequestModel.Service) {
    if (!service.name) return

    try {
      const res = await $fetchV2.Post<ResponseModel.Service>(`/services`, service)
      Services.value.set(res.uuid, res)
    } catch (err) {
      console.log(err)
    }
  }

  async function UpdatetService(service: RequestModel.Service, uuid: string) {
    if (!uuid) return

    try {
      await $fetchV2.Put(`/services/${uuid}`, service)
      const oldValue = Services.value.get(uuid)
      oldValue &&
        (((oldValue.name = service.name), (oldValue.groupId = service.groupId)),
        Services.value.set(uuid, oldValue))
    } catch (err) {
      console.log(err)
    }
  }

  async function DeletService(uuid: string) {
    if (!uuid) return

    try {
      await $fetchV2.Delete(`/services/${uuid}`)
      Services.value.delete(uuid)
    } catch (err) {
      console.log(err)
    }
  }

  function $reset() {
    Services.value.clear()
  }

  return {
    Services,
    GetServices,
    InsertService,
    UpdatetService,
    DeletService,
    $reset
  }
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useServicesStore, import.meta.hot))
}
