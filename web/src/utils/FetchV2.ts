export class FetchV2 {
  private ourBaseURL: string
  private ourRequestInterceptors: Array<(url: string, option: RequestInit) => void>
  private ourResponseInterceptors: Array<(response: Response) => void>
  /**
   * Http request library base on native fetch api
   * @param url A base api url
   */
  constructor(url: string) {
    this.ourBaseURL = url
    this.ourRequestInterceptors = []
    this.ourResponseInterceptors = []
  }

  /**
   * GET method of restful api
   * @param path An api path `ex: /user/Zak`
   * @returns A response data
   */
  public async Get<T>(path: string): Promise<T> {
    return this.handleRequest('GET', path, true)
  }

  /**
   * POST method of restful api
   * @param path An api path `ex: /user`
   * @param body An object of data
   * @returns A response data
   */
  public async Post<T>(path: string, body: any): Promise<T> {
    return this.handleRequest('POST', path, true, undefined, body)
  }

  /**
   * PUT method of restful api
   * @param path An api path `ex: /user`
   * @param body An object of data
   * @returns A response data
   */
  public async Put<T>(path: string, body: any): Promise<T> {
    return this.handleRequest('PUT', path, true, undefined, body)
  }

  /**
   * DELETE method of restful api
   * @param path An api path `ex: /user`
   * @returns A response data
   */
  public async Delete<T>(path: string): Promise<T> {
    return this.handleRequest('DELETE', path, true)
  }

  /**
   * Authorize method base on pure Fetch without any interceptors
   * @param path An api path `ex: /user`
   * @param authorization A token which is added in headers for authorization
   * @returns A response data
   */
  public async Auth<T>(path: string, authorization: string): Promise<T> {
    return this.handleRequest('POST', path, false, { Authorization: authorization })
  }

  /**
   * Add a request interceptor to instance
   * @param interceptor An interceptor solved before send request
   */
  public AddRequestInterceptor(interceptor: (url: string, option: RequestInit) => void) {
    this.ourRequestInterceptors.push(interceptor)
  }

  /**
   * Add a response interceptor to instance
   * @param interceptor An interceptor solved after recieve response
   */
  public AddResponseInterceptors(interceptor: (response: Response) => void) {
    this.ourResponseInterceptors.push(interceptor)
  }

  private runRequestInterceptor(url: string, option: RequestInit) {
    this.ourRequestInterceptors.forEach((interceptor) => interceptor(url, option))
  }

  private runResponsetInterceptor(response: Response) {
    this.ourResponseInterceptors.forEach((interceptor) => interceptor(response))
  }

  private async handleRequest<T>(
    method: string,
    path: string,
    requireInterceptor: boolean,
    header?: HeadersInit,
    body?: any
  ): Promise<T> {
    const apiUrl = this.ourBaseURL + path
    const requestOption: RequestInit = {
      method: method,
      headers: {
        ...header,
        'Content-Type': 'application/json'
      }
    }
    body && (requestOption.body = JSON.stringify(body))

    if (requireInterceptor) {
      this.runRequestInterceptor(apiUrl, requestOption)
    }
    const response = await fetch(apiUrl, requestOption)
    this.runResponsetInterceptor(response)
    if (!response.ok) {
      throw new Error(response.statusText)
    }
    return <T>await response.json()
  }
}
