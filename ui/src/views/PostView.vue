<script setup lang="ts">
import { type Ref, ref } from 'vue';
import {useRoute} from "vue-router";
const route = useRoute();
const url = "http://127.0.0.1:8080/api/posts/" + route.params.id;

type Post = {
	Id: number
	Title: string
	Body: string
}

const post: Ref<Post> = ref({Id: 0, Title: "", Body: ""});
fetch(url)
	.then(r => r.json())
	.then(d => {
		post.value = d
	});
</script>

<template>
	<h1>{{ post.Title }}</h1>
	{{ post.Body }}
</template>
