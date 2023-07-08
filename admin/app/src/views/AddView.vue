<script setup lang="ts">
import { ref } from 'vue'

const categories = ref([
  "Indica",
  "Sativa",
  "Hybrid",
])
const newStrain = ref({
  "name": "",
  "category": "",
  "genetics": "",
  "thc": 0,
  "terpenes": 0,
  "price": 0,
  "harvest_date": "",
})

function onSubmit() {
  
}

</script>

<template>
  <main>
    <h1 class="text-3xl space-y-1 font-bold leading-tight text-gray-700">
      Add a Strain
    </h1>
    <form @submit.prevent="onSubmit" class="w-full max-w-lg">

      <!-- Row 1-->
      <div class="flex flex-wrap -mx-3 mb-3">

        <!-- Name -->
        <div class="w-full md:w-1/2 px-3 mb-6 md:mb-0">
          <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-name">
            Name
          </label>
          <input v-model="newStrain.name" class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white" id="grid-name" type="text" placeholder="Green Blaze">
        </div>

        <!-- Category -->
        <div class="w-full md:w-1/2 px-3 mb-6 md:mb-0">
          <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-state">
            Category
          </label>
          <div class="relative">
            <select v-model="newStrain.category" class="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-state">
              <option v-for="option in categories" :key="option" :value="option">
                {{ option }}
              </option>
            </select>
            <div class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
              <svg class="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"/></svg>
            </div>
          </div>
        </div>

      </div>
  
      <!-- Row 2-->
      <div class="flex flex-wrap -mx-3 mb-3">
        <div class="w-full px-3">
          <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-genetics">
            Genetics
          </label>
          <textarea v-model="newStrain.genetics" class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-genetics"></textarea>
          <p class="text-gray-600 text-xs italic">Make it as long and as crazy as you'd like</p>
        </div>
      </div>

      <!-- Row 3-->
      <div class="flex flex-wrap -mx-3 mb-3">

        <!-- THC -->
        <div class="w-full md:w-1/2 px-3 mb-6 md:mb-0">
          <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-city">
            THC% (0-100)
          </label>
          <input v-model.number="newStrain.thc" class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-city" type="number" placeholder="25.0">
        </div>

        <!-- Terpenes -->
        <div class="w-full md:w-1/2 px-3 mb-6 md:mb-0">
          <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-city">
            Terpenes (0-100)
          </label>
          <input v-model.number="newStrain.terpenes" class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-city" type="number" placeholder="25.0">
        </div>

      </div>

      <!-- Row 4-->
      <div class="flex flex-wrap -mx-3 mb-3">

        <!-- Price -->
        <div class="w-full md:w-1/2 px-3 mb-6 md:mb-0">
          <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-city">
            Price (USD)
          </label>
          <input v-model.number="newStrain.price" class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-city" type="number" placeholder="25.0">
        </div>

        <!-- Harvest Date -->
        <div class="w-full md:w-1/2 px-3 mb-6 md:mb-0">
          <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-city">
            Harvest Date
          </label>
          <input v-model="newStrain.harvest_date" class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-city" type="date" placeholder="25.0">
        </div>

      </div>

    </form>
  </main>
</template>
