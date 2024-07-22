import { AxiosRequestConfig, AxiosResponse } from 'axios';

import { request } from './request';

export const get = <T, R = AxiosResponse<T>, D = unknown>(
  url: string,
  params?: AxiosRequestConfig['params'],
  config?: AxiosRequestConfig<D>,
) => {
  const commonParams = {
    ...config,
    params,
  };
  return request<T, R>({ method: 'get', url, ...commonParams });
};
