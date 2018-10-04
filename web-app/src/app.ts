import Vue from 'vue';
import VueTurbolinks from "./view/mixins/vue-turbolinks"
import InputValidator from './view/components/InputValidator.vue';

Vue.use(VueTurbolinks);

document.addEventListener('turbolinks:load', () => {
    new Vue({
        el: '#app',
        components: {
            'input-validator': InputValidator
        }
    });
});