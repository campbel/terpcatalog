<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import moment from 'moment'

class Producer {
  id: string = ''
  name: string = ''
  email: string = ''
  website: string = ''
  address: string = ''
  phone: string = ''
}

const producers = ref<Producer[]>([])

function load() {
  axios.get('/api/producers')
    .then((response) => {
      console.log(response.data)
      producers.value = response.data
    })
    .catch((error) => {
      console.log(error)
    })
}

function del(id: string) {
  axios.delete(`/api/producers?id=${id}`)
    .then((response) => {
      console.log(response.data)
      load()
    })
    .catch((error) => {
      console.log(error)
    })
}
</script>

<template>
  <div class="px-6 max-w-4xl mx-auto">

    <div class="flex justify-between items-end mb-6">
      <div>
        <h1 class="text-3xl font-bold leading-tight text-gray-700 mb-3">
          Producers
        </h1>
        <p class="text-xs text-gray-400">
          Your current producers are listed below.
        </p>
      </div>
      <div>
        <RouterLink to="/producers/new"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
          New Producer
        </RouterLink>
      </div>
  </div>

  <div v-if="producers && producers.length > 0"
      class="relative overflow-x-auto border border-slate-300 sm:rounded-lg">
  <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
    <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
      <tr class="bg-slate-300">
        <th scope="col" class="p-3">Name</th>
        <th scope="col" class="p-3">Email</th>
        <th scope="col" class="p-3">Address</th>
        <th scope="col" class="p-3">Phone</th>
        <th scope="col" class="p-3"></th>
      </tr>
    </thead>
    <tbody class="bg-white">
      <tr v-for="(producer, index) in producers"
        class="bg-white border-b hover:bg-gray-100">
        <td class="text-gray-700 p-3">{{ producer.name }}</td>
        <td class="text-gray-700 p-3">{{ producer.email }}</td>
        <td class="text-gray-700 p-3">{{ producer.address }}</td>
        <td class="text-gray-700 p-3">{{ producer.phone }}%</td>
        <td class="text-red-400 text-right p-3">
          <button @click="del(producer.id)" class="border px-2 rounded-xl border-red-400 hover:bg-red-400 hover:text-white">delete</button>
        </td>
      </tr>
    </tbody>
  </table>
  </div>
  <p v-else class="text-gray-400 text-md">
    We couldn't find any producers.
  </p>
</div>
</template>

<style>

</style>
