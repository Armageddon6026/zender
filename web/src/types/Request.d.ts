declare namespace RequestModel {
  type User = {
    name: string
    password: string
  }

  type Service = {
    name: string
    groupId: number
  }

  type Group = {
    name: string
    note: string
  }
}
