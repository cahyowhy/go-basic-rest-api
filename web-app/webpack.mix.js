const mix = require("laravel-mix");
const path = require("path");
const resolve = file => path.resolve(__dirname, file);
const isProduction = mix.inProduction();
var LiveReloadPlugin = require("webpack-livereload-plugin");
const environmentPath = resolve("src/view/config/environment");
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
  .scripts(
    ["../public/vendor/js/magnificPopUp.min.js"],
    "../public/vendor/js/non-priority.vendor.js"
  )
  .styles(
    ["../public/vendor/css/magnificPopUp.min.css"],
    "../public/vendor/css/non-priority.vendor.css"
  )
  .options({
    processCssUrls: false
  })
  .sass("src/sass/app.scss", "css")
  .js("src/app.ts", "js")
  .webpackConfig({
    output: {
      filename: '[name].js',
      path: path.resolve(__dirname, "dist"),
      publicPath: '/public/',
      chunkFilename: 'js/[chunkhash][name].js',
    },
    devtool: mix.inProduction() ? "" : "inline-source-map",
    module: {
      rules: [
        {
          test: /\.tsx?$/,
          loader: "ts-loader",
          options: {
            appendTsSuffixTo: [/\.vue$/]
          },
          exclude: /node_modules/
        }
      ]
    },
    plugins: [new LiveReloadPlugin()],
    resolve: {
      modules: [resolve("src"), resolve("node_modules")],
      extensions: ["*", ".js", ".jsx", ".vue", ".ts", ".tsx"],
      alias: {
        window: "window",
        $: "window.$",
        jQuery: "window.$",
        vue$: "vue/dist/vue.js",
        "typescript-ioc": "typescript-ioc/es6",
        annotation: path.resolve(__dirname, "src/view/util/Annotation.ts"),
        environment:
          environmentPath +
          `${isProduction ? "/production" : "/development"}.env.ts`
      }
    }
  })
  .copyDirectory("dist", "../public");
