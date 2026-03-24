import { AppProviders } from 'app/providers/AppProviders';
import { DevicesPage } from 'pages/devices/DevicesPage';
import React from 'react';
import { AppLayout } from 'shared/components/layout/AppLayout';

export const App: React.FC = () => {
    return (
        <AppProviders>
            <AppLayout>
                <DevicesPage />
            </AppLayout>
        </AppProviders>
    );
};
