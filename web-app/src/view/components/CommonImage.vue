<!--suppress ALL -->
<template>
  <img :src="src" :href="src" v-img-fallback="imageCallback">
</template>

<script lang="ts">
import { Vue, Component, Prop } from "annotation";

@Component
export default class CommonImage extends Vue {
  @Prop({ default: "" })
  private src: string;

  @Prop({ default: false })
  private isProfile: boolean;

  @Prop({ default: true })
  private isMagnified: boolean;

  private mounted() {
    (window as any).$(this.$el).magnificPopup({
      type: "image"
    });
  }

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
