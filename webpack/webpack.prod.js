const webpack = require('webpack');
const webpackMerge = require('webpack-merge');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const commonConfig = require('./webpack.common.js');
const helpers = require('../helpers');

module.exports = webpackMerge(commonConfig, {
    devtool: 'eval',
    entry: {
        bundle: [
            'babel-polyfill',
            './client/index.js',
        ],
    },
    output: {
        path: helpers.root('dist'),
        filename: '[name].[chunkhash].js',
    },
    plugins: [
        new webpack.optimize.CommonsChunkPlugin({
            names: ['vendor', 'manifest'],
        }),
        new CleanWebpackPlugin(['dist'], {
            root: helpers.root(),
            verbose: true,
            watch: true,
        }),
        new HtmlWebpackPlugin({
            template: './client/index.html',
        }),
        new webpack.optimize.UglifyJsPlugin(),
        new webpack.optimize.AggressiveMergingPlugin(),
    ],
});
