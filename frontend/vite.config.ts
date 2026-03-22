import react from '@vitejs/plugin-react';
import path from 'path';
import { defineConfig } from 'vite';
import viteTsconfigPaths from 'vite-tsconfig-paths';

const config = defineConfig({
    base: '/',
    build: {
        outDir: 'build',
        reportCompressedSize: false,
        manifest: true,
        sourcemap: false,
        minify: 'esbuild',
        target: 'es2020',
        cssCodeSplit: true
    },

    optimizeDeps: {
        include: ['react', 'react-dom'],
        force: true
    },

    resolve: {
        alias: {
            react: path.resolve(__dirname, 'node_modules/react'),
            'react-dom': path.resolve(__dirname, 'node_modules/react-dom'),
            'rc-util': path.resolve(__dirname, 'node_modules/rc-util')
        },
        dedupe: ['react', 'react-dom', 'rc-util']
    },

    plugins: [
        react({
            jsxRuntime: 'automatic'
        }),
        viteTsconfigPaths()
    ],

    server: {
        host: true,
        open: true,
        port: 3000
    }
});

export default config;
