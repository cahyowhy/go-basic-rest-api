const mix = require('laravel-mix');
const path = require('path');
const resolve = (file) => path.resolve(__dirname, file);
const isProduction = mix.inProduction();
const environmentPath = resolve('src/view/config/environment');
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
var LiveReloadPlugin = require('webpack-livereload-plugin');

mix
  .options({
    processCssUrls: false,
  })
  .sass('src/sass/app.scss', "dist/css")
  .js('src/app.ts', "dist/js")
  .copyDirectory("dist", "../public")
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
    plugins: [
      new LiveReloadPlugin()
    ],
    resolve: {
      modules: [resolve('src'), resolve('node_modules')],
      extensions: ['*', '.js', '.jsx', '.vue', '.ts', '.tsx'],
      alias: {
        'window': 'window',
        '$': "window.$",
        "jQuery": "window.$",
        'vue$': 'vue/dist/vue.js',
        'typescript-ioc': 'typescript-ioc/es6',
        'annotation': path.resolve(__dirname, 'src/view/util/Annotation.ts'),
        'environment': environmentPath + `${isProduction ? '/production' : '/development'}.env.ts`,
      }
    },
  });