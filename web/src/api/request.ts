import { AxiosRequestConfig, AxiosResponse } from 'axios';

import { createAxiosInstance } from 'api/createAxiosInstance';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const request = <T, R = AxiosResponse<T>, D = any>(config: AxiosRequestConfig<D>) => {
  const instance = createAxiosInstance();
  return instance<T, R, D>(config);
};
