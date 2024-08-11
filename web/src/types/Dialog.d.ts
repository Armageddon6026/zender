declare namespace Dialog {
  type ConfirmMessage = {
    visible: boolean
    header?: string
    body?: string
    footer?: string
    level?: 'Warning' | 'Normal' | 'Critic'
    args: any[]
  }

  type InputMessage = {
    visible: boolean
    header?: string
    args: any[]
  }

  type ServiceInputMessage = InputMessage & {
    service: RequestModel.Service
  }

  type GroupInputMessage = InputMessage & {
    group: RequestModel.Group
  }
}
