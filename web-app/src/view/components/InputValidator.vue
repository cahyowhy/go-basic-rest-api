<!--suppress ALL -->
<template>
    <input type="text" :class="`input ${isSuccess ? '' : 'is-danger'}`" :placeholder="placeholder" v-model="newValue">
</template>
<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";

@Component
export default class InputValidator extends Vue {
  @Prop({ default: true })
  private isSuccess: Boolean;

  @Prop({ default: "" })
  private value: String;

  @Prop({ default: "" })
  private placeholder: String;

  get newValue() {
    return this.value;
  }

  set newValue(value) {
    this.$emit("input", value);
  }
}
</script>

