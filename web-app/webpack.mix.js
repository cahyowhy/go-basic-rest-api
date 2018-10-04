const mix = require('laravel-mix');
const path = require('path');

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for the application as well as bundling up all the JS files.
 |
 */

mix.disableSuccessNotifications();

mix
  .options({
    processCssUrls: false,
  })
  .sass('src/sass/app.scss', "dist/css")
  .js('src/app.ts', "dist/js")
  .copyDirectory("dist","../public")
  .webpackConfig({
    devtool: mix.inProduction() ? '' : 'inline-source-map',
    module: {
      rules: [{
        test: /\.tsx?$/,
        loader: 'ts-loader',
        options: {
          appendTsSuffixTo: [/\.vue$/],
        },
        exclude: /node_modules/,
      }],
    },
    resolve: {
      extensions: ['*', '.js', '.jsx', '.vue', '.ts', '.tsx'],
      alias: {
        styles: path.resolve(__dirname, 'resources/src/sass')
      },
    },
  });