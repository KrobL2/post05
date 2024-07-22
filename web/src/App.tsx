import React, { FC } from 'react';
import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';

import { hot } from 'react-hot-loader/root';
import { ConfigProvider } from 'antd';

import { Layout } from 'layout';

import { SwitchRoutes } from 'pages/SwitchRoutes';

import { store } from 'store';

const App: FC = () => (
  <>
    <Provider store={store}>
      
        <ConfigProvider >
          <BrowserRouter>
              <Layout>
                <SwitchRoutes />
              </Layout>
          </BrowserRouter>
        </ConfigProvider>
     
    </Provider>
  
   
  </>
);

export default hot(App);
