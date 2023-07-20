import {AxiosError} from "axios";


export interface ResponseError{
  status: number | undefined;
  statusText: string | undefined;
  message: string | undefined;
}

export class RequestResponseApi<T> {
  error: ResponseError | null;
  data?: T;

  constructor (error: ResponseError | null, data?: T) {
    this.error = error;
    this.data = data;
  }
}
