<!--suppress ALL -->
<template>
    <section class="form-login-register">
        <section class="logo-wrp">
            <img class="is-rounded" src="/public/images/todo-logo.svg" alt="Todo">
        </section>
        <form v-on:submit.prevent="doLogin" method="POST">
            <b-field label="Username" :type="user.usernameFeedback().type" :message="user.usernameFeedback().error">
                <b-input v-model="user.username" maxlength="12"></b-input>
            </b-field>
            <b-field label="Password" :type="user.passwordFeedback().type" :message="user.passwordFeedback().error">
                <b-input type="password" v-model="user.password" password-reveal>
                </b-input>
            </b-field>
            <div class="has-text-centered field">
                <input type="submit" value="Login" :class="`button is-${user.validLogin() ? 'info' : 'danger'}`" />
            </div>
        </form>
        <section class="has-text-centered">
            <p class="help is-grey">{{isRegister ? 'Sudah Punya Akun ?' : 'Belum Punya Akun ?'}}</p>
            <p class="help">
                <a data-turbolinks-action="replace" href="/" v-if="isRegister">Login</a>
                <a data-turbolinks-action="replace" href="/?register=true" v-else>Daftar dulu</a>
            </p>
        </section>
    </section>
</template>

<script lang="ts">
import { Vue, Component, Inject } from "annotation";
import environment from "environment";
import { isEmpty } from "lodash";

import User from "../models/User";

import UserService from "../service/UserService";

@Component
export default class FormLogin extends Vue {
  @Inject private userService: UserService;

  private form: any = null;

  private user: User = new User();

  private get isRegister() {
    const query = (this as any).$root.route.query;

    return !isEmpty(query) && !isEmpty(query.register);
  }

  private mounted() {
      console.log(this.$root.$el);
  }

  private async doLogin() {
    if (this.user.validLogin()) {
      const { username, password } = this.user;
      const body = {
        username,
        password
      };
      const data = await this.userService.doLogin(body);
    }
  }
}
</script>