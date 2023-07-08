<script setup lang="ts">
import { ref } from 'vue'

const checked = ref(true)
const name = ref('')
</script>

<template>
  <main>
    <h1 class="text-3xl font-bold underline">
      Add a Strain
    </h1>

    <label for="name">Name</label>
    <input v-model="name" placeholder="Name" />

    <input type="checkbox" id="checkbox" v-model="checked" />
    <label for="checkbox">{{ checked }}</label>
  </main>
</template>
