<script setup lang="ts">
import { type Ref, ref } from 'vue';
const url = "http://127.0.0.1:8080/api/posts";

type Post = {
	Id: number
	Title: string
	Body: string
}

function postFetch() {
	const posts = ref([]);
	fetch(url)
		.then(r => r.json())
		.then(d => posts.value = d);
	return posts

}
const posts: Ref<Post[]>  = postFetch()

function like(postId: number) {
	console.log("Liking " + postId);
}
</script>

<template>
	<template v-for="post in posts" :key="post.Id">
		<h1><font-awesome-icon :icon="['far', 'newspaper']" /> {{ post.Title }}</h1><br>
		{{ post.Body.substring(0,100) }}
		<br>
		<button @click="like(post.Id)"><font-awesome-icon :icon="['far', 'thumbs-up']" /> Like</button>&nbsp;
		<router-link :to="{ name: 'post', params: { id: post.Id }}">Read more...</router-link>
	<br>
	</template>
</template>