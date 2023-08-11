<script setup lang="ts">
import { reactive } from 'vue';
import { useRouter, useRoute } from 'vue-router'
import { baseurl } from './config.js'
import { type Post } from './types.js'

const router = useRouter()

const p: Post = reactive({
    Id: 0,
	Title: "",
	Body: "",
	Rendered: "",
    Likes: 0
})

function submitPost() {
    fetch(baseurl + '/posts', {
        method: "POST",
        body: JSON.stringify(p)
    })
    .then(() =>
        router.push('/')
    )
}
</script>

<template>
    <input type="text" placeholder="Title" v-model="p.Title"/><br>
    <textarea placeholder="Post text" v-model="p.Body"></textarea><br>
    <button @click="submitPost">Submit</button>
</template>

<style scoped>
textarea {
    outline: none;
    width: 80%;
    height: 400px;
}

input {
    outline: none;
    width: 80%;
}

button {
    float: right;
    margin-right: 20%
}
</style>