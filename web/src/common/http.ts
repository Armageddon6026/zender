import router from '@/router'
import { FetchV2 } from '@/utils/FetchV2'
import { Enviroment } from '@/common/const'

const instance = new FetchV2(Enviroment.BASE_API_URL)

instance.AddResponseInterceptors((response) => {
  switch (response.status) {
    case 400:
      console.log('400 無效的請求')
      break
    case 401:
      console.log('401 無效的驗證')
      if (String(router.currentRoute) != 'login') {
        router.push({ name: 'LoginView' })
      }
      break
    case 403:
      console.log('403 權限不足')
      break
    case 404:
      console.log('404 請求失敗')
      break
    case 500:
      console.log('500 未預料的狀況')
      break
  }
})

export { instance as $fetchV2 }
