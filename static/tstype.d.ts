declare module '*.vue' {
    import Vue = require('vue');

    const value: Vue.ComponentOptions<Vue>;
    export default value;
}

declare module "turbolinks"
declare module "vue-turbolinks"