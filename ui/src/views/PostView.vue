<script setup lang="ts">
import { type Ref, ref } from 'vue';
import { useRoute } from "vue-router";
import { baseurl } from './config.js'
import { type Post, type Comment } from './types.js'

const route = useRoute();
const post: Ref<Post> = ref({ Id: 0, Title: "", Body: "", Rendered: "", Likes: 0 });
const comments: Ref<Comment[]> = commentsFetch()
const newComment = ref("")

function commentsFetch() {
	const comments = ref([]);
	fetch(baseurl + "/comments/" + route.params.id)
		.then(r => r.json())
		.then(d => comments.value = d);
	return comments
}


function submitComment() {
	const formData = new URLSearchParams({ Comment: newComment.value });
	fetch(baseurl + "/comments/" + route.params.id, {
		method: "POST",
		headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
		body: formData
	})
		.then(r => r.json())
		.then(d => {
			if (comments.value == null) {
				comments.value = []
			}
			comments.value.unshift(d);
			newComment.value = "";
		});
}

fetch(baseurl + "/posts/" + route.params.id)
	.then(r => r.json())
	.then(d => {
		post.value = d
	});
</script>

<template>
	<h1>{{ post.Title }}</h1>

	<span v-html="post.Rendered"></span>

	<hr>

	<input type="text" placeholder="Leave a comment" v-model="newComment">
	<button @click="submitComment">Submit</button>

	<div v-for="comment in comments" :key="comment.Id" class="comment">
		Comment {{ comment.Id }} - {{ comment.Comment }}
	</div>
</template>

<style scoped>
.comment {
	font-family: Verdana, Geneva, Tahoma, sans-serif;
	opacity: 70%;
}
</style>
