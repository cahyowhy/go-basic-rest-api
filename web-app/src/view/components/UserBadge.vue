<!--suppress ALL -->
<template>
    <div class="box">
        <article v-if="isValidUser" class="media">
            <figure class="media-left">
                <p class="image is-64x64 is-cover">
                    <common-image :src="user.image_profile" :isProfile="true" />
                </p>
            </figure>
            <div class="media-content">
                <div class="content">
                    <p>
                        <strong>{{user.name}}</strong>
                        <small>@{{user.username}}</small>
                        <br>
                        <span class="help is-grey">
                            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros
                        </span>
                    </p>
                </div>
                <slot/>
            </div>
        </article>
    </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "annotation";
import { Deserialize } from "cerialize";

import User from "../models/User";

@Component
export default class UserBadge extends Vue {
  @Prop({
    default: ""
  })
  private user: any;

  private get isValidUser() {
      return this.user instanceof User;
  }

  private created() {
    if (typeof this.user === "string") {
      try {
        const userJson = JSON.parse(this.user);

        this.user = Deserialize(userJson, User);
      } catch (err) {
        this.user = new User();
        console.log(err);
      }
    }
  }
}
</script>