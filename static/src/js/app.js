import Vue from "vue/dist/vue.esm";
import TurbolinksAdapter from "vue-turbolinks";
import Turbolinks from "turbolinks";

//component
import InputValidator from "./components/InputValidator.vue"

Turbolinks.start();

Vue.use(TurbolinksAdapter);

document.addEventListener("turbolinks:load", () => {
  var vueapp = new Vue({
    el: "#app",
    "components":{
      "input-validator": InputValidator
    }
  });
});