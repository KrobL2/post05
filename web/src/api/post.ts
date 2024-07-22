import { AxiosRequestConfig, AxiosResponse } from 'axios';

import { request } from './request';

export const post = <T = never, R = AxiosResponse<T>>(
  url: string,
  data?: T,
  config?: AxiosRequestConfig<T>,
) => {
  return request<T, R>({ method: 'post', url, data, ...config });
};
