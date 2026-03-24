import { AppProviders } from 'app/providers/AppProviders';
import { AuditPage } from 'pages/audit/AuditPage';
import { DevicesPage } from 'pages/devices/DevicesPage';
import React from 'react';
import { Route, Routes } from 'react-router-dom';
import { AppLayout } from 'shared/components/layout/AppLayout';

export const App: React.FC = () => {
    return (
        <AppProviders>
            <AppLayout>
                <Routes>
                    <Route path="/" element={<DevicesPage />} />
                    <Route path="/audit" element={<AuditPage />} />
                </Routes>
            </AppLayout>
        </AppProviders>
    );
};
