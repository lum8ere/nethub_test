import { App as AntdApp, ConfigProvider } from 'antd';
import ruRU from 'antd/locale/ru_RU';
import React from 'react';

interface AppProvidersProps {
    children: React.ReactNode;
}

export const AppProviders: React.FC<AppProvidersProps> = ({ children }) => {
    return (
        <ConfigProvider
            locale={ruRU}
            theme={{
                token: {
                    colorPrimary: '#FF793F',
                    borderRadius: 6,
                    fontFamily:
                        '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial'
                },
                components: {
                    Layout: {
                        headerBg: '#ffffff'
                    }
                }
            }}
        >
            <AntdApp>{children}</AntdApp>
        </ConfigProvider>
    );
};
