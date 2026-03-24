import {
    DesktopOutlined,
    GlobalOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    UserOutlined
} from '@ant-design/icons';
import { Avatar, Breadcrumb, Layout, Menu, Space, theme, Typography } from 'antd';
import React, { useState } from 'react';

const { Header, Sider, Content, Footer } = Layout;
const { Title, Text } = Typography;

interface AppLayoutProps {
    children: React.ReactNode;
}

export const AppLayout: React.FC<AppLayoutProps> = ({ children }) => {
    const [collapsed, setCollapsed] = useState(false);
    const {
        token: { colorBgContainer, borderRadiusLG }
    } = theme.useToken();

    return (
        <Layout style={{ minHeight: '100vh' }}>
            <Sider trigger={null} collapsible collapsed={collapsed} theme="dark" breakpoint="lg">
                <div
                    style={{
                        height: 64,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        background: 'rgba(255, 255, 255, 0.05)',
                        margin: '16px',
                        borderRadius: '8px'
                    }}
                >
                    <GlobalOutlined style={{ fontSize: '24px', color: '#1890ff' }} />
                    {!collapsed && (
                        <Title level={4} style={{ color: '#fff', margin: '0 0 0 12px' }}>
                            TesTTask
                        </Title>
                    )}
                </div>

                <Menu
                    theme="dark"
                    mode="inline"
                    defaultSelectedKeys={['1']}
                    items={[
                        {
                            key: '1',
                            icon: <DesktopOutlined />,
                            label: 'Устройства'
                        }
                    ]}
                />
            </Sider>

            <Layout>
                <Header
                    style={{
                        padding: 0,
                        background: colorBgContainer,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'space-between',
                        paddingRight: '24px',
                        boxShadow: '0 1px 4px rgba(0,21,41,.08)',
                        zIndex: 1
                    }}
                >
                    <div style={{ display: 'flex', alignItems: 'center' }}>
                        <div
                            style={{ padding: '0 24px', fontSize: '18px', cursor: 'pointer' }}
                            onClick={() => setCollapsed(!collapsed)}
                        >
                            {collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                        </div>

                        <Breadcrumb items={[{ title: 'Главная' }, { title: 'Устройства' }]} />
                    </div>

                    <Space size="middle">
                        <div style={{ textAlign: 'right', lineHeight: '1.2' }}>
                            <Text strong style={{ display: 'block' }}>
                                Admin
                            </Text>
                        </div>
                        <Avatar icon={<UserOutlined />} style={{ backgroundColor: '#1890ff' }} />
                    </Space>
                </Header>
                <Content
                    style={{
                        margin: '24px 16px',
                        padding: 24,
                        minHeight: 280,
                        background: colorBgContainer,
                        borderRadius: borderRadiusLG,
                        overflow: 'initial'
                    }}
                >
                    {children}
                </Content>
                <Footer style={{ textAlign: 'center', color: '#bfbfbf' }}>
                    lumiere MDM ©{new Date().getFullYear()} Created for Test Assignment
                </Footer>
            </Layout>
        </Layout>
    );
};
