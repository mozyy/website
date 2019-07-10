const path = require('path');
const fs = require('fs');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const WorkboxWebpackPlugin = require('workbox-webpack-plugin');

const appPath = {
  entry: './src/entries/'
}
const entries = fs.readdirSync(appPath.entry)
  .reduce((a,b)=>Object.assign(a,{[b.replace(/\.ts|\.js/,'')]: appPath.entry + b}),{})

console.log(entries)
module.exports = {
  entry: entries,
  output: {
    filename: '[name].js',
    path: path.resolve(__dirname, 'dist')
  },
  mode: 'development',
  devtool: 'inline-source-map',
  devServer: {
    // contentBase: './dist',
    historyApiFallback: {
      rewrites: 
      // [
      //   { from: /^\/clock.*$/, to: '/clock.html' },
      //   { from: /^\/lit.*$/, to: '/lit.html' },
      // ]
       Object.keys(entries).map(entry=>({ from: new RegExp(`^\\/${entry}.*$`), to: `/${entry}.html` }))
    }
  },
  resolve: {
    extensions: [ '.tsx', '.ts', '.js' ]
  },
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/
      },
      {
        test: /\.css$/,
        use: [
          'style-loader',
          'css-loader'
        ]
      },
      {
        test: /\.(png|svg|jpg|gif)$/,
        use: [
          'file-loader'
        ]
      },
      {
        test: /\.(woff|woff2|eot|ttf|otf)$/,
        use: [
          'file-loader'
        ]
      },
    ]
  },
  plugins: [
    new CleanWebpackPlugin(),
    ...(Object.keys(entries).map(entry=>new HtmlWebpackPlugin({
      filename: entry + '.html',
      chunks: [entry],
      inject:true,
      template: 'index.html'
    }))),
    new HtmlWebpackPlugin({
      filename: 'index.html',
      chunks: ['clock'],
      inject:true,
      template: 'index.html'
    })
    // new WorkboxWebpackPlugin.GenerateSW({
    //   clientsClaim: true,
    //   exclude: [/\.map$/, /asset-manifest\.json$/],
    //   importWorkboxFrom: 'cdn',
    //   navigateFallback: `${''}/index.html`,
    //   navigateFallbackBlacklist: [
    //     // Exclude URLs starting with /_, as they're likely an API call
    //     new RegExp('^/_'),
    //     // Exclude URLs containing a dot, as they're likely a resource in
    //     // public/ and not a SPA route
    //     new RegExp('/[^/]+\\.[^/]+$'),
    //   ],
    // }),
    // new WorkboxWebpackPlugin.GenerateSW({
    //   // these options encourage the ServiceWorkers to get in there fast
    //   // and not allow any straggling "old" SWs to hang around
    //   clientsClaim: true,
    //   skipWaiting: true
    // })
  ],
};