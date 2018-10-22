<template>
    <div :class="`empty-states ${isVertical && 'is-vertical'}`">
        <span><i class="mdi mdi-timer-sand-empty"></i></span>
        <p>
            Tidak ada lagi Konten Untuk Dimuat
        </p>
    </div>
</template>

<script lang="ts">
    import {
        Vue,
        Component,
        Prop
    } from "annotation";
    
    @Component
    export default class EmptyStates extends Vue {
        @Prop({
            default: false
        })
        private isVertical: boolean;
    }
</script>

