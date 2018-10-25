<!--suppress ALL -->
<template>
    <div>
        <div class="nav-wrp">
            <navigation-bar />
        </div>
        <div class="columns">
            <div class="column">
                <user-badge :user="user" />
            </div>
            <div class="column is-two-thirds">
                <div class="box">
                    <b-table :striped="true" :narrowed="true" :bordered="true" :data="userDatas" :columns="userColumns" @page-change="(page) => onPageChange(page, 'USER')" :total="20" backend-pagination paginated per-page="9" detailed detail-key="id" @details-open="(row, index) => $toast.open(`Expanded ${props.row.username}`)">
                        <template slot="detail" slot-scope="props">
                            <article class="media">
                                <figure class="media-left">
                                    <p class="image is-64x64">
                                        <common-image :src="props.row.path" />
                                    </p>
                                </figure>
                                <div class="media-content">
                                    <div class="content">
                                        <p>
                                            <strong>{{ props.row.name }}</strong>
                                            <small>@{{props.row.username }}</small>
                                            <br>
                                            Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                                            Proin ornare magna eros, eu pellentesque tortor vestibulum ut.
                                            Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.
                                        </p>
                                    </div>
                                </div>
                            </article>
                        </template>
                    </b-table>
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
import { cloneDeep } from "lodash";

import User from "../models/User";
import Todo from "../models/Todo";

@Component
export default class UserHome extends Vue {
  @Inject
  private userService: UserService;

  @Inject
  private commonService: CommonService;

  @Inject
  private todoService: TodoService;

  private user: User = new User();

  private users: Array<User> = [];

  private userColumns: Array<any> = User.columnName();

  private todos: Array<Todo> = [];

  private get userDatas(): Array<any> {
    return cloneDeep(this.users).map(item => item.table());
  }

  private mounted() {
    this.user = this.commonService.getUser();

    this.doFind();
    this.doFind(0, "USER");
  }

  private async doFind(offset: number = 0, type: string = "") {
    const param = { offset, limit: environment["LIMIT"] };

    switch (type) {
      case "USER":
        const userResponses = await this.userService.find(param);
        this.users = Array.isArray(userResponses) ? userResponses : [];
        break;
      default:
        const todoResponses = await this.todoService.find(param);
        this.todos = Array.isArray(todoResponses) ? todoResponses : [];
        break;
    }
  }

  private onPageChange(page, type: string) {
    this.doFind(page * environment["LIMIT"], type);
  }
}
</script>