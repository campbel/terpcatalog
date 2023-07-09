<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import moment from 'moment'

class Strain {
  name: string = ''
  category: string = ''
  genetics: string = ''
  thc: number = 0
  terpenes: number = 0
  price: number = 0
  harvest_date: string = ''
}

const tableHeaders = ref([
  "Name",
  "Category",
  "Genetics",
  "THC",
  "Terpenes",
  "Price",
  "Harvest Date",
])

const strains = ref<Strain[]>([])
loadStrains();

function formatDate(date: string) {
  return moment(date, 'YYYY-MM-DD').fromNow();
}

function loadStrains() {
  axios.get('/api/strains')
    .then((response) => {
      console.log(response.data)
      strains.value = response.data
    })
    .catch((error) => {
      console.log(error)
    })
}

function deleteStrain(id: number) {
  axios.delete(`/api/strains?id=${id}`)
    .then((response) => {
      console.log(response.data)
      loadStrains()
    })
    .catch((error) => {
      console.log(error)
    })
}
</script>

<template>
  <div class="px-6 max-w-4xl mx-auto">
  <h1 class="text-3xl font-bold leading-tight text-gray-700 mb-6">
    Strains
  </h1>
  <p class="text-xs text-gray-400 mb-6">
    Your current strains are listed below.
  </p>

  <div v-if="strains && strains.length > 0"
      class="relative overflow-x-auto border border-slate-300 sm:rounded-lg">
  <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
    <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
      <tr class="bg-slate-300">
        <th v-for="header in tableHeaders" 
          scope="col" class="p-3">{{ header }}</th>
        <th></th>
      </tr>
    </thead>
    <tbody class="bg-white">
      <tr v-for="(strain, index) in strains"
        class="bg-white border-b hover:bg-gray-100">
        <td class="text-gray-700 p-3">{{ strain.name }}</td>
        <td class="text-gray-700 p-3">{{ strain.category }}</td>
        <td class="text-gray-700 p-3">{{ strain.genetics }}</td>
        <td class="text-gray-700 p-3">{{ strain.thc }}%</td>
        <td class="text-gray-700 p-3">{{ strain.terpenes }}%</td>
        <td class="text-gray-700 p-3">${{ strain.price }}</td>
        <td class="text-gray-700 p-3">{{ formatDate(strain.harvest_date) }}</td>
        <td class="text-red-400 text-right p-3">
          <button @click="deleteStrain(index)" class="border px-2 rounded-xl border-red-400 hover:bg-red-400 hover:text-white">delete</button>
        </td>
      </tr>
    </tbody>
  </table>
  </div>
  <p v-else class="text-gray-400 text-md">
    We couldn't find any strains.
    <RouterLink to="/strains/new" class="text-blue-700 hover:text-blue-400">Add a Strain</RouterLink> 
    now.
  </p>
</div>
</template>

<style>

</style>
