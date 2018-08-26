const webpack = require("webpack");

module.exports = {
    entry: "./entry.point",
    output: {
        filename: "prod.inc.js",
        libraryTarget: "this"
    },
    plugins: [
        new webpack.DefinePlugin({
            'process.env': {
                'NODE_ENV': JSON.stringify('production')
            }
        })
    ],
    optimization: {
        minimize: true
    }
};
