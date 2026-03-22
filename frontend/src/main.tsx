import React from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';

import { App } from 'antd';
import './index.css';

const rootElement = document.getElementById('root');

const NetHubMdm: React.FC = () => (
    <BrowserRouter>
        <App />
    </BrowserRouter>
);

if (rootElement) {
    const container = document.getElementById('root');
    const root = createRoot(container!);
    root.render(<NetHubMdm />);
} else {
    // eslint-disable-next-line no-console
    console.error('Root element not found!');
}

export default NetHubMdm;
