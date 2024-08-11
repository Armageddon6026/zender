import { defineStore, acceptHMRUpdate } from 'pinia'
import { computed, ref } from 'vue'
import { fetchEventSource } from '@microsoft/fetch-event-source'
import { Enviroment } from '@/common/const'

export const useEventStore = defineStore('ourEventStore', () => {
  // state:
  const LoginServices = ref<EventModel.LoginService[]>([])

  // getters:
  const ErrorServiceUuids = computed(() => {
    const errorService = new Map<string, string>()
    LoginServices.value.forEach((service) => {
      if (
        service.scanApplications &&
        Object.values(service.scanApplications).length !=
          GetAliveApplicationCount(service.scanApplications)
      )
        errorService.set(service.uuid, 'application error')
    })
    return errorService
  })

  // actions:
  function StartLoginServiceListen(ctrl: AbortController) {
    return fetchEventSource(`${Enviroment.BASE_API_URL}/events/services-status`, {
      method: 'GET',
      openWhenHidden: true,
      signal: ctrl.signal,
      onmessage(ev) {
        LoginServices.value = <EventModel.LoginService[]>JSON.parse(ev.data)
      },
      onclose() {
        $reset()
        throw new Error('Stop Listen LoginService Status')
      },
      onerror(err) {
        6
        if (err instanceof Error) {
          throw err // rethrow to stop the operation
        } else {
          // do nothing to automatically retry. You can also
          // reurn a specific retry interval here.
        }
      }
    })
  }

  function GetAllScanFilesDataCount(fileInfo: Map<string, EventModel.ScanFile>): number {
    if (!fileInfo) return 0

    let count = 0
    Object.values(fileInfo).forEach((item: EventModel.ScanFile) => {
      item.dataCount && (count += item.dataCount)
    })
    return count
  }

  function GetLastScanFileTime(fileInfo: Map<string, EventModel.ScanFile>): string {
    if (!fileInfo) return ''

    let lastTime = 0
    Object.values(fileInfo).forEach((item: EventModel.ScanFile) => {
      if (item.lastDataTime && lastTime < item.lastDataTime) lastTime = item.lastDataTime
    })
    return new Date(lastTime * 1000).toLocaleString()
  }

  function GetAliveApplicationCount(appInfo: Map<string, EventModel.ScanApplication>): number {
    if (!appInfo) return 0

    const now = Date.now()
    const apps = Object.values(appInfo)
    let alive = 0
    apps.forEach((item: EventModel.ScanApplication) => {
      if (now - item.lastDataTime * 1000 <= Enviroment.SCAN_APP_DELAY_TIME * 60 * 1000) {
        alive++
      }
    })
    return alive
  }

  function $reset() {
    LoginServices.value = []
    ErrorServiceUuids.value.clear()
  }

  return {
    LoginServices,
    ErrorServiceUuids,
    StartLoginServiceListen,
    GetAllScanFilesDataCount,
    GetLastScanFileTime,
    GetAliveApplicationCount,
    $reset
  }
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useEventStore, import.meta.hot))
}
