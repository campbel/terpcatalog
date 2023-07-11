<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';

const strains = ref([]);
const producers = ref([]);
const cart = ref([]);

axios.get('/api/strains')
  .then((response) => {
    console.log(response.data)
    strains.value = response.data
  })
  .catch((error) => {
    console.log(error)
  })

axios.get('/api/strains')
  .then((response) => {
    console.log(response.data)
    producers.value = response.data
  })
  .catch((error) => {
    console.log(error)
  })

</script>

<template>
  <main>

    <div class="bg-white m-auto">
      <div class="mx-auto max-w-2xl px-4 py-16 sm:px-6 sm:py-24 lg:max-w-7xl lg:px-8">
        <h2 class="text-2xl font-bold tracking-tight text-gray-900">Available Strains</h2>

        <!-- Strain -->
        <div v-for="strain in strains" :key="strain.id"
          class="mt-6 grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
          <div class="group relative">
            <div class="aspect-h-1 aspect-w-1 w-full overflow-hidden rounded-md bg-gray-200 lg:aspect-none group-hover:opacity-75 lg:h-80">
              <img :src="strain.images[0]" class="h-full w-full object-cover object-center lg:h-full lg:w-full">
            </div>
            <div class="mt-4 flex justify-between">
              <div>
                <h3 class="text-sm text-gray-700">
                  <a href="#">
                    <span aria-hidden="true" class="absolute inset-0"></span>
                    {{ strain.name }}
                  </a>
                </h3>
                <p class="mt-1 text-sm text-gray-500">{{ strain.genetics }}</p>
              </div>
              <p class="text-sm font-medium text-gray-900">{{ strain.price }}</p>
            </div>
          </div>

        </div>
      </div>
    </div>
  </main>
</template>
