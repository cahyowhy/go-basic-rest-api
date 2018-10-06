<!--suppress ALL -->
<template>
  <div class="form-user">
    <div class="field">
      <label class="label">Name</label>
      <div class="control">
        <input-validator v-model="user.name" placeholder="Nama user" />
      </div>
      <div class="field">
        <div class="control">
          <a @click="doSave" class="button is-info">Save</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
  import {
    Vue,
    Component
  } from "vue-property-decorator";
  import User from '../models/User';
  
  @Component
  export default class FormUser extends Vue {
  
    private user: User = new User();

    private doSave() {
      
    }
  }
</script>
