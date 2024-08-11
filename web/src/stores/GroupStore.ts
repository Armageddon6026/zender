import { defineStore, acceptHMRUpdate } from 'pinia'
import { ref } from 'vue'
import { $fetchV2 } from '@/common/http'

export const useGroupStore = defineStore('ourGroupStore', () => {
  // state:
  const Groups = ref<Map<number, ResponseModel.Group>>(
    new Map([
      [
        -1,
        {
          id: -1,
          name: 'Other',
          note: 'asd56as4d654as6d 54dsa654dsa654asd564sadasd56as4d654as6d54 dsa654dsa654asd564sad'
        }
      ]
    ])
  )

  // getters:

  // actions:
  async function GetGroups() {
    try {
      const res = await $fetchV2.Get<ResponseModel.Group[]>('/groups')
      res.forEach((item) => {
        Groups.value.set(item.id, item)
      })
    } catch (err) {
      console.log(err)
    }
  }

  async function InsertGroup(group: RequestModel.Group) {
    if (!group.name) return

    try {
      const res = await $fetchV2.Post<ResponseModel.Group>(`/groups`, group)
      Groups.value.set(res.id, res)
    } catch (err) {
      console.log(err)
    }
  }

  async function UpdatetGroup(group: RequestModel.Group, groupId: number) {
    if (!groupId) return

    try {
      await $fetchV2.Put(`/groups/${groupId}`, group)
      const oldValue = Groups.value.get(groupId)
      oldValue &&
        (((oldValue.name = group.name), (oldValue.note = group.note)),
        Groups.value.set(groupId, oldValue))
    } catch (err) {
      console.log(err)
    }
  }

  async function DeleteGroup(groupId: number) {
    if (!groupId) return

    try {
      await $fetchV2.Delete(`/groups/${groupId}`)
      Groups.value.delete(groupId)
    } catch (err) {
      console.log(err)
    }
  }

  return { Groups, GetGroups, InsertGroup, UpdatetGroup, DeleteGroup }
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useGroupStore, import.meta.hot))
}
