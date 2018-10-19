<!--suppress ALL -->
<template>
  <a-row>
    <a-col :span="12">
      <a-form @submit="handleSubmit" :autoFormCreate="(form)=>{this.form = form}">
    <a-form-item
      label='Note'
      :labelCol="{ span: 5 }"
      :wrapperCol="{ span: 12 }"
      fieldDecoratorId="note"
      :fieldDecoratorOptions="{rules: [{ required: true, message: 'Please input your note!' }]}"
    >
      <a-input />
    </a-form-item>
    <a-form-item
      label='Gender'
      :labelCol="{ span: 5 }"
      :wrapperCol="{ span: 12 }"
      fieldDecoratorId="gender"
      :fieldDecoratorOptions="{rules: [{ required: true, message: 'Please select your gender!' }]}"
    >
      <a-select
        placeholder='Select a option and change input text above'
        @change="this.handleSelectChange"
      >
        <a-select-option value='male'>male</a-select-option>
        <a-select-option value='female'>female</a-select-option>
      </a-select>
    </a-form-item>
    <a-form-item
      :wrapperCol="{ span: 12, offset: 5 }"
    >
      <a-button type='primary' htmlType='submit'>
        Submit
      </a-button>
    </a-form-item>
  </a-form>
    </a-col>
  </a-row>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import User from "../models/User";

@Component
export default class FormUser extends Vue {
  private user: User = new User();

  private formLayout: string = "";

  private form: any = null;

  private handleSubmit(e) {
    e.preventDefault();
    (this as any).form.validateFields((err, values) => {
      if (!err) {
        console.log("Received values of form: ", values);
      }
    });
  }

  private handleSelectChange(value) {
    console.log(value);
    (this as any).form.setFieldsValue({
      note: `Hi, ${value === "male" ? "man" : "lady"}!`
    });
  }
}
</script>
