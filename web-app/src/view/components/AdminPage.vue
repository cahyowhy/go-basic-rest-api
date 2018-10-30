<!--suppress ALL -->
<template>
    <div>
        <div class="nav-wrp">
            <navigation-bar />
        </div>
        <div class="columns">
            <b-modal :active.sync="modalActive" class="modal-edit-user">
                <header class="modal-card-head">
                    <p class="modal-card-title is-small">Login</p>
                </header>
                <div class="modal-card">
                    <form method="POST" v-on:submit.prevent="doUpdate">
                        <b-field label="Name" :type="user.nameFeedback().type" :message="user.nameFeedback().error">
                            <b-input v-model="user.name" placeholder="e.g John Smith" maxlength="30"></b-input>
                        </b-field>
                        <b-field label="Username" :type="user.usernameFeedback().type" :message="user.usernameFeedback().error">
                            <b-input v-model="user.username" placeholder="e.g john_smith01" maxlength="30"></b-input>
                        </b-field>
                        <div class="has-text-right field">
                            <input type="submit" value="Simpan" :class="`button is-${user.valid() ? 'info' : 'danger'}`" />
                        </div>
                    </form>
                </div>
            </b-modal>
            <div class="column">
                <user-badge :user="userPreview" />
            </div>
            <div class="column is-two-thirds">
                <div class="box">
                    <p class="title is-6">Tabel User</p>
                    <common-table :perPage="paramTableUser.limit" :total="paramTableUser.total" detailKey="id"
                    defaultSort="username" :isDetail="true" :columnProps="userColumns" :datas="userDatas"
                    @pageChange="(page)=>onPageChange(page, 'USER')">
                        <!-- slot detail -->
                        <template slot="table-detail" slot-scope="{props}">
                            <article class="media">
                                <figure class="media-left">
                                    <p class="image is-64x64">
                                        <common-image :imageProfile="true" :src="props.row.path" />
                                    </p>
                                </figure>
                                <div class="media-content">
                                    <div class="content">
                                        <p>
                                            <strong>{{ props.row.name }}</strong>
                                            <small>@{{props.row.username }}</small>
                                            <br> Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.
                                        </p>
                                    </div>
                                </div>
                            </article>
                        </template>
                        <!-- slot action (custom) -->
                        <template v-for="(user, index) in userDatas" :slot="`column-aksi-${index}`" slot-scope="{row}">
                            <a class="button is-info is-small" 
                            :key="`user-action-${index}`" @click="doEdit(row)">
                                <i class="mdi mdi-account-edit"></i>
                            </a>
                        </template>
                    </common-table>
                </div>

                 <div class="box">
                    <p class="title is-6">Tabel User</p>
                    <common-table :perPage="paramTableTodo.limit" :total="paramTableTodo.total" 
                    :columnProps="todoColumns" :datas="todoDatas" @pageChange="(page)=>onPageChange(page, 'TODO')">
                    </common-table>
                 </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { Vue, Component, Prop, Inject } from "annotation";
import environment from "environment";
import TodoService from "../service/TodoService";
import UserService from "../service/UserService";
import CommonService from "../service/CommonService";
import { cloneDeep, isEmpty, isNil } from "lodash";
import Constant from "../config/Constant";

import User from "../models/User";
import Todo from "../models/Todo";

@Component
export default class UserHome extends Vue {
  @Inject private userService: UserService;

  @Inject private commonService: CommonService;

  @Inject private todoService: TodoService;

  private user: User = new User();

  private userPreview: User = new User();

  private users: Array<User> = [];

  private userColumns: Array<any> = User.columnName();

  private todoColumns: Array<any> = Todo.columnName();

  private todos: Array<Todo> = [];

  private paramGetUser: any = { username: "", limit: environment["LIMIT"] };

  private paramGetTodo: any = { user_id: "", limit: environment["LIMIT"] };

  private paramTableUser: any = { limit: environment["LIMIT"], total: 0 };

  private paramTableTodo: any = { limit: environment["LIMIT"], total: 0 };

  private modalActive: boolean = false;

  private get userDatas(): Array<any> {
    return cloneDeep(this.users).map(item => item.table());
  }

  private get todoDatas(): Array<any> {
    return cloneDeep(this.todos).map(item => item.table());
  }

  private mounted() {
    this.userPreview = this.commonService.getUser();

    this.doFind();
    this.doFind(0, "USER");
  }

  private doEdit(param: any) {
    const user = this.users.find((user: User) => user.id === (param || { id: 0 }).id);
    if (user) {
      this.modalActive = true;
      this.user = user;
    }
  }

  private async doFind(offset: number = 0, type: string = "") {
    const { paramGetTodo, paramGetUser, paramGetFormated } = this;
    const paramGet = paramGetFormated(type);
    paramGet.offset = offset;

    switch (type) {
      case "USER":
        const userResponses = await this.userService.find(paramGet, true);
        var { data, count } = userResponses;
        this.users = Array.isArray(data) ? data : [];

        if (count) {
          this.paramTableUser.total = count;
        }
        break;
      default:
        const todoResponses = await this.todoService.find(paramGet, true);
        var { data, count } = todoResponses;
        this.todos = Array.isArray(data) ? data : [];

        if (count) {
          this.paramTableTodo.total = count;
        }
        break;
    }
  }

  private onPageChange(page, type: string) {
    this.doFind((page - 1) * environment["LIMIT"], type);
  }

  private paramGetFormated(type: string = "") {
    if (type === "USER") {
      const { username, limit } = this.paramGetUser;
      if (isEmpty(username)) {
        return { limit };
      }

      return this.paramGetUser;
    }

    const { user_id, limit } = this.paramGetTodo;
    if (isEmpty(user_id)) {
      return { limit };
    }

    return this.paramGetTodo;
  }

  private async doUpdate() {
    this.userService.returnWithStatus = true;
    const { SAVE_SUCCESS, UPDATE_SUCCESS } = Constant.STATUS.API;
    const response = await this.userService.update(this.user, this.user.id.toString());

    if (response.status === UPDATE_SUCCESS) {
      this.user = new User();
      this.doFind(0, "USER");
      this.modalActive = false;
    }

    this.userService.returnWithStatus = false;
  }
}
</script>