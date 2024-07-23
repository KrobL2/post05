import React, { FC } from 'react';

import { Breadcrumb, Layout as AntLayout, Menu, theme } from 'antd';

import { items } from './items';

const { Content, Sider } = AntLayout;

export const Layout: FC = () => {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  const onClick = ({ item }: any) => {
    const { props } = item;
    console.log('item', item);

    const { link } = props;
    console.log('1', link);
  };

  return (
    <>
      <>
        <div className='demo-logo' />
      </>

      <AntLayout>
        <Sider width={200} style={{ background: colorBgContainer }}>
          <Menu
            mode='inline'
            defaultSelectedKeys={['1']}
            defaultOpenKeys={['sub1']}
            style={{ height: '100%', borderRight: 0 }}
            items={items}
            onClick={onClick}
          />
        </Sider>

        <AntLayout style={{ padding: '0 24px 24px' }}>
          <Breadcrumb style={{ margin: '16px 0' }}>
            <Breadcrumb.Item>Home</Breadcrumb.Item>
            <Breadcrumb.Item>List</Breadcrumb.Item>
            <Breadcrumb.Item>App</Breadcrumb.Item>
          </Breadcrumb>

          <Content
            style={{
              padding: 24,
              margin: 0,
              minHeight: 280,
              background: colorBgContainer,
              borderRadius: borderRadiusLG,
            }}
          >
            Content
          </Content>
        </AntLayout>
      </AntLayout>
    </>
  );
};
