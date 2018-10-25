<!--suppress ALL -->
<template>
    <section class="form-login-register">
        <b-notification v-if="isLoginFirst" type="is-info" has-icon>
            Anda harus masuk terlebih dahulu!
        </b-notification>
        <section class="logo-wrp">
            <img class="is-rounded" src="/public/images/todo-logo.svg" alt="Todo">
        </section>
        <form v-on:submit.prevent="doLoginRegister" method="POST">
            <b-field v-if="isRegister" label="Name" :type="user.nameFeedback().type" :message="user.nameFeedback().error">
                <b-input v-model="user.name" placeholder="e.g John Smith" maxlength="30"></b-input>
            </b-field>
            <b-field label="Username" :type="user.usernameFeedback().type" :message="user.usernameFeedback().error">
                <b-input v-model="user.username" placeholder="e.g john_smith01" maxlength="30"></b-input>
            </b-field>
            <b-field label="Password" :type="user.passwordFeedback().type" :message="user.passwordFeedback().error">
                <b-input type="password" placeholder="Fill with valid password" v-model="user.password" password-reveal>
                </b-input>
            </b-field>
            <b-field v-if="isRegister" label="Password Confirmation" :type="user.passwordConfirmFeedback().type" :message="user.passwordConfirmFeedback().error">
                <b-input type="password" placeholder="Must same as password above" v-model="user.passwordConfirm" password-reveal>
                </b-input>
            </b-field>
            <div class="has-text-centered field">
                <input type="submit" :value="isRegister ? 'Register' : 'Login'" :class="`button is-${(isRegister ? user.validRegister() : user.validLogin()) ? 'info' : 'danger'}`" />
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
import Constant from "../config/Constant";

import User from "../models/User";

import UserService from "../service/UserService";
import CommonService from "../service/CommonService";

@Component
export default class FormLogin extends Vue {
  @Inject private userService: UserService;
  
  @Inject private commonService: CommonService;

  private user: User = new User();

  private get isRegister() {
    const query = (this as any).$root.route.query;

    return !isEmpty(query) && !isEmpty(query.register);
  }

  private get isLoginFirst() {
    const query = (this as any).$root.route.query;
    
    return !isEmpty(query) && !isEmpty(query["login-first"]);
  }

  private async doLoginRegister() {
    const { isRegister } = this;
    this.userService.returnWithStatus = true;

    if (this.user.validLogin()) {
      const method = isRegister ? "save" : "doLogin";
      const payload = isRegister ? this.user : this.user.loginProperty();
      const data = await this.userService[method](payload);
      const status = (data || { status: "" }).status;

      if (isRegister && status === Constant.STATUS.API.SAVE_SUCCESS) {
        (window as any).Turbolinks.visit("/");
      }

      if (!isRegister && status === Constant.STATUS.API.LOGIN_SUCCESS) {
        (window as any).location = "home";
      }
    }

    this.userService.returnWithStatus = false;
  }
}
</script>