import React, { FC } from 'react';

export const Layout: FC<Props> = ({ children }) => {
  return (
    <Layout
     
      sider={{
       
    
        menu: [{ key: '1', title: '', items: menu }],
       
      }}
    >
      {children}
    </Layout>
  );
};

interface Props {
  children: ReactNode;
}
