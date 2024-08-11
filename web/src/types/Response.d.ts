declare namespace ResponseModel {
  type Service = {
    uuid: string
    name: string
    privateKey: string
    groupId: number
  }

  type Jwt = {
    token: string
  }

  type User = {
    account: string
    name: string
    date: number
  }

  type Group = {
    id: number
    name: string
    note: string
  }
}
