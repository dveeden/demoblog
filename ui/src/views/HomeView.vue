<script setup lang="ts">
import { ref } from 'vue';
const url = "http://127.0.0.1:8080/api/posts";

const posts = ref(null);
fetch(url)
	.then(r => r.json())
	.then(d => {
		posts.value = d
	});
</script>

<template>
	<template v-for="post in posts">
		<h1>{{ post.Title }}</h1><br>
		{{ post.Body.substring(0,100) }}
		<br>
		<router-link :to="{ name: 'post', params: { id: post.Id }}">Read more...</router-link>
	<br>
	</template>
</template>
