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

  <table v-if="strains" class="table-auto w-full">
    <thead class="border-b mb-3">
      <tr>
        <th v-for="header in tableHeaders" class="text-gray-500 text-left text-xs font-bold uppercase py-2">{{ header }}</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(strain, index) in strains" class="border-b mb-3 last:border-b-0">
        <td class="text-gray-700 text-left py-2">{{ strain.name }}</td>
        <td class="text-gray-700 text-left py-2">{{ strain.category }}</td>
        <td class="text-gray-700 text-left py-2">{{ strain.genetics }}</td>
        <td class="text-gray-700 text-left py-2">{{ strain.thc }}%</td>
        <td class="text-gray-700 text-left py-2">{{ strain.terpenes }}%</td>
        <td class="text-gray-700 text-left py-2">${{ strain.price }}</td>
        <td class="text-gray-700 text-left py-2">{{ formatDate(strain.harvest_date) }}</td>
        <td class="text-red-400 text-right py-2">
          <button @click="deleteStrain(index)" class="border px-2 rounded-xl border-red-400 hover:bg-red-400 hover:text-white">delete</button>
        </td>
      </tr>
    </tbody>
  </table>
  <p v-else class="text-gray-400 text-sm">No strains found.</p>
</div>
</template>

<style>

</style>
