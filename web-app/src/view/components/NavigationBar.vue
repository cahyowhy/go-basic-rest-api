<!--suppress ALL -->
<template>
    <nav class="navbar" role="navigation" aria-label="main navigation">
        <div class="navbar-brand">
            <a class="navbar-item" data-turbolinks-action="replace" href="/home">
                <img src="/public/images/todo-logo.svg" width="30" height="30">
                <span>
                    TodoApp
                </span>
            </a>
        </div>

        <div id="navbarBasicExample" class="navbar-menu">
            <div class="navbar-end">
                <a class="navbar-item" data-turbolinks-action="replace" href="/setting">
                    Setting
                </a>
                <div class="navbar-item">
                    <div class="buttons">
                        <a @click="doLogout" class="button is-danger">
                            Log out
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </nav>
</template>

<script lang="ts">
import { Vue, Component, Inject } from "annotation";
import CommonService from "../service/CommonService";

@Component
export default class NavigationBar extends Vue {
    @Inject private commonService: CommonService;

    private doLogout() {
        this.commonService.removeUser();
    }
}
</script>
