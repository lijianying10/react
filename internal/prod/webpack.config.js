const webpack = require("webpack");

module.exports = {
    entry: "./entry.point",
    output: {
        filename: "prod.inc.js",
        libraryTarget: "this"
    },
};
