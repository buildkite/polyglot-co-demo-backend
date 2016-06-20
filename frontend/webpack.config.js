var webpack = require('webpack');
var postcssImport = require('postcss-import');

module.exports = {
  entry: "./index.js",
  output: {
    filename: "polyglot-co.js"
  },
  module: {
    loaders: [
      {
        test: /\.js$/,
        loaders: ['babel']
      },
      {
        test: /\.css$/,
        loaders: ['style','css','postcss']
      },
      {
        test: /\.(png|svg|jpg|gif)$/i,
        loaders: ['url-loader','image-webpack']
      }
   ]
  },
  postcss: function (webpack) {
    return [postcssImport({addDependencyTo: webpack})];
  },
  plugins: [
    new webpack.DefinePlugin({
      'process.env':{
        'NODE_ENV': JSON.stringify(process.env.NODE_ENV)
      }
    })
  ]
}