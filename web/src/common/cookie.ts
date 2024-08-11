import { Enviroment } from '@/common/const'
import Cookies from 'js-cookie'

function ClearAllCookies() {
  Cookies.remove(Enviroment.ACCESS_TOKEN_HEADER_NAME)
}

export { ClearAllCookies }
