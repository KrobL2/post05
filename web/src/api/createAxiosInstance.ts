import axios, { isAxiosError } from 'axios';
import url from 'url';

import { RequestUrlsObjType } from 'api/types';

const redirectLocation = (path: string) => {
  const { location } = window;

  const redirectUrl = url.format({
    pathname: path,
    query: { next: location.pathname },
  });

  location.replace(redirectUrl);
};

interface ValidationError {
  message: string;
  errors: Record<string, string[]>;
}

const currentExecutingRequests: RequestUrlsObjType = {};

export const createAxiosInstance = () => {
  const instance = axios.create({
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
  });

  /* Перехватываем запросы с тем же URL и выполняем abort */
  instance.interceptors.request.use(config => {
    const { url } = config;

    if (url && currentExecutingRequests[url]) {
      currentExecutingRequests[url]?.abort();
      currentExecutingRequests[url] = null;
    }

    if (url) {
      currentExecutingRequests[url] = new AbortController();
      return {
        ...config,
        signal: currentExecutingRequests[url]?.signal,
      };
    } else {
      return config;
    }
  });

  /* Перехватываем ответы  */
  instance.interceptors.response.use(
    responce => responce,
    error => {
      if (isAxiosError<ValidationError, Record<string, unknown>>(error)) {
       
        const { response } = error;

        if (response) {
          const { statusText, status, data } = response;

        

          switch (status) {
            case 400:
              console.error('Bad request', data);
              break;

            case 401:
              // console.error('Unauthorised', error);
              // window.location.replace('/accounts/login/samolet');
              redirectLocation('/accounts/login/samolet');
              break;

            case 403:
              // console.error('Forbidden. Access denied', error);
              // window.location.replace('/accounts/login/samolet');
              redirectLocation('/accounts/login/samolet');
              break;

            case 404:
              console.error(statusText);
              break;

            case 500:
              console.error('Server-error', error);
              break;
          }
        } else if (error.request) {
          // This error is most commonly caused by a bad/spotty network,
          // a hanging backend that does not respond instantly to each request,
          // unauthorized or cross-domain requests, and lastly
          // if the backend API returns an error.

          // Отмененные запросы
          // Нет интернета
          console.log(error.request);
        } else {
          console.log('Error', error.message);
        }
      } else {
        console.error('Not Axios Error', error);
      }

      return Promise.reject(error);
    },
  );
  
  return instance;
};
