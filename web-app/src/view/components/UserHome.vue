<!--suppress ALL -->
<template>
    <div class="columns">
        <div class="column">
            <div class="box">
                <article class="media">
                    <figure class="media-left">
                        <p class="image is-64x64 is-cover">
                            <common-image :src="userData.image_profile" :isProfile="true" />
                        </p>
                    </figure>
                    <div class="media-content">
                        <div class="content">
                            <p>
                                <strong>{{userData.name}}</strong> <small>@{{userData.username}}</small>
                                <br> <span class="help is-grey">
                                                                                            Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
                                                                                        Proin ornare magna eros
                                                                                        </span>
                            </p>
                        </div>
                    </div>
                </article>
            </div>
            <div class="box">
                <p class="title is-6">Upload User Foto</p>
                <div class="file-upload-wrp">
                    <b-upload v-model="userPhoto.file">
                        <a class="button is-light">
                            <b-icon icon="upload"></b-icon>
                        </a>
                    </b-upload>
                    <span class="file-name">
                                                            {{ userPhoto.file ? userPhoto.file.name : 'Pilih file' }}
                                                        </span>
                </div>
                <a @click="doSave(true)" v-if="userPhoto.file" class="bt-upload button is-white">Upload</a>
            </div>
            <div class="box">
                <p class="title is-6">User Foto</p>
                <template v-if="userPhotos.length > 0">
                                        <div class="photo-view">
                                        <p v-for="(userPhoto, index) in userPhotos" :key="'user-photo-'+index" class="image is-cover">
                                            <common-image :src="userPhoto.path" />
                                        </p>
                                    </div>
                                        <a @click="() => hasMoreUserPhoto ? doFind(userPhotos.length) : null" class="button is-white bt-load">
                                                                    {{hasMoreUserPhoto ? 'Cari' : 'Tak ada lagi konten dimuat'}}
                                                                </a>
</template>
                <empty-states v-else :isVertical="true"/>
            </div>
        </div>
        <div class="column is-two-thirds">
            <common-editor v-model="todo.content" @update="({html})=> todo.content = html">
                <b-field slot="slot-before" label="Name" :type="todo.nameFeedback().type" :message="todo.nameFeedback().error">
                    <b-input v-model="todo.name" size="is-small" placeholder="e.g Kill Your self" maxlength="50"></b-input>
                </b-field>
                <template slot="slot-after">
                    <b-field label="Select a date">
                        <b-datepicker :min-date="todo.minDateDue" v-model="todo.due" size="is-small" placeholder="Click to select..." icon="calendar-today">
                        </b-datepicker>
                    </b-field>
                    <b-field>
                        <p class="control">
                            <a @click="doSave(false)" :class="`button is-${todo.valid() ? 'info' : 'danger'} has-icon`">
                                <span class="icon"><i class="mdi mdi-content-save"></i></span>
                                <span>Save</span>
                            </a>
                        </p>
                    </b-field>
                </template>
            </common-editor>
            <div class="todo-item box"  v-for="(todo, index) in todos" :key="'todo-item-'+index">
                <div class="content">
                    <div>
                        <p class="title is-5">{{todo.name}}</p>
                        <div class="content-todo" v-html="todo.content"></div>
                        <div class="field">
                            <p class="help is-grey">{{todo.user.name}}</p>
                            <span>•</span>
                            <p class="help is-grey">{{todo.created_at}}</p>
                        </div>
                    </div>
                </div>
            </div>
    
            <a @click="() => hasMoreTodo ? doFind(todos.length) : null" class="button is-white bt-load">
                                                {{hasMoreTodo ? 'Cari' : 'Tak ada lagi konten dimuat'}}
                                            </a>
        </div>
    </div>
</template>

<script lang="ts">
    import {
        Vue,
        Component,
        Prop,
        Inject
    } from "annotation";
    import {
        Deserialize
    } from "cerialize";
    import environment from "environment";
    import {
        isEmpty,
        isNil
    } from "lodash";
    import Constant from "../config/Constant";
    
    import TodoService from "../service/TodoService";
    import UserPhotoService from "../service/UserPhotoService";
    
    import User from "../models/User";
    import Todo from "../models/Todo";
    import UserPhoto from "../models/UserPhoto";
    
    @Component
    export default class UserHome extends Vue {
        @Prop({
            default: ""
        })
        private user: string;
    
        @Inject private todoService: TodoService;
        @Inject private userPhotoService: UserPhotoService;
    
        private hasMoreTodo: boolean = true;
    
        private hasMoreUserPhoto: boolean = true;
    
        private todo: Todo = new Todo();
    
        private todos: Array < Todo > = [];
    
        private userPhoto: UserPhoto = new UserPhoto();
    
        private userPhotos: Array < UserPhoto > = [];
    
        private userData: User = new User();
    
        private created() {
            try {
                const userJson = JSON.parse(this.user);
    
                this.userData = Deserialize(userJson, User);
                this.userPhoto.user_id = this.todo.user_id = this.userData.id;
            } catch (err) {
                console.log(err);
            }
        }
    
        private async mounted() {
            await this.doFind();
            await this.doFind(0, true);
        }
    
        private async doFind(offset: number = 0, isPhoto: boolean = false) {
            const user_id = this.userData.id;
            const param = {
                offset,
                limit: environment["LIMIT"],
                user_id
            };
    
            if (isPhoto) {
                const userPhotos = await this.userPhotoService.find(param);
                this.userPhotos = Array.isArray(userPhotos) ? userPhotos : [];
                this.hasMoreUserPhoto =
                    userPhotos.length % environment["LIMIT"] === 0 &&
                    userPhotos.length !== 0;
            } else {
                const todos = await this.todoService.find(param);
                this.todos = Array.isArray(todos) ? todos : [];
                this.hasMoreTodo =
                    todos.length % environment["LIMIT"] === 0 && todos.length !== 0;
            }
        }
    
        private resetEntity() {
            this.userPhoto = new UserPhoto();
            this.todo = new Todo();
            this.userPhoto.user_id = this.todo.user_id = this.userData.id;
        }
    
        private async doSave(isUpload: boolean = false) {
            this.todoService.returnWithStatus = true;
            this.userPhotoService.returnWithStatus = true;
            let response = null;
    
            if (this.userPhoto.file && isUpload) {
                response = await this.userPhotoService.save(this.userPhoto);
            } else if (!isUpload && this.todo.valid()) {
                response = await this.todoService.save(this.todo);
            }
    
            if (response.status === Constant.STATUS.API.SAVE_SUCCESS) {
                this.doFind(0, !isNil(this.userPhoto.file));
            }
    
            this.todoService.returnWithStatus = false;
            this.userPhotoService.returnWithStatus = false;
        }
    }
</script>