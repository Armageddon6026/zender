import { defineStore, acceptHMRUpdate } from 'pinia'
import { ref, computed } from 'vue'
import { $fetchV2 } from '@/common/http'

export const useUserStore = defineStore('ourUserStore', () => {
  // state:
  const UserInfo = ref<ResponseModel.User>()

  // getters:
  const IsUserLogin = computed(() => {
    return UserInfo.value !== undefined
  })

  // actions:
  async function Login(account: string, password: string): Promise<string> {
    if (!account || !password) {
      throw new Error('帳號密碼不能為空!!')
    }
    try {
      const res = await $fetchV2.Auth<ResponseModel.Jwt>(
        '/login',
        `Basic ${btoa(account + ':' + password)}`
      )
      return res.token
    } catch (err) {
      console.log(err)
      $reset()
      throw new Error('帳號密碼錯誤!!')
    }
  }

  async function Logout() {
    try {
      await $fetchV2.Post<ResponseModel.Jwt>('/logout', '')
    } catch (err) {
      console.log(err)
      $reset()
    }
  }

  /** Set User Information in state : `UserInfo` */
  async function GetUser() {
    try {
      const res = await $fetchV2.Get<ResponseModel.User>('/user')
      UserInfo.value = res
    } catch (err) {
      $reset()
      console.log(err)
    }
  }

  async function UpdateUser(data: RequestModel.User) {
    if (!UserInfo.value) return
    if (!data.name && !data.password) return

    try {
      await $fetchV2.Put(`/user`, data)
      data.name && (UserInfo.value.name = data.name)
    } catch (err) {
      console.log(err)
    }
  }

  function $reset() {
    UserInfo.value = undefined
  }
  return { UserInfo, IsUserLogin, Login, Logout, GetUser, UpdateUser, $reset }
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useUserStore, import.meta.hot))
}
