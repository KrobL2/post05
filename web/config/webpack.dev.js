/* eslint-disable @typescript-eslint/no-var-requires */
const DefinePlugin = require('webpack').DefinePlugin;
const ESLintPlugin = require('eslint-webpack-plugin');
const ProvidePlugin = require('webpack').ProvidePlugin;
const CopyPlugin = require('copy-webpack-plugin');

const { merge } = require('webpack-merge');
const commonConfig = require('./webpack.common');

const proxy = require('./proxy');

const isDev = process.env.STAND === 'dev';

module.exports = merge(commonConfig, {
  mode: 'development',
  target: 'web',
  plugins: [
    new ESLintPlugin({
      extensions: ['ts', 'tsx', 'js'],
      formatter: 'compact',
    }),
    new ProvidePlugin({
      process: 'process/browser',
    }),
    new DefinePlugin({
      'process.env.NODE_ENV': JSON.stringify('development'),
      'process.env.STAND': JSON.stringify('dev'),
    }),
    new CopyPlugin({
      patterns: [
        {
          from: 'public/env.config.js',
          // Подмена динамических переменных окружения
          // для локальной разработки
          transform: function (buffer) {
            if (isDev) return buffer;
            const content = buffer
              .toString()
              .replace(
                '$SAMOLET_OAUTH2_PKCE_CLIENT_ID',
                'FDuFhN9sBrNvp6F4fKXNw2tManBm6IMeubEMHBkO',
              );

            return content;
          },
        },
      ],
    }),
  ],
  devtool: 'eval-source-map',
  devServer: {
    client: {
      logging: 'error',
      overlay: {
        errors: false,
        warnings: false,
        runtimeErrors: false,
      },
      progress: false,
    },
    compress: true,
    historyApiFallback: true,
    hot: true,
    open: true,
    port: 3000,
    proxy: proxy,
  },
  cache: {
    type: 'filesystem',
    buildDependencies: {
      config: [__filename],
    },
  },
});
