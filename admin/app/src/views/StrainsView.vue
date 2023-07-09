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
axios.get('/api/strains')
  .then((response) => {
    console.log(response.data)
    strains.value = response.data
  })
  .catch((error) => {
    console.log(error)
  })

function formatDate(date: string) {
  return moment(date, 'YYYY-MM-DD').fromNow();
}
</script>

<template>
  <div class="px-6 max-w-4xl mx-auto">
  <h1 class="text-3xl font-bold leading-tight text-gray-700 mb-6">
    Strains
  </h1>
  <p class="text-xs text-gray-400 mb-6">
    The current strains you have listed as available are below. You can edit or delete them from here.
  </p>

  <table v-if="strains" class="table-auto w-full">
    <thead class="border-b mb-3">
      <tr>
        <th v-for="header in tableHeaders" class="text-gray-700 text-xs font-bold uppercase py-2">{{ header }}</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="strain in strains" class="border-b mb-3 last:border-b-0">
        <td class="text-gray-700 py-2">{{ strain.name }}</td>
        <td class="text-gray-700 py-2">{{ strain.category }}</td>
        <td class="text-gray-700 py-2">{{ strain.genetics }}</td>
        <td class="text-gray-700 py-2">{{ strain.thc }}%</td>
        <td class="text-gray-700 py-2">{{ strain.terpenes }}%</td>
        <td class="text-gray-700 py-2">${{ strain.price }}</td>
        <td class="text-gray-700 py-2">{{ formatDate(strain.harvest_date) }}</td>
        <td class="text-gray-700 py-2">edit</td>
      </tr>
    </tbody>
  </table>
  <p v-else class="text-gray-400 text-sm">No strains found.</p>
</div>
</template>

<style>
th {
  text-align: left;
}
</style>
