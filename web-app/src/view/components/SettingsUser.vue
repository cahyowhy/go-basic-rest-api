<!--suppress ALL -->
<template>
    <div>
        <div class="nav-wrp">
            <navigation-bar />
        </div>
        <div class="columns">
            <template v-if="isValidUser">
                <div class="column">
                    <user-badge :user="user">
                        <nav class="level is-mobile">
                            <div class="level-left">
                                <b-upload @input="onUploadImageProfile" v-model="userData.file">
                                    <a class="button is-dark is-small">
                                        <i class="mdi mdi-image-area"></i>
                                    </a>
                                </b-upload>
                            </div>
                        </nav>
                    </user-badge>
                </div>
                <div class="column is-two-thirds">
                    <div class="box">
                        <p class="title is-5">Data diri</p>
                        <form v-on:submit.prevent="doSave(false)" method="POST">
                            <b-field label="Name" :type="userData.nameFeedback().type" :message="userData.nameFeedback().error">
                                <b-input v-model="userData.name" placeholder="e.g John Smith" maxlength="30"></b-input>
                            </b-field>
                            <b-field label="Username" :type="userData.usernameFeedback().type" :message="userData.usernameFeedback().error">
                                <b-input v-model="userData.username" placeholder="e.g john_smith01" maxlength="30"></b-input>
                            </b-field>
                            <div class="has-text-right field">
                                <input type="submit" value="Update" :class="`button is-${userData.valid() ? 'info' : 'danger'}`" />
                            </div>
                        </form>
                    </div>

                    <div class="box">
                        <p class="title is-5">Keamanan</p>
                        <form v-on:submit.prevent="doSave(true)" method="POST">
                            <b-field label="Password lama" :type="userData.passwordOldFeedback().type" :message="userData.passwordOldFeedback().error">
                                <b-input type="password" placeholder="Fill with valid old password" v-model="userData.passwordOld" password-reveal>
                                </b-input>
                            </b-field>
                            <b-field label="Password" :type="userData.passwordFeedback().type" :message="userData.passwordFeedback().error">
                                <b-input type="password" placeholder="Fill with valid password" v-model="userData.password" password-reveal>
                                </b-input>
                            </b-field>
                            <b-field v-if="isRegister" label="Password Confirmation" :type="userData.passwordConfirmFeedback().type" :message="userData.passwordConfirmFeedback().error">
                                <b-input type="password" placeholder="Must same as password above" v-model="userData.passwordConfirm" password-reveal>
                                </b-input>
                            </b-field>
                            <div class="has-text-right field">
                                <input type="submit" value="Update Password" :class="`button is-${userData.validUpdatePassword() ? 'info' : 'danger'}`" />
                            </div>
                        </form>
                    </div>
                </div>
            </template>
        </div>
    </div>
</template>

<script lang="ts">
import { Vue, Component, Inject, Prop } from "annotation";
import { Deserialize } from "cerialize";
import { isNil } from "lodash";
import UserService from "../service/UserService";
import CommonService from "../service/CommonService";
import Constant from "../config/Constant";

import UserPhoto from "../models/UserPhoto";
import User from "../models/User";

@Component
export default class SettingsUser extends Vue {
  @Inject private userService: UserService;

  @Inject private commonService: CommonService;

  @Prop({
    default: ""
  })
  private user: any;

  private userData: User = new User();

  private get isValidUser() {
    return this.user instanceof User;
  }

  private created() {
    if (typeof this.user === "string") {
      try {
        const userJson = JSON.parse(this.user);

        this.user = Deserialize(userJson, User);
        this.userData = Deserialize(userJson, User);
      } catch (err) {
        this.user = new User();
        console.log(err);
      }
    }
  }

  private onUploadImageProfile(files: any) {
    const userPhoto = new UserPhoto();
    userPhoto.user_id = this.userData.id;
    userPhoto.file = files;

    this.doSave(false, userPhoto);
  }

  private async doSave(isPassword: boolean, userPhoto: any = null) {
    this.userService.returnWithStatus = true;
    let response = null;

    if (!isNil(userPhoto)) {
      response = await this.userService.uploadImageProfile(userPhoto);
    } else if (isPassword) {
      const { passwordOld, password, id } = this.userData;
      const param = { passwordOld, password, id };

      response = this.userService.updateUserPassword(param);
    } else {
      const id = this.userData.id;
      response = await this.userService.update(this.userData, id.toString());
    }

    this.userService.returnWithStatus = false;
    
    if (
      response.status === Constant.STATUS.API.UPDATE_SUCCESS &&
      !isNil(response.data)
    ) {
      this.commonService.removeUser();
    }
  }
}
</script>