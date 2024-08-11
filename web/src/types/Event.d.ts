declare namespace EventModel {
  type LoginService = {
    uuid: string
    name: string
    ip: string
    scanFiles: Map<string, ScanFile>
    scanApplications: Map<string, ScanApplication>
  }

  type ScanFile = {
    path: string
    dataCount: number
    lastDataTime: number
  }

  type ScanApplication = {
    name: string
    lastDataTime: number
  }
}
