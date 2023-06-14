<script setup lang="ts">
import axios from "axios";
import { defineComponent, onMounted, ref } from "vue";
import { NButton, NInput, NList, NListItem, NCard, NIcon } from "naive-ui";

const todos = ref([] as any[]);
const content = ref("");

onMounted(() => {
  loadTodos();
});

function loadTodos() {
  axios.get("http://127.0.0.1:8080/todo/").then((res) => {
    console.log(res.data["list"]);
    todos.value = res.data["list"];
  });
}
function addTodo() {
  // if (!text.value) return;
  console.log(content.value);
  // axios
  //   .post("http://127.0.0.1:8080/todo/add", {
  //     title: text.value,
  //   })
  //   .then((_) => {
  //     loadTodos();
  //   });
}
function done(todo: object) {
  console.log(todo);
  axios.post("http://127.0.0.1:8080/todo/update", todo).then((_) => {
    loadTodos();
  });
}

defineComponent({
  components: {
    NButton,
    NInput,
    NList,
    NListItem,
    NCard,
    NIcon,
  },
});
</script>

<template>
  <n-card>
    count is {{ todos.length }}
    <n-list>
      <n-list-item>
        <n-input v-model="content" />
        <n-button @click="addTodo">添加</n-button></n-list-item
      >
      <n-list-item v-for="todo in todos" :key="todo.id">
        {{ todo.id }}:{{ todo.title }}
        <n-button @click="done(todo)">完成</n-button>
      </n-list-item>
    </n-list>
  </n-card>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
