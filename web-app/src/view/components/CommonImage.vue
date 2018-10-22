<!--suppress ALL -->
<template>
<img :src="src" v-img-fallback="imageCallback">
</template>

<script lang="ts">
import { Vue, Component, Prop } from "annotation";

@Component
export default class CommonImage extends Vue {
  @Prop({ default: "" })
  private src: string;

  @Prop({ default: false })
  private isProfile: boolean;

  private get imageCallback() {
    const error = `/public/images/${
      this.isProfile ? "default-user.jpg" : "no-img-found.png"
    }`;
    
    return {
      loading: "/public/images/loading.gif",
      error
    };
  }
}
</script>
